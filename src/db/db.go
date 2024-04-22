package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Stat struct {
	Id      string
	Queries int
}

func Incr(dborm *gorm.DB, key string) error {
	var line Stat
	result := dborm.FirstOrInit(&line, Stat{Id: key, Queries: 0})
	if result.Error != nil {
		return result.Error
	}

	line.Queries += 1

	result = dborm.Save(line)
	if result.Error != nil {
		return result.Error
	}
	return nil

}

func ConnectToMariaDB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:mypass@tcp(host.docker.internal:3306)/fizzDB?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetMaxQuery(db *gorm.DB) (*Stat, error) {
	var result Stat
	db.Raw("SELECT id, queries FROM stats GROUP BY id ORDER BY queries DESC LIMIT 1").Scan(&result)

	return &result, nil

}
