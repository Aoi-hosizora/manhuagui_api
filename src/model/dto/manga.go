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
				goapidoc.NewProperty("mname", "string", true, "manga name"),
				goapidoc.NewProperty("cover", "string", true, "manga cover"),
				goapidoc.NewProperty("url", "string", true, "manga link"),
				goapidoc.NewProperty("publish_year", "string", true, "manga publish year"),
				goapidoc.NewProperty("alphabet_index", "string", true, "manga alphabet index"),
				goapidoc.NewProperty("type", "string", true, "manga type"),
				goapidoc.NewProperty("author_name", "string", true, "manga author name"),
				goapidoc.NewProperty("alias", "string", true, "manga alias name"),
				goapidoc.NewProperty("finished", "boolean", true, "manga is finished"),
				goapidoc.NewProperty("newest_chapter", "string", true, "manga last update chapter"),
				goapidoc.NewProperty("newest_date", "string", true, "manga last update date"),
				goapidoc.NewProperty("introduction", "string", true, "manga introduction"),
				goapidoc.NewProperty("rank", "string", true, "manga rank"),
				goapidoc.NewProperty("average_score", "number#float", true, "manga average score"),
				goapidoc.NewProperty("score_count", "integer#int32", true, "manga score count"),
				goapidoc.NewProperty("per_scores", "number#float[]", true, "manga per scores, from 0 to 5"),
				goapidoc.NewProperty("groups", "MangaChapterGroupDto[]", true, "manga chapter groups"),
			),

		goapidoc.NewDefinition("MangaChapterDto", "Mange chapter response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("mname", "string", true, "manga name"),
				goapidoc.NewProperty("cid", "integer#int64", true, "manga chapter id"),
				goapidoc.NewProperty("cname", "string", true, "manga chapter name"),
				goapidoc.NewProperty("url", "string", true, "manga chapter link"),
				goapidoc.NewProperty("pages", "string[]", true, "manga chapter pages"),
				goapidoc.NewProperty("next_cid", "integer#int64", true, "manga next chapter id"),
				goapidoc.NewProperty("prev_cid", "integer#int64", true, "manga prev chapter id"),
			),

		goapidoc.NewDefinition("MangaPageLinkDto", "Manga page link response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("mname", "string", true, "manga name"),
				goapidoc.NewProperty("cover", "string", true, "manga cover"),
				goapidoc.NewProperty("url", "string", true, "manga link"),
				goapidoc.NewProperty("finished", "boolean", true, "manga is finished"),
				goapidoc.NewProperty("newest_chapter", "string", true, "manga last update chapter"),
			),

		goapidoc.NewDefinition("MangaChapterLinkDto", "Manga chapter link response").
			Properties(
				goapidoc.NewProperty("cid", "integer#int64", true, "manga chapter id"),
				goapidoc.NewProperty("cname", "string", true, "manga chapter name"),
				goapidoc.NewProperty("page_count", "integer#int32", true, "manga chapter page number"),
				goapidoc.NewProperty("url", "string", true, "manga chapter link"),
				goapidoc.NewProperty("new", "boolean", true, "manga chapter is uploaded newly"),
			),

		goapidoc.NewDefinition("MangaPageGroupDto", "Mange page group response").
			Properties(
				goapidoc.NewProperty("title", "string", true, "manga group title"),
				goapidoc.NewProperty("mangas", "MangaChapterLinkDto[]", true, "manga group mangas"),
			),

		goapidoc.NewDefinition("MangaChapterGroupDto", "Mange chapter group response").
			Properties(
				goapidoc.NewProperty("title", "string", true, "chapter group title"),
				goapidoc.NewProperty("chapters", "MangaChapterLinkDto[]", true, "chapter group chapters"),
			),

		goapidoc.NewDefinition("MangaGroupListDto", "Manga group list").
			Properties(
				goapidoc.NewProperty("title", "string", true, "manga group title"),
				goapidoc.NewProperty("top_group", "MangaPageGroupDto", true, "manga top page group"),
				goapidoc.NewProperty("groups", "MangaPageGroupDto[]", true, "manga page groups"),
				goapidoc.NewProperty("other_groups", "MangaPageGroupDto[]", true, "manga other page groups"),
			),
	)
}

// 漫画页的完整信息 vo.MangaPage
type MangaPageDto struct {
	Mid           uint64                  `json:"mid"`
	Mname         string                  `json:"mname"`
	Cover         string                  `json:"cover"`
	Url           string                  `json:"url"`
	PublishYear   string                  `json:"publish_year"`
	Zone          string                  `json:"zone"`
	AlphabetIndex string                  `json:"alphabet_index"`
	Type          string                  `json:"type"`
	AuthorName    string                  `json:"author_name"`
	Alias         string                  `json:"alias"`
	Finished      bool                    `json:"finished"`
	NewestChapter string                  `json:"newest_chapter"`
	NewestDate    string                  `json:"newest_date"`
	Introduction  string                  `json:"introduction"`
	Rank          string                  `json:"rank"`
	AverageScore  float32                 `json:"average_score"`
	ScoreCount    int32                   `json:"score_count"`
	PerScores     [6]float32              `json:"per_scores"`
	Groups        []*MangaChapterGroupDto `json:"groups"`
}

