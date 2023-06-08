package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	_ "github.com/lib/pq"
)

// ConnectionURLBuilder func for building url connection.
func ConnectionURLBuilder(str string) (string, error) {
	// define URL to connection
	var url string

	// switch given names.
	switch str {
	case "postgres":
		// url for postgre connection
		if os.Getenv("STAGE_STATUS") == "docker" {
			url = fmt.Sprintf("postgresql://%s:%s@postgres:%s/%s?sslmode=%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_SSL_MODE"))
		} else {
			url = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_SSL_MODE"))
		}
	case "redis":
		// url for redis connection
		if os.Getenv("STAGE_STATUS") == "docker" {
			url = fmt.Sprintf("redis:%s", os.Getenv("REDIS_PORT"))
		} else {
			url = fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT"))
		}
	case "fiber":
		// url for fiber connection
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		)
	default:
		// Return error message.
		return "", fmt.Errorf("connection name '%v' is not supported", str)
	}

	return url, nil
}

func CalculateMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)[:16]
	return hex.EncodeToString(hashInBytes), nil
}
