package dto

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/internal/model/object"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("AuthorDto", "Small author response").
			Properties(
				goapidoc.NewProperty("aid", "integer#int64", true, "author id"),
				goapidoc.NewProperty("name", "string", true, "author name"),
				goapidoc.NewProperty("alias", "string", true, "author alias"),
				goapidoc.NewProperty("zone", "string", true, "author zone"),
				goapidoc.NewProperty("cover", "string", true, "author cover"),
				goapidoc.NewProperty("url", "string", true, "author url"),
				goapidoc.NewProperty("manga_count", "integer#int32", true, "author manga count"),
				goapidoc.NewProperty("newest_manga_id", "integer#int32", true, "author newest manga id"),
				goapidoc.NewProperty("newest_manga_title", "string", true, "author newest manga title"),
				goapidoc.NewProperty("newest_manga_url", "string", true, "author newest manga url"),
				goapidoc.NewProperty("newest_date", "string", true, "author update newest date"),
				goapidoc.NewProperty("highest_manga_id", "integer#int32", true, "author highest manga id"),
				goapidoc.NewProperty("highest_manga_title", "string", true, "author highest manga title"),
				goapidoc.NewProperty("highest_manga_url", "string", true, "author highest manga url"),
				goapidoc.NewProperty("highest_score", "number#float", true, "author highest score"),
				goapidoc.NewProperty("average_score", "number#float", true, "author average score"),
				goapidoc.NewProperty("popularity", "number#int32", true, "author popularity"),
				goapidoc.NewProperty("introduction", "string", true, "author introduction"),
				goapidoc.NewProperty("related_authors", "TinyZonedAuthorDto[]", true, "author related authors"),
			),

		goapidoc.NewDefinition("SmallAuthorDto", "Small author response").
			Properties(
				goapidoc.NewProperty("aid", "integer#int64", true, "author id"),
				goapidoc.NewProperty("name", "string", true, "author name"),
				goapidoc.NewProperty("zone", "string", true, "author zone"),
				goapidoc.NewProperty("cover", "string", true, "author cover"),
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

		goapidoc.NewDefinition("TinyZonedAuthorDto", "Tiny author with zone response").
			Properties(
				goapidoc.NewProperty("aid", "integer#int64", true, "author id"),
				goapidoc.NewProperty("name", "string", true, "author name"),
				goapidoc.NewProperty("url", "string", true, "author url"),
				goapidoc.NewProperty("zone", "string", true, "author zone"),
			),
	)
}

type AuthorDto struct {
	Aid               uint64                `json:"aid"`
	Name              string                `json:"name"`
	Alias             string                `json:"alias"`
	Zone              string                `json:"zone"`
	Cover             string                `json:"cover"`
	Url               string                `json:"url"`
	MangaCount        int32                 `json:"manga_count"`
	NewestMangaId     uint64                `json:"newest_manga_id"`
	NewestMangaTitle  string                `json:"newest_manga_title"`
	NewestMangaUrl    string                `json:"newest_manga_url"`
	NewestDate        string                `json:"newest_date"`
	HighestMangaId    uint64                `json:"highest_manga_id"`
	HighestMangaTitle string                `json:"highest_manga_title"`
	HighestMangaUrl   string                `json:"highest_manga_url"`
	HighestScore      float32               `json:"highest_score"`
	AverageScore      float32               `json:"average_score"`
	Popularity        int32                 `json:"popularity"`
	Introduction      string                `json:"introduction"`
	RelatedAuthors    []*TinyZonedAuthorDto `json:"related_authors"`
}

func BuildAuthorDto(author *object.Author) *AuthorDto {
	return &AuthorDto{
		Aid:               author.Aid,
		Name:              author.Name,
		Alias:             author.Alias,
		Zone:              author.Zone,
		Cover:             author.Cover,
		Url:               author.Url,
		MangaCount:        author.MangaCount,
		NewestMangaId:     author.NewestMangaId,
		NewestMangaTitle:  author.NewestMangaTitle,
		NewestMangaUrl:    author.NewestMangaUrl,
		NewestDate:        author.NewestDate,
		HighestMangaId:    author.HighestMangaId,
		HighestMangaTitle: author.HighestMangaTitle,
		HighestMangaUrl:   author.HighestMangaUrl,
		HighestScore:      author.HighestScore,
		AverageScore:      author.AverageScore,
		Popularity:        author.Popularity,
		Introduction:      author.Introduction,
		RelatedAuthors:    BuildTinyZonedAuthorDtos(author.RelatedAuthors),
	}
}

func BuildAuthorDtos(authors []*object.Author) []*AuthorDto {
	out := make([]*AuthorDto, len(authors))
	for idx, author := range authors {
		out[idx] = BuildAuthorDto(author)
	}
	return out
}

type SmallAuthorDto struct {
	Aid        uint64 `json:"aid"`
	Name       string `json:"name"`
	Zone       string `json:"zone"`
	Cover      string `json:"cover"`
	Url        string `json:"url"`
	MangaCount int32  `json:"manga_count"`
	NewestDate string `json:"newest_date"`
}

func BuildSmallAuthorDto(author *object.SmallAuthor) *SmallAuthorDto {
	return &SmallAuthorDto{
		Aid:        author.Aid,
		Name:       author.Name,
		Zone:       author.Zone,
		Cover:      author.Cover,
		Url:        author.Url,
		MangaCount: author.MangaCount,
		NewestDate: author.NewestDate,
	}
}

func BuildSmallAuthorDtos(authors []*object.SmallAuthor) []*SmallAuthorDto {
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

func BuildTinyAuthorDto(author *object.TinyAuthor) *TinyAuthorDto {
	return &TinyAuthorDto{
		Aid:  author.Aid,
		Name: author.Name,
		Url:  author.Url,
	}
}

func BuildTinyAuthorDtos(authors []*object.TinyAuthor) []*TinyAuthorDto {
	out := make([]*TinyAuthorDto, len(authors))
	for idx, author := range authors {
		out[idx] = BuildTinyAuthorDto(author)
	}
	return out
}

type TinyZonedAuthorDto struct {
	Aid  uint64 `json:"aid"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Zone string `json:"zone"`
}

func BuildTinyZonedAuthorDto(author *object.TinyZonedAuthor) *TinyZonedAuthorDto {
	return &TinyZonedAuthorDto{
		Aid:  author.Aid,
		Name: author.Name,
		Url:  author.Url,
		Zone: author.Zone,
	}
}

func BuildTinyZonedAuthorDtos(authors []*object.TinyZonedAuthor) []*TinyZonedAuthorDto {
	out := make([]*TinyZonedAuthorDto, len(authors))
	for idx, author := range authors {
		out[idx] = BuildTinyZonedAuthorDto(author)
	}
	return out
}
