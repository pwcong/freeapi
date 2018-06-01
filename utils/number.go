package utils

import (
	"strconv"
)

func ConvertHexadecimal(number string, from int, to int) (string, error) {
	base, err := strconv.ParseInt(number, from, 10)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(base, to), nil
}
func BToX(b string) (string, error) {
	return ConvertHexadecimal(b, 2, 16)
}
func XToB(x string) (string, error) {
	return ConvertHexadecimal(x, 16, 2)
}
func BToO(b string) (string, error) {
	return ConvertHexadecimal(b, 2, 8)
}
func OToB(o string) (string, error) {
	return ConvertHexadecimal(o, 8, 2)
}
func BToD(b string) (string, error) {
	return ConvertHexadecimal(b, 2, 10)
}
func DToB(d string) (string, error) {
	return ConvertHexadecimal(d, 10, 2)
}
func OToD(o string) (string, error) {
	return ConvertHexadecimal(o, 8, 10)
}
func DToO(d string) (string, error) {
	return ConvertHexadecimal(d, 10, 8)
}
func OToX(o string) (string, error) {
	return ConvertHexadecimal(o, 8, 16)
}
func XToO(x string) (string, error) {
	return ConvertHexadecimal(x, 16, 8)
}
func DToX(d string) (string, error) {
	return ConvertHexadecimal(d, 10, 16)
}
func XToD(x string) (string, error) {
	return ConvertHexadecimal(x, 16, 10)
}
