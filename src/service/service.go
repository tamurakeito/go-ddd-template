package service

type TokenGenerator interface {
	GenerateToken(userId string) (string, error)
}
