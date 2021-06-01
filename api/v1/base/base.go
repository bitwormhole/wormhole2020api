package base

type VO struct {
	Error   string
	Message string
	Date    string
	Token   string

	Params map[string]string

	Status    int
	Timestamp int64
}

type DTO struct {
	Type string
	ID   string
}
