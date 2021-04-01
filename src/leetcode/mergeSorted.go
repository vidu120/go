package main

//the function to implement
func merge(nums1 []int, m int, nums2 []int, n int)  {

	counter1 := 0
	counter2 := 0

	parsedFirst := 0

	for parsedFirst < m && counter2 < n{
		if nums1[counter1] <= nums2[counter2]{
			counter1++
			parsedFirst++
		}else{
			save := nums2[counter2]
			for i := counter1;i <= counter1 + m - parsedFirst;i++{
				temp := nums1[i]
				nums1[i] = save
				save = temp
			}
			counter1++
			counter2++
		}
	}
	if counter2 < n{
		copy(nums1[counter1:] , nums2[counter2:])
	}
}
