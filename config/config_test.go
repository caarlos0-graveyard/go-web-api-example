package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadDefaultConfig(t *testing.T) {
	var assert = assert.New(t)
	var cfg = Get()
	assert.Equal("3000", cfg.Port)
	assert.Equal("postgres://localhost:5432/beers?sslmode=disable", cfg.DatabaseURL)
}
