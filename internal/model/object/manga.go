package object

// 漫画页的完整信息
type Manga struct {
	Mid               uint64               // 漫画编号
	Title             string               // 漫画标题
	Cover             string               // 漫画封面
	Url               string               // 漫画链接
	PublishYear       string               // 出品年代
	MangaZone         string               // 漫画地区
	AlphabetIndex     string               // 字母索引
	Genres            []*Category          // 漫画剧情
	Authors           []*TinyAuthor        // 漫画作者
	Alias             string               // 漫画别名 (x)
	AliasTitle        string               // 标题别名 (x)
	Aliases           []string             // 漫画别名
	Finished          bool                 // 是否完结
	NewestChapter     string               // 最新章节
	NewestDate        string               // 更新时间
	BriefIntroduction string               // 简要介绍
	Introduction      string               // 漫画介绍
	MangaRank         string               // 漫画排名
	ScoreCount        int32                // 给分人数
	AverageScore      float32              // 平均给分
	PerScores         [6]string            // 具体给分
	Banned            bool                 // 是否屏蔽 (x)
	Downed            bool                 // 是否下架
	Copyright         bool                 // 拥有版权
	Violent           bool                 // 色情暴力
	Lawblocked        bool                 // 法律屏蔽
	ChapterGroups     []*MangaChapterGroup // 章节链接
}

// 漫画章节的完整信息
type MangaChapter struct {
	Cid        uint64   `json:"cid"`      // 章节编号
	Title      string   `json:"cname"`    // 章节标题
	Mid        uint64   `json:"bid"`      // 漫画编号
	MangaTitle string   `json:"bname"`    // 漫画标题
	MangaCover string   `json:"bpic"`     // 漫画封面
	MangaUrl   string   `json:"-"`        // 漫画链接
	Url        string   `json:"url"`      // 章节链接
	Pages      []string `json:"files"`    // 每页链接
	Finished   bool     `json:"finished"` // 是否完结
	PageCount  int32    `json:"len"`      // 总共页数
	Path       string   `json:"path"`     // 链接路径
	NextId     int32    `json:"nextId"`   // 下一章节
	PrevId     int32    `json:"prevId"`   // 上一章节
	Copyright  bool     `json:"-"`        // 拥有版权
	Sl         *struct {
		E int64  `json:"e"` // e
		M string `json:"m"` // m
	} `json:"sl"` // 查询加密
}

// 漫画页的部分信息 (Small)
type SmallManga struct {
	Mid               uint64        // 漫画编号
	Title             string        // 漫画标题
	Cover             string        // 漫画封面
	Url               string        // 漫画链接
	PublishYear       string        // 出品年代
	MangaZone         string        // 漫画地区
	Genres            []*Category   // 漫画剧情
	Authors           []*TinyAuthor // 漫画作者
	Finished          bool          // 是否完结
	NewestChapter     string        // 最新章节
	NewestDate        string        // 更新时间
	BriefIntroduction string        // 简要介绍
}

// 漫画页的部分信息 (Smaller)
type SmallerManga struct {
	Mid           uint64   // 漫画编号
	Title         string   // 漫画标题
	Cover         string   // 漫画封面
	Url           string   // 漫画链接
	Finished      bool     // 是否完结
	Authors       []string // 漫画作者
	Genres        []string // 漫画分类
	NewestChapter string   // 最新章节
	NewestDate    string   // 更新时间
}

// 漫画页的部分信息 (Tiny)
type TinyManga struct {
	Mid           uint64 // 漫画编号
	Title         string // 漫画标题
	Cover         string // 漫画封面
	Url           string // 漫画链接
	Finished      bool   // 是否完结
	NewestChapter string // 最新章节
	NewestDate    string // 更新时间
}

// 漫画章节的部分信息 (Tiny)
type TinyMangaChapter struct {
	Cid       uint64 // 章节编号
	Title     string // 章节标题
	Mid       uint64 // 漫画编号
	Url       string // 章节链接
	PageCount int32  // 章节页数
	IsNew     bool   // 最近发布
	Group     string // 漫画分组
	Number    int32  // 章节顺序
}

// 漫画页的部分信息 (TinyBlock)
type TinyBlockManga struct {
	Mid           uint64 // 漫画编号
	Title         string // 漫画标题
	Cover         string // 漫画封面
	Url           string // 漫画链接
	Finished      bool   // 是否完结
	NewestChapter string // 最新章节
}

// 随机漫画信息
type RandomMangaInfo struct {
	Mid uint64 `json:"mid"`
	Url string `json:"url"`
}

// 漫画页分组 (Group)
type MangaGroup struct {
	Title  string            // 分组标题
	Mangas []*TinyBlockManga // 漫画集合
}

// 漫画章节分组 (Group)
type MangaChapterGroup struct {
	Title    string              // 分组标题
	Chapters []*TinyMangaChapter // 章节集合
}

// 主页的漫画列表
type MangaGroupList struct {
	Title       string        // 列表标题
	TopGroup    *MangaGroup   // 置顶分组
	Groups      []*MangaGroup // 类别分组
	OtherGroups []*MangaGroup // 其他分组
}

// 主页的三个漫画列表等数据
type HomepageMangaGroupList struct {
	Serial *MangaGroupList // 热门连载
	Finish *MangaGroupList // 经典完结
	Latest *MangaGroupList // 最新上架
	Daily  []*MangaRank    // 日排行榜
	Genres []*Category     // 漫画类别-剧情
	Zones  []*Category     // 漫画类别-地区
	Ages   []*Category     // 漫画类别-受众
}

// 漫画排名
type MangaRank struct {
	Mid           uint64        // 漫画编号
	Title         string        // 漫画标题
	Cover         string        // 漫画封面
	Url           string        // 漫画链接
	Finished      bool          // 是否完结
	Authors       []*TinyAuthor // 漫画作者
	NewestChapter string        // 最新章节
	NewestDate    string        // 更新时间
	Order         int8          // 漫画排名
	Score         float64       // 排名评分
	Trend         uint8         // 排名趋势
}

// 书架漫画
type ShelfManga struct {
	Mid            uint64 // 漫画编号
	Title          string // 漫画标题
	Cover          string // 漫画封面
	Url            string // 漫画链接
	NewestChapter  string // 最新章节
	NewestDuration string // 更新时间
	LastChapter    string // 最近阅读
	LastDuration   string // 最近时间
}
