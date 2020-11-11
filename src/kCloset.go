package main

// https://leetcode-cn.com/problems/k-closest-points-to-origin/

func kClosest(points [][]int, K int) [][]int {
	result := [][]int{}
	empty := []int{0, 0}

	for _, pv := range points {
		if len(result) == 0 {
			result = append(result, pv)
			continue
		}
		for ri, rv := range result {
			p := (pv[0] * pv[0]) + (pv[1] * pv[1])
			r := (rv[0] * rv[0]) + (rv[1] * rv[1])
			if p <= r {
				result = append(result, empty)
				copy(result[ri+1:], result[ri:])
				result[ri] = pv
				break
			}

			if len(result)-1 == ri {
				result = append(result, pv)
			}
		}
	}

	return result[:K]
}
