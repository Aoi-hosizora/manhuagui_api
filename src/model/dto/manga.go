package dto

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-backend/src/model/vo"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("MangaPageDto", "Mange page response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("title", "string", true, "manga name"),
				goapidoc.NewProperty("cover", "string", true, "manga cover"),
				goapidoc.NewProperty("url", "string", true, "manga link"),
				goapidoc.NewProperty("publish_year", "string", true, "manga publish year"),
				goapidoc.NewProperty("manga_zone", "string", true, "manga zone"),
				goapidoc.NewProperty("genres", "CategoryDto[]", true, "manga genres"),
				goapidoc.NewProperty("authors", "TinyAuthorDto[]", true, "manga authors"),
				goapidoc.NewProperty("alias", "string", true, "manga alias name"),
				goapidoc.NewProperty("finished", "boolean", true, "manga is finished"),
				goapidoc.NewProperty("newest_chapter", "string", true, "manga last update chapter"),
				goapidoc.NewProperty("newest_date", "string", true, "manga last update date"),
				goapidoc.NewProperty("brief_introduction", "string", true, "manga brief introduction"),
				goapidoc.NewProperty("introduction", "string", true, "manga introduction"),
				goapidoc.NewProperty("manga_rank", "string", true, "manga rank"),
				goapidoc.NewProperty("average_score", "number#float", true, "manga average score"),
				goapidoc.NewProperty("score_count", "integer#int32", true, "manga score count"),
				goapidoc.NewProperty("per_scores", "number#float[]", true, "manga per scores, from 0 to 5"),
				goapidoc.NewProperty("chapter_groups", "MangaChapterGroupDto[]", true, "manga chapter groups"),
			),

		goapidoc.NewDefinition("MangaChapterDto", "Mange chapter response").
			Properties(
				goapidoc.NewProperty("cid", "integer#int64", true, "chapter id"),
				goapidoc.NewProperty("title", "string", true, "chapter name"),
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("manga_title", "string", true, "manga name"),
				goapidoc.NewProperty("url", "string", true, "chapter link"),
				goapidoc.NewProperty("pages", "string[]", true, "chapter pages"),
				goapidoc.NewProperty("page_count", "integer#int32", true, "chapter pages count"),
				goapidoc.NewProperty("next_cid", "integer#int64", true, "next chapter id"),
				goapidoc.NewProperty("prev_cid", "integer#int64", true, "prev chapter id"),
			),

		goapidoc.NewDefinition("SmallMangaPageDto", "Small mange page response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("title", "string", true, "manga name"),
				goapidoc.NewProperty("cover", "string", true, "manga cover"),
				goapidoc.NewProperty("url", "string", true, "manga link"),
				goapidoc.NewProperty("publish_year", "string", true, "manga publish year"),
				goapidoc.NewProperty("manga_zone", "string", true, "manga zone"),
				goapidoc.NewProperty("genres", "CategoryDto[]", true, "manga genres"),
				goapidoc.NewProperty("authors", "TinyAuthorDto[]", true, "manga authors"),
				goapidoc.NewProperty("finished", "boolean", true, "manga is finished"),
				goapidoc.NewProperty("newest_chapter", "string", true, "manga last update chapter"),
				goapidoc.NewProperty("newest_date", "string", true, "manga last update date"),
				goapidoc.NewProperty("brief_introduction", "string", true, "manga brief introduction"),
			),

		goapidoc.NewDefinition("TinyMangaPageDto", "Tiny manga page response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("title", "string", true, "manga name"),
				goapidoc.NewProperty("cover", "string", true, "manga cover"),
				goapidoc.NewProperty("url", "string", true, "manga link"),
				goapidoc.NewProperty("finished", "boolean", true, "manga is finished"),
				goapidoc.NewProperty("newest_chapter", "string", true, "manga last update chapter"),
				goapidoc.NewProperty("newest_date", "string", true, "manga last update date"),
			),

		goapidoc.NewDefinition("TinyMangaChapterDto", "Tiny manga chapter response").
			Properties(
				goapidoc.NewProperty("cid", "integer#int64", true, "chapter id"),
				goapidoc.NewProperty("title", "string", true, "chapter name"),
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("url", "string", true, "chapter link"),
				goapidoc.NewProperty("page_count", "integer#int32", true, "chapter pages count"),
				goapidoc.NewProperty("is_new", "boolean", true, "chapter is uploaded newly"),
			),

		goapidoc.NewDefinition("MangaPageGroupDto", "Mange page group response").
			Properties(
				goapidoc.NewProperty("title", "string", true, "group title"),
				goapidoc.NewProperty("mangas", "TinyMangaPageDto[]", true, "group mangas"),
			),

		goapidoc.NewDefinition("MangaChapterGroupDto", "Mange chapter group response").
			Properties(
				goapidoc.NewProperty("title", "string", true, "group title"),
				goapidoc.NewProperty("chapters", "TinyMangaChapterDto[]", true, "group chapters"),
			),

		goapidoc.NewDefinition("MangaPageGroupListDto", "Manga page group list").
			Properties(
				goapidoc.NewProperty("title", "string", true, "list title"),
				goapidoc.NewProperty("top_group", "MangaPageGroupDto", true, "manga top page group"),
				goapidoc.NewProperty("groups", "MangaPageGroupDto[]", true, "manga page groups"),
				goapidoc.NewProperty("other_groups", "MangaPageGroupDto[]", true, "manga other page groups"),
			),
	)
}

