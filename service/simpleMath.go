package service

func Sum(a int, b int) []int {
	var total []int
	sum := a + b
	divide := 2 * (b / a)
	multi := divide * 2
	sum2 := sum + multi
	subtract := sum2 - multi

	total = append(total, sum, divide, multi, sum2, subtract)

	return total
}
