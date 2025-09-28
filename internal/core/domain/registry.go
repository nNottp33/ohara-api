package domain

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

var (
	modelRegistry []interface{}
	registryLock  sync.Mutex
)

func RegisterModel(model interface{}) {
	registryLock.Lock()
	defer registryLock.Unlock()
	modelRegistry = append(modelRegistry, model)
}

func GetModelRegistry() []interface{} {
	registryLock.Lock()
	defer registryLock.Unlock()
	return modelRegistry
}

func AutomigrateAll(db *gorm.DB) error {
	models := GetModelRegistry()
	if len(models) == 0 {
		return errors.New("no model found")
	}

	if err := db.AutoMigrate(models...); err != nil {
		return err
	}

	return nil
}
