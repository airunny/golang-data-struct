package binary_search

/**
 ** 在有序数组中查找元素
 *
 * 时间复杂度：O(logn)
 */
func BinarySearch(in []int, n int) int {
	s := 0
	e := len(in) - 1
	for s <= e {
		m := s + (e-s)>>1
		if in[m] == n {
			return m
		} else if in[m] > n {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return -1
}

/**
 ** 在有序数组中查找第一个出现的元素
 *
 * 时间复杂度：O(logn)
 */
func BinarySearchFirst(in []int, n int) int {
	s := 0
	e := len(in) - 1
	for s <= e {
		m := s + (e-s)>>1
		if in[m] > n {
			e = m - 1
		} else if in[m] < n {
			s = m + 1
		} else {
			if m == 0 || in[m-1] != n {
				return m
			}
			e = m - 1
		}
	}
	return -1
}

/**
 ** 在有序数组中查找最后一个出现的元素
 *
 * 时间复杂度：O(logn)
 */
func BinarySearchLast(in []int, n int) int {
	s := 0
	e := len(in) - 1
	for s <= e {
		m := s + (e-s)>>1
		if in[m] > n {
			e = m - 1
		} else if in[m] < n {
			s = m + 1
		} else {
			if m == len(in)-1 || in[m+1] != n {
				return m
			}
			s = m + 1
		}
	}
	return -1
}

/**
 ** 在有序数组中查找最后一个小于等于n的元素
 *
 * 时间复杂度：O(logn)
 */
func BinarySearchLastLessThan(in []int, n int) int {
	s := 0
	e := len(in) - 1
	for s <= e {
		m := s + (e-s)>>1
		if in[m] > n {
			e = m - 1
		} else {
			if m == len(in)-1 || in[m+1] > n {
				return m
			}
			s = m + 1
		}
	}
	return -1
}

/**
 ** 在有序数组中查找第一个大于等于n的元素
 *
 * 时间复杂度：O(logn)
 */
func BinarySearchFirstGreaterThan(in []int, n int) int {
	s := 0
	e := len(in) - 1
	for s <= e {
		m := s + (e-s)>>1
		if in[m] < n {
			s = m + 1
		} else {
			if m == 0 || in[m-1] < n {
				return m
			}
			e = m - 1
		}
	}
	return -1
}
