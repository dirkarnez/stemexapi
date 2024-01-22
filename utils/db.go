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

func GenerateServerPhysicalFileName(originalPhysicalFileName string) string {
	extension := filepath.Ext(originalPhysicalFileName)
	return fmt.Sprintf("%s-%d%s", originalPhysicalFileName, time.Now().UnixNano(), extension)
}

func SaveUpload(fileHeader *multipart.FileHeader, db *gorm.DB, ctx iris.Context) (*model.File, error) {
	if fileHeader == nil {
		return nil, fmt.Errorf("nil fileHeader")
	}
	serverPhysicalFileName := GenerateServerPhysicalFileName(fileHeader.Filename)

	multipartFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer multipartFile.Close()
	err = NewStemexS3Client().UploadFile(fmt.Sprintf("%s/%s", "Course Resources", serverPhysicalFileName), multipartFile)
	// _, err := ctx.SaveFormFile(fileHeader, )
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
