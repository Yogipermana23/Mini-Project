package main

import "os"

func Getenv(key string, fallback string) string{
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
func initConfig(){
	os.Setenv("pqs_url", Getenv("pqs_url", "http://172.18.102.193:8765"))
}