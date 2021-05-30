package api

type Context interface {
	Get(name string) (interface{}, error)
	Set(name string, obj interface{})
}
