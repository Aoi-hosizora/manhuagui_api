package dto

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("CategoryDto", "Category response").
			Properties(
				goapidoc.NewProperty("name", "string", true, "category name"),
				goapidoc.NewProperty("title", "string", true, "category title"),
				goapidoc.NewProperty("url", "string", true, "category link"),
			),
	)
}

// 索引项 vo.Category
type CategoryDto struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

func BuildCategoryDto(category *vo.Category) *CategoryDto {
	return &CategoryDto{
		Name:  category.Name,
		Title: category.Title,
		Url:   category.Url,
	}
}
func BuildCategoryDtos(categories []*vo.Category) []*CategoryDto {
	out := make([]*CategoryDto, len(categories))
	for idx, category := range categories {
		out[idx] = BuildCategoryDto(category)
	}
	return out
}
