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
				``,
				courseTypeDTO.ID,
				``,
				``,
				``,
				[]dto.CurriculumCourseBlogEntries{
					{Title: ``, ExternalURL: ``},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: ``},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    ``,
						Title:       ``,
						Description: ``,
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
				``,
				courseTypeDTO.ID,
				``,
				``,
				``,
				[]dto.CurriculumCourseBlogEntries{
					{Title: ``, ExternalURL: ``},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: ``},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    ``,
						Title:       ``,
						Description: ``,
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
				``,
				courseTypeDTO.ID,
				``,
				``,
				``,
				[]dto.CurriculumCourseBlogEntries{
					{Title: ``, ExternalURL: ``},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: ``},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    ``,
						Title:       ``,
						Description: ``,
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
				``,
				courseTypeDTO.ID,
				``,
				``,
				``,
				[]dto.CurriculumCourseBlogEntries{
					{Title: ``, ExternalURL: ``},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: ``},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    ``,
						Title:       ``,
						Description: ``,
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
				``,
				courseTypeDTO.ID,
				``,
				``,
				``,
				[]dto.CurriculumCourseBlogEntries{
					{Title: ``, ExternalURL: ``},
				},
				[]dto.CurriculumCourseYoutubeVideoEntries{
					{URL: ``},
				},
				[]dto.CurriculumCourseLevels{
					{
						Name:        "A",
						IconPath:    ``,
						Title:       ``,
						Description: ``,
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
		}

		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "Coding Roblox", "Coding Roblox", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)
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
