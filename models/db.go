package models

import (
	"PaperSystem/utils"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

var Db *xorm.Engine

func init(){
	fmt.Println("init")
	utils.InitConf(`C:\Users\YngziJ\go\src\PaperSystem\conf\config.toml`)
	switch utils.DB_Driver {
	case "mysql":
	default:
		log.Fatalln("Wrong db driver name :",utils.DB_Driver)
	}
	fmt.Println("Done")


	db,err:=xorm.NewEngine(utils.DB_Driver,utils.DB_Connect)
	fmt.Println(db)
	if err != nil {
		log.Fatalln("New engine error :",err.Error())
	}
	err=db.Ping()
	if err != nil {
		log.Fatalln("Cannot connect to database ",err.Error())
	}
	err=db.Sync2(new (User),new(Session),new(Test),new(Paper))
	if err != nil {
		log.Fatalln("failed to sync database struct :",err.Error())
	}
	Db=db
}



func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}


func PaperByUUID(uuid string) (paper Paper, err error) {
	fmt.Println(uuid)
	_, err = Db.Where("uuid=?", uuid).Get(&paper)
	return
}


func UserByUUID(uuid string) (user User, err error) {
	_, err = Db.Where("uuid=?", uuid).Get(&user)
	return
}



func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}


