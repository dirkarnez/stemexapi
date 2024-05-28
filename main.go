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
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, `AppInventor`, "AppInventor Mobile Apps", `Level 1-min.png`)
			if err != nil {
				log.Fatalln(err)
				return
			}

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`AppInventor\Introductory`,
				courseTypeDTO.ID,
				"AppInventor Mobile Apps Development Introductory",
				"icon.png",
				"App Inventor Intro Curriculum Guide.pdf",
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

			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`AppInventor\Intermediate`,
				courseTypeDTO.ID,
				"AppInventor Mobile Apps Development Intermediate",
				"Level 2-Intermediate-min.png",
				"App Inventor Intermediate A Curriculum Guide.pdf",
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
						Description: "In this project you will create a virtual “Snow Globe” that displays snowflakes falling randomly on New York City at night whenever you shake your Android device. You will be introduced to the “Any Component” advanced feature in App Inventor which is used to give collective behaviors to components of the same type",
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
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "Coding Python", "Coding Python", "Level 1-min.png")
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
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "VEX Robotics", "Vex Robotics", "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			log.Printf("%s", courseTypeDTO.Description)
		}
		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "A.I. & Machine Learning", `A.I. & Machine Learning`, "Level 1-min.png")
			if err != nil {
				log.Fatalln(err)
				return
			}
			//
			_, err = migration.AddCourse(
				q,
				s3,
				prefix,
				`A.I. & Machine Learning\Introductory`,
				courseTypeDTO.ID,
				"AppInventor Mobile Apps Development Introductory",
				"icon.png",
				"App Inventor Intro Curriculum Guide.pdf",
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
		}
		{
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "3D_Design_Printing", `3D Design & Printing`, "Level 1-min.png")
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
			courseTypeDTO, err := migration.AddCourseType(q, s3, prefix, "AR_VR", "AR/VR", "Level 1-min.png")
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

		// err = redumpparents.RedumpParents(q)
		// if err != nil {
		// 	log.Fatalln(err)
		// 	return
		// }
	}

	// if !dbInstance.Migrator().HasTable(&model.Role{}) {
	// 	// log.Println("has `users table, dropping...`")
	// 	// db.Migrator().DropTable(&User{})
	//
	// 	if err := dbInstance.Debug().Migrator().CreateTable(&model.Role{}); err != nil {
	// 		log.Fatalln(err)
	// 		return
	// 	}

	// 	if !dbInstance.Migrator().HasTable(&model.User{}) {
	// 		// log.Println("has `users table, dropping...`")
	// 		// db.Migrator().DropTable(&User{})
	// 		if err := dbInstance.Debug().Migrator().CreateTable(&model.User{}); err != nil {
	// 			log.Fatalln(err)
	// 			return
	// 		}
	// 		dbInstance.AutoMigrate()
	// 	}

	// }

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
