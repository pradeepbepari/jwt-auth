package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBAddress     string
	S3Bucket      string
	S3Region      string
	S3Key         string
	S3Secret      string
	Port          string
	Jwt_SecretKey string
	TimeFormat    string
}

var Env = NewConfig()

func NewConfig() Config {
	godotenv.Load()
	return Config{
		DBHost:        getEnvs("PUBLIC_HOST", "	localhost"),
		DBPort:        getEnvs("DB_PORT", "3306	"),
		DBUser:        getEnvs("DB_USER", "root"),
		DBPassword:    getEnvs("DB_PASSWORD", "root"),
		DBAddress:     fmt.Sprintf("%s:%s", getEnvs("DB_HOST", "mysql"), getEnvs("DB_PORT", "3306")),
		DBName:        getEnvs("DB_NAME", "users"),
		Port:          getEnvs("PORT", "8000"),
		Jwt_SecretKey: getEnvs("JWT_SECRETKEY", "645dfdhs887d8c8s7dshn"),
		TimeFormat:    getEnvs("TIME_FORMAT", "2006-01-02 15:04:05 IST"),
		// S3Bucket:   getEnvs("S3_BUCKET", ""),
		// S3Region:   getEnvs("S3_REGION", "us-west-2"),
		// S3Key:      getEnvs("S3_KEY", ""),
		// S3Secret:   getEnvs("S3_SECRET", ""),
	}

}
func getEnvs(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback

}
