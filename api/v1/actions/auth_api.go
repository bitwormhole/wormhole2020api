package actions

import "github.com/bitwormhole/wormhole2020api/api/v1/vo"

// path=/api/v1/auth
type IAuth interface {

	// method=POST
	Login(param *vo.Auth) (*vo.Auth, error)
}
