package middleware

import (
	"github.com/google/uuid"
	"github.com/nNottp33/ohara-api/internal/config/env"
)

func GenerateRequestIdMiddleware() string {
	appEnv := env.Get[env.AppConfig]("App")
	uniqueString := uuid.New().String()

	return appEnv.PrefixRequestId + "-" + uniqueString
}
