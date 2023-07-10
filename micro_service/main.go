package main

import (
	"flag"
	"fmt"
	. "github.com/Sadio94/micro_service/internal/intialize"
	"github.com/Sadio94/micro_service/internal/models"
	"os"
)

func main() {
	var cfn string
	flag.StringVar(&cfn, "conf", "./conf/config.yaml", "指定配置文件路径")
	flag.Parse()

	// parse config
	configErr := Init(cfn)
	if configErr != nil {
		fmt.Println(configErr.Error())
		os.Exit(-1)
	}

	// init db
	modelErr := models.InitDb(Conf.MySQLConfig)
	if modelErr != nil {
		Logger.Error(modelErr.Error())
		os.Exit(-1)
	}

	// gateway start

}
