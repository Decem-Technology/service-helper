package helpers

import (
	"math"

	responsePB "github.com/Decem-Technology/service-helper/proto/response"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	"gorm.io/gorm"
)

type (
	// FetchQuery fetch query options
	FetchQuery struct {
		Type  string
		Query string
		Args  []interface{}
	}

	// PagingConfig paging config
	PagingConfig struct {
		DB         *gorm.DB
		FetchQuery []*FetchQuery
		Page       int
		PerPage    int
		OrderBy    []string
		ShowSQL    bool
	}

	// Paginator paging response
	Paginator struct {
		TotalRecord int         `json:"total_record"`
		TotalPage   int         `json:"total_page"`
		Records     interface{} `json:"data"`
		Offset      int         `json:"offset"`
		PerPage     int         `json:"PerPage"`
		Page        int         `json:"page"`
		PrevPage    int         `json:"prev_page"`
		NextPage    int         `json:"next_page"`
	}
)

// ToProtobuf parse paging to meta response
func (h *Paginator) ToProtobuf() (*responsePB.Meta, error) {
	meta := responsePB.Meta{
		Page:       &wrappers.Int32Value{Value: int32(h.Page)},
		PerPage:    &wrappers.Int32Value{Value: int32(h.PerPage)},
		PageCount:  &wrappers.Int32Value{Value: int32(h.TotalPage)},
		TotalCount: &wrappers.Int32Value{Value: int32(h.TotalRecord)},
	}
	return &meta, nil
}

// Paging query data with paging
func Paging(p *PagingConfig, result interface{}) (*Paginator, error) {
	db := p.DB

	if p.ShowSQL {
		db = db.Debug()
	}
	if p.Page < 1 {
		p.Page = 1
	}
	if p.PerPage == 0 {
		p.PerPage = 10
	}

	var paginator Paginator
	var count int64
	var offset int

	if err := db.Model(result).Count(&count).Error; err != nil {
		return nil, err
	}

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.PerPage
	}

	if len(p.FetchQuery) > 0 {
		for _, v := range p.FetchQuery {
			switch queryType := v.Type; queryType {
			case "select":
				db = db.Select(v.Query, v.Args...)
			case "join":
				db = db.Joins(v.Query, v.Args...)
			case "order":
				db = db.Order(v.Query)
			default:
			}
		}
	}

	if len(p.OrderBy) > 0 {
		for _, o := range p.OrderBy {
			db = db.Order(o)
		}
	}
	if err := db.Limit(p.PerPage).Offset(offset).Find(result).Error; err != nil {
		return nil, err
	}

	paginator.TotalRecord = int(count)
	paginator.Records = result
	paginator.Page = p.Page
	paginator.Offset = offset
	paginator.PerPage = p.PerPage
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.PerPage)))

	if p.Page > 1 {
		paginator.PrevPage = p.Page - 1
	} else {
		paginator.PrevPage = p.Page
	}

	if p.Page == paginator.TotalPage {
		paginator.NextPage = p.Page
	} else {
		paginator.NextPage = p.Page + 1
	}
	return &paginator, nil
}
