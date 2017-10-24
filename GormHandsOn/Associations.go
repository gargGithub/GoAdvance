package main

import(
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"log"
	"fmt"

)

func main() {
	db,err:=gorm.Open("mysql","root:password@tcp(127.0.0.1:3306)/gorm_db?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		log.Println("Error Connecting to a database")
	}

	defer db.Close()

	err = db.DB().Ping()

	fmt.Println(err)

    db.DropTableIfExists(&User{},&Profile{})
	db.CreateTable(&User{},&Profile{})


	user1:=User{
		Profile:Profile{
			Name:"UserProfile",
		},
           Name:"Shubham",

	}
	user2:=User{
		Profile:Profile{
			Name:"UserProfile2",
		},
		Name:"Aman",
	}


    var user []User
	db.Create(&user1)
	db.Create(&user2)
	//db.Create(&profile)

	db.Debug().Preload("Profile").Find(&user)
    fmt.Println(user)

}

type User struct {
	gorm.Model
	Profile Profile
	ProfileID int
	Name string
}

type Profile struct {
	gorm.Model
	Name string
}
