package config

type PostgresConf struct {
	Host     string `toml:"host"`
	Port     uint   `toml:"port"`
	Database string `toml:"database"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

type Auth struct {
	Username string
	APIToken string
}

type AppConfig struct {
	Postgres PostgresConf
	Auth     Auth
	URL      string
}

//func FromEnv() (*AppConfig, error) {
//	var err error
//
//	config := AppConfig{}
//	config.Postgres.Host = os.Getenv("POSTGRES_HOST")
//	config.Postgres.Database = os.Getenv("POSTGRES_DATABASE")
//	config.Postgres.Username = os.Getenv("POSTGRES_USERNAME")
//	config.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
//
//	port, err := strconv.ParseUint(os.Getenv("POSTGRES_PORT"), 10, 64)
//	if err != nil {
//		return nil, err
//	}
//	config.Postgres.Port = uint(port)
//
//	return &config, nil
//}
