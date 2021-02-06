package config

type app struct {
	Port     string
	name     string
	env      string
	debug    string
	timezone string
	locale   string
}

var App = app{
	/*
	   |--------------------------------------------------------------------------
	   | Application Port
	   |--------------------------------------------------------------------------
	   |
	   | This value is the port of your application.
	   |
	*/
	Port: "3001",
	/*
	   |--------------------------------------------------------------------------
	   | Application Name
	   |--------------------------------------------------------------------------
	   |
	   | This value is the name of your application. This value is used when the
	   | framework needs to place the application's name in a notification or
	   | any other location as required by the application or its packages.
	   |
	*/
	name: env("APP_NAME", "Fiber boilerplate"),
	/*
	   |--------------------------------------------------------------------------
	   | Application Environment
	   |--------------------------------------------------------------------------
	   |
	   | This value determines the "environment" your application is currently
	   | running in. This may determine how you prefer to configure various
	   | services the application utilizes. Set this in your ".env" file.
	   |
	*/
	env: env("APP_ENV", "production"),
	/*
	   |--------------------------------------------------------------------------
	   | Application Locale Configuration
	   |--------------------------------------------------------------------------
	   |
	   | The application locale determines the default locale that will be used
	   | by the translation service provider. You are free to set this value
	   | to any of the locales which will be supported by the application.
	   |
	*/
	locale: "en",
}
