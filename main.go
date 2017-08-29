package main

import (
	"fmt"
)
var i int64
	var j int64
	var k int64
	var minLength [3]int64
	var min int64
	var sum int64 = 0
	var num int64 = 0

func main() {
	//var numOfBulbs int = 8
	//var glowState [8]int = [8]int{0,1,0,0,1,1,0,0}
	//var distanceOfBulb [8]int = [8] int{3,5,10,12,13,23,30,38}

	//var numOfBulbs int = 7
	//var glowState [7]int = [7]int{1,0,1,1,0,1,1}
	//var distanceOfBulb [7]int = [7] int{1,5,6,7,8,9,17}

	var numOfBulbs int64 = 3
	var glowState [3]int64 = [3]int64{1,0,0}
	var distanceOfBulb [3]int64 = [3] int64{1,5,6}

	value :=minWireLength(numOfBulbs,glowState,distanceOfBulb)
	fmt.Println(value)

}

func minWireLength(numOfBulbs int64,glowState [3]int64,distanceOfBulb [3]int64) int64{
	return sumValue(numOfBulbs,glowState,distanceOfBulb)
}

func sumValue(numOfBulbs int64,glowState [3]int64,distanceOfBulb [3]int64) int64{
	for i=0;i<numOfBulbs;i++{
		if glowState[i] == 0 {
			num++
			for j=0;j<numOfBulbs;j++ {
				if glowState[j] == 1 || (j+1!=numOfBulbs && glowState[j+1] == 1) {
					minLength[j] = distanceOfBulb[j] - distanceOfBulb[i]
					if minLength[j] < 0 {
						minLength[j] = - minLength[j]
					}
				}else {
					minLength[j] = 0
				}
			}

			for k=0;k<j;k++ {
				if minLength[k] != 0 {
					min = minLength[k]
					break
				}
			}

			for k=0;k<j;k++{
					temp := minLength[k]
					if temp != 0 {
						if min > temp {
							min = temp
						}
					}
			}
			fmt.Println(min)
			sum += min
			glowState[i] = 1
		}
	}
	return sum
}
