package main

import(
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"time"
)
func main() {

	//establishing connection with database
	db,err:= gorm.Open("mysql","root:password@tcp(127.0.0.1:3306)/gorm_db?charset=utf8&parseTime=True&loc=Local")
	if err!=nil{
		log.Println("Error Connecting Database")
	}
	defer db.Close()
	err2:= db.DB().Ping()   //to check if database connected or not

	fmt.Println(err2)  // if connected successfully it will print <nil>

	//db.SingularTable(true) //to remove pluralisation convention
	//mapping structure(Owner) to database

//	db.DropTableIfExists(&Owner{},&Book{},&Author{})  //to drop existing table
	//db.CreateTable(&Owner{},&Book{},&Author{}) //passing empty interface; by default table name = owners


    // to create record

    owner:=Owner{
    	FirstName:"Shubham",
    	LastName: "Garg",
	}
	db.Debug().Model(owner).Where("id=?",4).Update("first_name","kiran","last_name","kanade")
	db.Debug().Model(owner).Where("id=?",4).Update("last_name","kanade")
//	db.Create(&owner)


	/*owner2:=Owner{
		FirstName:"Aman",
		LastName: "Patel",
	}

	db.Create(&owner2)
*/
	//to update data
	//*owner.FirstName = "Shubham"
	//db.Debug().Save(&owner)  //debug() is used to see log level

	// to delete first record and usage of deleted at field in gorm.model


		/*var o Owner

		db.First(&o,"id=?",2) // fetches the record according to where clause
		fmt.Println(o)*/


	//db.Delete(&owner)

	/*o = Owner{}
	db.Debug().First(&o)  // fetches first record and stores in o
	fmt.Println(o)*/
}


//Owner and Book struct implementing has-many relationship
type Owner struct {
	gorm.Model   //I way, more convention way
	//ID uint   // II way ,gorm will treat this as primary key
	FirstName string
	LastName string
	Books []Book
}

//to change table name
/*
func (o *Owner) TableName() string{
	return "users"
}
*/

type Book struct {
	gorm.Model
	Name string
	PublishDate time.Time
	OwnerID uint `sql:"index"`     //foreign key association with owner table, `sql:"index"` is used to create index
    Authors []Author  `gorm:"many2many:books_authors"`  //many 2 many relationship will create join table
}

type Author struct {
	gorm.Model
	FirstName string
	LastName string
}




