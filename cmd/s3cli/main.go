package main

import (
	"fmt"

	"github.com/dirkarnez/stemexapi/utils"
)

func main() {
	s3 := utils.NewStemexS3Client()
	objs, _ := s3.ListObjects()
	for _, obj := range objs {
		err := s3.DeleteFile(*obj.Key)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	objs, _ = s3.ListObjects()
	fmt.Println("Deleted all!", len(objs))
}
