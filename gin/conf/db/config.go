package config

import "time"

type Config struct {
	DB DBCOnfig
}

type DBCOnfig struct {
	Url          string
	Driver       string
	UserName     string
	Password     string
	PoolSize     int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
