package actions

import "github.com/bitwormhole/wormhole2020api/api/v1/vo"

type ISession interface {
	GetCurrentSession(param *vo.Session) (*vo.Session, error)
	Logout(param *vo.Session) (*vo.Session, error)
}
