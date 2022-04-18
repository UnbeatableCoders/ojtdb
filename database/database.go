package database

import (
	"fmt"
	"task_management/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var DSN = "host=ec2-3-229-161-70.compute-1.amazonaws.com user=diemaezjbbqrna password=19642e056f5798055e67ad72aa5405ec100c63f7fcd799b11f207491f6d758dc dbname=d19m46uuv65c9l port=5432 "

func Migration() {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to the databse")
	DB.AutoMigrate(&model.User{}, &model.Preference{}, &model.Team{}, &model.Workspace{}, &model.Task{})
}
