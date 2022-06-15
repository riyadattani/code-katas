package gameofthrees

func Count(n int) []int {
	var nums []int

	for {
		nums = append(nums, n)
		if n == 1 {
			break
		}

		remainder := n % 3

		if remainder%3 == 1 {
			n = n - 1
		}

		if remainder%3 == 2 {
			n = n + 1
		}

		if remainder%3 == 0 {
			n = n / 3
		}
	}

	return nums
}
