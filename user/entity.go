package user

import "time"

type User struct{
	ID string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	PasswordHash string `json:"password_hash"`
	AvatarFileName string `json:"avatar_file_name"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}