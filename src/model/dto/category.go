package dto

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/vo"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("CategoryDto", "Category response").
			Properties(
				goapidoc.NewProperty("name", "string", true, "category name"),
				goapidoc.NewProperty("title", "string", true, "category title"),
				goapidoc.NewProperty("url", "string", true, "category link"),
				goapidoc.NewProperty("cover", "string", true, "category cover"),
			),

		goapidoc.NewDefinition("CategoryListDto", "Category list response").
			Properties(
				goapidoc.NewProperty("genres", "CategoryDto[]", true, "genre categories"),
				goapidoc.NewProperty("zones", "CategoryDto[]", true, "zone categories"),
				goapidoc.NewProperty("ages", "CategoryDto[]", true, "age categories"),
			),
	)
}

// 索引项 vo.Category
type CategoryDto struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Cover string `json:"cover"`
}

func BuildCategoryDto(category *vo.Category) *CategoryDto {
	return &CategoryDto{
		Name:  category.Name,
		Title: category.Title,
		Url:   category.Url,
		Cover: category.Cover,
	}
}
func BuildCategoryDtos(categories []*vo.Category) []*CategoryDto {
	out := make([]*CategoryDto, len(categories))
	for idx, category := range categories {
		out[idx] = BuildCategoryDto(category)
	}
	return out
}

// 索引列表 vo.CategoryList
type CategoryListDto struct {
	Genres []*CategoryDto `json:"genres"`
	Zones  []*CategoryDto `json:"zones"`
	Ages   []*CategoryDto `json:"ages"`
}

func BuildCategoryListDto(lists *vo.CategoryList) *CategoryListDto {
	return &CategoryListDto{
		Genres: BuildCategoryDtos(lists.Genres),
		Zones:  BuildCategoryDtos(lists.Zones),
		Ages:   BuildCategoryDtos(lists.Ages),
	}
}
func BuildCategoryListDtos(lists []*vo.CategoryList) []*CategoryListDto {
	out := make([]*CategoryListDto, len(lists))
	for idx, list := range lists {
		out[idx] = BuildCategoryListDto(list)
	}
	return out
}
