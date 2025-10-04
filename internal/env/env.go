package env

import (
	"log"
	"os"
	"strconv"
)

func GetEnvString(Key, defaultValue string) string {
	if value ,ok := os.LookupEnv(Key); ok{
		return value
	}

	return defaultValue
}

func GetEnvInt(Key string, defaultValue int) int {
	if value ,ok := os.LookupEnv(Key); ok{
		answerVal , err := strconv.Atoi(value)

		if(err != nil){
			log.Fatal(err)
		}

		return answerVal
	}

	return defaultValue
}