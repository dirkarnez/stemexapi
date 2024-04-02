package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/query"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const PrefixCourseResourses = "Course Resources"
const PrefixStudentResourses = "Student Resources"

func GenerateObjectKey(prefixes []string, fileName string) string {
	extension := filepath.Ext(fileName)
	return strings.Join(append(prefixes, fmt.Sprintf("%s-%d%s", fileName[:len(fileName)-len(extension)], time.Now().UnixNano(), extension)), "/")
}

func SaveUpload(fileHeader *multipart.FileHeader, prefixes []string, s3 *StemexS3Client, db *gorm.DB, ctx iris.Context) (*model.File, error) {
	if fileHeader == nil {
		return nil, fmt.Errorf("nil fileHeader")
	}

	objectKey := GenerateObjectKey(prefixes, fileHeader.Filename)

	multipartFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer multipartFile.Close()
	err = s3.UploadFile(objectKey, multipartFile)
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

func SaveUploadV2(fileHeader *multipart.FileHeader, uuidptr *model.UUIDEx, prefixes []string, s3 *StemexS3Client, txOrQ *query.Query) (*model.File, error) {
	if fileHeader == nil {
		return nil, fmt.Errorf("nil fileHeader")
	}

	objectKey := GenerateObjectKey(prefixes, fileHeader.Filename)

	multipartFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer multipartFile.Close()
	err = s3.UploadFile(objectKey, multipartFile)
	if err != nil {
		return nil, err
	}
	var file = model.File{FileNameUploaded: fileHeader.Filename, ObjectKey: objectKey}

	if uuidptr != nil {
		if (*uuidptr) != model.UUIDEx(uuid.Nil) {
			file.ID = *uuidptr
		}
	}

	err = q.File.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&file)

	if err != nil {
		return nil, err
	} else {
		return &file, nil
	}
}
