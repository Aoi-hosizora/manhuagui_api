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
				goapidoc.NewProperty("chapters", "MangaChapterGroupDto[]", true, "manga chapters"),
			),

		goapidoc.NewDefinition("MangaChapterGroupDto", "Mange chapter group response").
			Properties(
				goapidoc.NewProperty("title", "string", true, "chapter group title"),
				goapidoc.NewProperty("links", "MangaChapterLinkDto[]", true, "chapter group links"),
			),

		goapidoc.NewDefinition("MangaChapterLinkDto", "Manga chapter group url response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("mname", "string", true, "manga name"),
				goapidoc.NewProperty("cid", "integer#int64", true, "manga chapter id"),
				goapidoc.NewProperty("cname", "string", true, "manga chapter title"),
				goapidoc.NewProperty("pages", "integer#int32", true, "manga chapter page number"),
				goapidoc.NewProperty("url", "string", true, "manga chapter link"),
				goapidoc.NewProperty("new", "boolean", true, "manga chapter is uploaded newly"),
			),

		goapidoc.NewDefinition("MangaChapterDto", "Mange chapter response").
			Properties(
				goapidoc.NewProperty("cid", "integer#int64", true, "manga chapter id"),
				goapidoc.NewProperty("cname", "string", true, "manga chapter name"),
				goapidoc.NewProperty("url", "string", true, "manga chapter link"),
				goapidoc.NewProperty("pages", "string[]", true, "manga pages"),
				goapidoc.NewProperty("next_cid", "integer#int64", true, "manga next chapter id"),
				goapidoc.NewProperty("prev_cid", "integer#int64", true, "manga prev chapter id"),
			),
	)
}

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
	Chapters      []*MangaChapterGroupDto `json:"chapters"`
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
		Chapters:      BuildMangaChapterGroupDtos(page.Chapters),
	}
}

func BuildMangaPageDtos(pages []*vo.MangaPage) []*MangaPageDto {
	out := make([]*MangaPageDto, len(pages))
	for idx, page := range pages {
		out[idx] = BuildMangaPageDto(page)
	}
	return out
}

type MangaChapterGroupDto struct {
	Title string                 `json:"title"`
	Links []*MangaChapterLinkDto `json:"list"`
}

type MangaChapterLinkDto struct {
	Cid   uint64 `json:"cid"`
	Cname string `json:"cname"`
	Url   string `json:"url"`
	Pages int32  `json:"pages"`
	New   bool   `json:"new"`
}

func BuildMangaChapterGroupDto(group *vo.MangaChapterGroup) *MangaChapterGroupDto {
	links := make([]*MangaChapterLinkDto, len(group.Links))
	for idx, link := range group.Links {
		links[idx] = &MangaChapterLinkDto{
			Cid:   link.Cid,
			Cname: link.Cname,
			Url:   link.Url,
			Pages: link.Pages,
			New:   link.New,
		}
	}
	return &MangaChapterGroupDto{
		Title: group.Title,
		Links: links,
	}
}

func BuildMangaChapterGroupDtos(groups []*vo.MangaChapterGroup) []*MangaChapterGroupDto {
	out := make([]*MangaChapterGroupDto, len(groups))
	for idx, group := range groups {
		out[idx] = BuildMangaChapterGroupDto(group)
	}
	return out
}

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
