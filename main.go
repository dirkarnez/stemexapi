package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	casbin "github.com/casbin/casbin/v2"
	casbinModel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/dirkarnez/stemexapi/api"
	"github.com/dirkarnez/stemexapi/db"
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/migration"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/dirkarnez/stemexapi/utils"
	"github.com/gorilla/securecookie"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"gopkg.in/yaml.v2"
)

var (
	mode string
)

func middlewareAuthorizedSPA(ctx iris.Context) {
	requestPath := ctx.Path()

	if !strings.HasPrefix(requestPath, "/api/") && !strings.Contains(requestPath, ".") {
		auth, _ := sessions.Get(ctx).GetBoolean("authenticated")

		if !auth && requestPath != "/login" && requestPath != "/register" && requestPath != "/activation" && strings.HasPrefix(requestPath, "/curriculum-embeded") == false {
			ctx.Redirect("/login")
		}
		/*else if auth && requestPath == "/login" {
			ctx.Redirect("/")
		}*/
	}
	ctx.Next()
}

func middlewareAuthorizedAPI(ctx iris.Context) {
	session := sessions.Get(ctx)

	if session != nil {
		auth, _ := session.GetBoolean("authenticated")
		if !auth {
			ctx.StopWithStatus(iris.StatusUnauthorized)
		}
	}

	ctx.Next()
}

