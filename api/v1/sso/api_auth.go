package sso

import "github.com/bitwormhole/wormhole2020api/api/v1/base"

type AuthVO struct {
	base.VO
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
	Login(context base.Context, param *AuthVO) (*AuthVO, error)
}
