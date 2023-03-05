package dto

import (
	"github.com/Aoi-hosizora/goapidoc"
	"github.com/Aoi-hosizora/manhuagui-api/src/model/vo"
)

func init() {
	goapidoc.AddDefinitions(
		goapidoc.NewDefinition("MangaDto", "Mange page response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("title", "string", true, "manga name"),
				goapidoc.NewProperty("cover", "string", true, "manga cover"),
				goapidoc.NewProperty("url", "string", true, "manga link"),
				goapidoc.NewProperty("publish_year", "string", true, "manga publish year"),
				goapidoc.NewProperty("manga_zone", "string", true, "manga zone"),
				goapidoc.NewProperty("genres", "CategoryDto[]", true, "manga genres"),
				goapidoc.NewProperty("authors", "TinyAuthorDto[]", true, "manga authors"),
				goapidoc.NewProperty("alias", "string", true, "manga alias name (deprecated)"),
				goapidoc.NewProperty("alias_title", "string", true, "manga alias title name (deprecated)"),
				goapidoc.NewProperty("aliases", "string", true, "manga aliases"),
				goapidoc.NewProperty("finished", "boolean", true, "manga is finished"),
				goapidoc.NewProperty("newest_chapter", "string", true, "manga last update chapter"),
				goapidoc.NewProperty("newest_date", "string", true, "manga last update date"),
				goapidoc.NewProperty("brief_introduction", "string", true, "manga brief introduction"),
				goapidoc.NewProperty("introduction", "string", true, "manga introduction"),
				goapidoc.NewProperty("manga_rank", "string", true, "manga rank"),
				goapidoc.NewProperty("score_count", "integer#int32", true, "manga score count"),
				goapidoc.NewProperty("average_score", "number#float", true, "manga average score"),
				goapidoc.NewProperty("per_scores", "string[]", true, "manga per scores, skip 0, from 1 to 5"),
				goapidoc.NewProperty("banned", "boolean", true, "manga is banned"),
				goapidoc.NewProperty("copyright", "boolean", true, "has copyright"),
				goapidoc.NewProperty("chapter_groups", "MangaChapterGroupDto[]", true, "manga chapter groups"),
			),

		goapidoc.NewDefinition("MangaChapterDto", "Mange chapter response").
			Properties(
				goapidoc.NewProperty("cid", "integer#int64", true, "chapter id"),
				goapidoc.NewProperty("title", "string", true, "chapter name"),
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("manga_title", "string", true, "manga name"),
				goapidoc.NewProperty("manga_cover", "string", true, "manga cover"),
				goapidoc.NewProperty("manga_url", "string", true, "manga link"),
				goapidoc.NewProperty("url", "string", true, "chapter link"),
				goapidoc.NewProperty("pages", "string[]", true, "chapter pages"),
				goapidoc.NewProperty("page_count", "integer#int32", true, "chapter pages count"),
				goapidoc.NewProperty("next_cid", "integer#int64", true, "next chapter id"),
				goapidoc.NewProperty("prev_cid", "integer#int64", true, "prev chapter id"),
				goapidoc.NewProperty("copyright", "boolean", true, "has copyright"),
			),

		goapidoc.NewDefinition("SmallMangaDto", "Small mange page response").
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

		goapidoc.NewDefinition("TinyMangaDto", "Tiny manga response").
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
				goapidoc.NewProperty("group", "string", true, "chapter group"),
				goapidoc.NewProperty("number", "boolean", true, "chapter number"),
			),

		goapidoc.NewDefinition("TinyBlockMangaDto", "Tiny block manga response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("title", "string", true, "manga name"),
				goapidoc.NewProperty("cover", "string", true, "manga cover"),
				goapidoc.NewProperty("url", "string", true, "manga link"),
				goapidoc.NewProperty("finished", "boolean", true, "manga is finished"),
				goapidoc.NewProperty("newest_chapter", "string", true, "manga last update chapter"),
			),

		goapidoc.NewDefinition("RandomMangaInfoDto", "Random manga info response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("url", "integer#int64", true, "manga link"),
			),

		goapidoc.NewDefinition("MangaGroupDto", "Mange page group response").
			Properties(
				goapidoc.NewProperty("title", "string", true, "group title"),
				goapidoc.NewProperty("mangas", "TinyBlockMangaDto[]", true, "group mangas"),
			),

		goapidoc.NewDefinition("MangaChapterGroupDto", "Mange chapter group response").
			Properties(
				goapidoc.NewProperty("title", "string", true, "group title"),
				goapidoc.NewProperty("chapters", "TinyMangaChapterDto[]", true, "group chapters"),
			),

		goapidoc.NewDefinition("MangaGroupListDto", "manga group list").
			Properties(
				goapidoc.NewProperty("title", "string", true, "list title"),
				goapidoc.NewProperty("top_group", "MangaGroupDto", true, "manga top page group"),
				goapidoc.NewProperty("groups", "MangaGroupDto[]", true, "manga groups"),
				goapidoc.NewProperty("other_groups", "MangaGroupDto[]", true, "manga other page groups"),
			),

		goapidoc.NewDefinition("HomepageMangaGroupListDto", "manga group list").
			Properties(
				goapidoc.NewProperty("serial", "MangaGroupListDto", true, "homepage serial manga group list"),
				goapidoc.NewProperty("finish", "MangaGroupListDto", true, "homepage finish group list"),
				goapidoc.NewProperty("latest", "MangaGroupListDto", true, "homepage latest manga group list"),
				goapidoc.NewProperty("daily", "MangaRankDto[]", true, "manga daily ranking for homepage"),
				goapidoc.NewProperty("genres", "CategoryDto[]", true, "manga all genres for homepage"),
				goapidoc.NewProperty("zones", "CategoryDto[]", true, "manga all zones for homepage"),
				goapidoc.NewProperty("ages", "CategoryDto[]", true, "manga all ages for homepage"),
			),

		goapidoc.NewDefinition("MangaRankDto", "Mange rank result response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "rank manga id"),
				goapidoc.NewProperty("title", "string", true, "rank manga title"),
				goapidoc.NewProperty("url", "string", true, "rank manga url"),
				goapidoc.NewProperty("finished", "boolean", true, "rank manga is finished"),
				goapidoc.NewProperty("authors", "string", true, "rank manga authors"),
				goapidoc.NewProperty("newest_chapter", "string", true, "rank manga newest chapter"),
				goapidoc.NewProperty("newest_date", "string", true, "rank manga newest date"),
				goapidoc.NewProperty("order", "integer#int32", true, "rank order"),
				goapidoc.NewProperty("score", "number#float", true, "rank manga score"),
				goapidoc.NewProperty("trend", "integer#int32", true, "rank trend, 0: None, 1: Up, 2: Down"),
			),

		goapidoc.NewDefinition("ShelfMangaDto", "Shelf manga response").
			Properties(
				goapidoc.NewProperty("mid", "integer#int64", true, "manga id"),
				goapidoc.NewProperty("title", "string", true, "manga name"),
				goapidoc.NewProperty("cover", "string", true, "manga cover"),
				goapidoc.NewProperty("url", "string", true, "manga link"),
				goapidoc.NewProperty("newest_chapter", "string", true, "manga last update chapter"),
				goapidoc.NewProperty("newest_duration", "string", true, "manga last update duration"),
				goapidoc.NewProperty("last_chapter", "string", true, "manga last read chapter"),
				goapidoc.NewProperty("last_duration", "string", true, "manga last read duration"),
			),
	)
}

