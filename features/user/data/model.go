package data

import (
	"be22/clean-arch/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// ID        string `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Phone     string
	Address   string
	StoreName string
	Role      string
}

func CoreToModel(entity user.Core) User {
	return User{
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		Phone:     entity.Phone,
		Address:   entity.Address,
		StoreName: entity.StoreName,
		Role:      entity.Role,
	}
}

func (u User) ModelToCore() user.Core {
	return user.Core{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Password:  u.Password,
		Phone:     u.Phone,
		Address:   u.Address,
		StoreName: u.StoreName,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ListModelToCore(data []User) []user.Core {
	var results []user.Core
	for _, v := range data {
		results = append(results, v.ModelToCore())
	}
	return results
}
