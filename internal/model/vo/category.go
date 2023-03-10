package vo

// 索引项
type Category struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Url   string `json:"url"`
	Cover string `json:"cover"`
}

// 索引列表
type CategoryList struct {
	Genres []*Category
	Zones  []*Category
	Ages   []*Category
}