// 漫画页的完整信息 vo.Manga
type MangaDto struct {
	Mid               uint64                  `json:"mid"`
	Title             string                  `json:"title"`
	Cover             string                  `json:"cover"`
	Url               string                  `json:"url"`
	PublishYear       string                  `json:"publish_year"`
	MangaZone         string                  `json:"manga_zone"`
	Genres            []*CategoryDto          `json:"genres"`
	Authors           []*TinyAuthorDto        `json:"authors"`
	Alias             string                  `json:"alias"`
	AliasTitle        string                  `json:"alias_title"`
	Aliases           []string                `json:"aliases"`
	Finished          bool                    `json:"finished"`
	NewestChapter     string                  `json:"newest_chapter"`
	NewestDate        string                  `json:"newest_date"`
	BriefIntroduction string                  `json:"brief_introduction"`
	Introduction      string                  `json:"introduction"`
	MangaRank         string                  `json:"manga_rank"`
	ScoreCount        int32                   `json:"score_count"`
	AverageScore      float32                 `json:"average_score"`
	PerScores         [6]string               `json:"per_scores"`
	Banned            bool                    `json:"banned"`
	Copyright         bool                    `json:"copyright"`
	ChapterGroups     []*MangaChapterGroupDto `json:"chapter_groups"`
}

