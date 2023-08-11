package model

type Config struct {
	Database Database
	Server   Server
}

type Database struct {
	Host            string
	Port            string
	DBName          string
	User            string
	Password        string
	SSLMode         string
	TimeZone        string
	MaxCons         int
	MinCons         int
	MaxConnIdleTime int
}

type Server struct {
	Host string
	Port string
}
