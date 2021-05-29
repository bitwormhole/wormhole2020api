package v1

import "github.com/bitwormhole/wormhole2020api/api"

type AuthVO struct {
	BaseVO
	Auth AuthBody
}

type AuthBody struct {
	Account   string
	Secret    string
	Mechanism string
}

// path=/api/v1/auth
type IAuth interface {

	// method=POST
	Login(context api.Context, param *AuthVO) (*AuthVO, error)
}
