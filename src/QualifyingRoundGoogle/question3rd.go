package main

//YAAY PROBLEM DONE

//this method finds a particular slice that matches the cost given
func findParticularSlice(allottedRange, costRequired int) []int {
	costPosition := allottedRange - 1
	//by default even if our array is sorted we will get a cost of allottedRange which is the minimum cost that we can get in a reverse operation
	costRequired -= allottedRange - 1

	//our slice of elements
	array := make([]int, allottedRange)

	//boolean to keep track of the direction of filling the slice of integer
	leftToRight := true
	numberToFill := 1
	rightBoundary := allottedRange - 1
	leftBoundary := 0

	for numberToFill <= allottedRange {
		if costRequired >= costPosition {
			if leftToRight {
				array[leftBoundary+costPosition] = numberToFill
				rightBoundary--
			} else {
				array[rightBoundary-costPosition] = numberToFill
				leftBoundary++
			}
		} else {
			if leftToRight {
				array[leftBoundary+costRequired] = numberToFill
				temp := leftBoundary + costRequired
				for temp > leftBoundary {
					numberToFill++
					temp--
					array[temp] = numberToFill
				}
				temp = leftBoundary + costRequired
				for numberToFill < allottedRange {
					numberToFill++
					temp++
					array[temp] = numberToFill
				}
			} else {
				array[rightBoundary-costRequired] = numberToFill
				temp := rightBoundary - costRequired
				for temp < rightBoundary {
					numberToFill++
					temp++
					array[temp] = numberToFill
				}
				temp = rightBoundary - costRequired
				for numberToFill < allottedRange {
					numberToFill++
					temp--
					array[temp] = numberToFill
				}
			}
			break
		}
		leftToRight = !leftToRight
		costRequired -= costPosition
		numberToFill++
		costPosition--
	}
	return array
}

//check to see if the solution exists for the particular cost
func solutionExists(allottedRange, cost int) bool {
	if cost < allottedRange-1 || cost > ((allottedRange*(allottedRange+1))/2)-1 {
		return false
	}
	return true
}

//
//func main() {
//	var testCases , allottedRange, cost int
//	_, err := fmt.Scan(&testCases)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for i := 1; i <= testCases; i++ {
//		_ , err = fmt.Scanf("%d %d" , &allottedRange, &cost)
//		if err != nil{
//			log.Fatal(err)
//		}
//		if solutionExists(allottedRange , cost){
//			fmt.Printf("Case #%d: ",i)
//			for _ , value := range findParticularSlice(allottedRange , cost){
//				fmt.Printf("%d " , value)
//			}
//			fmt.Println();
//		}else{
//			fmt.Printf("Case #%d: IMPOSSIBLE\n" , i)
//		}
//	}
//}
