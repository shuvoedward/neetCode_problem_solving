package tree

func RightSideViewBFS(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := []int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		rightSide := 0
		qLen := len(q)

		for range qLen {
			node := q[0]
			q = q[1:]

			if node != nil {
				rightSide = node.Val
				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}
			}
		}
		res = append(res, rightSide)
	}
	return res
}

func RightSideViewDFS(root *TreeNode) []int {
	var res []int

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if depth == len(res) {
			res = append(res, node.Val)
		}
		dfs(node.Right, depth+1)
		dfs(node.Left, depth+1)
	}

	dfs(root, 0)
	return res
}
