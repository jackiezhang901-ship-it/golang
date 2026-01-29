package config

type Config struct {
	DB DBCOnfig
}

type DBCOnfig struct {
	Url      string
	Driver   string
	UserName string
	Password string
}



