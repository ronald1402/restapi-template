package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewRepository(t *testing.T) {
	repo := NewRepository()
	assert.NotNil(t, repo, "NewLoanRepository() should return a non-nil repository")
}
