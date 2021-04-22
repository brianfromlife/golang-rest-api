package data

type IUserProvider interface {
	GetUsers() string
}

func (d DataProvider) GetUsers() string {
	return "string"
}
