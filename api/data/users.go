package data

type IUserProvider interface {
	CreateAcount() string
	FindUserById() string
}

func (d DataProvider) CreateAcount() string {
	return "string"
}

func (d DataProvider) FindUserById() string {
	return "string"
}
