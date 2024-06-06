package functions

import "os"

func GetEnvValueWithDefault(key, defaultValue string) string{
	value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
