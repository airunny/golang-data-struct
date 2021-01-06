package sort

/**
 ** 冒泡排序
 *
 * 最好情况(完全有序情况下)：O(n)
 * 最坏情况（完全逆序）：O(n2)
 * 平均时间复杂度：O(n2)
 * 是否稳定：true
 * 是否原地排序：true
 */
func BubbleSort(in []int) {
	for i := 0; i < len(in); i++ {
		flag := false
		for j := 0; j < len(in)-i-1; j++ {
			if in[j] > in[j+1] {
				in[j], in[j+1] = in[j+1], in[j]
				flag = true
			}
		}

		if !flag {
			return
		}
	}
}

/**
 ** 插入排序
 *
 * 最好情况(完全有序情况下)：O(n)
 * 最坏情况（完全逆序）：O(n2)
 * 平均时间复杂度：O(n2)
 * 是否稳定：true
 * 是否原地排序：true
 */
func InsertSort(in []int) {
	for i := 1; i < len(in); i++ {
		j := i - 1
		value := in[i]
		for j >= 0 {
			if in[j] > value {
				in[j+1] = in[j]
				j--
				continue
			}
			break
		}
		in[j+1] = value
	}
}

/**
 ** 选择排序(因为排序过程中的数据比较跟交换跟数据的逆序度无关；所以任何情况下的时间复杂度都是n2)
 *
 * 最好情况(完全有序情况下)：O(n2)
 * 最坏情况（完全逆序）：O(n2)
 * 平均时间复杂度：O(n2)
 * 是否稳定：false
 * 是否原地排序：true
 */
func SelectSort(in []int) {
	for i := 0; i < len(in)-1; i++ {
		index := i
		value := in[i]
		for j := i + 1; j < len(in); j++ {
			if in[j] < value {
				index = j
				value = in[j]
			}
		}
		if index != i {
			in[index], in[i] = in[i], in[index]
		}
	}
}

/**
 ** 归并排序
 *
 * 最好情况(完全有序情况下)：O(nlogn)
 * 最坏情况（完全逆序）：O(nlogn)
 * 平均时间复杂度：O(nlogn)
 * 是否稳定：true
 * 是否原地排序：false,空间复杂度 logn
 */
func MergeSort(in []int) {
	MergeS(in, 0, len(in)-1)
}

func MergeS(in []int, s, e int) {
	if s >= e {
		return
	}

	mid := s + (e-s)>>1
	MergeS(in, s, mid)
	MergeS(in, mid+1, e)
	Merge(in, s, mid, e)
}

func Merge(in []int, s, m, e int) {
	i := s
	j := m + 1

	temp := make([]int, 0, e-s+1)
	for i <= m && j <= e {
		if in[i] < in[j] {
			temp = append(temp, in[i])
			i++
		} else {
			temp = append(temp, in[j])
			j++
		}
	}

	ss := i
	ee := m
	if j <= e {
		ss = j
		ee = e
	}

	for i := ss; i <= ee; i++ {
		temp = append(temp, in[i])
	}

	for i := s; i <= e; i++ {
		in[i] = temp[i-s]
	}
}

/**
 ** 快速排序
 *
 * 最好情况(完全有序情况下)：O(nlogn)
 * 最坏情况（完全逆序,每次分区极度不均匀）：O(n2)
 * 平均时间复杂度：O(nlogn)
 * 是否稳定：false
 * 是否原地排序：true
 */
func QuickSort(in []int) {
	QuickSortS(in, 0, len(in)-1)
}

func QuickSortS(in []int, s, e int) {
	if s >= e {
		return
	}

	pivot := Partition(in, s, e)
	QuickSortS(in, 0, pivot-1)
	QuickSortS(in, pivot+1, e)
}

func Partition(in []int, s, e int) int {
	i := s
	pivot := in[e]
	for j := s; j <= e; j++ {
		if in[j] < pivot {
			in[i], in[j] = in[j], in[i]
			i++
		}
	}
	in[i], in[e] = in[e], in[i]
	return i
}

/*我是分割线============================================*/
// 1、在O(n)时间复杂度内找到数组中第K大元素
/**
 * 分析：
 *  partition分区时
 *  第一次遍历元素个数为 n/2
 *  第一次遍历元素个数为 n/2/2 = n/4
 *  第一次遍历元素个数为 n/2/2/2 = n /8
 *  ...
 * 所以总的次数为 等比数列的和 2n-1,忽略到常数时间复杂度 O(n)
 */
func FindKNumber(in []int, k int) int {
	if len(in) <= 0 || k <= 0 || k > len(in) {
		return -1
	}

	return FindNumber(in, k, 0, len(in)-1)
}

func FindNumber(in []int, k, s, e int) int {
	if s > e {
		return -1
	}

	pivot := Partition(in, s, e)
	if pivot+1 == k {
		return in[pivot]
	} else if pivot+1 > k {
		return FindNumber(in, k, 0, pivot-1)
	} else {
		return FindNumber(in, k, pivot+1, len(in)-1)
	}
}
