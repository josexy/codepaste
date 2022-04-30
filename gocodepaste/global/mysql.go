package global

import (
	"github.com/josexy/gocodepaste/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLConfig struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	DSN      string
}

func InitMySQL(config *MySQLConfig) {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		Logger.Errorln(err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		Logger.Errorln(err)
		panic(err)
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	DB = db.Set("gorm:table_options", "ENGINE=Innodb DEFAULT CHARSET=utf8mb4")
	DB = db

	migration()

	pubUser := model.User{
		Nickname: "PUBLIC",
		Username: "$PUBLIC$",
	}

	err = DB.Transaction(func(tx *gorm.DB) error {
		var count int64
		if e := tx.Model(model.User{}).
			Where("`username` = ? and `nickname` = ?", "$PUBLIC$", "PUBLIC").
			Count(&count).Error; e != nil {
			return e
		}
		if count <= 0 {
			return tx.Create(&pubUser).Error
		}
		return nil
	})
	if err != nil {
		Logger.Errorln(err)
		panic(err)
	}
}
