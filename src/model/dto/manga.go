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
			),

		goapidoc.NewDefinition("MangaChapterDto", "Mange chapter response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("mname", "string", true, "manga name"),
				goapidoc.NewProperty("cid", "integer#int64", true, "manga chapter id"),
				goapidoc.NewProperty("cname", "string", true, "manga chapter name"),
				goapidoc.NewProperty("pages", "string[]", true, "manga pages"),
				goapidoc.NewProperty("next_cid", "integer#int64", true, "manga next chapter id"),
				goapidoc.NewProperty("prev_cid", "integer#int64", true, "manga prev chapter id"),
			),
	)
}

type MangaPageDto struct {
	Mid uint64 `json:"mid"`
}

func BuildMangaPageDto(page *vo.MangaPage) *MangaPageDto {
	return &MangaPageDto{
		Mid: page.Bid,
	}
}

func BuildMangaPageDtos(pages []*vo.MangaPage) []*MangaPageDto {
	out := make([]*MangaPageDto, len(pages))
	for idx, page := range pages {
		out[idx] = BuildMangaPageDto(page)
	}
	return out
}

type MangaChapterDto struct {
	Mid     uint64   `json:"mid"`
	Mname   string   `json:"mname"`
	Cid     uint64   `json:"cid"`
	Cname   string   `json:"cname"`
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
