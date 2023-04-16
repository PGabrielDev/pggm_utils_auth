package interfaces

type Token interface {
	SaveToken(token, minitoken string) error
	UpdateToken(token, minitoken string) error
	GetToken(minitoken string) (string, error)
}
