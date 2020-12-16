func isMirror(n1 *TreeNode, n2 *TreeNode) bool {
	switch {
	case n1 == nil && n2 == nil:
		return true
	case n1 == nil || n2 == nil:
		return false
	default:
		return n1.Val == n2.Val && isMirror(n1.Right, n2.Left) && isMirror(n1.Left, n2.Right)
	}
}

 func isSymmetric(root *TreeNode) bool {
	 if root == nil {
		 return true
	 }
	 return isMirror(root.Left, root.Right)
 }