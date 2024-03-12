package utils

import (
	"net/http"

	"github.com/gin-gonic/gin/binding"
)

func FormMultipartParse(req *http.Request, stuctPointer any) error {
	return binding.FormMultipart.Bind(req, stuctPointer)
}
