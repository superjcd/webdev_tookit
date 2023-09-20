package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDb(t *testing.T) {
	options := NewDbOptions("root", "root", "192.168.0.77", "3306", "outperform")

	_, err := options.NewDB()

	assert.Nil(t, err)
}
