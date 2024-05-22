package main

import (
	"fmt"
	"log"

	"github.com/dirkarnez/stemexapi/db"
	"github.com/dirkarnez/stemexapi/query"
	C:\Users\stude\Downloads\stemexapi\cmd\redumpparents\\redumpparents.go
	"github.com/dirkarnez/stemexapi/cmd/redumpparents/redumpparents"
)

func main() {
	dbInstance, dbInstanceErr := db.InitConntection()
	if dbInstanceErr != nil {
		log.Fatal(dbInstanceErr.Error())
	}
	if dbInstance != nil {
		fmt.Println("Connected!")
	}

	dbInstance = dbInstance.Debug()
	var q = query.Use(dbInstance)

	if err != nil {
		fmt.Println("err!")
	}

}
