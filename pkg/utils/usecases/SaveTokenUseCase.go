package usecases

import (
	"github.com/pgabrieldev/pggm_utils_auth/pkg/utils/interfaces"
)

type SaveTokenUseCase struct {
	repository interfaces.Token
}

func NewSaveTokenUseCase(repository interfaces.Token) *SaveTokenUseCase {
	return &SaveTokenUseCase{
		repository: repository,
	}
}
func (u *SaveTokenUseCase) Save(token string) error {
	err := u.repository.SaveToken(token)
	if err != nil {
		return err
	}
	return nil
}
