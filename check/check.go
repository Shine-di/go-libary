/**
 * @author: D-S
 * @date: 2021/1/23 11:40 上午
 */
// 参数校验
package check

import (
	"errors"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

// tag  名称 校验多个用`,`分割
const (
	tag    = "check"
	intMax = "max" //int最大值
	intMin = "min" //int最小值
)

// 目前支持的校验类型
var (
	restrictedTags = map[string]struct{}{
		intMax: {},
		intMin: {},
	}
)

type Check struct {
	checkFields map[string][]string      //校验的字段名和所对应的范围
	fieldValue  map[string]reflect.Value // 字段名和所对应的值
}

func New() *Check {
	return &Check{
		fields:     map[string][]string{},
		fieldValue: map[string]reflect.Value{},
	}
}

// 校验struct
func (l *Check) Struct(data interface{}) error {
	fields := reflect.ValueOf(data).Elem()
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Type().Field(i)
		tags := field.Tag.Get(tag)
		if tags == "" {
			continue
		}
		//缓存字段名和所对应的值
		l.fieldValue[field.Name] = fields.FieldByName(field.Name)
		//缓存 需要校验的字段
		for _, e := range strings.Split(tags, ",") {
			l.checkFields[field.Name] = append(l.checkFields[field.Name], e)
		}
	}
	return nil
}

// 校验字段
func (l *Check) Field(field) {

}

//todo gin 中间件
func (l *Check) GINMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

func (l *Check) doCheck() error {
	// 没有校验的字段
	if len(l.checkFields) == 0 {
		return nil
	}
	if len(l.checkFields) != len(l.fieldValue) {
		return errors.New("校验的字段数量和所对应的数据值数量不匹配")
	}

	return nil
}

//校验字段和值
func (l *Check) fieldCheck(field string, tags []string, value reflect.Value) error {
	return nil
}
