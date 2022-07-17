package main

import "fmt"

func main() {
	println(printNumWith2(3.1415))
	println(printBytes([]byte{97, 98}))
}

// 输出两位小数
func printNumWith2(float642 float64) string {
	return fmt.Sprintf("%.2f", float642)
}

func printBytes(data []byte) string {
	return fmt.Sprintf("%s", data)
}
