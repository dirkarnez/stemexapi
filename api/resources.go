package api

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/dirkarnez/stemexapi/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
)

func GetResourceByID(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		ex, _ := os.Getwd() //use os.Executable() in the future

		id := ctx.URLParam("id")

		var file = model.File{}
		if len(id) > 0 {
			idUUIDex, err := model.ValidUUIDExFromIDString(id)
			if err != nil {
				ctx.StopWithError(iris.StatusInternalServerError, err)
				return
			}
			if err := dbInstance.First(&file, "`id` = ?", idUUIDex).Error; err != nil {
				ctx.StopWithError(iris.StatusInternalServerError, err)
				return
			}
		}

		// var idUUIDex model.UUIDEx

		// 	, err : model.ValidUUIDExFromIDString(id)
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		// idUUID, _ := uuid.Parse(id)

		// param := model.File{}
		// param.ID = idUUIDex

		// file := model.File{}
		// if err := dbInstance.
		// 	Model(&model.File{}).
		// 	Where(&param).
		// 	First(&file).Error; err != nil {
		// 	ctx.StopWithStatus(http.StatusNotFound)
		// 	return
		// }

		path := []string{ex, "uploads"}
		path = append(path, strings.Split(file.ServerPhysicalFileName, "/")...)
		ctx.ServeFile(filepath.Join(path...))

		// go func(filename string, ctx iris.Context) {
		// 	// linux only
		// 	// file, err := os.OpenFile(filepath.Join(path...), os.O_RDONLY|os.O_NONBLOCK, 0)
		// 	// if err != nil {
		// 	// 	// Handle the error case
		// 	// 	ctx.StatusCode(iris.StatusInternalServerError)
		// 	// 	ctx.WriteString("Error serving file")
		// 	// 	return
		// 	// }
		// 	// defer file.Close()

		// 	// Perform any necessary file I/O operations asynchronously
		// 	// ...

		// 	// Serve the file
		// 	ctx.ServeFile(filename)
		// }(filepath.Join(path...), ctx)
	}
}
