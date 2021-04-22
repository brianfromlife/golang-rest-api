package data

type IAccountProvider interface {
	CreateAccount() string
}

func (d DataProvider) CreateAccount() string {
	return "shd"
}