func main() {
	flag.StringVar(&mode, "mode", "update", `mode: "reinit" or "update"`)
	flag.Parse()

	httpClient := &http.Client{}

	s3 := utils.NewStemexS3Client()

	app := iris.New()
	app.Use(iris.Compression)
	//	app.Use(iris.NoCache)

	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowedMethods:   []string{"GET", "HEAD", "OPTIONS", "PUT", "PATCH", "POST", "DELETE"},
		ExposedHeaders:   []string{"X-Header"},
		MaxAge:           int((24 * time.Hour).Seconds()),
		AllowCredentials: true,
		Debug:            true,
	}))

	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.Writef("404 not found here")
	})

	hashKey := securecookie.GenerateRandomKey(64)
	blockKey := securecookie.GenerateRandomKey(32)

	sess := sessions.New(sessions.Config{
		Cookie:          "stemexsessionid",
		AllowReclaim:    true,
		CookieSecureTLS: true,
		Encoding:        securecookie.New(hashKey, blockKey),
		Expires:         30 * time.Minute,
	})

	app.Use(sess.Handler())

	var withCookieOptions = func(ctx iris.Context) {
		ctx.AddCookieOptions(iris.CookieAllowReclaim())
		ctx.AddCookieOptions(iris.CookieSecure)
		ctx.AddCookieOptions(iris.CookieHTTPOnly(true))
		ctx.AddCookieOptions(iris.CookieSameSite(http.SameSiteStrictMode))
		ctx.Next()
	}

	app.Use(withCookieOptions)

	dbInstance, dbError := db.InitConntection()
	utils.CheckError(dbError)

	dbInstance = dbInstance.Debug()

	adapter, adapterErr := gormadapter.NewAdapterByDBUseTableName(dbInstance, "", "rules")
	utils.CheckError(adapterErr)

	m := casbinModel.NewModel()
	m.AddDef("r", "r", "sub, obj, act")
	m.AddDef("p", "p", "sub, obj, act")
	m.AddDef("g", "g", "_, _")
	m.AddDef("e", "e", "some(where (p.eft == allow))")
	m.AddDef("m", "m", "g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act")

	e, enforcerErr := casbin.NewSyncedEnforcer(m, adapter)
	utils.CheckError(enforcerErr)

	e.EnableAutoSave(true)
	e.LoadPolicy()

	//factoryInstance := bo.NewFactory(dbInstance, e)

	// admin
	// sale
	// instructor
	// parent
	// prospect

	_, enforcerErr = e.AddPolicy("admin", "_", "read")
	utils.CheckError(enforcerErr)

	_, enforcerErr = e.AddPolicy("partner", "_", "read")
	utils.CheckError(enforcerErr)

	_, enforcerErr = e.AddPolicy("sales", "_", "read")
	utils.CheckError(enforcerErr)

	_, enforcerErr = e.AddPolicy("instructor", "_", "read")
	utils.CheckError(enforcerErr)

	_, enforcerErr = e.AddPolicy("parent", "_", "read")
	utils.CheckError(enforcerErr)

	_, enforcerErr = e.AddPolicy("prospect", "_", "read")
	utils.CheckError(enforcerErr)

	// t, enforcerErr = e.AddRoleForUser("alice" /* uuid pk for user */, "admin")
	// if enforcerErr != nil {
	// 	log.Fatal(enforcerErr.Error())
	// }

	//fmt.Println(t)

	fmt.Println(e.GetAllRoles())

	// // Check if a user has a role
	// hasRole, enforcerErr := e.HasRoleForUser("alice", "admin")
	// if enforcerErr != nil {
	// 	log.Fatal(enforcerErr)
	// }
	// if hasRole {
	// 	log.Println("alice has the admin role")
	// } else {
	// 	log.Println("alice does not have the admin role")
	// }

	// authorized, _ := e.Enforce("alice", "data1", "read")

	// roles, _ := e.GetRolesForUser("alice")
	// log.Println("!!!!!!done", roles)

	if mode == "reinit" {

		dbInstance.Migrator().DropTable(model.AllTables...)
		dbInstance.AutoMigrate(model.AllTables...)

		var sales = model.Role{Name: "sales"}
		if err := dbInstance.Create(&sales).Error; err != nil {
			log.Fatalln(err)
			return
		}

		var prospect = model.Role{Name: "prospect"}
		if err := dbInstance.Create(&prospect).Error; err != nil {
			log.Fatalln(err)
			return
		}

		var instructor = model.Role{Name: "instructor"}
		if err := dbInstance.Create(&instructor).Error; err != nil {
			log.Fatalln(err)
			return
		}

		// var partner = model.Role{Name: "partner"}
		// if err := dbInstance.Create(&partner).Error; err != nil {
		// 	log.Fatalln(err)
		// 	return
		// }

		// if err := dbInstance.Create(&model.User{FullName: "Singapore company 1", UserName: "singapore1", Password: "stemex", RoleID: &partner.ID, IsActivated: true}).Error; err != nil {
		// 	log.Println("?????????????????????????????")
		// 	log.Fatalln(err)
		// 	return
		// }

		// if err := dbInstance.Create(&model.User{FullName: "Singapore company 2", UserName: "singapore2", Password: "stemex", RoleID: &partner.ID, IsActivated: true}).Error; err != nil {
		// 	log.Println("?????????????????????????????")
		// 	log.Fatalln(err)
		// 	return
		// }

		// if err := dbInstance.Create(&model.User{FullName: "Singapore company 3", UserName: "singapore3", Password: "stemex", RoleID: &partner.ID, IsActivated: true}).Error; err != nil {
		// 	log.Println("?????????????????????????????")
		// 	log.Fatalln(err)
		// 	return
		// }

		// if err := dbInstance.Create(&model.User{FullName: "Singapore company 4", UserName: "singapore4", Password: "stemex", RoleID: &partner.ID, IsActivated: true}).Error; err != nil {
		// 	log.Println("?????????????????????????????")
		// 	log.Fatalln(err)
		// 	return
		// }

		// if err := dbInstance.Create(&model.User{FullName: "Singapore company 5", UserName: "singapore5", Password: "stemex", RoleID: &partner.ID, IsActivated: true}).Error; err != nil {
		// 	log.Println("?????????????????????????????")
		// 	log.Fatalln(err)
		// 	return
		// }

		if err := dbInstance.Create(&model.User{FullName: "Jovy", UserName: "jovy", Password: "stemex", Email: "jovy@stemex.org", RoleID: &sales.ID, IsActivated: true}).Error; err != nil {
			log.Println("?????????????????????????????")
			log.Fatalln(err)
			return
		}

		var admin = model.Role{Name: "admin"}
		if err := dbInstance.Create(&admin).Error; err != nil {
			log.Fatalln(err)
			return
		}

		if err := dbInstance.Create(&model.User{FullName: "Joe", UserName: "joe", Password: "stemex", Email: "joe@stemex.org", RoleID: &admin.ID, IsActivated: true}).Error; err != nil {
			log.Println("?????????????????????????????")
			log.Fatalln(err)
			return
		}

		if err := dbInstance.Create(&model.User{FullName: "prospect123", UserName: "prospect123", Email: "prospect123@stemex.org", Password: "stemex", RoleID: &prospect.ID, IsActivated: true}).Error; err != nil {
			log.Println("?????????????????????????????")
			log.Fatalln(err)
			return
		}

		var parent = model.Role{Name: "parent"}
		if err := dbInstance.Create(&parent).Error; err != nil {
			log.Fatalln(err)
			return
		}

		// if err := dbInstance.Create(&model.User{FullName: "Loretta Leung", UserName: "leungloretta", Password: "stemex", RoleID: &parent.ID, IsActivated: true}).Error; err != nil {
		// 	log.Println("?????????????????????????????")
		// 	log.Fatalln(err)
		// 	return
		// }

		if err := dbInstance.Create(&model.CurriculumCourseLessonResourceType{Name: "presentation_notes"}).Error; err != nil {
			log.Println("?????????????????????????????")
			log.Fatalln(err)
			return
		}

		if err := dbInstance.Create(&model.CurriculumCourseLessonResourceType{Name: "student_notes"}).Error; err != nil {
			log.Println("?????????????????????????????")
			log.Fatalln(err)
			return
		}

		if err := dbInstance.Create(&model.CurriculumCourseLessonResourceType{Name: "teacher_notes"}).Error; err != nil {
			log.Println("?????????????????????????????")
			log.Fatalln(err)
			return
		}

		if err := dbInstance.Create(&model.CurriculumCourseLessonResourceType{Name: "misc_materials"}).Error; err != nil {
			log.Println("?????????????????????????????")
			log.Fatalln(err)
			return
		}

		q := query.Use(dbInstance)

		prefix := fmt.Sprintf(`%s\Downloads\stemex-curriculum`, os.Getenv("USERPROFILE"))
		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "3D_Design_Printing", `3D Design & Printing`, "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`3D_Design_Printing\Introductory`,
				courseTypeDTO.ID,
				`3D Design & Printing Introductory`,
				"Level 2-Introductory-min.jpg",
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: "Why Take Up STEM?", ExternalURL: "https://hk.stemex.org/why-take-up-stem/"},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: "https://www.youtube.com/watch?v=2JnQIQFUaUw"},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    "Level 3-Introductory-min.png",
						Title:       `3D Design & Printing Introductory`,
						Description: `Learn to create your very own 3D structures that can be printed in the future. At the same time, students will be able to learn about how to use TinkerCAD and its various tools, such as alignment tools and hole generation.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`3D_Design_Printing\Intermediate`,
				courseTypeDTO.ID,
				"3D Design & Printing Intermediate",
				"Level 2-Intermediate-min.jpg",
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: "Why Take Up STEM?", ExternalURL: "https://hk.stemex.org/why-take-up-stem/"},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: "https://www.youtube.com/watch?v=2JnQIQFUaUw"},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Superhero-min.png`,
						Title:       `Superhero`,
						Description: `In this topic, student will create their superhero with their designed costume, wings and decoration through learning mirror tool, rotation, wokrplane. Student will also revisiting Boolean addition, duplication, scaling and grouping`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Intermediate B-min.png`,
						Title:       `Architecture`,
						Description: `In this topic of 3D printing, students will learn more about the functions of TinkerCAD, such as scaling and aligning objects in architecture design. Through learning different style of famous architectures. student will have an opportunity to create a Japanese style architecture and Roman Dome with columns architecture for their masterpieces.`,
					},
					{
						Name:        "C",
						IconPath:    `Level 3-Transport-min.png`,
						Title:       `Transport`,
						Description: `In this topic, students will consolidate their knowledge gained in Tickercad to create their own car using the balloon connector and to design the best boat by exploring buoyancy designed Sea Craft. Student will unleash their creativities from planning, designing and to the building process.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "A.I. & Machine Learning", `A.I. & Machine Learning`, "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`A.I. & Machine Learning\Introductory`,
				courseTypeDTO.ID,
				`A.I. & Machine Learning Introductory`,
				`Level 2-Introductory-min.png`,
				utils.GetStringPointer(`AI Machine Learning Introductory Curriculum Guide.pdf`),
				[]dto.CurriculumCourseBlogEntries{
					{Title: `Why Take Up STEM?`, ExternalURL: `https://hk.stemex.org/why-take-up-stem/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=n-IOSJCYJyM`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Introductory A-min.png`,
						Title:       `Chameleon`,
						Description: `In this project kids will make a chameleon that changes color to match its background using a webcam to take pictures of different colors, then use machine learning with those examples to train the chameleon to recognize colors.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Introductory B-min.png`,
						Title:       `Shoot your goal`,
						Description: `In this project you will train a computer to play a simple arcade game. The game is based on shooting balls at a target. You can't aim at the target directly because there is a wall in the way, so you need to bounce the ball off a wall to do it. You will teach the computer to be able to play this game by collecting examples of shots that hit and miss, so that it can learn to make predictions about the shots it can take.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`A.I. & Machine Learning\Intermediate`,
				courseTypeDTO.ID,
				`A.I. & Machine Learning Intermediate`,
				`Level 2-Intermediate-min.png`,
				utils.GetStringPointer(`AI Machine Learning Intermediate A Curriculum Guide.pdf`),
				[]dto.CurriculumCourseBlogEntries{
					{Title: `Why Take Up STEM?`, ExternalURL: `https://hk.stemex.org/why-take-up-stem/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=n-IOSJCYJyM`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Intermediate A-min.png`,
						Title:       `Chatbots`,
						Description: `In this project you will make a chatbot that can answer questions about a topic of your choice.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Intermediate B-min.png`,
						Title:       `Zombie Escape!`,
						Description: `You are trapped in a huge hotel that has been overrun by zombies! To help you escape, you have a small remote-controlled robot.There's no point trying to use it to memorize where the zombies are - there are too many rooms and too many zombies, and they're all moving around the hotel too much anyway. You need to make your robot learn.`,
					},
					{
						Name:        "C",
						IconPath:    `Level 3-Inermediate C-min.png`,
						Title:       `Secret Code`,
						Description: `In this project you will train the computer to understand secret code Words. You'll use that to say commands to a spy to guide it around a town.`,
					},
					{
						Name:        "D",
						IconPath:    `Level 3-Intermediate D-min.png`,
						Title:       `Laser Eyes`,
						Description: `In this project you will make voice-powered laser eyes! This is a game where you will see laser beams shoot from your eyes in your computer's webcam. You will use these to shoot at bottles. Your laser eyes will be voice-activated, so you will have to shout “laser eyes” to make them shoot. You will be using two kinds of machine learning model. Speech recognition to activate the lasers and face detection to aim them!`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`A.I. & Machine Learning\Advanced`,
				courseTypeDTO.ID,
				`A.I. & Machine Learning Advanced`,
				`Level 2-Advanced-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `Why Take Up STEM?`, ExternalURL: `https://hk.stemex.org/why-take-up-stem/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=n-IOSJCYJyM`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Advanced A-min.png`,
						Title:       `Top Trumps`,
						Description: `In this project you will train a computer to play a card game. For some values on the cards, you win by having the highest number. For others, you win by having the lowest. The range of numbers for different values will vary. The aim will be for the computer to learn how to play the game well without you having to give it a list of all the cards or tell it the rules.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Advanced B-min.png`,
						Title:       `Phishing`,
						Description: `People are sent links to these fake phishing websites in emails or instant messages. How can they know if a link is safe to click on? In this project, you will learn about the research that is being done to train machine learning systems to predict if a link is to a phishing website or a legitimate website.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, `AppInventor`, "AppInventor Mobile Apps", `Level 1-min.png`)
			if err != nil {
				log.Fatalln(err)
				return
			}

			curriculumPlanPath := "App Inventor Intro Curriculum Guide.pdf"
			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`AppInventor\Introductory`,
				courseTypeDTO.ID,
				"AppInventor Mobile Apps Development Introductory",
				"icon.png",
				&curriculumPlanPath,
				[]dto.CurriculumCourseBlogEntries{
					{Title: "從小培養孩子的自控能力 3款提升自控能力的電子應用程式", ExternalURL: "https://hk.stemex.org/self-control-app/"},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: "https://www.youtube.com/watch?v=zbpzr_hYwtg"},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    "level_a_icon.png",
						Title:       "HelloPurr: Tap the Kitty, Hear Him Meow",
						Description: `HelloPurr is a simple app that you can build in a very fun way. You will create a button that has a picture of your favorite cat on it, and then program the button so that when it is clicked a "meow" sound plays with some vibrations.`,
					},
					{
						Name:        "B",
						IconPath:    "level_b_icon.png",
						Title:       "Piccall",
						Description: "PicCall shows you how you can use App Inventor to make apps that do actual things, like calling friends. We will learn about how real-life applications work and are programmed.",
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			curriculumPlanPath = "App Inventor Intermediate A Curriculum Guide.pdf"
			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`AppInventor\Intermediate`,
				courseTypeDTO.ID,
				"AppInventor Mobile Apps Development Intermediate",
				"Level 2-Intermediate-min.png",
				&curriculumPlanPath,
				[]dto.CurriculumCourseBlogEntries{
					{Title: "從小培養孩子的自控能力 3款提升自控能力的電子應用程式", ExternalURL: "https://hk.stemex.org/self-control-app/"},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: "https://www.youtube.com/watch?v=zbpzr_hYwtg"},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    "Level 3-Intermediate A-min.png",
						Title:       "QuizMe",
						Description: `You will design the quiz game so that the user proceeds from question to question by clicking a Next button and receives simple correct/incorrect feedback on each answer`,
					},
					{
						Name:        "B",
						IconPath:    "Level 3-Intermediate B-min.png",
						Title:       "Snow Globe",
						Description: `In this project you will create a virtual "Snow Globe" that displays snowflakes falling randomly on New York City at night whenever you shake your Android device. You will be introduced to the “Any Component” advanced feature in App Inventor which is used to give collective behaviors to components of the same type`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`AppInventor\Advanced`,
				courseTypeDTO.ID,
				"AppInventor Mobile Apps Development Advanced",
				"Level 2-Advanced-min.png",
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: "從小培養孩子的自控能力 3款提升自控能力的電子應用程式", ExternalURL: "https://hk.stemex.org/self-control-app/"},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: "https://www.youtube.com/watch?v=zbpzr_hYwtg"},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    "Level 3-Advance A-min.png",
						Title:       "Android, Where's My Car?",
						Description: `Time to use the advanced features for app inventors to remember where you parked your car in case you go to a new location and are not familiar with it. With your very own app and your mobile device we can pinpoint and remember it using the sensors in our devices.`,
					},
					{
						Name:        "B",
						IconPath:    "Level 3-Advance B-min.png",
						Title:       "Firebase Authentication in App Inventor Using Javascript",
						Description: `The kids will learn what firebase is and set up for it and how we use it for authentication purposes in google and update any number of apps with fresh data, how data is managed in it.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "AR_VR", "AR/VR", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`AR_VR\Introductory`,
				courseTypeDTO.ID,
				`AR/VR Introductory`,
				`Level 2-Introductory-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `Why Take Up STEM?`, ExternalURL: `https://hk.stemex.org/why-take-up-stem/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=3Umu5vidiGw`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Introductory A-min.png`,
						Title:       `Interesting Zoo`,
						Description: `Our AR/VR course keeps student abreast of the latest technology in STEM learning. Students will have knowledges about the basic applications of AR/VR technologies. They will engage in creations of three-dimensional scenes and even games with Cospaces, in which they can develop their spatial sense and design thinking skills. Here students will have a basic understanding of VR and AR and make a zoo with a variety kinds of animals that are animated`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Introductory B-min.png`,
						Title:       `Comfort Home`,
						Description: `Students consolidate their knowledge in CoSpaces through creating a special meal in special occasion of their own choice. They will be thrilled to have their own AR project and share with their family members or friends. Here students apply their knowledge in manipulating objects to design and create a comfort home on a floor plan like an interior designer`,
					},
					{
						Name:        "C",
						IconPath:    `Level 3-Introductory C-min.png`,
						Title:       `Creative Story Remix`,
						Description: `This course is suitable for kids who have experinece in CoSpaces and would like to challenge themselves. They need to rewrite novel stories with creativity and illustrate in an immersive VR environment.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`AR_VR\Intermediate`,
				courseTypeDTO.ID,
				`AR/VR Intermediate`,
				`Level 2-Intermediate-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `Why Take Up STEM?`, ExternalURL: `https://hk.stemex.org/why-take-up-stem/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=3Umu5vidiGw`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Intermediate A-min.png`,
						Title:       `Interactive Art (Starry Night)`,
						Description: `For these course students should have previous experience in working in CoSpaces. They are going to further explore on the potentials of VR/AR in different areas. Introduce some background about Van Gogh and his last drawing Starry Night. Student depict the artworks with creativity in form of interactive Art.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Intermediate B.png`,
						Title:       `Interactive Art (Pumpkin)`,
						Description: ` Introduce some background about Japanese Artist Yayoi Kusama. Her iconic dots in every pieces of her works. She said that " Keep creating artworks make me happy"`,
					},
					{
						Name:        "C",
						IconPath:    `Level 3-Intermediate C-min.png`,
						Title:       `Interactive Art (Self Portrait)`,
						Description: `This is an extension for those who challenge themselves and have great interest in exploring the possibilities of using VR integrating into artworks. They need to review some self portraits of great artists and create a VR of their own.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "Coding Minecraft", "Coding Minecraft", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Minecraft\Elementary`,
				courseTypeDTO.ID,
				`Coding Minecraft Elementary`,
				`Level 2-Elementary-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `JAVA YOUR WAY THROUGH MINECRAFT!`, ExternalURL: `https://hk.stemex.org/java-your-way-through-minecraft/`},
					{Title: `Minecraft 編程 - 怎樣提升孩子`, ExternalURL: `https://hk.stemex.org/minecraft-kids/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=0SLnKsFWwFA`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `elementry1.png`,
						Title:       `Medieval Machinations Redstone`,
						Description: `This course will introduce students to use Redstone, electrical circuitry, in a Medieval Theme. Students make mine carts to gather resources, collaborate to build their kingdom and to defend their castle. They are going to experience a lot of creation, adventure and exploration.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Elementary B.png`,
						Title:       `Theme Park`,
						Description: `Everyone loves amusement theme parks. Students will have to navigate a number of engineering and teamwork challenges. They draft blueprints and plan for their parks build and create it collaboratively. They will play around and make it as much like the process of designing a real amusement park.`,
					},
					{
						Name:        "C",
						IconPath:    `Level 3-Elementary C.png`,
						Title:       `Travelling Into the Future`,
						Description: `Minecraft is a versatile and fantasy game in which players are immersed in a world made up of various kinds of blocks. In order to use blocks, players must gather resources from the world they are in and can use them to craft new materials, tools or potions. In this lesson, students will be introduced to Minecraft in a future world that will teach them the basics of playing Minecraft and will teach them to work as a team to overcome obstacles and build a survival area in a new area.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Minecraft\Introductory`,
				courseTypeDTO.ID,
				`Coding Minecraft Introductory`,
				`Level 2-Introductory-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `JAVA YOUR WAY THROUGH MINECRAFT!`, ExternalURL: `https://hk.stemex.org/java-your-way-through-minecraft/`},
					{Title: `Minecraft 編程 - 怎樣提升孩子`, ExternalURL: `https://hk.stemex.org/minecraft-kids/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=vMkAZw6nFK4`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Introductory-min.jpg`,
						Title:       `Sheep City`,
						Description: `Changing arrow's explosive power, bounciness of golden block, game mode and difficulties…from basic programming components to more complex changes, students will have fun changing/programming the Minecraft worlds to their like.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Minecraft\Intermediate`,
				courseTypeDTO.ID,
				`Coding Minecraft Intermediate`,
				`Level 2-Intermediate-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `JAVA YOUR WAY THROUGH MINECRAFT!`, ExternalURL: `https://hk.stemex.org/java-your-way-through-minecraft/`},
					{Title: `Minecraft 編程 - 怎樣提升孩子`, ExternalURL: `https://hk.stemex.org/minecraft-kids/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=DofLMIvBQ5k`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Intermediate A.jpg`,
						Title:       `Heist and Seek`,
						Description: `Be the player to mine the most blocks in this timed hunt for stolen goods! Watch out, there's a bank robber in town and they've hidden their stolen goods all over the map! It's your job to go head to head against the other players and find the most boxes to win! Use loops, conditionals, and timers to add players into different teams and add different rounds into a treasure hunt game.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Intermediate B-min.jpg`,
						Title:       `Spartan School`,
						Description: `Build an infinite mob arena game, to fight alone or with friends! Oh, and did we mention you'll be fighting blazes whilst you do it? Want to be the greatest Minecraft Spartan warrior of all time? Learn programming basics while creating wave after wave of mobs to fight in a Spartan training arena. Learn programming basics such as loops, methods and variables to create this mob fighting mini-game. Battle increasingly harder waves of enemies that multiply every round.`,
					},
					{
						Name:        "C",
						IconPath:    `Level 3-Intermediate C-min.jpg`,
						Title:       `EggWars`,
						Description: `Battle your friends to have the last egg still intact and become champion of the server! Details to tend to include notification on eggs being placed and eggs being broken, signals on game start, building base for eggs, etc. This course will allow you to expand your knowledge of adding rules to a PVP game, as well as learn more about structure generation and for loops.`,
					},
					{
						Name:        "D",
						IconPath:    `Level 3-Intermediate D-min.jpg`,
						Title:       `Rise Of The Werebunnies`,
						Description: `Create a multiplayer fight for survival, complete with exploding eggs! Beware! The moon is full and the werebunnies are out…This course shows you how to split players into different teams and add a scoring system. This course uses loops and conditionals to split players into different teams with different characteristics, then gives players scores when they defeat their enemy.`,
					},
					{
						Name:        "E",
						IconPath:    `Level 3-Intermediate E-min.png`,
						Title:       `Hungary Games`,
						Description: `Build your own hungry games style server to make the game how you want it to be! Adding game rules, timers and a shrinking border with code to make your Hungry Games world perilous for any victims trapped inside! This course will show you how to develop a PVP server. You will use conditionals and loops to create gamephrases so you can control the fate of your players in the arena.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Minecraft\Advanced`,
				courseTypeDTO.ID,
				`Coding Minecraft Advanced`,
				`Level 2-Advanced-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `JAVA YOUR WAY THROUGH MINECRAFT!`, ExternalURL: `https://hk.stemex.org/java-your-way-through-minecraft/`},
					{Title: `Minecraft 編程 - 怎樣提升孩子`, ExternalURL: `https://hk.stemex.org/minecraft-kids/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=DofLMIvBQ5k`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Advanced A-min.png`,
						Title:       `Haunted Mansion`,
						Description: `Troll your friend with three spooky mode to scare them in a haunted house! Feeling scary wary and want things a little odd? Come on in and join the fun and make this little mod. It's not one for the cowards, it's a spooktacular affair! With zombies, traps and so much more, you'll be a horror extraordinaire! Create three troll mods using more complex programming constructs including intervals and return types. Make a zombie track you, a leaky lava room and an infinite staircase in this spooky single player mod.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Advanced B-min.png`,
						Title:       `Flappy Block`,
						Description: `Rack up a high score and try not touch the lava in our 3D version of Floppy Bird in Minecraft. Think you've mastered Floppy Bird? Think again! Try our 3D Minecraft version, complete with the original flapping motion. This course shows you how to create unlimited obstacle courses using methods, variables and for loops, so do your best to rack up the highest score possible without touching the lava!`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Minecraft\Master`,
				courseTypeDTO.ID,
				`Coding Minecraft Master`,
				`Level 2-Master-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `JAVA YOUR WAY THROUGH MINECRAFT!`, ExternalURL: `https://hk.stemex.org/java-your-way-through-minecraft/`},
					{Title: `Minecraft 編程 - 怎樣提升孩子`, ExternalURL: `https://hk.stemex.org/minecraft-kids/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=DofLMIvBQ5k`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Master A-min.png`,
						Title:       `Dances to the Moosic`,
						Description: `Create a cow themed nightclub where your players have to dance in time to the moo-sic! Anyone who says going to a disco full of cows isn't fun has either never tried it, or is just wrong. In this course, you will be using Java syntax code to make music and getting your players to dance to the beat till the cows come home. This course involves randomizing events and adding a score system to detect the movement of players in a cow-themed dancing game.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Master B-min.png`,
						Title:       `Swoop De Loop`,
						Description: `Learn how to create rings using mathematical knowledge which give players a superboost when gliding. Swoop around the map using elytra and get a much needed boost every time you surge through a hoop! Learn how to use Java syntax to code hoops that you can place in any world!`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "Coding Python", "Coding Python", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Python\Introductory`,
				courseTypeDTO.ID,
				`Coding Python Introductory`,
				`Level 2-Introductory-min.png`,
				nil, // Python Introductory A (Level 1) Curriculum Guide.pdf and Python Introductory B (Level 2) Curriculum Guide.pdf should be combined
				[]dto.CurriculumCourseBlogEntries{
					{Title: `甚麼是Python? Python 容易被學習嗎?`, ExternalURL: `https://hk.stemex.org/?s=python`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=hh3W_tjPGlI`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Introductory A-min.png`,
						Title:       `Python Turtle`,
						Description: `What is AI? How does a machine learn? Let your kids start to know more about Python. This course is suitable for students who have a little experience in coding. They will learn and understand the Python turtle library and graphics. They will be challenged to animate a clock and to control a spinner enhancing their creativity.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Introductory B-min.png`,
						Title:       `Python Game Design`,
						Description: `Students will step into simple game design using Python. They will learn how to interact with the computer on screen or using keyboard. They are going to apply coding to design racing, word guess game plus other challenges! Sharing amongst friends and further exploration are encouraged.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Python\Intermediate`,
				courseTypeDTO.ID,
				`Coding Python Intermediate`,
				`Level 2-Intermediate-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `甚麼是Python? Python 容易被學習嗎?`, ExternalURL: `https://hk.stemex.org/?s=python`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=hh3W_tjPGlI`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Intermediate A-min.png`,
						Title:       `Python Simon Says`,
						Description: `Students are going to have a comprehensive understanding on the data structure in Python. They will advance their skills in using Python editor. Apply their understanding in data structure and engineering design process to create games as challenges.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Intermediate B-min.png`,
						Title:       `Python Flappy & Snake Games`,
						Description: `Students learn how to utilize other python resources. They will concentrate on coding the movement of objects using vectors and control. A lot of logic training will be involved in understanding the conditions in gaming and how to solve them in coding.`,
					},
					{
						Name:        "C",
						IconPath:    `Level 3-Intermediate C-min.png`,
						Title:       `Python Ping Pong`,
						Description: `In this course, students learn complicated interaction in game design. They are going to apply knowledge in python to solve geniune challenges and stunning geometric dancing figures. At the end they simulate the classic game of life revealing special patterns from simple rules.`,
					},
					{
						Name:        "D",
						IconPath:    `Level 3-Intermediate D-min.png`,
						Title:       `Python Hangman and Pacman`,
						Description: `In this course, students learn complicated interaction in game design. They are going to apply knowledge in python to solve geninue challenges and stunning geometric dancing figures. At the end they simulate the classic game of life revealing speical patterns from simple rules.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Python\Advanced`,
				courseTypeDTO.ID,
				`Coding Python Advanced`,
				`Level 2-Advance-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `甚麼是Python? Python 容易被學習嗎?`, ExternalURL: `https://hk.stemex.org/?s=python`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=hh3W_tjPGlI`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Advanced A-min.png`,
						Title:       `Python Game Concept`,
						Description: `With their deeper understanding of Python and its various structures, students will take a deep dive into the pygame module and other modules, as well as a deeper understanding into data types, to learn more about how they can be used to create more complex program without having to write everything from scratch.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Advanced B-min.png`,
						Title:       `Python Discord Bot`,
						Description: `Students will learn to creating a discord bot and how to do with language processing. In the course, students will have a chance to do some works related to the concept of neural network.`,
					},
					{
						Name:        "C",
						IconPath:    `Level 3-Advanced C-min.png`,
						Title:       `Python Complex Game Design`,
						Description: `Students will learn the most popular python module pygame. Through designing the classic space invader game, they will consolidate the skills in game design. A lot of revisions of the python coding structure and logic deductions in solving problems involved. Use of different platforms in coding, how to utilize resources in machine learning.`,
					},
					{
						Name:        "D",
						IconPath:    `Level 3-Advanced D-min.png`,
						Title:       `Python Machine Learning`,
						Description: `Students will learn how to use tensflow and the use of the module keraus. They are going to learn how to build up a machine learning module to identify different clothing from thousands of images.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Python\Master`,
				courseTypeDTO.ID,
				`Coding Python Master`,
				`Level 2-Master-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `甚麼是Python? Python 容易被學習嗎?`, ExternalURL: `https://hk.stemex.org/?s=python`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=hh3W_tjPGlI`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Master-min.png`,
						Title:       `Yolo`,
						Description: `We will learn how to use the yolov3 model to detect objects present in an image. It will help differentiate different objects.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 1-min.png`,
						Title:       `Tic tac toe`,
						Description: `We will learn about graphical user interface in this project, learn about Tkinter and use it to make a game GUI of tic tac toe and we will learn logic of the game design.`,
					},
					{
						Name:        "C",
						IconPath:    `Level 1-min.png`,
						Title:       `Brick Breaker game`,
						Description: `The kids will learn about game design theory. We will start with a simple Brick breaker game in which there is a ball that bounces off a platform to break a brick wall and the player has to keep the ball going by making sure the paddle is always there to bounce off the ball back.`,
					},
					{
						Name:        "D",
						IconPath:    `Level 1-min.png`,
						Title:       `Photo Editing App`,
						Description: `We will learn about the pillow library which is used for image processing, we will learn about how we can edit images using python.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "Coding Roblox", "Coding Roblox", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Roblox\Introductory`,
				courseTypeDTO.ID,
				`Coding Roblox Introductory`,
				`Level 2- Introductory-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `甚麼是Roblox? Coding Roblox是怎麼一回事?`, ExternalURL: `https://hk.stemex.org/%e7%94%9a%e9%ba%bc%e6%98%afroblox%ef%bc%9fcoding-roblox%e6%98%af%e6%80%8e%e9%ba%bc%e4%b8%80%e5%9b%9e%e4%ba%8b%ef%bc%9f/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=CW0RrMdB0kQ`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Introductory-min.jpg`,
						Title:       `Prison Escape`,
						Description: `Duck through the maze, avoid the flashing lasers and break out of prison armed with your trusty dynamite slingshot - just try not to get blown up! This course is great for beginners. Get to grips with the code editor and learn how to make your Roblox game from scratch using functions and conditionals. Killer Lasers - teach how to make deadly lasers Flashing Lasers - teach how to make lasers that flash on and off. Dynamite Singshot - teach how to throw dynamite and make it exploded.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`Coding Roblox\Intermediate`,
				courseTypeDTO.ID,
				`Coding Roblox Intermediate`,
				`Level 2-Intermediate-min.png`,
				nil,
				[]dto.CurriculumCourseBlogEntries{
					{Title: `甚麼是Roblox? Coding Roblox是怎麼一回事?`, ExternalURL: `https://hk.stemex.org/%e7%94%9a%e9%ba%bc%e6%98%afroblox%ef%bc%9fcoding-roblox%e6%98%af%e6%80%8e%e9%ba%bc%e4%b8%80%e5%9b%9e%e4%ba%8b%ef%bc%9f/`},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: `https://www.youtube.com/watch?v=CW0RrMdB0kQ`},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    `Level 3-Intermediate A-min.png`,
						Title:       `Box Car Racer`,
						Description: `Create a box racing game where the fastest racer wins! Students learn how to code checkpoints, a finish line and write clean code in this fast-paced box racing game! This Roblox course shows you how to structure code well and add a finish line to the track using inheritance.`,
					},
					{
						Name:        "B",
						IconPath:    `Level 3-Intermediate B-min.jpg`,
						Title:       `Wrestleball`,
						Description: `Players must face off in the wrestling arena, firing balls to knock each other out of the centre. The player who stays in the centre the longest the winner! In this course you'll learn how to make an arena-based PVP game using vectors, loops and a score system. It's perfect for confident Roblox fans who are familiar with the Code Editor.`,
					},
					{
						Name:        "C",
						IconPath:    `Level 3-Intermediate C-min.jpg`,
						Title:       `Ninja Obby`,
						Description: `Work on your particular skills while navigating this tricky obstacle course! Players must use their top ninja skills to navigate their way through the course, avoiding both the deadly obstacles and trying not to fall off the platforms. The player who manages to make it to the end can unlock a ninja outfit as a reward! With this course, you will learn some of the more complex Lua coding constructs such as vectors. You'll also improve your knowledge of the basics of coding. This course also covers game development concepts such as procedural generation.`,
					},
					{
						Name:        "D",
						IconPath:    `Level 3-Intermediate D-min.jpg`,
						Title:       `Riddle Ruin`,
						Description: `Crack the codes inside the pyramid to set yourself free! Code a series of puzzles including a secret door, combination lock and blocked corridors to trap players inside a forbidden template! This course shows you how to use code to create and change behaviour in your Roblox map. You’ll also get introduced to some key game design concepts that you can apply to future games.`,
					},
					{
						Name:        "E",
						IconPath:    `Level 3-Intermediate E-min.jpg`,
						Title:       `Platform Game Design`,
						Description: `Avoid red obstacles and collect coins in this 2D platform Roblox game. Platform Game Design is structured in a slightly different way to a lot of the other Roblox courses as it primarily focuses on game design, as opposed to practical coding. Some coding is required in the course. but we would recommend it as most suitable for confident budding programmers or experienced Roblox players with an interest in game design. This course is designed to teach you how to make a 2D platform game and learn game design theories that can be applied to any game you make in future.`,
					},
				},
			)

			if err != nil {
				log.Fatalln(err)
				return
			}
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "Coding Scratch", "Coding Scratch", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "Cyber Virtual Robotics", "Cyber Virtual Robotics", "Level 1.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "Lego Robotics", "Lego Robotics", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "Micro_bit", "Micro:bits", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "Smart City", "Smart City", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "VEX Robotics", "Vex Robotics", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)
		}

		// err = redumpparents.RedumpParents(q)
		// if err != nil {
		// 	log.Fatalln(err)
		// 	return
		// }
	}

	app.Use(func(ctx iris.Context) {
		mimeOverrides := loadMIMEOverrides()
		extension := filepath.Ext(ctx.RequestPath(true))
		mimeForExtension, ok := mimeOverrides[extension]
		if ok {
			log.Println(extension, mimeForExtension, "user-defined")
		} else {
			mimeForExtension := mime.TypeByExtension(extension)
			log.Println(extension, mimeForExtension, "iris built-in")
		}

		ctx.ContentType(mimeForExtension)
		ctx.Next()
	})

	app.Use(middlewareAuthorizedSPA)

	app.PartyFunc("/api", func(party iris.Party) {
		// party.Any("/", middlewareAuthorizedSPA, func(ctx iris.Context) {
		// 	ctx.WriteString("it works")
		// })

		party.Post("/login", api.Login(dbInstance))
		party.Post("/logout", middlewareAuthorizedAPI, api.Logout)

		party.Post("/register", api.Register(dbInstance))
		party.Post("/activation", api.Activation(dbInstance))

		party.Post("/user", middlewareAuthorizedAPI, api.CreateOrUpdateUser(dbInstance))

		party.Get("/users", middlewareAuthorizedAPI, api.GetAllUsers(dbInstance))
		party.Get("/partners", middlewareAuthorizedAPI, api.GetAllPartners(dbInstance))

		//party.Post("/users", middlewareAuthorizedAPI, api.CreateUser(factoryInstance.GetUsersBO()))

		party.Get("/roles", middlewareAuthorizedAPI, api.GetAllRoles(dbInstance))

		party.Get("/curriculum-tree", api.GetCurriculumTree(dbInstance))

		party.Post("/curriculum-course", middlewareAuthorizedAPI, api.CreateOrUpdateCurriculumCourse(s3, dbInstance))
		party.Get("/curriculum-course", api.GetCurriculumCourse(s3, dbInstance))

		party.Post("/curriculum-course-type", middlewareAuthorizedAPI, api.CreateOrUpdateCurriculumCourseType(s3, dbInstance))
		party.Get("/curriculum-course-type", api.GetCurriculumCourseType(dbInstance))

		// party.Get("/curriculum-courses", middlewareAuthorizedAPI, api.GetCurriculumCourses(dbInstance))

		party.Get("/prospect-activity", middlewareAuthorizedAPI, api.GetProspectActivities(dbInstance))
		party.Get("/parent-activity", middlewareAuthorizedAPI, api.GetParentActivities(dbInstance))
		party.Get("/internal-user-activity", middlewareAuthorizedAPI, api.GetInternalUserActivities(dbInstance))
		party.Get("/resourse", middlewareAuthorizedAPI, api.GetResourceByID(s3, dbInstance))

		party.Get("/resourse-list", middlewareAuthorizedAPI, api.GetFiles(dbInstance))
		party.Post("/upload", middlewareAuthorizedAPI, api.UploadFile(s3, dbInstance))

		party.Get("/deals/getDeal", middlewareAuthorizedAPI, api.GetDeals(httpClient))

		party.Get("/students-to-user", middlewareAuthorizedAPI, api.GetStudentsToUser(httpClient, dbInstance))
		party.Get("/student-deals", middlewareAuthorizedAPI, api.SearchDealIDList(httpClient, dbInstance))

		party.Get("/student-deal-attachments", middlewareAuthorizedAPI, api.GetAttachment(httpClient))

		party.Get("/secret", middlewareAuthorizedAPI, func(ctx iris.Context) {
			userName := sessions.Get(ctx).GetString("user_name")
			ctx.WriteString(fmt.Sprintf("Hi %s!", userName))
		})

		party.Get("/init", middlewareAuthorizedAPI, func(ctx iris.Context) {
			userName := sessions.Get(ctx).GetString("user_name")

			var user model.User
			if err := dbInstance.Where(&model.User{UserName: userName, IsActivated: true}).First(&user).Error; err != nil {
				ctx.WriteString(fmt.Sprintf("Hi %s!", userName))
				return
			}

			var rule model.Role
			var id = user.RoleID
			if err := dbInstance.First(&rule, "id = ?", id).Error; err != nil {
				ctx.JSON(iris.Map{
					"user_name": userName,
					"role":      "",
				})
				return
			} else {
				log.Println(rule.Name)

				ctx.JSON(iris.Map{
					"user_name": user.FullName,
					"role":      rule.Name,
				})
				return
			}

			// if err := dbInstance.Debug().Create(&sales).Error; err != nil {
			// 	log.Fatalln(err)
			// 	return
			// }
			// get user role
			// get

			//roles, _ := e.GetRolesForUser("alice")

			// e.Enforce("{role}", "USERS_BO", "read")
		})

		party.Any("/{any:path}", func(ctx iris.Context) {
			ctx.StatusCode(iris.StatusNotFound)
			ctx.WriteString("404 Not Found")
		})
	})

	app.HandleDir("/", iris.Dir("./public"), iris.DirOptions{
		IndexName: "index.html",
		SPA:       true,
	})

	var port int

	if runtime.GOOS == "windows" {
		port = 4437 // local development
	} else {
		port = 443
	}

	log.Printf("Listening on %d\n", port)

	//
	err := app.Run(
		// AutoTLS
		iris.TLS(fmt.Sprintf(":%d", port), "server.crt", "server.key"),
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithPostMaxMemory(32*iris.MB),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func loadMIMEOverrides() map[string]string {
	m := make(map[string]string)
	buf := new(bytes.Buffer)
	file, err := os.Open("mime.yaml")
	if err != nil {
		return m
	}
	defer file.Close()
	_, err = buf.ReadFrom(file)
	if err != nil {
		return m
	}
	err = yaml.Unmarshal(buf.Bytes(), &m)
	if err != nil {
		return m
	}
	return m
}
