package stack

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	stk := New()

	stk.Push(10)

	stk.Push(5)

	stk.Push(11)

	fmt.Println(stk.Peek())

	stk.Pop()

}
