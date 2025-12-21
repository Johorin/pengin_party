package app

import "os"

func IsLocal() bool {
	return os.Getenv("APP_ENV") == "local"
}

func IsDevelopment() bool {
	return os.Getenv("APP_ENV") == "development"
}

func IsProduction() bool {
	return os.Getenv("APP_ENV") == "production"
}

func IsTest() bool {
	return os.Getenv("APP_ENV") == "test"
}

func IsLocalOrTest() bool {
	return IsLocal() || IsTest()
}

func IsDevelopmentOrProduction() bool {
	return IsDevelopment() || IsProduction()
}