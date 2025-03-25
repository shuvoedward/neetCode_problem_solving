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

func interativeDFSMaximumDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	/*
		 Step 1: Define the struct (optional — for clarity)
		type NodeDepth struct {
		    node  *TreeNode
		    depth int
		}

		 Step 2: Initialize the slice
		stack := []NodeDepth{
		    {node: root, depth: 1}, // or {root, 1}
		}
	*/

	stack := []struct {
		node  *TreeNode
		depth int
	}{
		{root, 1},
	}

	// stack := [] *treenode {root}
	res := 0

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.node != nil {
			res = max(res, current.depth)

			// pushing the children onto the stack
			stack = append(stack, struct {
				node  *TreeNode
				depth int
			}{current.node.Left, current.depth + 1})

			stack = append(stack, struct {
				node  *TreeNode
				depth int
			}{current.node.Right, current.depth + 1})

		}
	}
	return res

}
