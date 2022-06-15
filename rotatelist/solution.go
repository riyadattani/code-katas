package rotatelist

func Rotate(list []int, num int) []int {
	if num > len(list) || num < 0 {
		return []int{}
	}
	return append(list[num:], list[0:num]...)
}