func BuildMangaPageDto(page *vo.MangaPage) *MangaPageDto {
	return &MangaPageDto{
		Mid:           page.Bid,
		Mname:         page.Bname,
		Cover:         page.Bpic,
		Url:           page.Url,
		PublishYear:   page.PublishYear,
		Zone:          page.Zone,
		AlphabetIndex: page.AlphabetIndex,
		Type:          page.Type,
		AuthorName:    page.AuthorName,
		Alias:         page.Alias,
		Finished:      page.Finished,
		NewestChapter: page.NewestChapter,
		NewestDate:    page.NewestDate,
		Introduction:  page.Introduction,
		Rank:          page.Rank,
		AverageScore:  page.AverageScore,
		ScoreCount:    page.ScoreCount,
		PerScores:     page.PerScores,
		Groups:        BuildMangaChapterGroupDtos(page.Groups),
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
	Mid     uint64   `json:"mid"`
	Mname   string   `json:"mname"`
	Cid     uint64   `json:"cid"`
	Cname   string   `json:"cname"`
	Url     string   `json:"url"`
	Pages   []string `json:"pages"`
	NextCid int32    `json:"next_cid"`
	PrevCid int32    `json:"prev_cid"`
}

func BuildMangaChapterDto(chapter *vo.MangaChapter) *MangaChapterDto {
	return &MangaChapterDto{
		Mid:     chapter.Bid,
		Mname:   chapter.Bname,
		Cid:     chapter.Cid,
		Cname:   chapter.Cname,
		Url:     chapter.Url,
		Pages:   chapter.Files,
		NextCid: chapter.NextId,
		PrevCid: chapter.PrevId,
	}
}

func BuildMangaChapterDtos(chapters []*vo.MangaChapter) []*MangaChapterDto {
	out := make([]*MangaChapterDto, len(chapters))
	for idx, chapter := range chapters {
		out[idx] = BuildMangaChapterDto(chapter)
	}
	return out
}

// 漫画页的链接 vo.MangaPageLink
type MangaPageLinkDto struct {
	Mid           uint64 `json:"mid"`
	Mname         string `json:"mname"`
	Cover         string `json:"cover"`
	Url           string `json:"url"`
	Finished      bool   `json:"finished"`
	NewestChapter string `json:"newest_chapter"`
}

func BuildMangaPageLinkDto(link *vo.MangaPageLink) *MangaPageLinkDto {
	return &MangaPageLinkDto{
		Mid:           link.Bid,
		Mname:         link.Bname,
		Cover:         link.Bpic,
		Url:           link.Url,
		Finished:      link.Finished,
		NewestChapter: link.NewestChapter,
	}
}

func BuildMangaPageLinkDtos(links []*vo.MangaPageLink) []*MangaPageLinkDto {
	out := make([]*MangaPageLinkDto, len(links))
	for idx, link := range links {
		out[idx] = BuildMangaPageLinkDto(link)
	}
	return out
}

// 漫画章节的链接 vo.MangaChapterLink
type MangaChapterLinkDto struct {
	Cid       uint64 `json:"cid"`
	Cname     string `json:"cname"`
	Url       string `json:"url"`
	PageCount int32  `json:"page_count"`
	New       bool   `json:"new"`
}

func BuildMangaChapterLinkDto(link *vo.MangaChapterLink) *MangaChapterLinkDto {
	return &MangaChapterLinkDto{
		Cid:       link.Cid,
		Cname:     link.Cname,
		Url:       link.Url,
		PageCount: link.PageCount,
		New:       link.New,
	}
}

func BuildMangaChapterLinkDtos(links []*vo.MangaChapterLink) []*MangaChapterLinkDto {
	out := make([]*MangaChapterLinkDto, len(links))
	for idx, link := range links {
		out[idx] = BuildMangaChapterLinkDto(link)
	}
	return out
}

// 漫画页分组 vo.MangaPageGroup
type MangaPageGroupDto struct {
	Title  string              `json:"title"`
	Mangas []*MangaPageLinkDto `json:"mangas"`
}

func BuildMangaPageGroupDto(group *vo.MangaPageGroup) *MangaPageGroupDto {
	return &MangaPageGroupDto{
		Title:  group.Title,
		Mangas: BuildMangaPageLinkDtos(group.Mangas),
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
	Chapters []*MangaChapterLinkDto `json:"chapters"`
}

func BuildMangaChapterGroupDto(group *vo.MangaChapterGroup) *MangaChapterGroupDto {
	return &MangaChapterGroupDto{
		Title:    group.Title,
		Chapters: BuildMangaChapterLinkDtos(group.Chapters),
	}
}

func BuildMangaChapterGroupDtos(groups []*vo.MangaChapterGroup) []*MangaChapterGroupDto {
	out := make([]*MangaChapterGroupDto, len(groups))
	for idx, group := range groups {
		out[idx] = BuildMangaChapterGroupDto(group)
	}
	return out
}

// 主页的漫画列表 vo.MangaGroupList
type MangaGroupListDto struct {
	Title       string               `json:"title"`
	TopGroup    *MangaPageGroupDto   `json:"top_group"`
	Groups      []*MangaPageGroupDto `json:"groups"`
	OtherGroups []*MangaPageGroupDto `json:"other_groups"`
}

func BuildMangaGroupListDto(list *vo.MangaGroupList) *MangaGroupListDto {
	return &MangaGroupListDto{
		Title:       list.Title,
		TopGroup:    BuildMangaPageGroupDto(list.TopGroup),
		Groups:      BuildMangaPageGroupDtos(list.Groups),
		OtherGroups: BuildMangaPageGroupDtos(list.OtherGroups),
	}
}

func BuildMangaGroupListDtos(lists []*vo.MangaGroupList) []*MangaGroupListDto {
	out := make([]*MangaGroupListDto, len(lists))
	for idx, list := range lists {
		out[idx] = BuildMangaGroupListDto(list)
	}
	return out
}
