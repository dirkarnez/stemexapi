package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/albrow/forms"
	"github.com/gin-gonic/gin/binding"
)

/*
Not working for nested structure
*/
func FormMultipartParse(req *http.Request, stuctPointer any) error {
	return binding.FormMultipart.Bind(req, stuctPointer)
}

func FormMultipartParseV2[T any](req *http.Request, stuctPointer *T) error {
	// Parse request data.
	userData, err := forms.Parse(req)
	if err != nil {
		// Handle err
		// ...
		// ...
		return nil
	}

	// Validate
	val := userData.Validator()
	val.Require("username")
	val.LengthRange("username", 4, 16)
	val.Require("email")
	val.MatchEmail("email")
	val.Require("password")
	val.MinLength("password", 8)
	val.Require("confirmPassword")
	val.Equal("password", "confirmPassword")
	val.RequireFile("profileImage")
	val.AcceptFileExts("profileImage", "jpg", "png", "gif")
	if val.HasErrors() {
		// Write the errors to the response
		// Maybe this means formatting the errors as json
		// or re-rendering the form with an error message
		// ...
		return nil
	}
	return nil

	// Use data to create a user object
	// user := &models.User{
	// 	Username:       userData.Get("username"),
	// 	Email:          userData.Get("email"),
	// 	HashedPassword: hash(userData.Get("password")),
	// }

	// // Continue by saving the user to the database and writing
	// // to the response
	// // ...

	// // Get the contents of the profileImage file
	// imageBytes, err := userData.GetFileBytes("profileImage")
	// if err != nil {
	// 	// Handle err
	// }
	// Now you can either copy the file over to your server using io.Copy,
	// upload the file to something like amazon S3, or do whatever you want
	// with it.
}

// func FileToMultipartFileHeader(filePath string) (*multipart.FileHeader, error) {
// 	os.Open(filePath)

// 	result := multipart.FileHeader{
// 		Filename: "",
// 		Header:   textproto.MIMEHeader{},
// 		Size:     0,
// 	}

// 	return &result, nil
// }

// func byteToFileHeader(filePath string, filename string) (multipart.FileHeader, error) {
// 	var data []byte
// 	file := bytes.NewReader(data)

// 	multipart.NewReader(msg.Body, params["boundary"]
// 	// return multipart.FileHeader{
// 	// 	Filename: filename,
// 	// 	Size:     int64(len(data)),
// 	// 	Header: map[string][]string{
// 	// 		"Content-Type": {http.DetectContentType(data)},
// 	// 	},
// 	// 	Content: file,
// 	// }, nil
// }

func CreateMultipartFileHeader(filePath string) (*multipart.FileHeader, error) {
	// open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// create a buffer to hold the file in memory
	var buff bytes.Buffer
	buffWriter := io.Writer(&buff)

	// create a new form and create a new file field
	formWriter := multipart.NewWriter(buffWriter)
	formPart, err := formWriter.CreateFormFile("file", filepath.Base(file.Name()))
	if err != nil {
		return nil, err
	}

	// copy the content of the file to the form's file field
	if _, err := io.Copy(formPart, file); err != nil {
		return nil, err
	}

	// close the form writer after the copying process is finished
	// I don't use defer in here to avoid unexpected EOF error
	formWriter.Close()

	// transform the bytes buffer into a form reader
	buffReader := bytes.NewReader(buff.Bytes())
	formReader := multipart.NewReader(buffReader, formWriter.Boundary())

	// read the form components with max stored memory of 1MB
	multipartForm, err := formReader.ReadForm(1 << 20)
	if err != nil {
		return nil, err
	}

	// return the multipart file header
	files, exists := multipartForm.File["file"]
	if !exists || len(files) == 0 {
		return nil, fmt.Errorf("multipart file not exists")
	}

	return files[0], nil
}
