package handlers

import "github.com/KutsDenis/anti-smoking/internal/domain/service/user"

type RegistrationHandler struct {
	service user.RegistrationService
}

func NewRegistrationHandler(service user.RegistrationService) *RegistrationHandler {
	return &RegistrationHandler{service: service}
}

func (h RegistrationHandler) Handle(u user.User) error {
	if err := h.service.RegisterUser(u); err != nil {
		return err
	}
	return nil
}
