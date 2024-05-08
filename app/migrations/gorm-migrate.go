package migrations

import (
	_userData "be22/clean-arch/features/user/data"

	"gorm.io/gorm"
)

func InitDBMigration(db *gorm.DB) {
	db.AutoMigrate(&_userData.User{})
}
