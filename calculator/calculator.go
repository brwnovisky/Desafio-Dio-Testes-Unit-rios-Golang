package calculator

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}
func Subtrai(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}
func Multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}
func Divide(i ...int) int {
	total := 1
	for _, v := range i {
		total = v / total
	}
	return total
}
