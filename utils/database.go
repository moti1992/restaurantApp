package utils

import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
)

func ConnectToDatabase() {
	var err error
	username := "root"
	password := "root"
	host := "127.0.0.1:3306"
	database := "squadcast"
	connectionString := username + ":" + password + "@tcp(" + host + ")/" + database

	_, err = orm.GetDB("default")
	if err != nil {
		err = orm.RegisterDataBase("default", "mysql", connectionString)
	}

	if err != nil {
		log.Println("Connection error:", err)
		return
	}

	err = MysqlTest("default")
	if err != nil {
		log.Println("Connection error:", err)
		return
	}

	orm.DefaultTimeLoc, _ = time.LoadLocation("Asia/Kolkata")
	orm.Debug = true

	log.Println("DB connected successfully!")
}

func MysqlTest(alias string) error {
	o := orm.NewOrm()
	o.Using(alias)
	_, err := o.Raw("SELECT 1").Exec()
	return err
}
