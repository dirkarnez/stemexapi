package api

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/dirkarnez/stemexapi/model"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
)

func GetResourceByID(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		ex, _ := os.Getwd() //use os.Executable() in the future

		id := ctx.URLParam("id")
		if len(id) < 1 {
			ctx.StopWithStatus(http.StatusNotFound)
			return
		}

		idUUID, _ := uuid.Parse(id)

		param := model.File{}
		param.ID = model.UUIDEx(idUUID)

		file := model.File{}
		if err := dbInstance.
			Model(&model.File{}).
			Where(&param).
			First(&file).Error; err != nil {
			ctx.StopWithStatus(http.StatusNotFound)
			return
		}

		path := []string{ex, "uploads"}
		path = append(path, strings.Split(file.ServerPhysicalFileName, "/")...)
		ctx.ServeFile(filepath.Join(path...))
	}
}
