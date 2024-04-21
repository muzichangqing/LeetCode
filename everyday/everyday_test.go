package everyday

import "testing"

func TestDecodeTheMessage(t *testing.T) {
	var key, message, correct string

	key = "the quick brown fox jumps over the lazy dog"
	message = "vkbs bs t suepuv"
	correct = "this is a secret"
	if result := decodeMessage(key, message); result != correct {
		t.Fatalf("For key: %s, message: %s: expect \"%s\", got \"%s\"",
			key, message, correct, result)
	}
}

func TestShortestAlternatingPaths(t *testing.T) {
	t.Log(shortestAlternatingPaths(3, [][]int{{0, 1}, {0, 2}}, [][]int{}))
}

func TestMaximumConsecutive(t *testing.T) {
	var (
		coins   []int
		correct int
	)

	coins, correct = []int{1, 3}, 2
	if ans := getMaximumConsecutive(coins); ans != correct {
		t.Errorf("for %v, expect %d, got %d", coins, correct, ans)
	}
	coins, correct = []int{1, 4, 1, 1}, 8
	if ans := getMaximumConsecutive(coins); ans != correct {
		t.Errorf("for %v, expect %d, got %d", coins, correct, ans)
	}
	coins, correct = []int{1, 4, 10, 3, 1}, 20
	if ans := getMaximumConsecutive(coins); ans != correct {
		t.Errorf("for %v, expect %d, got %d", coins, correct, ans)
	}
}

func TestAlertNames(t *testing.T) {
	keyNames := []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"}
	keyTimes := []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"}

	t.Log(alertNames(keyNames, keyTimes))
}

func TestNumTimesAllBlue(t *testing.T) {
	flips := []int{3, 2, 4, 1, 5}
	correct := 2
	if ans := numTimesAllBlue(flips); ans != correct {
		t.Errorf("for %v, expect %d, got %d", flips, correct, ans)
	}
}

func TestCanMakePaliQueries(t *testing.T) {
	queris := [][]int{
		{3, 3, 0},
		{1, 2, 0},
		{0, 3, 1},
		{0, 3, 2},
		{0, 4, 1},
	}
	s := "abcda"
	correct := []bool{true, false, false, true, true}
	ans := canMakePaliQueries(s, queris)
	for i, v := range ans {
		if v != correct[i] {
			t.Errorf("query %v expect %v, got %v", queris[i], correct[i], v)
		}
	}
}

func TestMinNumberOfSemesters(t *testing.T) {
	n := 4
	relations := [][]int{{2, 1}, {3, 1}, {1, 4}}
	k := 2
	correct := 3
	if res := minNumberOfSemesters(n, relations, k); res != correct {
		t.Errorf("expect %d, got %d", correct, res)
	}
}

func TestReverseOddLevels(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 8},
			Right: &TreeNode{Val: 13},
		},
		Right: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 21},
			Right: &TreeNode{Val: 34},
		},
	}
	root = reverseOddLevels(root)
	printBinaryTreeByLevel(root, t)
}

func printBinaryTreeByLevel(root *TreeNode, t *testing.T) {
	if root == nil {
		t.Log("root = nil")
		return
	}
	args := []any{}

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		newQueue := []*TreeNode{}
		hasNotNil := false
		for _, node := range queue {
			if node == nil {
				args = append(args, "null")
				newQueue = append(newQueue, nil, nil)
				continue
			}
			args = append(args, node.Val)
			newQueue = append(newQueue, node.Left, node.Right)
			if node.Left != nil || node.Right != nil {
				hasNotNil = true
			}
		}
		if !hasNotNil {
			break
		}
		queue = newQueue
	}
	t.Log(args...)
}

func TestCombinationSum3(t *testing.T) {
	k, n := 3, 7
	res := combinationSum3(k, n)
	t.Log(res)
}
