package apis

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	// DefaultPageSize 返回列表的默认长度
	DefaultPageSize = 10
	// PageMaxCount 返回列表的最大长度
	PageMaxCount = 50
)

// ListOption 列表选项
type ListOption struct {
	Count   int64  `json:"count" form:"count"`
	Offset  int64  `json:"offset,omitempty" form:"offset"`
	Limit   int64  `json:"limit,omitempty" form:"limit"`
	ListAll bool   `json:"list_all,omitempty" form:"list_all"`
	Sort    string `json:"sort,omitempty" form:"sort"`
}

// Check 列表选项检查
func (o *ListOption) Check() error {
	if o.Offset <= 0 {
		o.Offset = 0
	}
	if o.Limit <= 0 || o.Limit > PageMaxCount {
		o.Limit = DefaultPageSize
	}
	return nil
}

type checker interface {
	Check() error
}

// Query get条件
func Query(ctx *gin.Context, v interface{}) error {
	if err := ctx.BindQuery(v); err != nil {
		return err
	}

	if ck, ok := v.(checker); ok {
		if err := ck.Check(); err != nil {
			return err
		}
	}
	return nil
}

// QueryID 获取路径或者query中的id
func QueryID(ctx *gin.Context, idName ...string) uint {
	pName := "id"
	if len(idName) > 0 && idName[0] != "" {
		pName = idName[0]
	}

	idstr := ctx.Param(pName)
	if idstr == "" {
		idstr = ctx.Query(pName)
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		logrus.Warnf("get param %s:%s failed, %s", pName, idstr, err)
		return 0
	}
	return uint(id)
}
