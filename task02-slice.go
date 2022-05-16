package homework

func reverse(input []int64) (result []int64) {
	var b []int64
	for i := 0; i < len(input); i++ {
		b = append(b, input[len(input)-(1+i)])
	}
	//Place your code here
	return b
}
