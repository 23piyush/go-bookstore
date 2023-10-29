package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "piyush:password@tcp(localhost:3306)/mydb?charset=utf8&parseTime=True&loc=Local")
	// username:password/tablename?<some_inputs_required_by_mysql>
	// username:password@tcp(hostname:port)/database?charset=utf8mb4&parseTime=True&loc=Local
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}