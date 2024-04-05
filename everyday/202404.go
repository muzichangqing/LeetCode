package everyday

// 20240405 1026. 节点与其祖先之间的最大差值
func maxAncestorDiff(root *TreeNode) int {
	getDiff := func(val1, val2 int) int {
		if val1 > val2 {
			return val1 - val2
		}
		return val2 - val1
	}
	getMax := func(val1, val2 int) int {
		if val1 > val2 {
			return val1
		}
		return val2
	}
	getMin := func(val1, val2 int) int {
		if val1 > val2 {
			return val2
		}
		return val1
	}
	var find func(node *TreeNode) (min, max, maxDiff int)
	find = func(node *TreeNode) (min int, max int, maxDiff int) {
		if node == nil {
			panic("node should't be nil")
		}
		min, max = node.Val, node.Val
		maxDiff = 0
		if node.Left != nil {
			minL, maxL, maxDiffL := find(node.Left)
			maxDiff = getMax(maxDiff, maxDiffL)
			maxDiff = getMax(maxDiff, getDiff(node.Val, minL))
			maxDiff = getMax(maxDiff, getDiff(node.Val, maxL))
			min = getMin(min, minL)
			max = getMax(max, maxL)
		}
		if node.Right != nil {
			minR, maxR, maxDiffR := find(node.Right)
			maxDiff = getMax(maxDiff, maxDiffR)
			maxDiff = getMax(maxDiff, getDiff(node.Val, minR))
			maxDiff = getMax(maxDiff, getDiff(node.Val, maxR))
			min = getMin(min, minR)
			max = getMax(max, maxR)
		}
		return
	}
	_, _, maxDiff := find(root)
	return maxDiff
}
