package service

import (
	"belajar-unit-test/entity"
	"belajar-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.GetCategoryById("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "1",
		Name: "PC",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.GetCategoryById("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Name, result.Name)
}
