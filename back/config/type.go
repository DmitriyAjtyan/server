package config

// ConfigurationDataStruct is a data type for server configuration data
type ConfigurationDataStruct struct {
	Port      int
	Host      string
	KeyJWT    string
	JWTSecret string
	Postgres  Postgres
	Redis     Redis
}

// Postgres is a data type for postgresql configuration data
type Postgres struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	MaxConnect int
}

// Redis is a data type for redis configuration data
type Redis struct {
	Host        string
	Port        int
	Password    string
	DBName      string
	MaxConnect  int
	Prefix      string
	MaxAge      int
	IdleTimeout int
}
