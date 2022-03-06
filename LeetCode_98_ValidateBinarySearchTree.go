package main

// 判断二叉树是不是二叉搜索树
// 二叉搜索树中序遍历为有序
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	} else if root.Left == nil && root.Right == nil {
		return true
	}
	res := []int{}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			// 中序，向左递归
			dfs(root.Left)
		}
		res = append(res, root.Val)
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	flag := false
	for i := 1; i <= len(res)-1; i++ {
		if res[i-1] < res[i] {
			flag = true
			continue
		} else {
			flag = false
			break
		}
	}
	return flag
}
