package main

import (
	"fmt"
	"log"

	"github.com/dirkarnez/stemexapi/cmd/redumpparents/redumpparents"
	"github.com/dirkarnez/stemexapi/db"
	"github.com/dirkarnez/stemexapi/query"
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

	err := redumpparents.RedumpParents(q)
	if err != nil {
		fmt.Println("err!")
	}
}
