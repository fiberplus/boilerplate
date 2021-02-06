package config

func init() {
	envLoader()
}

type mysql struct {
	driver   string
	url      string
	host     string
	port     string
	database string
	username string
	password string
}

/*
   |--------------------------------------------------------------------------
   | Default Database Connection Name
   |--------------------------------------------------------------------------
   |
   | Here you may specify which of the database connections below you wish
   | to use as your default connection for all database work. Of course
   | you may use many connections at once using the Database library.
   |
*/

// Connection ...
Config  = connections{
	mysql: mysql{
		driver:   env("DB_CONNECTION", "mysql"),
		url:      env("DB_URL", ""),
		host:     env("DB_HOST", "127.0.0.1"),
		port:     env("DB_PORT", "3306"),
		database: env("DB_DATABASE", "mysql"),
		username: env("DB_USERNAME", "mysql"),
		password: env("DB_PASSWORD", "mysql"),
	},
}
