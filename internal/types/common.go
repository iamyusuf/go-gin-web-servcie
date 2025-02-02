package types

var defaultSize = 20

type Pagination struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func (p Pagination) Offset() int {
	if p.Page == 0 {
		return 0
	}

	return (p.Page - 1) * p.Size
}

func (p Pagination) Limit() int {
	if p.Size == 0 {
		return defaultSize
	}
	return p.Size
}
