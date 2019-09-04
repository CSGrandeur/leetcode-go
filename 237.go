//https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
// 学习GO记录的第一道题，这题直接给要删除的节点，只能把这个节点变成“下一个节点”，然后把“下一个节点”删除。
// 题目说了给定节点为非末尾节点，所以这样操作安全。
// 暂时不知道GO语言如何删除节点(手动释放内存)，所以……
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/
func deleteNode(node *ListNode) {
	*node = *(node.Next)
}