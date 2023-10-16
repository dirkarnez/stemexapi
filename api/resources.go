package api

import (
	"log"
	"os"
	"path/filepath"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
)

func GetResourceByID(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		log.Println("*****GetResourceByID")

		ex, err := os.Getwd() //use os.Executable() in the future
		log.Println(ex)

		if err != nil {
			log.Println(">>>>>>>ERR")
		}

		// id := ctx.URLParam("id")
		// if len(id) < 1 {
		// 	ctx.StopWithStatus(http.StatusForbidden)
		// 	return
		// }filepath.Join(exPath, ""), "client.zip")
		//users.Get("/{id:int}/profile", userProfileHandler)
		log.Println("!!!!!!!!!!!!!GetResourceByID")
		ctx.ServeFile(filepath.Join(ex, "uploads", "upcoming-schedule", "appInventorMobileApps.png"))
	}
}
