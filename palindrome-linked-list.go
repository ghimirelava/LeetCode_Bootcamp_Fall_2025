/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {

    values := []int{}
    for head != nil {
        values = append(values, head.Val)
        head = head.Next
    }
    for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
        if values[i] != values[j] {
            return false
        }
    }
    return true
    
    
}