// 漫画页的完整信息 vo.MangaPage
type MangaPageDto struct {
	Mid               uint64                  `json:"mid"`
	Title             string                  `json:"title"`
	Cover             string                  `json:"cover"`
	Url               string                  `json:"url"`
	PublishYear       string                  `json:"publish_year"`
	MangaZone         string                  `json:"manga_zone"`
	Genres            []*CategoryDto          `json:"genres"`
	Authors           []*TinyAuthorDto        `json:"authors"`
	Alias             string                  `json:"alias"`
	Finished          bool                    `json:"finished"`
	NewestChapter     string                  `json:"newest_chapter"`
	NewestDate        string                  `json:"newest_date"`
	BriefIntroduction string                  `json:"brief_introduction"`
	Introduction      string                  `json:"introduction"`
	MangaRank         string                  `json:"manga_rank"`
	AverageScore      float32                 `json:"average_score"`
	ScoreCount        int32                   `json:"score_count"`
	PerScores         [6]float32              `json:"per_scores"`
	ChapterGroups     []*MangaChapterGroupDto `json:"chapter_groups"`
}

func BuildMangaPageDto(manga *vo.MangaPage) *MangaPageDto {
	return &MangaPageDto{
		Mid:               manga.Mid,
		Title:             manga.Title,
		Cover:             manga.Cover,
		Url:               manga.Url,
		PublishYear:       manga.PublishYear,
		MangaZone:         manga.MangaZone,
		Genres:            BuildCategoryDtos(manga.Genres),
		Authors:           BuildTinyAuthorDtos(manga.Authors),
		Alias:             manga.Alias,
		Finished:          manga.Finished,
		NewestChapter:     manga.NewestChapter,
		NewestDate:        manga.NewestDate,
		Introduction:      manga.Introduction,
		BriefIntroduction: manga.BriefIntroduction,
		MangaRank:         manga.MangaRank,
		AverageScore:      manga.AverageScore,
		ScoreCount:        manga.ScoreCount,
		PerScores:         manga.PerScores,
		ChapterGroups:     BuildMangaChapterGroupDtos(manga.ChapterGroups),
	}
}

func BuildMangaPageDtos(mangas []*vo.MangaPage) []*MangaPageDto {
	out := make([]*MangaPageDto, len(mangas))
	for idx, manga := range mangas {
		out[idx] = BuildMangaPageDto(manga)
	}
	return out
}

