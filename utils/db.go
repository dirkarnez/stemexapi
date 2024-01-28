package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/dirkarnez/stemexapi/model"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

const PrefixCourseResourses = "Course Resources"
const PrefixStudentResourses = "Student Resources"

func GenerateObjectKey(fileName string) string {
	extension := filepath.Ext(fileName)
	return fmt.Sprintf("%s-%d%s", fileName, time.Now().UnixNano(), extension)
}

func SaveUpload(fileHeader *multipart.FileHeader, prefixes []string, s3 *StemexS3Client, db *gorm.DB, ctx iris.Context) (*model.File, error) {
	if fileHeader == nil {
		return nil, fmt.Errorf("nil fileHeader")
	}
	objectKey := GenerateObjectKey(fileHeader.Filename)

	multipartFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer multipartFile.Close()
	err = s3.UploadFile(strings.Join(append(prefixes, objectKey), "/"), multipartFile)
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
