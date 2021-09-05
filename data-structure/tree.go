package data_structure

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

type traverseFunc = func(p *TreeNode)

func beforeTraverse(root *TreeNode, f traverseFunc) {
	stack := make([]*TreeNode, 0)
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			if f != nil {
				f(p)
			}
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = top.Right
		}
	}
}

func middleTraverse(root *TreeNode, f traverseFunc) {
	stack := make([]*TreeNode, 0)
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if f != nil {
				f(top)
			}
			p = top.Right
		}
	}
}

func postOrderTraverse(root *TreeNode, f traverseFunc) {
	stack := []*TreeNode{root}
	var (
		prev *TreeNode
		p    *TreeNode
	)
	for len(stack) > 0 {
		p = stack[len(stack)-1]
		if p.Left == nil && p.Right == nil || // p is leaf node
			(prev != nil && (prev == p.Left || prev == p.Right)) { // or p's child node has traversed
			if f != nil {
				f(p)
			}
			prev = p
			stack = stack[:len(stack)-1]
		} else {
			if p.Right != nil {
				stack = append(stack, p.Right)
			}
			if p.Left != nil {
				stack = append(stack, p.Left)
			}
		}
	}
}

func buildTreeList(vs ...int) *TreeNode {
	arr := make([]TreeNode, len(vs))
	for i, v := range vs {
		arr[i].Value = v
		l := i*2 + 1
		if l < len(vs) {
			arr[i].Left = &arr[l]
		}
		r := i*2 + 2
		if r < len(vs) {
			arr[i].Right = &arr[r]
		}
	}
	return &arr[0]
}

func TreeRun() {
	root := buildTreeList(1, 2, 3, 4, 5, 6, 7)
	postOrderTraverse(root, func(p *TreeNode) {
		fmt.Println(p.Value, "  ")
	})
}