// 漫画章节的完整信息 vo.MangaChapter
type MangaChapterDto struct {
	Cid        uint64   `json:"cid"`
	Title      string   `json:"title"`
	Mid        uint64   `json:"mid"`
	MangaTitle string   `json:"manga_title"`
	Url        string   `json:"url"`
	Pages      []string `json:"pages"`
	PageCount  int32    `json:"page_count"`
	NextCid    int32    `json:"next_cid"`
	PrevCid    int32    `json:"prev_cid"`
}

func BuildMangaChapterDto(chapter *vo.MangaChapter) *MangaChapterDto {
	return &MangaChapterDto{
		Cid:        chapter.Cid,
		Title:      chapter.Title,
		Mid:        chapter.Mid,
		MangaTitle: chapter.MangaTitle,
		Url:        chapter.Url,
		Pages:      chapter.Pages,
		PageCount:  chapter.PageCount,
		NextCid:    chapter.NextId,
		PrevCid:    chapter.PrevId,
	}
}

func BuildMangaChapterDtos(chapters []*vo.MangaChapter) []*MangaChapterDto {
	out := make([]*MangaChapterDto, len(chapters))
	for idx, chapter := range chapters {
		out[idx] = BuildMangaChapterDto(chapter)
	}
	return out
}

// 漫画页的部分信息 vo.SmallMangaPage
type SmallMangaPageDto struct {
	Mid               uint64           `json:"mid"`
	Title             string           `json:"title"`
	Cover             string           `json:"cover"`
	Url               string           `json:"url"`
	PublishYear       string           `json:"publish_year"`
	MangaZone         string           `json:"manga_zone"`
	Genres            []*CategoryDto   `json:"genres"`
	Authors           []*TinyAuthorDto `json:"authors"`
	Finished          bool             `json:"finished"`
	NewestChapter     string           `json:"newest_chapter"`
	NewestDate        string           `json:"newest_date"`
	BriefIntroduction string           `json:"brief_introduction"`
}

func BuildSmallMangaPageDto(manga *vo.SmallMangaPage) *SmallMangaPageDto {
	return &SmallMangaPageDto{
		Mid:               manga.Mid,
		Title:             manga.Title,
		Cover:             manga.Cover,
		Url:               manga.Url,
		PublishYear:       manga.PublishYear,
		MangaZone:         manga.MangaZone,
		Genres:            BuildCategoryDtos(manga.Genres),
		Authors:           BuildTinyAuthorDtos(manga.Authors),
		Finished:          manga.Finished,
		NewestChapter:     manga.NewestChapter,
		NewestDate:        manga.NewestDate,
		BriefIntroduction: manga.BriefIntroduction,
	}
}

func BuildSmallMangaPageDtos(mangas []*vo.SmallMangaPage) []*SmallMangaPageDto {
	out := make([]*SmallMangaPageDto, len(mangas))
	for idx, manga := range mangas {
		out[idx] = BuildSmallMangaPageDto(manga)
	}
	return out
}

// 漫画页的部分信息 vo.TinyMangaPage
type TinyMangaPageDto struct {
	Mid           uint64 `json:"mid"`
	Title         string `json:"title"`
	Cover         string `json:"cover"`
	Url           string `json:"url"`
	Finished      bool   `json:"finished"`
	NewestChapter string `json:"newest_chapter"`
	NewestDate    string `json:"newest_date"`
}

func BuildTinyMangaPageDto(manga *vo.TinyMangaPage) *TinyMangaPageDto {
	return &TinyMangaPageDto{
		Mid:           manga.Mid,
		Title:         manga.Title,
		Cover:         manga.Cover,
		Url:           manga.Url,
		Finished:      manga.Finished,
		NewestChapter: manga.NewestChapter,
		NewestDate:    manga.NewestDate,
	}
}

