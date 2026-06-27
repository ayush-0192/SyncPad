type Config struct {
	PORT  string
	DBHOST string
	DBPORT string
	DBUSER string
	DBPASS string
	DBNAME string
}

func Load() *Config {
	return &Config {
		DBHOST: os.Getenv("DB_HOST")
		DBPORT: os.Getenv("DB_PORT")
		DBUSER: os.Getenv("DB_USER")
		DBPASS: os.Getenv("DB_PASSWORD")
		DBNAME: os.Getenv("DB_NAME")
	}
}
