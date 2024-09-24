package env

import (
	"os"
	"strconv"
)

func GetString(key string) string {
	value := os.Getenv(key)

	return value
}

func GetInt(key string) int {
	value := os.Getenv(key)

	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}

	return intValue
}

func GetBool(key string) bool {
	value := os.Getenv(key)

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		panic(err)
	}

	return boolValue
}
