package dto

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("AuthorDto", "Author response").
			Properties(
				goapidoc.NewProperty("aid", "integer#int64", true, "author id"),
				goapidoc.NewProperty("name", "string", true, "author name"),
				goapidoc.NewProperty("url", "string", true, "author url",
				),
			),

		goapidoc.NewDefinition("TinyAuthorDto", "Tiny author response").
			Properties(
				goapidoc.NewProperty("aid", "integer#int64", true, "author id"),
				goapidoc.NewProperty("name", "string", true, "author name"),
				goapidoc.NewProperty("url", "string", true, "author url",
				),
			),
	)
}

type AuthorDto struct {
	Aid  uint64 `json:"aid"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func BuildAuthorDto(author *vo.Author) *AuthorDto {
	return &AuthorDto{
		Aid:  author.Aid,
		Name: author.Name,
		Url:  author.Url,
	}
}

func BuildAuthorDtos(authors []*vo.Author) []*AuthorDto {
	out := make([]*AuthorDto, len(authors))
	for idx, author := range authors {
		out[idx] = BuildAuthorDto(author)
	}
	return out
}

type TinyAuthorDto struct {
	Aid  uint64 `json:"aid"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func BuildTinyAuthorDto(author *vo.TinyAuthor) *TinyAuthorDto {
	return &TinyAuthorDto{
		Aid:  author.Aid,
		Name: author.Name,
		Url:  author.Url,
	}
}

func BuildTinyAuthorDtos(authors []*vo.TinyAuthor) []*TinyAuthorDto {
	out := make([]*TinyAuthorDto, len(authors))
	for idx, author := range authors {
		out[idx] = BuildTinyAuthorDto(author)
	}
	return out
}
