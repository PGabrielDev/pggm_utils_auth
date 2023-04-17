package interfaces

type Token interface {
	SaveToken(token string) error
	UpdateToken(token, minitoken string) error
	GetToken(minitoken string) (string, error)
}
