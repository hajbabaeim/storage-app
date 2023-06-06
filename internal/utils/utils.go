package utils

import (
	"encoding/csv"
	"fmt"
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
		url = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"), os.Getenv("POSTGRES_SSL_MODE"))
	case "redis":
		// url for redis connection
		url = fmt.Sprintf(
			"%s:%s",
			os.Getenv("REDIS_HOST"),
			os.Getenv("REDIS_PORT"),
		)
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

func ReadCSVFromFile(filePath string) (*csv.Reader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(file)
	return r, nil
}

// func ReadCSV(client *ent.Client, filePath string) error {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	lines, err := csv.NewReader(file).ReadAll()
// 	if err != nil {
// 		return err
// 	}

// 	for _, line := range lines {
// 		price, err := strconv.ParseFloat(line[1], 64)
// 		if err != nil {
// 			return err
// 		}

// 		expirationDate, err := time.Parse(time.RFC3339, line[2])
// 		if err != nil {
// 			return err
// 		}

// 		_, err = client.Promotion.
// 			Create().
// 			SetUUID(line[0]).
// 			SetPrice(price).
// 			SetExpirationDate(expirationDate).
// 			Save(context.Background())

// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
