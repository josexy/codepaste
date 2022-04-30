package global

import "github.com/josexy/gocodepaste/model"

func migration() {
	_ = DB.AutoMigrate(
		&model.User{},
		&model.Paste{},
	)
}
