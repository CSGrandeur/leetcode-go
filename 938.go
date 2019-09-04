// 二叉树常规操作。
// None，NULL，nullptr，nil。。。不同语言的“空”真是五花八门啊
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rangeSumBST(root *TreeNode, L int, R int) int {
    ret := 0
    if root.Left != nil && root.Val > L {
        ret += rangeSumBST(root.Left, L, R)
    }
    if root.Right != nil && root.Val < R {
        ret += rangeSumBST(root.Right, L, R)
    }
    if L <= root.Val && root.Val <= R {
        ret += root.Val
    }
    return ret
}