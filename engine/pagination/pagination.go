/*
 Package pagination compute the sum of pages and index of data start\end of exact page
*/
package pagination

import (
	"math"
)

type Pagination struct {
	Limit int //data limitaion of each page

	PageCount int //sum of pages
	Total     int //data counts
}

func (p *Pagination) Init(limit int, total int) {
	p.Limit = limit

	//compute

	p.Total = total

	pageCount := math.Ceil(float64(total) / float64(limit))
	p.PageCount = int(pageCount)

}

func (p *Pagination) GetPage(page int) (s int, e int) {
	//get exact page
	if page > p.PageCount {
		page = p.PageCount
	}
	if page < 0 {
		page = 1
	}

	//always start at 1
	page -= 1

	//compute
	start := page * p.Limit
	end := start + p.Limit

	if start > p.Total {
		return 0, p.Total - 1
	}
	if end > p.Total {
		end = p.Total
	}

	return start, end

}
