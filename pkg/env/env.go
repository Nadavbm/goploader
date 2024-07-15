package env

import "os"

var ServerUrl string = getEnvVar("SERVER_URL", "http://localhost:8080/upload")
var FilesDir string = getEnvVar("GOPLOADER_FILES_DIR", "../files/")

func getEnvVar(key, defaultValue string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultValue
	}
	return env
}
