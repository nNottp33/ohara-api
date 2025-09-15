package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppEnv          string
	AppPort         string
	PrefixRequestId string
}

type DBConfig struct {
	MongoUser string
	MongoPass string
	MongoURI  string
}

type Env struct {
	App AppConfig
	Db  DBConfig
}

func getOrDefault(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return fallback
}

var store = make(map[string]interface{})

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	store["App"] = &AppConfig{
		AppEnv:          getOrDefault("APP_ENV", "local"),
		AppPort:         getOrDefault("APP_PORT", "3033"),
		PrefixRequestId: getOrDefault("PREFIX_REQUEST_ID", "requestId"),
	}

	store["Db"] = &DBConfig{
		MongoUser: getOrDefault("MONGO_INIT_USER", "root"),
		MongoPass: getOrDefault("MONGO_INIT_PASSWORD", "password"),
		MongoURI:  getOrDefault("MONGO_URI", "mongodb://localhost:27017/?authSource=admin"),
	}
}

func Get[T any](key string) *T {
	if val, ok := store[key]; ok {
		if casted, ok := val.(*T); ok {
			return casted
		}
		log.Fatalf("Config key %s exists but type mismatch", key)
	} else {
		log.Fatalf("Config key %s not found", key)
	}
	return nil
}
