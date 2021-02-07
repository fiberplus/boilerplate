package main

import (
	"log"
	"testing"
)

func TestApp(t *testing.T) {

	//	value := os.Getenv("DB_CONNECTION")
	//	fmt.Print(value)

	log.Println(Connection.mysql.driver)

}
