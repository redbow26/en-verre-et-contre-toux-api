package database

import "gorm.io/gorm"

var DB *gorm.DB

func Init(dialect gorm.Dialector, opts ...*gorm.Config) error {
	var err error
	DB, err = gorm.Open(dialect, opts[0])
	if err != nil {
		return err
	}
	return nil
}
