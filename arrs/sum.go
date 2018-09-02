package arrs

func Sum(nums []int) int {
	sum := 0

	for _, number := range nums {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	//lengthOfNumbers := len(numbersToSum)
	//sums := make([]int, lengthOfNumbers)
	var sums []int

	for _, nums := range numbersToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))

		}

	}
	return sums
}
