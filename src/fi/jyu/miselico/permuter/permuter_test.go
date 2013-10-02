package permuter

import (
	"container/list"
	"fmt"
	"testing"
)

func ExamplePermute() {
	myList := []interface{}{1, 2, 3}
	Permute(myList, func(l *list.List) {
		fmt.Print("[")
		for e := l.Front(); e != nil; e = e.Next() {
			fmt.Print(e.Value)
			if e.Next() != nil {
				fmt.Print(",")
			}
		}
		fmt.Print("]")
	})
	//Output:[1,2,3][2,1,3][2,3,1][1,3,2][3,1,2][3,2,1]
}

func TestEmptyPermute(t *testing.T) {
	original := []interface{}{}
	called := false
	Permute(original, func(permutation *list.List) {
		if called {
			t.Error("Permute with emty list calls sink more as once")
		}
		called = true
		if permutation.Len() != 0 {
			t.Error("Permute with empty list has wrong size expected 0, got ", permutation.Len())
		}
	})
	if !called {
		t.Error("sink did not get called for empty list")
	}
}

// var lenTests = []struct {
// 	length int
// 	//the sum of all numbers appearing in each position
// 	sum int
// }{
// 	{ 1, },
// }

func testLength() int {
	if testing.Short() {
		return 5
	} else {
		return 11
	}
}

func TestPermutationCount(t *testing.T) {
	maxLength := testLength()
	for length := 0; length <= maxLength; length++ {
		theList := ListOfFirstNPositiveNumbers(length)
		expectedCount := factorial(length)
		actual := 0
		Permute(theList, func(*list.List) {
			actual++
		})
		if actual != expectedCount {
			t.Errorf("Expected %d permutations for a list %s with length %d, got %d instead",
				expectedCount, theList, length, actual)
		}
	}
}


func TestPermuteSum(t *testing.T) {
	maxLength := testLength()
	for length := 0; length <= maxLength; length++ {
		theList := ListOfFirstNPositiveNumbers(length)
		actual := make([]int, length)
		Permute(theList, func(theList *list.List) {
			counter := 0
			for e := theList.Front(); e != nil; e = e.Next() {
				v := (e.Value).(int)
				actual[counter] += v
				counter++
			}
		})
		expectedSum := expectedSum(length)
		for _, value := range actual {
			if value != expectedSum {
				t.Errorf("Expected %d as the sum  for a list %s with length %d, got %d instead",
					expectedSum, theList, length, value)
			}
		}
	}
}


func TestPermuteMult(t *testing.T) {
	maxLength := 4 //higher number will overflow int64 !!
	
	for length := 0; length <= maxLength; length++ {
		theList := ListOfFirstNPositiveNumbers(length)
		actual := make([]uint64, length)
		for index := range actual{
			actual[index] = 1
		}
		Permute(theList, func(theList *list.List) {
			counter := 0
			for e := theList.Front(); e != nil; e = e.Next() {
				v := (e.Value).(int)
				actual[counter] *= uint64(v)
				counter++
			}
		})
		expectedProduct := intPow(factorial(length), factorial(length - 1))
		for _, value := range actual {
			if value != expectedProduct {
				t.Errorf("Expected %d as the product  for a list %s with length %d, got %d instead",
					expectedProduct, theList, length, value)
			}
		}
	}
}

func intPow(base, exp int) uint64{
	var result uint64 = 1
	
	base64 := uint64(base)
	for i:= 0; i < exp; i++{
		result *= base64
	}
	return result
}

func testSumOfPermutationElement(t *testing.T) {
	maxLength := testLength()
	for length := 0; length <= maxLength; length++ {
		theList := ListOfFirstNPositiveNumbers(length)
		actual := 0
		Permute(theList, func(theList *list.List) {
			for e := theList.Front(); e != nil; e = e.Next() {
				v := (e.Value).(int)
				actual += v
			}
		})
		expectedSum := ((length + 1) * length) / 2

		if actual != expectedSum {
			t.Errorf("Expected %d as the sum  for a list %s with length %d, got %d instead",
				expectedSum, theList, length, actual)
		}
	}
}

func ListOfFirstNPositiveNumbers(n int) []interface{} {
	theList := make([]interface{}, n)
	for i := 0; i < n; i++ {
		theList[i] = i + 1
	}
	return theList
}

func listToString(theList *list.List) string {

	s := fmt.Sprint("[")
	for e := theList.Front(); e != nil; e = e.Next() {
		s += fmt.Sprint(e.Value)
		if e.Next() != nil {
			s += fmt.Sprint(",")
		}
	}
	s += fmt.Sprint("]")
	return s
}

func expectedSum(length int) int {
	return ((length + 1) * length) / 2 * factorial(length-1)
}

func factorial(length int) int {
	result := 1
	for i := length; i > 1; i-- {
		result *= i
	}
	return result
}

/*
func expectedNumber(length int) int {
	return factorial(length)
}
*/
