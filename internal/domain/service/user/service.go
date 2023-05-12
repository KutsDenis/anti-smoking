package user

import "github.com/KutsDenis/anti-smoking/internal/domain/repository/user"

type RegistrationService interface {
	RegisterUser(user User) error
}

type registrationService struct {
	repository user.Repository
}

func NewRegistrationService(repository user.Repository) RegistrationService {
	return &registrationService{repository: repository}
}

func (s registrationService) RegisterUser(u User) error {
	if err := s.repository.CreateUser(u); err != nil {
		return err
	}

	return nil
}