func BuildTinyMangaPageDtos(mangas []*vo.TinyMangaPage) []*TinyMangaPageDto {
	out := make([]*TinyMangaPageDto, len(mangas))
	for idx, manga := range mangas {
		out[idx] = BuildTinyMangaPageDto(manga)
	}
	return out
}

// 漫画章节的部分信息 vo.TinyMangaChapter
type TinyMangaChapterDto struct {
	Cid       uint64 `json:"cid"`
	Title     string `json:"title"`
	Mid       uint64 `json:"mid"`
	Url       string `json:"url"`
	PageCount int32  `json:"page_count"`
	IsNew     bool   `json:"is_new"`
}

func BuildTinyMangaChapterDto(chapter *vo.TinyMangaChapter) *TinyMangaChapterDto {
	return &TinyMangaChapterDto{
		Cid:       chapter.Cid,
		Title:     chapter.Title,
		Mid:       chapter.Mid,
		Url:       chapter.Url,
		PageCount: chapter.PageCount,
		IsNew:     chapter.IsNew,
	}
}

func BuildTinyMangaChapterDtos(chapters []*vo.TinyMangaChapter) []*TinyMangaChapterDto {
	out := make([]*TinyMangaChapterDto, len(chapters))
	for idx, link := range chapters {
		out[idx] = BuildTinyMangaChapterDto(link)
	}
	return out
}

// 漫画页分组 vo.MangaPageGroup
type MangaPageGroupDto struct {
	Title  string              `json:"title"`
	Mangas []*TinyMangaPageDto `json:"mangas"`
}

func BuildMangaPageGroupDto(group *vo.MangaPageGroup) *MangaPageGroupDto {
	return &MangaPageGroupDto{
		Title:  group.Title,
		Mangas: BuildTinyMangaPageDtos(group.Mangas),
	}
}

func BuildMangaPageGroupDtos(groups []*vo.MangaPageGroup) []*MangaPageGroupDto {
	out := make([]*MangaPageGroupDto, len(groups))
	for idx, group := range groups {
		out[idx] = BuildMangaPageGroupDto(group)
	}
	return out
}

// 漫画章节分组 vo.MangaChapterGroup
type MangaChapterGroupDto struct {
	Title    string                 `json:"title"`
	Chapters []*TinyMangaChapterDto `json:"chapters"`
}

func BuildMangaChapterGroupDto(group *vo.MangaChapterGroup) *MangaChapterGroupDto {
	return &MangaChapterGroupDto{
		Title:    group.Title,
		Chapters: BuildTinyMangaChapterDtos(group.Chapters),
	}
}

func BuildMangaChapterGroupDtos(groups []*vo.MangaChapterGroup) []*MangaChapterGroupDto {
	out := make([]*MangaChapterGroupDto, len(groups))
	for idx, group := range groups {
		out[idx] = BuildMangaChapterGroupDto(group)
	}
	return out
}

// 主页的漫画列表 vo.MangaPageGroupList
type MangaPageGroupListDto struct {
	Title       string               `json:"title"`
	TopGroup    *MangaPageGroupDto   `json:"top_group"`
	Groups      []*MangaPageGroupDto `json:"groups"`
	OtherGroups []*MangaPageGroupDto `json:"other_groups"`
}

func BuildMangaPageGroupListDto(list *vo.MangaPageGroupList) *MangaPageGroupListDto {
	return &MangaPageGroupListDto{
		Title:       list.Title,
		TopGroup:    BuildMangaPageGroupDto(list.TopGroup),
		Groups:      BuildMangaPageGroupDtos(list.Groups),
		OtherGroups: BuildMangaPageGroupDtos(list.OtherGroups),
	}
}

func BuildMangaPageGroupListDtos(lists []*vo.MangaPageGroupList) []*MangaPageGroupListDto {
	out := make([]*MangaPageGroupListDto, len(lists))
	for idx, list := range lists {
		out[idx] = BuildMangaPageGroupListDto(list)
	}
	return out
}
