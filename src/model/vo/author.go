package vo

// 作者的完整信息
type Author struct {
	Aid              uint64  // 作者编号
	Name             string  // 作者名称
	Zone             string  // 作者地区
	Cover            string  // 作者封面
	Url              string  // 作者链接
	MangaCount       int32   // 漫画数量
	NewestMangaId    uint64  // 最新漫画
	NewestMangaTitle string  // 漫画标题
	NewestDate       string  // 更新时间
	Alias            string  // 作者别名
	AverageScore     float32 // 平均得分
	Introduction     string  // 作者简介
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
