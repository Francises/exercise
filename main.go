package main

import (
	"math"
	"sort"
	"fmt"
)

func main() {

}

//1
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	} else {
		lenth := 1
		temp := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] == temp {
				continue
			} else {
				temp = nums[i]
				nums[lenth] = nums[i]
				lenth++
			}
		}
		return lenth
	}
}

//2
func maxProfit(prices []int) int {
	output := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			output += prices[i] - prices[i-1]
		}
	}
	return output
}

//3
func rotate(nums []int, k int) {
	for n := k % len(nums); n > 0; n-- {
		move(nums)
	}
}

func move(nums []int) {
	temp := nums[0]
	bemp := 0
	for i := 1; i < len(nums); i++ {
		bemp = nums[i]
		nums[i] = temp
		temp = bemp
	}
	nums[0] = temp
}

//4.1 quick sort

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}

//4.2 java:hash set;go:map
func containsDuplicate2(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		_, b := m[num]
		if b {
			return true
		}
		m[num] = 1
	}
	return false
}

func contains3(nums []int) bool {
	m := make(map[int]int, len(nums))
	for _, num := range nums { //ignore index
		_, ok := m[num] //_:ignore this variability. here is value
		if ok {
			return true
		}
		m[num] = 0
	}
	return false
}

//5.1 dictionary
func singleNumber(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		_, ok := m[num]
		if ok {
			delete(m, num)
		} else {
			m[num] = 1
		}
	}
	for k, _ := range m {
		return k
	}
	return 0
}

//5.2 lie biao

//5.3 yi huo yun suan
func singleNumber3(nums []int) int {
	re := nums[0]
	for i := 1; i < len(nums); i++ {
		re ^= nums[i]
	}
	return re
}

//6.1
func intersect(nums1 []int, nums2 []int) []int {
	var r []int
	for _, num1 := range nums1 {
		for i, num2 := range nums2 {
			if num1 == num2 {
				r = append(r, num1)
				nums2 = append(nums2[:i], nums2[i+1:]...) //delete nums2[i]
				break
			}
		}
	}
	return r
}

//6.2
func intersect2(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	var r []int
	if len(nums1) <= len(nums2) {
		for _, num1 := range nums1 {
			_, ok := m[num1]
			if ok {
				m[num1]++
			} else {
				m[num1] = 1
			}
		}
		for _, num2 := range nums2 {
			v, ok := m[num2]
			if ok && v > 0 {
				r = append(r, num2)
				m[num2]--
			}
		}
	} else {
		for _, num2 := range nums2 {
			_, ok := m[num2]
			if ok {
				m[num2]++
			} else {
				m[num2] = 1
			}
		}
		for _, num1 := range nums1 {
			v, ok := m[num1]
			if ok && v > 0 {
				r = append(r, num1)
				m[num1]--
			}
		}
	}
	return r
}

//7
func plusOne(digits []int) []int {
	for l := len(digits) - 1; l >= 0; l-- {
		if digits[l] == 9 && l != 0 {
			digits[l] = 0
		} else if digits[l] == 9 && l == 0 {
			digits[l] = 0
			digits = append([]int{1}, digits...)
		} else {
			a := digits[l]
			digits[l] = a + 1
			return digits
		}
	}
	return digits
}

//8
func moveZeroes(nums []int) {
	var n int
	var j int
	for i, num := range nums {
		if num == 0 && i != len(nums)-1 && i != 0 {
			nums = append(nums[:j], nums[j+1:]...)
			n++
			j--
		} else if num == 0 && i == 0 {
			nums = nums[1:] //delete the first number
			n++
			j--
		}
		j++
		fmt.Println(nums)
		fmt.Println(j)
	}
	fmt.Println(n)
	for n > 0 {
		nums = append(nums, 0)
		n--
		fmt.Println(nums)
	}
}

//8.2
func moveZeroes2(nums []int) { //把所有非零的数字左移，i寻找非零数字，j记录应该存到哪个位置
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}
	for j < len(nums) { //把数组剩余位 置为零
		nums[j] = 0
		j++
	}
}

//9
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var a int
	var r []int
	for i, num := range nums {
		v, ok := m[num]
		if !ok {
			m[num] = i
			a = target - num
			v, ok := m[a]
			if ok && v != i {
				r = append(r, i)
				r = append(r, v)
				break
			}
		} else if ok && (num+num) == target {
			r = append(r, v)
			r = append(r, i)
			break
		}

	}
	return r
}

//10
func isValidSudoku(board [][]byte) bool {
	var index, temp int
	var box, line, column [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if string(board[i][j]) != "." {
				index = i/3*3 + j/3
				temp = int(board[i][j]) - 48 - 1
				if box[index][temp] == 1 || line[i][temp] == 1 || column[j][temp] == 1 {
					return false
				} else {
					box[index][temp] = 1
					line[i][temp] = 1
					column[j][temp] = 1
				}
			}
		}
	}
	return true
}

//11
func rotate11(matrix [][]int) { //先翻转，再行转列，两步操作均可通过数组元素位置互换完成
	n := len(matrix) - 1
	var temp int
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix); j++ {
			temp = matrix[i][j]
			matrix[i][j] = matrix[n][j]
			matrix[n][j] = temp
		}
		n--
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			temp = matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}

//12
func reverse(x int) int {
	var push int
	max := int(math.Pow(2, 31) / 10)
	for x != 0 {
		pop := x % 10
		x = x / 10
		if (push == max && pop > 7) || push > max || (push == -max && pop < -8) || push < -max {
			return 0
		}
		push = push*10 + pop
	}
	return push
}

//2.9
//func longestCommonPrefix(strs []string) string {
//	for i,str:=range strs{
//
//	}
//}



























