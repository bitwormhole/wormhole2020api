package v1

type BaseVO struct {
	Error   string
	Message string
	Date    string
	Token   string

	Params map[string]string

	Status    int
	Timestamp int64
}

type BaseDTO struct {
	Type string
	ID   string
}
