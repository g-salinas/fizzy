package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Stat struct {
	Id      string `gorm:"primaryKey"`
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
	bin, err := os.ReadFile("/run/secrets/db-password")
	if err != nil {
		return nil, err
	}

	sqlDB := mysql.Open(fmt.Sprintf("root:%s@tcp(db:3306)/buzzDB", string(bin)))

	db, err := gorm.Open(sqlDB, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Stat{})
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
