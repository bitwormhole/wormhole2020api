package actions

import "github.com/bitwormhole/wormhole2020api/api/v1/vo"

type IAuth interface {
	Login(param *vo.Auth) (*vo.Auth, error)
}
