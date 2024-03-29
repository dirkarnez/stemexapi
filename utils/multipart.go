package utils

import (
	"net/http"

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
