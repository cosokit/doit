package utils

import (
	"fmt"
	"math"
	"strconv"
)

/**
 * 判断两个浮点数是否相等
 */
func IsEqualFloat(f1, f2 float64, prec float64) bool {
	if prec <= 0.00000001 {
		prec = 0.00000001
	}

	if f1 > f2 {
		return f1-f2 < prec
	}

	return f2-f1 < prec
}

/**
 * 浮点数精度调整
 * value 要调整的浮点数
 * fixed 要保留的小数位数
 */
func Decimal(value float64, fixed int) float64 {
	if math.IsInf(value, 0) || math.IsNaN(value) {
		return 0.0
	}
	strFormat := fmt.Sprintf("%%.%df", fixed)
	value, _ = strconv.ParseFloat(fmt.Sprintf(strFormat, value), 64)

	return value
}
