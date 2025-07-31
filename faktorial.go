package main
import (
	"fmt"
	"math"
)
func hitung_faktorial(angka int) int {
	nilai := 1
	for i := 2; i <= angka; i++ {
		nilai *= i
	}
	return nilai
}
func hitung_rumus(angka int) int {
	faktorial := float64(hitung_faktorial(angka))
	pangkatDua := math.Pow(2, float64(angka))
	hasilBagi := faktorial / pangkatDua
	return int(math.Ceil(hasilBagi))
}
func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("f(%d) = %d\n", i, hitung_rumus(i))
	}
}