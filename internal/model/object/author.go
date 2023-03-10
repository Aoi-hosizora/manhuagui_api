package object

// 作者的完整信息
type Author struct {
	Aid               uint64             // 作者编号
	Name              string             // 作者名称
	Alias             string             // 作者别名
	Zone              string             // 作者地区
	Cover             string             // 作者封面
	Url               string             // 作者链接
	MangaCount        int32              // 漫画数量
	NewestMangaId     uint64             // 最新漫画
	NewestMangaTitle  string             // 最新漫画标题
	NewestMangaUrl    string             // 最新漫画链接
	NewestDate        string             // 收录时间
	HighestMangaId    uint64             // 最热漫画
	HighestMangaTitle string             // 最热漫画标题
	HighestMangaUrl   string             // 最热漫画链接
	HighestScore      float32            // 最热评分
	AverageScore      float32            // 平均得分
	Popularity        int32              // 人气指数
	Introduction      string             // 作者简介
	RelatedAuthors    []*TinyZonedAuthor // 相关作者
}

// 作者的部分信息 (Small)
type SmallAuthor struct {
	Aid        uint64 // 作者编号
	Name       string // 作者名称
	Zone       string // 作者地区
	Cover      string // 作者封面
	Url        string // 作者链接
	MangaCount int32  // 漫画数量
	NewestDate string // 更新时间
}

// 作者的部分信息 (Tiny)
type TinyAuthor struct {
	Aid  uint64 // 作者编号
	Name string // 作者名称
	Url  string // 作者链接
}

// 作者的部分信息 (Tiny+)
type TinyZonedAuthor struct {
	Aid  uint64 // 作者编号
	Name string // 作者名称
	Url  string // 作者链接
	Zone string // 作者地区
}
