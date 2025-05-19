package myUtils

import "os"

func IsProd() bool {
	return os.Getenv("ENV") == "prod"
}
