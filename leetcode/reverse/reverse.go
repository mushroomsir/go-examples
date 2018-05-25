package reverse

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -1 * x
	}
	res := 0
	for x > 0 {
		temp := x % 10
		res = res*10 + temp
		x = x / 10
	}
	res = sign * res
	return res
}
