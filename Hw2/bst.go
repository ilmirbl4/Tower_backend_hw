package main

import "fmt"

type bstnode struct {
	val   int
	l_son *bstnode
	r_son *bstnode
}

func (node *bstnode) add_element(x int) {
	if node.val > x {
		if node.l_son == nil {
			var a bstnode = bstnode{
				val:   x,
				l_son: nil,
				r_son: nil,
			}
			node.l_son = &a
		} else {
			node.l_son.add_element(x)
		}
	} else if node.val < x {
		if node.r_son == nil {
			var a bstnode = bstnode{
				val:   x,
				l_son: nil,
				r_son: nil,
			}
			node.r_son = &a
		} else {
			node.r_son.add_element(x)
		}
	}
}
func (node *bstnode) Isexist(x int) bool {
	var res bool
	if node.val == x {
		res = true
	} else if node.val > x {
		if node.l_son == nil {
			res = false
		} else {
			return node.l_son.Isexist(x)
		}
	} else if node.val < x {
		if node.r_son == nil {
			res = false
		} else {
			return node.r_son.Isexist(x)
		}
	}
	return res
}
func find_max(node *bstnode) *bstnode {
	if node.r_son == nil {
		return node
	} else {
		node = find_max(node.r_son)
	}
	return node
}
func (node *bstnode) remove(x int) *bstnode {
	if node.val == x {
		if node.l_son == nil {
			right := node.r_son
			node = nil
			return right
		} else if node.r_son == nil {
			left := node.l_son
			node = nil
			return left
		} else if node.l_son != nil && node.r_son != nil {
			node_max := find_max(node.l_son)
			node.val = node_max.val
			node.l_son = node.l_son.remove(node_max.val)
		}
	} else if node.val > x {
		node = node.l_son.remove(x)
	} else if node.val < x {
		node = node.r_son.remove(x)
	}
	return node
}
func main() {
	x := 15
	var root bstnode = bstnode{
		val:   x,
		l_son: nil,
		r_son: nil,
	}
	fmt.Println(root)
	root.add_element(1)
	root.add_element(19)
	root.add_element(2)
	root.add_element(21)
	root.add_element(24)
	root.add_element(17)
	fmt.Println(root)
	fmt.Println(root.l_son.val, root.r_son.val, root.r_son.r_son.val)
	fmt.Println(root.Isexist(1))
	root.remove(21)
	root.remove(24)
	fmt.Print(root.r_son.val)
}
