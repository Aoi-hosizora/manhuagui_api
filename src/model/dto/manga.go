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
				goapidoc.NewProperty("alphabet_index", "string", true, "manga alphabet index"),
				goapidoc.NewProperty("category", "string", true, "manga category"),
				goapidoc.NewProperty("author_name", "string", true, "manga author name"),
				goapidoc.NewProperty("alias", "string", true, "manga alias name"),
				goapidoc.NewProperty("finished", "boolean", true, "manga is finished"),
				goapidoc.NewProperty("newest_chapter", "string", true, "manga last update chapter"),
				goapidoc.NewProperty("newest_date", "string", true, "manga last update date"),
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

		goapidoc.NewDefinition("TinyMangaPageDto", "Manga page link response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("title", "string", true, "manga name"),
				goapidoc.NewProperty("cover", "string", true, "manga cover"),
				goapidoc.NewProperty("url", "string", true, "manga link"),
				goapidoc.NewProperty("finished", "boolean", true, "manga is finished"),
				goapidoc.NewProperty("newest_chapter", "string", true, "manga last update chapter"),
				goapidoc.NewProperty("newest_date", "string", true, "manga last update date"),
			),

		goapidoc.NewDefinition("TinyMangaChapterDto", "Manga chapter link response").
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
				goapidoc.NewProperty("mangas", "TinyMangaChapterDto[]", true, "group mangas"),
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
	Mid           uint64                  `json:"mid"`
	Title         string                  `json:"title"`
	Cover         string                  `json:"cover"`
	Url           string                  `json:"url"`
	PublishYear   string                  `json:"publish_year"`
	MangaZone     string                  `json:"manga_zone"`
	AlphabetIndex string                  `json:"alphabet_index"`
	Category      string                  `json:"category"`
	AuthorName    string                  `json:"author_name"`
	Alias         string                  `json:"alias"`
	Finished      bool                    `json:"finished"`
	NewestChapter string                  `json:"newest_chapter"`
	NewestDate    string                  `json:"newest_date"`
	Introduction  string                  `json:"introduction"`
	MangaRank     string                  `json:"manga_rank"`
	AverageScore  float32                 `json:"average_score"`
	ScoreCount    int32                   `json:"score_count"`
	PerScores     [6]float32              `json:"per_scores"`
	ChapterGroups []*MangaChapterGroupDto `json:"chapter_groups"`
}

func BuildMangaPageDto(page *vo.MangaPage) *MangaPageDto {
	return &MangaPageDto{
		Mid:           page.Mid,
		Title:         page.Title,
		Cover:         page.Cover,
		Url:           page.Url,
		PublishYear:   page.PublishYear,
		MangaZone:     page.MangaZone,
		AlphabetIndex: page.AlphabetIndex,
		Category:      page.Category,
		AuthorName:    page.AuthorName,
		Alias:         page.Alias,
		Finished:      page.Finished,
		NewestChapter: page.NewestChapter,
		NewestDate:    page.NewestDate,
		Introduction:  page.Introduction,
		MangaRank:     page.MangaRank,
		AverageScore:  page.AverageScore,
		ScoreCount:    page.ScoreCount,
		PerScores:     page.PerScores,
		ChapterGroups: BuildMangaChapterGroupDtos(page.ChapterGroups),
	}
}

func BuildMangaPageDtos(pages []*vo.MangaPage) []*MangaPageDto {
	out := make([]*MangaPageDto, len(pages))
	for idx, page := range pages {
		out[idx] = BuildMangaPageDto(page)
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

// 漫画页的链接 vo.TinyMangaPage
type TinyMangaPageDto struct {
	Mid           uint64 `json:"mid"`
	Title         string `json:"title"`
	Cover         string `json:"cover"`
	Url           string `json:"url"`
	Finished      bool   `json:"finished"`
	NewestChapter string `json:"newest_chapter"`
	NewestDate    string `json:"newest_date"`
}

func BuildTinyMangaPageDto(page *vo.TinyMangaPage) *TinyMangaPageDto {
	return &TinyMangaPageDto{
		Mid:           page.Mid,
		Title:         page.Title,
		Cover:         page.Cover,
		Url:           page.Url,
		Finished:      page.Finished,
		NewestChapter: page.NewestChapter,
		NewestDate:    page.NewestDate,
	}
}

func BuildTinyMangaPageDtos(pages []*vo.TinyMangaPage) []*TinyMangaPageDto {
	out := make([]*TinyMangaPageDto, len(pages))
	for idx, link := range pages {
		out[idx] = BuildTinyMangaPageDto(link)
	}
	return out
}

// 漫画章节的链接 vo.TinyMangaChapter
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
