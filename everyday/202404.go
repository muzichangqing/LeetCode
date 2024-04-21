package everyday

import (
	"container/list"
)

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

// 705. 设计哈希集合
const base = 768

type MyHashSet struct {
	data []list.List
}

func MyHashSetConstructor() MyHashSet {
	return MyHashSet{make([]list.List, base)}

}

func (this *MyHashSet) hash(key int) int {
	return key % base
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		h := this.hash(key)
		this.data[h].PushBack(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			this.data[h].Remove(e)
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}

// 39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	size := len(candidates)
	var dfs func(combina []int, sum, index int)
	dfs = func(combina []int, sum, index int) {
		temp := make([]int, len(combina))
		copy(temp, combina)
		combina = temp
		if sum == target {
			copyRow := make([]int, len(combina))
			copy(copyRow, combina)
			res = append(res, copyRow)
			return
		}
		if index >= size {
			return
		}
		for sum <= target {
			dfs(combina, sum, index+1)
			sum += candidates[index]
			combina = append(combina, candidates[index])
		}
	}
	dfs([]int{}, 0, 0)
	return res
}

// 216. 组合总和 III
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var dfs func(row *[]int, rk, rn, num int)
	dfs = func(row *[]int, rk, rn, num int) {
		if rk == 0 && rn == 0 {
			temp := make([]int, len(*row))
			copy(temp, *row)
			res = append(res, temp)
			return
		}
		if rk <= 0 || rn < num || num > 9 {
			return
		}
		newRow := append(*row, num)
		row = &newRow
		dfs(row, rk-1, rn-num, num+1)
		newRow = (*row)[:k-rk]
		row = &newRow
		dfs(row, rk, rn, num+1)
	}
	dfs(&[]int{}, k, n, 1)
	return res
}
