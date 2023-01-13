package main

import (
	"gopkg.in/ini.v1"
	"log"
	"webframe/datebases"
	"webframe/routes"
)

func main() {
	conf, err := ini.Load("./config/my.ini")
	if err != nil {
		log.Fatal(err)
	}
	router := routes.InitRouter()
	err = router.Run(conf.Section("server").Key("http_port").String())
	if err != nil {
		log.Fatal(err)
	}
	db := datebases.InitMysql()
	
}
