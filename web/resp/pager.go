package resp

// Pager represents a response object of pagination
type Pager struct {
	Offset int `json:"offset"`
	Take   int `json:"take"`
	Total  int `json:"total"`
}

// NewPager returns a new Pager
func NewPager(n, skip, limit int) *Pager {
	p := Pager{
		Offset: skip,
		Take:   limit,
		Total:  n,
	}
	if p.Offset >= p.Total {
		p.Take = 0
		return &p
	}
	if p.Total-p.Offset < p.Take {
		p.Take = p.Total - p.Offset
	}
	return &p
}
