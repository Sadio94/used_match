package main

import (
	"fmt"
	"github.com/Sadio94/go_micro/internal/models"
	"os"
)

func main() {
	// parse config

	// init db
	err := models.InitDb()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// gateway start

	// proto micro

}
