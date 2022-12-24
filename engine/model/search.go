package model

type SearchRequest struct {
	Query       string `json:"query,omitempty" form:"database"` // 搜索关键词
	Order       string `json:"order,omitempty" form:"database"` // 排序类型
	Page        int    `json:"page,omitempty" form:"database"`  // 页码
	Limit       int    `json:"limit,omitempty" form:"database"` // 每页大小，最大1000，超过报错
	SortType    string `json:"sorttype,omitempty" form:"sorttype"`
	FilterWords string `json:"filterwords,omitempty" form:"filterwords"`
	Database    string `json:"database" form:"database"` // 数据库名字
}

func (s *SearchRequest) GetAndSetDefault() *SearchRequest {

	if s.Limit == 0 {
		s.Limit = 100
	}
	if s.Page == 0 {
		s.Page = 1
	}

	if s.Order == "" {
		s.Order = "desc"
	}

	return s
}

type SearchResult struct {
	Time      float64       `json:"time,omitempty"`      //查询用时
	Total     int           `json:"total"`               //总数
	PageCount int           `json:"pageCount"`           //总页数
	Page      int           `json:"page,omitempty"`      //页码
	Limit     int           `json:"limit,omitempty"`     //页大小
	Documents []ResponseDoc `json:"documents,omitempty"` //文档
}
