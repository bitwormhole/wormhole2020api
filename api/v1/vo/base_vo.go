package vo

type BaseVO struct {
	Params map[string]string

	Status    int
	Error     string
	Message   string
	Date      string
	Timestamp int64
}
