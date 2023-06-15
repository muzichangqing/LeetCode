package everyday

// 1177. 构建回文检测

func canMakePaliQueries(s string, queries [][]int) []bool {
	n := len(s)
	count := make([]int, n+1)
	for i := 0; i < n; i++ {
		count[i+1] = count[i] ^ (1 << (s[i] - 'a'))
	}
	res := make([]bool, len(queries))
	for i, query := range queries {
		l := query[0]
		r := query[1]
		k := query[2]
		bits := 0
		x := count[r+1] ^ count[l]
		for x > 0 {
			x &= x - 1
			bits++
		}
		res[i] = bits <= k*2+1
	}
	return res
}
