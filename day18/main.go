package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
)

//type SNumber interface {
//	Parent() SNumber
//	Left() SNumber
//	Right() SNumber
//	Value() SNumber
//	IsValue() bool
//}
//
//type SNPair struct {
//	parent, left, right SNumber
//}
//
//func (s *SNPair) Parent() SNumber {
//	return s.parent
//}
//
//func (s *SNPair) Left() SNumber {
//	return s.left
//}
//
//func (s *SNPair) Right() SNumber {
//	return s.right
//}
//
//func (s *SNPair) Value() int {
//	return -1
//}
//
//func (s *SNPair) IsValue() bool {
//	return false
//}
//
//
//type SNEnd struct {
//	parent SNumber
//	value int
//}
//
//func (s *SNEnd) Parent() SNumber {
//	return s.parent
//}
//
//func (s *SNEnd) Left() SNumber {
//	return nil
//}
//
//func (s *SNEnd) Right() SNumber {
//	return nil
//}
//
//func (s *SNEnd) Value() int {
//	return s.value
//}
//
//func (s *SNEnd) IsValue() bool {
//	return true
//}

const IsDebug = false

func Debug(objs ...interface{}){
	if IsDebug {
		fmt.Println(objs...)
	}
}

type Number struct {
	Parent      *Number
	Left, Right *Number
	Value       int
}

func NewNumber(parent *Number) *Number {
	return &Number{
		Parent: parent,
		Left:   nil,
		Right:  nil,
		Value:  -1,
	}
}

func NewRegularNumber(parent *Number, value int) *Number {
	return &Number{
		Parent: parent,
		Left:   nil,
		Right:  nil,
		Value:  value,
	}
}

func (n Number) String() string {
	if n.Left == nil && n.Right == nil {
		return strconv.Itoa(n.Value)
	}
	return fmt.Sprintf("[%s,%s]", n.Left, n.Right)
}

func (n Number) Magnitude() int {
	if n.Left == nil && n.Right == nil {
		return n.Value
	}

	return 3 * n.Left.Magnitude() + 2 * n.Right.Magnitude()
}

func (n *Number) Add(addition *Number) *Number {
	parent := NewNumber(nil)
	parent.Left = n.deepCopy(parent)
	parent.Right = addition.deepCopy(parent)

	parent.reduce()
	return parent
}

func (n *Number) countParents() int {
	total := 0
	parent := n.Parent
	for parent != nil {
		total++
		parent = parent.Parent
	}
	return total
}

func (n *Number) getChildren() []*Number {
	// if self is Value
	if n.Left == nil && n.Right == nil {
		return []*Number{n}
	}

	return append(n.Left.getChildren(), n.Right.getChildren()...)
}

func (n *Number) selfAndParent() (*Number, *Number) {
	return n, n.Parent
}

func (n *Number) getLeftSibling() (*Number, bool) {
	current, parent := n.selfAndParent()
	// drill up
	for parent.Left == current {
		current, parent = parent.selfAndParent()
		if parent == nil {
			return nil, false// this is the left most pair, ignore lefty
		}
	}
	current = parent.Left
	// drill down
	for current.Right != nil {
		current = current.Right
	}
	return current, true
}

func (n *Number) getRightSibling() (*Number, bool) {
	current, parent := n.selfAndParent()
	// while current is the right most child, go up
	for parent.Right == current {
		current, parent = parent.selfAndParent()

		// if parent is nil, this is the root node,
		// thus there is no right more child. return nil
		if parent == nil {
			return nil, false
		}
	}

	current = parent.Right
	for current.Left != nil {
		current = current.Left
	}

	return current, true
}

func (n *Number) root() *Number {
	current := n
	for current.Parent != nil {
		current = current.Parent
	}
	return current
}

func (n *Number) deepCopy(parent *Number) *Number {
	if n.Left == nil && n.Right == nil {
		return NewRegularNumber(parent, n.Value)
	}

	cp := NewNumber(parent)
	cp.Left = n.Left.deepCopy(cp)
	cp.Right = n.Right.deepCopy(cp)
	return cp
}

func (n *Number) explode() {
	// add value left
	if sibling, ok := n.getLeftSibling(); ok {
		sibling.Value += n.Left.Value
	}

	// add value right
	if sibling, ok := n.getRightSibling(); ok {
		//Debug("Right Sib", sibling)
		sibling.Value += n.Right.Value
		//Debug("Expl. Right", n.root())
	}

	// Reduce this node to ValueNode of 0
	n.Left, n.Right = nil, nil
	n.Value = 0
}

func (n *Number) split() {
	if n.Value < 10 {
		Debug(n.root())
		Debug(n)
		panic("val < 10")
	}

	val := n.Value
	n.Value = -1

	left := val / 2
	n.Left = NewNumber(n)
	n.Left.Value = left

	right := val - left
	n.Right = NewNumber(n)
	n.Right.Value = right
}

func (n *Number) reduceOne() bool {
	for _, c := range n.getChildren() {
		count := c.countParents()
		//Debug("Parents: ", count, c)
		if count > 4 {
			c.Parent.explode()
			Debug("After Explode\t", n.root())
			return true
		}
	}

	for _, c := range n.getChildren() {
		if c.Value > 9 {
			c.split()
			Debug("After Split\t", n.root())
			return true
		}
	}
	return false
}

func (n *Number) reduce() {
	for n.reduceOne() {}
}

func main(){
	files := []string{
		"demo.txt",
		"input.txt",
	}

	for _, file := range files {
		fpath := path.Join("inputs", file)
		fmt.Println("File: ", fpath)
		if err := exec(fpath); err != nil {
			panic(fmt.Errorf("%s: %v", fpath, err))
		}
	}
}

func solve1(nums []*Number) int {
	sum := nums[0]
	Debug("Before", sum)
	for i := 1; i < len(nums); i++ {
		sum = sum.Add(nums[i])
		Debug("After Addition\t", sum)
	}

	//fmt.Println(sum)

	return sum.Magnitude()
}

func solve2(nums []*Number) int {
	max := 0
	for i := 0; i < len(nums) - 1; i++ {
		for j := i+1; j < len(nums); j++ {
			mag1 := nums[i].Add(nums[j]).Magnitude()
			if mag1 > max {
				max = mag1
			}

			mag2 := nums[j].Add(nums[i]).Magnitude()
			if mag2 > max {
				max = mag2
			}
		}
	}

	return max
}

func exec(in string) error {
	nums, err := readInput(in)
	if err != nil {
		return err
	}

	for _, n := range nums {
		n.reduce()
	}

	fmt.Println("solve1:", solve1(nums))
	fmt.Println("solve2:", solve2(nums))

	return nil
}

func parseNumber(in string) *Number {
	root := NewNumber(nil)
	ptr := root

	for _, c := range in {
		switch c {
		case '[':
			if ptr.Left != nil {
				ptr.Right = NewNumber(ptr)
				ptr = ptr.Right
			} else {
				ptr.Left = NewNumber(ptr)
				ptr = ptr.Left
			}
		case ',':
			ptr = ptr.Parent
			ptr.Right = NewNumber(ptr)
			ptr = ptr.Right
		case ']':
			ptr = ptr.Parent
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			ptr.Value = int(c - '0')
		}
	}

	return root
}

func readInput(in string) ([]*Number, error) {
	f, err := os.Open(in)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	numbers := make([]*Number, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		number := parseNumber(text)
		numbers = append(numbers, number)
	}
	return numbers, nil
}
