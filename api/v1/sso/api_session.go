package sso

import "github.com/bitwormhole/wormhole2020api/api/v1/base"

type SessionVO struct {
	base.VO
	Session SessionBody
}

type SessionBody struct {
	Account      string
	Nickname     string
	CreationDate int64
	Authorized   bool
}

// path=/api/v1/session
type ISession interface {

	// method=GET
	GetCurrentSession(context base.Context, param *SessionVO) (*SessionVO, error)

	// method=DELETE
	Logout(context base.Context, param *SessionVO) (*SessionVO, error)
}
