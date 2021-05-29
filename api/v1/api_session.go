package v1

import "github.com/bitwormhole/wormhole2020api/api"

type SessionVO struct {
	BaseVO
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
	GetCurrentSession(context api.Context, param *SessionVO) (*SessionVO, error)

	// method=DELETE
	Logout(context api.Context, param *SessionVO) (*SessionVO, error)
}
