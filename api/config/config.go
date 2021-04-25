package config

type Settings struct {
	DbUser     string
	DbPassword string
	DbName     string
}

func New() *Settings {
	return &Settings{
		DbUser:     "root",
		DbPassword: "password",
		DbName:     "golang_ecs",
	}
}
