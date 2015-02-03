//Package stack contains implementation of common Stacks
package stack

type node struct {
	item string
	next *node
}

type StringsStack struct {
	first *node
}

func New() *StringsStack {

	ss := new(StringsStack)
	return ss
}

func (ss *StringsStack) IsEmpty() bool {

	return (ss.first == nil)
}

func (ss *StringsStack) Push(item string) {

	oldFirst := ss.first
	ss.first = new(node)
	ss.first.item = item
	ss.first.next = oldFirst
}

func (ss *StringsStack) Pop() string {

	item := ss.first.item
	ss.first = ss.first.next
	return item
}