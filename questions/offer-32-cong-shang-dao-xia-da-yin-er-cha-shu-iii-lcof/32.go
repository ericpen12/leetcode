package offer_32_cong_shang_dao_xia_da_yin_er_cha_shu_iii_lcof

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var n int
	var result [][]int
	for len(queue) > 0 {
		size := len(queue)
		var temp []int
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if n%2 != 0 {
			for i := 0; i < len(temp)/2; i++ {
				temp[i], temp[len(temp)-1-i] = temp[len(temp)-1-i], temp[i]
			}
		}
		result = append(result, temp)
		n++
	}
	return result
}
