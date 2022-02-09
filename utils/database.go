package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var (
	host = GetEnvValue("DB_HOST")
	port = GetEnvValue("DB_PORT")
	user  = GetEnvValue("DB_USER")
	password = GetEnvValue("DB_PASSWORD")
	database = GetEnvValue("DB_DATABASE")
)

func checkError (err error) bool{
	if err != nil {
		return err!=nil
	}
	return false
}

func GetEnvValue (key string) string {
	err := godotenv.Load();
	if err != nil {
		log.Fatal ("Error loading .env file")
	}
	return os.Getenv(key);
}

func ConnectDb () (*gorm.DB, error){

	//Converts string port from env to port number
	var _, err= strconv.Atoi(port);
	if err != nil {
		log.Fatal("Failed to convert port to string")
	}
	
	//Postres Connection details
	psqlInfo :=  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,database)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	
	if checkError(err) {
		return nil, err
	}
	return db, nil;
}

