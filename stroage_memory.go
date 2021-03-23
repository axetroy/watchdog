package watchdog

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var Store Storage

func init() {
	s := NewStorageMemory()

	createStorage(s)
}

func createStorage(s Storage) {
	Store = s
}

type Storage interface {
	SetItem(key string, value []ServiceStatus)
	GetItem(key string) *[]ServiceStatus
	RemoveItem(key string)
	Clear()
}

type StorageMemory struct {
	c *cache.Cache
}

func NewStorageMemory() *StorageMemory {
	return &StorageMemory{
		c: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (s *StorageMemory) SetItem(key string, value []ServiceStatus) {
	maxStoreItemsNumber := 100

	list := s.GetItem(key)
	if list == nil {
		s.c.Set(key, value, time.Hour*24)
	} else {
		result := *list

		result = append(result, value...)

		if len(result) > maxStoreItemsNumber {
			diff := len(result) - maxStoreItemsNumber

			result = result[diff:]
		}

		s.c.Set(key, result, time.Hour*24)
	}
}

func (s *StorageMemory) GetItem(key string) *[]ServiceStatus {
	value, found := s.c.Get(key)
	if found {
		v := value.([]ServiceStatus)
		return &v
	} else {
		return nil
	}
}

func (s *StorageMemory) RemoveItem(key string) {
	s.c.Delete(key)
}

func (s *StorageMemory) Clear() {
	s.c.Flush()
}
