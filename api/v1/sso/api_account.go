package sso

import (
	"github.com/bitwormhole/starter-security/security"
	"github.com/bitwormhole/wormhole2020api/api/v1/base"
)

type AccountVO struct {
	base.VO
	Account AccountBody
}

type AccountBody struct {
	Name     string
	FullName string
	Email    string
	Phone    string
}

type IAccount interface {
	Demo(context security.Context, param *AccountVO) (*AccountVO, error)
}
