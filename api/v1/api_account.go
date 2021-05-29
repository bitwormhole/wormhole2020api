package v1

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
	Demo(param *AccountVO) (*AccountVO, error)
}
