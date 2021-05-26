package actions

import "github.com/bitwormhole/wormhole2020api/api/v1/vo"

// path=/api/v1/session
type ISession interface {

	// method=GET
	GetCurrentSession(param *vo.Session) (*vo.Session, error)

	// method=DELETE
	Logout(param *vo.Session) (*vo.Session, error)
}