func BuildMangaDto(manga *vo.Manga) *MangaDto {
	return &MangaDto{
		Mid:               manga.Mid,
		Title:             manga.Title,
		Cover:             manga.Cover,
		Url:               manga.Url,
		PublishYear:       manga.PublishYear,
		MangaZone:         manga.MangaZone,
		Genres:            BuildCategoryDtos(manga.Genres),
		Authors:           BuildTinyAuthorDtos(manga.Authors),
		Alias:             manga.Alias,
		AliasTitle:        manga.AliasTitle,
		Aliases:           manga.Aliases,
		Finished:          manga.Finished,
		NewestChapter:     manga.NewestChapter,
		NewestDate:        manga.NewestDate,
		Introduction:      manga.Introduction,
		BriefIntroduction: manga.BriefIntroduction,
		MangaRank:         manga.MangaRank,
		ScoreCount:        manga.ScoreCount,
		AverageScore:      manga.AverageScore,
		PerScores:         manga.PerScores,
		Banned:            manga.Banned,
		Copyright:         manga.Copyright,
		ChapterGroups:     BuildMangaChapterGroupDtos(manga.ChapterGroups),
	}
}

func BuildMangaDtos(mangas []*vo.Manga) []*MangaDto {
	out := make([]*MangaDto, len(mangas))
	for idx, manga := range mangas {
		out[idx] = BuildMangaDto(manga)
	}
	return out
}

// 漫画章节的完整信息 vo.MangaChapter
type MangaChapterDto struct {
	Cid        uint64   `json:"cid"`
	Title      string   `json:"title"`
	Mid        uint64   `json:"mid"`
	MangaTitle string   `json:"manga_title"`
	MangaCover string   `json:"manga_cover"`
	MangaUrl   string   `json:"manga_url"`
	Url        string   `json:"url"`
	Pages      []string `json:"pages"`
	PageCount  int32    `json:"page_count"`
	NextCid    int32    `json:"next_cid"`
	PrevCid    int32    `json:"prev_cid"`
	Copyright  bool     `json:"copyright"`
}

func BuildMangaChapterDto(chapter *vo.MangaChapter) *MangaChapterDto {
	return &MangaChapterDto{
		Cid:        chapter.Cid,
		Title:      chapter.Title,
		Mid:        chapter.Mid,
		MangaTitle: chapter.MangaTitle,
		MangaCover: chapter.MangaCover,
		MangaUrl:   chapter.MangaUrl,
		Url:        chapter.Url,
		Pages:      chapter.Pages,
		PageCount:  chapter.PageCount,
		NextCid:    chapter.NextId,
		PrevCid:    chapter.PrevId,
		Copyright:  chapter.Copyright,
	}
}

func BuildMangaChapterDtos(chapters []*vo.MangaChapter) []*MangaChapterDto {
	out := make([]*MangaChapterDto, len(chapters))
	for idx, chapter := range chapters {
		out[idx] = BuildMangaChapterDto(chapter)
	}
	return out
}

