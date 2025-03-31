package max

import "fmt"

func Exercise() {
	nums := []int{16, 8, 42, 4, 23, 15}
	result := nums[0]

	for _, num := range nums {
		if num > result {
			result = num
		}
	}

	fmt.Println(result)
}
