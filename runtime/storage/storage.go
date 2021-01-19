package storage

import "sync"

type StorageController struct{}

var instance *StorageController
var once sync.Once

func GetInstance() *StorageController {
	once.Do(func() {
		instance = &StorageController{}
	})
	return instance
}

func (s *StorageController) ReadTables() {}
