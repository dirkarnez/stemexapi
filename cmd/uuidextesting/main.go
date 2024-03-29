package main

import (
	"fmt"

	"github.com/dirkarnez/stemexapi/model"
	"github.com/google/uuid"
)

func main() {
	uuidex1 := model.NewUUIDEx()
	uuidex2 := model.UUIDEx(uuid.Nil)
	isEmpty := uuidex2.IsEmpty()
	fmt.Println(uuidex1, isEmpty)
}
