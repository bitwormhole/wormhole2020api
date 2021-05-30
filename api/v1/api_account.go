package v1

import "github.com/bitwormhole/wormhole2020api/api"

type AccountVO struct {
	BaseVO
	Account AccountBody
}

type AccountBody struct {
	Name     string
	FullName string
	Email    string
	Phone    string
}

type IAccount interface {
	Demo(context api.Context, param *AccountVO) (*AccountVO, error)
}
