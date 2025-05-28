package category

import (
	"cable-parser/cable-parse-from-file/dataproviders/categoryRepository"
	"context"
)

type ICategory interface {
	// GetCategory получение категории
	GetCategory(ctx context.Context, slug string) (GetCategoryDto, error)
	// CreateCategory создание категории
	CreateCategory(ctx context.Context, category AddCategoryDto) error
}

type category struct {
	categoryRepository categoryRepository.ICategory
}

func NewCategory(categoryRepository categoryRepository.ICategory) ICategory {
	return &category{categoryRepository: categoryRepository}
}

// GetCategory получение категории
func (c *category) GetCategory(ctx context.Context, slug string) (GetCategoryDto, error) {

	ct, err := c.categoryRepository.GetCategory(ctx, slug)
	if err != nil {
		return GetCategoryDto{}, err
	}

	return GetCategoryDto{
		CategoryId: ct.CategoryId,
		Slug:       ct.Slug,
	}, nil
}

// CreateCategory создание категории
func (c *category) CreateCategory(ctx context.Context, category AddCategoryDto) error {

	err := c.categoryRepository.CreateCategory(ctx, categoryRepository.AddCategoryDto{
		Slug:         category.Slug,
		PluralName:   category.PluralName,
		SingularName: category.SingularName,
		Description:  category.Description,
		Img:          category.Img,
	})

	if err != nil {
		return err
	}

	return nil
}
