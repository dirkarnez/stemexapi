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

func GenerateObjectKey(originalPhysicalFileName string) string {
	extension := filepath.Ext(originalPhysicalFileName)
	return fmt.Sprintf("%s-%d%s", originalPhysicalFileName, time.Now().UnixNano(), extension)
}

func SaveUpload(fileHeader *multipart.FileHeader, s3 *StemexS3Client, db *gorm.DB, ctx iris.Context) (*model.File, error) {
	if fileHeader == nil {
		return nil, fmt.Errorf("nil fileHeader")
	}
	objectKey := GenerateObjectKey(fileHeader.Filename)

	multipartFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer multipartFile.Close()
	err = s3.UploadFile(fmt.Sprintf("%s/%s", "Course Resources", objectKey), multipartFile)
	// _, err := ctx.SaveFormFile(fileHeader, )
	if err != nil {
		return nil, err
	}
	file := model.File{FileNameUploaded: fileHeader.Filename, ObjectKey: objectKey}
	if err := db.
		Create(&file).Error; err != nil {
		return nil, err
	} else {
		return &file, nil
	}
}
