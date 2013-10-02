//Package permuter provides a utility for permuting lists
package permuter

import (
	"container/list"

//	"fmt"
//	"log"
)

//Permute permutates the original list in the sink function
//the sink must not modify the list
func Permute(original []interface{}, sink func(permutation *list.List)) {
	length := len(original)
	p := newPermuter(original)
	permutationCount := factorial(length) //using factorial from test file
	//send the first permutation before looping
	sink(p.permutation)
	for i := 1; i < permutationCount; i++ {
		p.next()
		sink(p.permutation)
	}
}

/*
func factorial(length int) int {
	result := 1
	for i := length; i > 1; i-- {
		result *= i
	}
	return result
}
*/
type permuter struct {
	permutation  *list.List
	mainPermuter IPermuteHelper
}

func newPermuter(original []interface{}) permuter {
	permutation := list.New()
	helper := newPermuteHelper(original, permutation)
	p := permuter{permutation, helper}
	return p
}

func (this permuter) next() {
	this.mainPermuter.permute()
}

type IPermuteHelper interface {
	permute()
}

type permuteHelper struct {
	myElement       *list.Element
	recursionHelper IPermuteHelper
	permutation     *list.List
}

func newPermuteHelper(original []interface{}, permutation *list.List) IPermuteHelper {

	length := len(original)
	if length == 0 {
		return noOpPermuteHelper{}
	}
	myElement := permutation.PushBack(original[0])
	if length == 1 {
		return noOpPermuteHelper{}
	}
	recursionHelper := newPermuteHelper(original[1:], permutation)
	return &permuteHelper{myElement, recursionHelper, permutation}
}

func (this *permuteHelper) permute() {
	// precondition: the permutation list has length size, the owned
	// element is in the list
	next := this.myElement.Next()
	myValue := this.permutation.Remove(this.myElement)

	if next == nil {
		//endOfList
		this.recursionHelper.permute()
		this.myElement = this.permutation.PushFront(myValue)
	} else {
		this.myElement = this.permutation.InsertAfter(myValue, next)
	}
	//log.Println("at end myElement =", this.myElement, " in list" , &this.permutation)

	// postcondition the permutation list has length size, the owned
	// element is in the list at the next position
}

type noOpPermuteHelper struct{}

func (this noOpPermuteHelper) permute() {
	//log.Println("NoOpPermuteHelper")
}
