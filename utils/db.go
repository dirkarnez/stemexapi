package utils

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/dirkarnez/stemexapi/model"
)

func SaveUpload(fileHeader *multipart.FileHeader) (model.File, error) {
	extension := filepath.Ext(fileHeader.Filename)
	serverPhysicalFileName := fmt.Sprintf("%d.%s", time.Now().UnixNano(), extension)
	_, err = ctx.SaveFormFile(form.IconFile, fmt.Sprintf("./%s/%s", "uploads", serverPhysicalFileName))
	if err != nil {
		ctx.WriteString("Failed to save file: " + form.IconFile.Filename)
		return err
	}

	file := model.File{OriginalPhysicalFileName: form.IconFile.Filename, ServerPhysicalFileName: serverPhysicalFileName}
	if err := tx.
		Create(&file).Error; err != nil {
		return err
	}
}
