package leetcode

//通过扫描法来寻找最大和,O(n)
//	    |-- numbers[i] 如果i==0或者f(i-1)<0
//f(i)=>|
//      |-- f(i-1) + numbers[i] 如果f(i-1)>0
func MaxSubSum(numbers []int) int {
	if numbers == nil || len(numbers) <= 0 {
		return 0
	}
	var max1 int = 0
	var sum1 int = 0
	for _, value := range numbers {
		if sum1 < 0 {
			sum1 = value
		} else {
			sum1 += value
		}
		if max1 < sum1 {
			max1 = sum1
		}
	}
	return max1
}

func DivideAndConquerSum(numbers []int, low, high int) int {
	if numbers == nil || len(numbers) <= 0 {
		return 0
	}

	if low == high {
		return numbers[low]
	}

	middle := (low + high) / 2
	leftMax := DivideAndConquerSum(numbers, low, middle)
	rightMax := DivideAndConquerSum(numbers, middle+1, high)
	middleMax := divideSum(numbers, low, middle, high)
	return max(leftMax, middleMax, rightMax)
}

func divideSum(numbers []int, low, middle, high int) int {
	leftSum := numbers[middle]
	temp := 0
	for i := middle; i >= low; i-- {
		temp += numbers[i]
		if leftSum < temp {
			leftSum = temp
		}
	}

	rightSum := numbers[middle]
	temp = 0
	for i := middle + 1; i <= high; i++ {
		temp += numbers[i]
		if rightSum < temp {
			rightSum = temp
		}
	}
	return leftSum + rightSum
}

func max(left, middle, right int) int {
	if left > middle && left > right {
		return left
	}
	if right > left && right > middle {
		return right
	}
	return middle
}