// 漫画页的部分信息 vo.SmallManga
type SmallMangaDto struct {
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

func BuildSmallMangaDto(manga *vo.SmallManga) *SmallMangaDto {
	return &SmallMangaDto{
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

func BuildSmallMangaDtos(mangas []*vo.SmallManga) []*SmallMangaDto {
	out := make([]*SmallMangaDto, len(mangas))
	for idx, manga := range mangas {
		out[idx] = BuildSmallMangaDto(manga)
	}
	return out
}

// 漫画页的部分信息 vo.TinyManga
type TinyMangaDto struct {
	Mid           uint64 `json:"mid"`
	Title         string `json:"title"`
	Cover         string `json:"cover"`
	Url           string `json:"url"`
	Finished      bool   `json:"finished"`
	NewestChapter string `json:"newest_chapter"`
	NewestDate    string `json:"newest_date"`
}

func BuildTinyMangaDto(manga *vo.TinyManga) *TinyMangaDto {
	return &TinyMangaDto{
		Mid:           manga.Mid,
		Title:         manga.Title,
		Cover:         manga.Cover,
		Url:           manga.Url,
		Finished:      manga.Finished,
		NewestChapter: manga.NewestChapter,
		NewestDate:    manga.NewestDate,
	}
}

func BuildTinyMangaDtos(mangas []*vo.TinyManga) []*TinyMangaDto {
	out := make([]*TinyMangaDto, len(mangas))
	for idx, manga := range mangas {
		out[idx] = BuildTinyMangaDto(manga)
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
	Group     string `json:"group"`
	Number    int32  `json:"number"`
}

func BuildTinyMangaChapterDto(chapter *vo.TinyMangaChapter) *TinyMangaChapterDto {
	return &TinyMangaChapterDto{
		Cid:       chapter.Cid,
		Title:     chapter.Title,
		Mid:       chapter.Mid,
		Url:       chapter.Url,
		PageCount: chapter.PageCount,
		IsNew:     chapter.IsNew,
		Group:     chapter.Group,
		Number:    chapter.Number,
	}
}

func BuildTinyMangaChapterDtos(chapters []*vo.TinyMangaChapter) []*TinyMangaChapterDto {
	out := make([]*TinyMangaChapterDto, len(chapters))
	for idx, link := range chapters {
		out[idx] = BuildTinyMangaChapterDto(link)
	}
	return out
}

// 漫画页的部分信息 vo.TinyBlockManga
type TinyBlockMangaDto struct {
	Mid           uint64 `json:"mid"`
	Title         string `json:"title"`
	Cover         string `json:"cover"`
	Url           string `json:"url"`
	Finished      bool   `json:"finished"`
	NewestChapter string `json:"newest_chapter"`
}

func BuildTinyBlockMangaDto(manga *vo.TinyBlockManga) *TinyBlockMangaDto {
	return &TinyBlockMangaDto{
		Mid:           manga.Mid,
		Title:         manga.Title,
		Cover:         manga.Cover,
		Url:           manga.Url,
		Finished:      manga.Finished,
		NewestChapter: manga.NewestChapter,
	}
}

func BuildTinyBlockMangaDtos(mangas []*vo.TinyBlockManga) []*TinyBlockMangaDto {
	out := make([]*TinyBlockMangaDto, len(mangas))
	for idx, manga := range mangas {
		out[idx] = BuildTinyBlockMangaDto(manga)
	}
	return out
}

// 随机漫画信息 vo.RandomMangaInfo
type RandomMangaInfoDto struct {
	Mid uint64 `json:"mid"`
	Url string `json:"url"`
}

func BuildRandomMangaInfoDto(info *vo.RandomMangaInfo) *RandomMangaInfoDto {
	return &RandomMangaInfoDto{
		Mid: info.Mid,
		Url: info.Url,
	}
}

func BuildRandomMangaInfoDtos(infos []*vo.RandomMangaInfo) []*RandomMangaInfoDto {
	out := make([]*RandomMangaInfoDto, len(infos))
	for idx, info := range infos {
		out[idx] = BuildRandomMangaInfoDto(info)
	}
	return out
}

// 漫画页分组 vo.MangaGroup
type MangaGroupDto struct {
	Title  string               `json:"title"`
	Mangas []*TinyBlockMangaDto `json:"mangas"`
}

func BuildMangaGroupDto(group *vo.MangaGroup) *MangaGroupDto {
	return &MangaGroupDto{
		Title:  group.Title,
		Mangas: BuildTinyBlockMangaDtos(group.Mangas),
	}
}

func BuildMangaGroupDtos(groups []*vo.MangaGroup) []*MangaGroupDto {
	out := make([]*MangaGroupDto, len(groups))
	for idx, group := range groups {
		out[idx] = BuildMangaGroupDto(group)
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

// 主页的漫画列表 vo.MangaGroupList
type MangaGroupListDto struct {
	Title       string           `json:"title"`
	TopGroup    *MangaGroupDto   `json:"top_group"`
	Groups      []*MangaGroupDto `json:"groups"`
	OtherGroups []*MangaGroupDto `json:"other_groups"`
}

func BuildMangaGroupListDto(list *vo.MangaGroupList) *MangaGroupListDto {
	return &MangaGroupListDto{
		Title:       list.Title,
		TopGroup:    BuildMangaGroupDto(list.TopGroup),
		Groups:      BuildMangaGroupDtos(list.Groups),
		OtherGroups: BuildMangaGroupDtos(list.OtherGroups),
	}
}

func BuildMangaGroupListDtos(lists []*vo.MangaGroupList) []*MangaGroupListDto {
	out := make([]*MangaGroupListDto, len(lists))
	for idx, list := range lists {
		out[idx] = BuildMangaGroupListDto(list)
	}
	return out
}

// 主页的三个漫画列表等数据 vo.HomepageMangaGroupList
type HomepageMangaGroupListDto struct {
	Serial *MangaGroupListDto `json:"serial"`
	Finish *MangaGroupListDto `json:"finish"`
	Latest *MangaGroupListDto `json:"latest"`
	Daily  []*MangaRankDto    `json:"daily"`
	Genres []*CategoryDto     `json:"genres"`
	Zones  []*CategoryDto     `json:"zones"`
	Ages   []*CategoryDto     `json:"ages"`
}

func BuildHomepageMangaGroupListDto(list *vo.HomepageMangaGroupList) *HomepageMangaGroupListDto {
	return &HomepageMangaGroupListDto{
		Serial: BuildMangaGroupListDto(list.Serial),
		Finish: BuildMangaGroupListDto(list.Finish),
		Latest: BuildMangaGroupListDto(list.Latest),
		Daily:  BuildMangaRankDtos(list.Daily),
		Genres: BuildCategoryDtos(list.Genres),
		Zones:  BuildCategoryDtos(list.Zones),
		Ages:   BuildCategoryDtos(list.Ages),
	}
}

func BuildHomepageMangaGroupListDtos(lists []*vo.HomepageMangaGroupList) []*HomepageMangaGroupListDto {
	out := make([]*HomepageMangaGroupListDto, len(lists))
	for idx, list := range lists {
		out[idx] = BuildHomepageMangaGroupListDto(list)
	}
	return out
}

// 漫画排名 vo.MangaRank
type MangaRankDto struct {
	Mid           uint64           `json:"mid"`
	Title         string           `json:"title"`
	Cover         string           `json:"cover"`
	Url           string           `json:"url"`
	Finished      bool             `json:"finished"`
	Authors       []*TinyAuthorDto `json:"authors"`
	NewestChapter string           `json:"newest_chapter"`
	NewestDate    string           `json:"newest_date"`
	Order         int8             `json:"order"`
	Score         float64          `json:"score"`
	Trend         uint8            `json:"trend"`
}

func BuildMangaRankDto(rank *vo.MangaRank) *MangaRankDto {
	return &MangaRankDto{
		Mid:           rank.Mid,
		Title:         rank.Title,
		Cover:         rank.Cover,
		Url:           rank.Url,
		Finished:      rank.Finished,
		Authors:       BuildTinyAuthorDtos(rank.Authors),
		NewestChapter: rank.NewestChapter,
		NewestDate:    rank.NewestDate,
		Order:         rank.Order,
		Score:         rank.Score,
		Trend:         rank.Trend,
	}
}

func BuildMangaRankDtos(ranks []*vo.MangaRank) []*MangaRankDto {
	out := make([]*MangaRankDto, len(ranks))
	for idx, rank := range ranks {
		out[idx] = BuildMangaRankDto(rank)
	}
	return out
}

// 书架漫画 vo.ShelfManga
type ShelfMangaDto struct {
	Mid            uint64 `json:"mid"`
	Title          string `json:"title"`
	Cover          string `json:"cover"`
	Url            string `json:"url"`
	NewestChapter  string `json:"newest_chapter"`
	NewestDuration string `json:"newest_duration"`
	LastChapter    string `json:"last_chapter"`
	LastDuration   string `json:"last_duration"`
}

func BuildShelfMangaDto(manga *vo.ShelfManga) *ShelfMangaDto {
	return &ShelfMangaDto{
		Mid:            manga.Mid,
		Title:          manga.Title,
		Cover:          manga.Cover,
		Url:            manga.Url,
		NewestChapter:  manga.NewestChapter,
		NewestDuration: manga.NewestDuration,
		LastChapter:    manga.LastChapter,
		LastDuration:   manga.LastDuration,
	}
}

func BuildShelfMangaDtos(mangas []*vo.ShelfManga) []*ShelfMangaDto {
	out := make([]*ShelfMangaDto, len(mangas))
	for idx, manga := range mangas {
		out[idx] = BuildShelfMangaDto(manga)
	}
	return out
}
