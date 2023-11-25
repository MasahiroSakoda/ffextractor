// Package cache provides basic cache method
package cache

import (
	"encoding/json"
	"io"
	"time"

	"github.com/mitchellh/mapstructure"
)

// Cache is container struct for cache
type Cache struct {
	Expiration time.Time   `json:"expiration"`
	Data       interface{} `json:"data"`
}

// New returns new cache container
func New(ttl time.Duration, data interface{}) *Cache {
	return &Cache{
		Expiration: time.Now().Add(ttl),
		Data:       data,
	}
}

// Expired returns whether cache is expired
func (c *Cache) Expired() bool {
	return time.Now().After(c.Expiration)
}

// Bind provides data binding method
func (c *Cache) Bind(dst interface{}) error {
	if err := mapstructure.Decode(c.Data, dst); err != nil {
		return err
	}
	return nil
}

// Write provides saving function to json
func (c *Cache) Write(w io.Writer) error {
	buf, err := json.Marshal(c)
	if err != nil {
		return err
	}

	if _, err := w.Write(buf); err != nil {
		return err
	}

	return nil
}
