package handler

import "be22/clean-arch/features/user"

type UserRequest struct {
	Name      string `json:"name" form:"name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
	Phone     string `json:"phone" form:"phone"`
	Address   string `json:"address" form:"address"`
	StoreName string `json:"store_name" form:"store_name"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(input UserRequest) user.Core {
	return user.Core{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Phone:     input.Phone,
		Address:   input.Address,
		StoreName: input.StoreName,
	}
}
