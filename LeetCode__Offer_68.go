package main

// 236 https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

// 法一 记录祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 记录遍历的父节点
	parent := map[int]*TreeNode{}
	visit := map[int]bool{}
	var dfs func(*TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		if r.Left != nil {
			// 左节点的父节点是 r
			parent[r.Left.Val] = r
			dfs(r.Left)
		}
		if r.Right != nil {
			// 右节点的父节点是 r
			parent[r.Right.Val] = r
			dfs(r.Right)
		}
	}
	dfs(root)

	for p != nil {
		visit[p.Val] = true
		// 递归回溯找父节点
		p = parent[p.Val]
	}
	for q != nil {
		if visit[q.Val] {
			// 访问值交错，找到父节点
			return q
		}
		q = parent[q.Val]
	}
	return nil
}

// 递归
func lowest(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowest(root.Left, p, q)
	right := lowest(root.Right, p, q)
	// 左右递归 都不为空，即在在该节点的左右都找到了目标节点
	if left != nil && right != nil {
		// 当前节点就是公共父节点
		return root
	}
	if left == nil {
		return right
	}
	return left
}
