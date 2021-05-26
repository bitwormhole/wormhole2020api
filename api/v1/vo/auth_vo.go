package vo

type Auth struct {
	BaseVO

	Account   string
	Secret    string
	Mechanism string
}
