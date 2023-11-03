package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/dirkarnez/stemexapi/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func SaveUpload(fileHeader *multipart.FileHeader, db *gorm.DB, ctx iris.Context) (*model.File, error) {
	if fileHeader == nil {
		return nil, fmt.Errorf("nil fileHeader")
	}
	extension := filepath.Ext(fileHeader.Filename)
	serverPhysicalFileName := fmt.Sprintf("%d.%s", time.Now().UnixNano(), extension)
	_, err := ctx.SaveFormFile(fileHeader, fmt.Sprintf("./%s/%s", "uploads", serverPhysicalFileName))
	if err != nil {
		return nil, err
	}

	file := model.File{OriginalPhysicalFileName: fileHeader.Filename, ServerPhysicalFileName: serverPhysicalFileName}
	if err := db.
		Create(&file).Error; err != nil {
		return nil, err
	} else {
		return &file, nil
	}
}
