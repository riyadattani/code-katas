package gameofthrees

func DivideBy3(num int) []int {
	var list []int

	for {
		list = append(list, num)
		if num == 1 {
			break
		}

		next := num % 3

		switch next {
		case 1:
			num -= 1
		case 2:
			num += 1
		case 0:
			num = num / 3
		}
	}

	return list
}
