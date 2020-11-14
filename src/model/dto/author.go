package dto

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("SmallAuthorDto", "Small author response").
			Properties(
				goapidoc.NewProperty("aid", "integer#int64", true, "author id"),
				goapidoc.NewProperty("name", "string", true, "author name"),
				goapidoc.NewProperty("zone", "string", true, "author zone"),
				goapidoc.NewProperty("url", "string", true, "author url"),
				goapidoc.NewProperty("manga_count", "integer#int32", true, "author manga count"),
				goapidoc.NewProperty("newest_date", "string", true, "author update newest date"),
			),

		goapidoc.NewDefinition("TinyAuthorDto", "Tiny author response").
			Properties(
				goapidoc.NewProperty("aid", "integer#int64", true, "author id"),
				goapidoc.NewProperty("name", "string", true, "author name"),
				goapidoc.NewProperty("url", "string", true, "author url"),
			),
	)
}

type SmallAuthorDto struct {
	Aid        uint64 `json:"aid"`
	Name       string `json:"name"`
	Zone       string `json:"zone"`
	Url        string `json:"url"`
	MangaCount int32  `json:"manga_count"`
	NewestDate string `json:"newest_date"`
}

func BuildSmallAuthorDto(author *vo.SmallAuthor) *SmallAuthorDto {
	return &SmallAuthorDto{
		Aid:        author.Aid,
		Name:       author.Name,
		Zone:       author.Zone,
		Url:        author.Url,
		MangaCount: author.MangaCount,
		NewestDate: author.NewestDate,
	}
}

func BuildSmallAuthorDtos(authors []*vo.SmallAuthor) []*SmallAuthorDto {
	out := make([]*SmallAuthorDto, len(authors))
	for idx, author := range authors {
		out[idx] = BuildSmallAuthorDto(author)
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
