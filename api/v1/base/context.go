package base

type Context interface {
	GetAttribute(name string) interface{}
	SetAttribute(name string, value interface{})
}
