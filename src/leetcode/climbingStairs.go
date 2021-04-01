package main

//this is the factorial of the number n
func dynamicSolution(n int) int{
	if n == 0{
		return 0
	}
	if n == 1{
		return 1
	}
	pre := 1
	prePre := 1
	answer := 0
	for i := 2; i <= n;i++{
		answer = pre + prePre
		prePre = pre
		pre = answer
	}
	return answer
}

//this is the number of ways we can climb the stairs
func climbStairs(n int) int {
	return dynamicSolution(n)
}