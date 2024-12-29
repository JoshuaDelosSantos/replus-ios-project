package auth

type TokenValidator interface {
	ValidateToken(token string) (*Claims, error)
}