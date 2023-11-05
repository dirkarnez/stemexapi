package api

import (
	"github.com/dirkarnez/stemexapi/dto"
	"github.com/dirkarnez/stemexapi/model"
	"github.com/dirkarnez/stemexapi/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"gorm.io/gorm"
)

func GetFiles(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		from, err := ctx.URLParamInt64("from")
		if err != nil {
			ctx.StopWithStatus(iris.StatusBadRequest)
			return
		}
		to, err := ctx.URLParamInt64("to")
		if err != nil {
			ctx.StopWithStatus(iris.StatusBadRequest)
			return
		}

		var files []dto.File
		var count int64
		if err := dbInstance.
			Model(&model.File{}).
			Where("`seq_no` BETWEEN ? AND ?", from, to).
			Find(&files).Error; err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}

		if err := dbInstance.
			Model(&model.File{}).
			Count(&count).Error; err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			return
		}

		ctx.JSON(dto.FileManagement{
			Files:              files,
			FromSeqNoInclusive: from,
			ToSeqNoExclusive:   to,
			TotalCount:         count,
		})
	}
}

// files, err := ioutil.ReadDir("./uploads")
// if err != nil {
// 	log.Fatal(err)
// }

// var fileNames []string
// for _, file := range files {
// 	if !file.IsDir() {
// 		fileNames = append(fileNames, file.Name())
// 	}
// }

func UploadFile(dbInstance *gorm.DB) context.Handler {
	return func(ctx iris.Context) {
		uploadedFiles, _, err := ctx.UploadFormFiles("./uploads")
		if err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		}

		var fileModels []model.File
		for _, uploadedFile := range uploadedFiles {
			fileModels = append(fileModels, model.File{
				OriginalPhysicalFileName: uploadedFile.Filename,
				ServerPhysicalFileName:   utils.GenerateServerPhysicalFileName(uploadedFile.Filename),
			})
		}

		if err := dbInstance.
			CreateInBatches(&fileModels, 100).Error; err != nil {
			ctx.StopWithStatus(iris.StatusInternalServerError)
			return
		} else {
			ctx.JSON(iris.Map{
				"status": 200,
			})
			return
		}

		// // Get the file from the dropzone request
		// file, info, err := ctx.FormFile("file")
		// if err != nil {
		// 	ctx.StatusCode(iris.StatusInternalServerError)
		// 	ctx.Application().Logger().Warnf("Error while uploading: %v", err.Error())
		// 	return
		// }

		// defer file.Close()
		// fname := info.Filename

		// // Create a file with the same name
		// // assuming that you have a folder named 'uploads'
		// out, err := os.OpenFile("/uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
		// if err != nil {
		// 	ctx.StatusCode(iris.StatusInternalServerError)
		// 	ctx.Application().Logger().Warnf("Error while preparing the new file: %v", err.Error())
		// 	return
		// }
		// defer out.Close()

		// io.Copy(out, file)

		// ctx.WriteString("Files uploaded successfully")

	}
}
