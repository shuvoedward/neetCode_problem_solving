package tree

/*
h
tree := TreeNode{
	Val : 1,
	Left : &TreeNode{ // go creates a struct TreeNode and assigns its address to Left
		Val : 2,
		Left : nil,
		Right: nil,
	}
	Right : &TreeNode{
		Val : 3,
		Left : nil,
		Right : nil,
	}

}

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursiveDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(recursiveDFS(root.Left), recursiveDFS(root.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
Time complexity O(n)
Space complexity O(n)
*/

// type TreeNode struct {
//     Val   int
//     Left  *TreeNode
//     Right *TreeNode
// }

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type NodeDepth struct {
		node  *TreeNode
		depth int
	}

	stack := []NodeDepth{{root, 1}}
	maxDepth := 0

	for len(stack) > 0 {
		// Pop the top element
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.node != nil {
			if current.depth > maxDepth {
				maxDepth = current.depth
			}
			// Push right and left children onto the stack with incremented depth
			stack = append(stack, NodeDepth{current.node.Left, current.depth + 1})
			stack = append(stack, NodeDepth{current.node.Right, current.depth + 1})
		}
	}

	return maxDepth
}
