package utils

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDecimal(t *testing.T) {
	Convey("测试浮点数精度的转换", t, func() {
		target1 := Decimal(37.625, 2)
		fmt.Println(target1)
		//So(target1, ShouldEqual, "20190302")

	})
}
