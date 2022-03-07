package main

/*
	自底向上的递归，后序遍历
	先处理根的左右子节点，左右子树事平衡二叉树，
	再判断以该根为数的是不是平衡二叉树，向上递归。
*/

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}
func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := height(root.Left)
	right := height(root.Right)
	if left == -1 || right == -1 || abs(left-right) > 1 {
		return -1
	}
	return max(left, right) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
