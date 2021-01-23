/**
 * @author: D-S
 * @date: 2021/1/23 11:40 上午
 */
// 参数校验
package check

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"strings"
)

// tag  名称 校验多个用`,`分割
const (
	tag    = "check"
	intMax = "max"  //int最大值 max=10
	intMin = "min"  //int最小值 min=20
	enum   = "enum" //枚举值只能是 enum=[1,2,3,2]
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
	result := &Check{}
	result.defaultCheck()
	return result
}

func (l *Check) defaultCheck() {
	l.checkFields = map[string][]string{}
	l.fieldValue = map[string]reflect.Value{}
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
			_, ok := restrictedTags[e]
			if !ok {
				return errors.New(fmt.Sprintf("暂不支持校验类型:%v", e))
			}
			l.checkFields[field.Name] = append(l.checkFields[field.Name], e)
		}
	}
	return nil
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
	for k, v := range l.checkFields {
		value := l.fieldValue[k]
		if err := l.fieldCheck(k, v, value); err != nil {
			return err
		}
	}
	return nil
}

//校验字段和值
func (l *Check) fieldCheck(field string, tags []string, value reflect.Value) error {
	for _, e := range tags {
		switch e {
		case intMax:
			//eg max=10
			checkValues := strings.Split(e, "=")
			if len(checkValues) != 2 {
				return errors.New(fmt.Sprintf("校验tag参数不满足,原始数据%v", e))
			}
			valid, _ := strconv.ParseInt(checkValues[1], 10, 64)
			v := value.Int()
			if !l.fieldIntMax(valid, v) {
				return errors.New(field + "不能大于" + strconv.Itoa(int(v)))
			}
		case intMin:
			//eg min=10
			checkValues := strings.Split(e, "=")
			if len(checkValues) != 2 {
				return errors.New(fmt.Sprintf("校验tag参数不满足,原始数据%v", e))
			}
			valid, _ := strconv.ParseInt(checkValues[1], 10, 64)
			v := value.Int()
			if !l.fieldIntMin(valid, v) {
				return errors.New(field + "不能小于" + strconv.Itoa(int(v)))
			}
		case enum:
			//eg
			checkValues := strings.Split(e, "=")
			if len(checkValues) != 2 {
				return errors.New(fmt.Sprintf("校验tag参数不满足,原始数据%v", e))
			}
			enum := checkValues[1]
			enum = strings.ReplaceAll(enum, "[", "")
			enum = strings.ReplaceAll(enum, "]", "")
			valid := strings.Split(enum, ",")
			v := value.String()
			if !l.fieldEnum(valid, v) {
				return errors.New(fmt.Sprintf("%v不在枚举值:[%s]之中,实际值%v", field, enum, v))
			}
		default:
			return errors.New("暂不支持的校验类型" + e)
		}
	}
	return nil
}

// 校验字段 大于  valid 目标  real实际
func (l *Check) fieldIntMax(valid int64, real int64) bool {
	if real > valid {
		return false
	}
	return true
}

// 校验字段 小于  valid 目标  real实际
func (l *Check) fieldIntMin(valid int64, real int64) bool {
	if real < valid {
		return false
	}
	return true
}

// 校验字段 枚举值  valid 目标  real实际 实际值只能是枚举值中的几个
func (l *Check) fieldEnum(valid []string, real string) bool {
	for _, e := range valid {
		if e == real {
			return true
		}
	}
	return false
}
