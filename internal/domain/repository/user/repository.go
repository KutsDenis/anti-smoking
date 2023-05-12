package user

import "github.com/KutsDenis/anti-smoking/internal/domain/service/user"

type Repository interface {
	CreateUser(user user.User) error
}
