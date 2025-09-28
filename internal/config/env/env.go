package env

import (
	"log"
	"os"
	"strconv"
	"unsafe"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppEnv          string
	AppPort         string
	PrefixRequestId string
}

type DBConfig struct {
	Host              string
	Port              int
	User              string
	Password          string
	DBName            string
	MaxConnection     int
	MaxIdleConnection int
}

type Env struct {
	App AppConfig
	Db  DBConfig
}

var store = make(map[string]interface{})

func bitSizeOf[T ~string |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~float32 | ~float64]() int {

	var zero T
	switch any(zero).(type) {
	case uint8, int8:
		return 8
	case uint16, int16:
		return 16
	case uint32, int32, float32:
		return 32
	case uint64, int64, float64:
		return 64
	case uint:
		return int(unsafe.Sizeof(uint(0)) * 8)
	case int:
		return int(unsafe.Sizeof(int(0)) * 8)
	default:
		return 0
	}
}

func getOrDefault[T ~string |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~float32 | ~float64](key string, fallback string, opts ...int) T {

	valueStr, exists := os.LookupEnv(key)
	if !exists {
		valueStr = fallback
	}

	var zero T
	switch any(zero).(type) {
	case string:
		return *(*T)(unsafe.Pointer(&valueStr))

	case uint8, uint16, uint32, uint64, uint:
		bitSize := bitSizeOf[T]()
		if len(opts) > 0 {
			bitSize = opts[0]
		}
		v, err := strconv.ParseUint(valueStr, 10, bitSize)
		if err != nil {
			v, _ = strconv.ParseUint(fallback, 10, bitSize)
		}
		return T(v)

	case int8, int16, int32, int64, int:
		bitSize := bitSizeOf[T]()
		if len(opts) > 0 {
			bitSize = opts[0]
		}
		v, err := strconv.ParseInt(valueStr, 10, bitSize)
		if err != nil {
			v, _ = strconv.ParseInt(fallback, 10, bitSize)
		}
		return T(v)

	case float32, float64:
		bitSize := bitSizeOf[T]()
		if len(opts) > 0 {
			bitSize = opts[0]
		}
		v, err := strconv.ParseFloat(valueStr, bitSize)
		if err != nil {
			v, _ = strconv.ParseFloat(fallback, bitSize)
		}
		return *(*T)(unsafe.Pointer(&v))

	default:
		panic("unsupported type")
	}
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	store["App"] = &AppConfig{
		AppEnv:          getOrDefault[string]("APP_ENV", "local"),
		AppPort:         getOrDefault[string]("APP_PORT", "3033"),
		PrefixRequestId: getOrDefault[string]("PREFIX_REQUEST_ID", "requestId"),
	}

	store["Db"] = &DBConfig{
		Host:              getOrDefault[string]("DB_HOST", "localhost"),
		Port:              getOrDefault[int]("DB_PORT", "5432"),
		User:              getOrDefault[string]("DB_USER", "postgres"),
		Password:          getOrDefault[string]("DB_PASSWORD", "postgres"),
		DBName:            getOrDefault[string]("DB_NAME", "postgres"),
		MaxConnection:     getOrDefault[int]("DB_MAX_CONNECTION", "10"),
		MaxIdleConnection: getOrDefault[int]("DB_MAX_IDLE_CONNECTION", "10"),
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
