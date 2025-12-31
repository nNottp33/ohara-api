package domain

import (
	"errors"
	"sync"
	"time"

	"gorm.io/gorm"
)

var (
	modelRegistry []interface{}
	registryLock  sync.Mutex
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime;default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

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
