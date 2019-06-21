package services

import (
	"errors"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostService(t *testing.T) {
	dao := newMockPostDAO()
	s := NewPostService(dao)
	assert.Equal(t, dao, s.dao)
}

func newMockPostDAO() postDAO {
	return &mockPostDAO{
		records: []models.Post{
			{Model: gorm.Model{ID: 1}, Title: "Test Title", Text: "Test Text."},
			{Model: gorm.Model{ID: 2}, Title: "Test Title 2", Text: "Test Text 2."},
		},
	}
}

func (m *mockPostDAO) Get(id uint) (*models.Post, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

type mockPostDAO struct {
	records []models.Post
}