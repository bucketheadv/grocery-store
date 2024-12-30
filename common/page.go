package common

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/schema"
	"strconv"
)

type BasePage struct {
	PageNo   int `json:"page" default:"1"`
	PageSize int `json:"pageSize" default:"10"`
}

type Page[T schema.Tabler] struct {
	PageNo   int `json:"page"`
	PageSize int `json:"pageSize"`
	Records  []T `json:"records"`
}

func (p *BasePage) Offset() int {
	var offset int
	if p.PageNo == 0 || p.PageNo == 1 {
		return 0
	}
	offset = (p.PageNo - 1) * p.PageSize
	return offset
}

func (p *Page[T]) SetRecords(records []T) {
	p.Records = records
}

func ParsePageParams(c *gin.Context) BasePage {
	var page = 1
	var limit = 10
	if pageNo, success := c.GetQuery("pageNo"); success {
		if p, e := strconv.Atoi(pageNo); e == nil {
			page = p
		}
	}

	if pageSize, success := c.GetQuery("pageSize"); success {
		if p, e := strconv.Atoi(pageSize); e == nil {
			limit = p
		}
	}
	return BasePage{
		PageNo: page, PageSize: limit,
	}
}
