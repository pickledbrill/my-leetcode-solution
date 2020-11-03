package main

// https://leetcode-cn.com/problems/valid-mountain-array/

func validMountainArray(A []int) bool {
	s := len(A)
	isValid := false
	if s <= 2 {
		return false
	}

	alen := len(A)
	blen := 2

	trendArray := []int{}
	peakArray := []bool{}

	increase := 0
	decrease := 1
	equal := -1

	for i := 0; i < alen; i += blen {
		if i == alen-1 {
			break
		}
		if alen <= i+blen {
			m := A[i:]
			if m[0] > m[1] {
				trendArray = append(trendArray, decrease)
				continue
			} else if m[0] < m[1] {
				trendArray = append(trendArray, increase)
				continue
			} else {
				trendArray = append(trendArray, equal)
				continue
			}
		} else {
			m := A[i : i+blen+1]
			if m[0] < m[1] && m[1] > m[2] {
				trendArray = append(trendArray, increase)
				trendArray = append(trendArray, decrease)
				peakArray = append(peakArray, true)
				continue
			} else if m[0] < m[1] && m[1] < m[2] {
				trendArray = append(trendArray, increase)
				continue
			} else if m[0] > m[1] && m[1] > m[2] {
				trendArray = append(trendArray, decrease)
				continue
			} else {
				println(m[0])
				println(m[1])
				println(m[2])

				trendArray = append(trendArray, equal)
				continue
			}
		}
	}

	if len(peakArray) > 1 {
		return false
	}

	for i := 0; i < len(trendArray); i++ {
		if equal == trendArray[i] {
			isValid = false
			break
		}
		if i == 0 && trendArray[i] == decrease {
			isValid = false
			break
		}

		if trendArray[i] == decrease {
			left := trendArray[:i]
			right := trendArray[i:]

			leftOk := true
			rightOk := true
			for _, lv := range left {
				if lv == decrease || lv == equal {
					leftOk = false
					break
				}
			}
			for _, rv := range right {
				if rv == increase || rv == equal {
					rightOk = false
					break
				}
			}

			isValid = leftOk && rightOk
			break
		}
	}

	return isValid
}
