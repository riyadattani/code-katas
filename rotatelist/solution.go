package rotatelist

func Rotate(list []int, num int) []int {
	if num > len(list) || num < 0 {
		return []int{}
	}

	first := list[0:num]
	second := list[num:]
	second = append(second, first...)
	return second
}
