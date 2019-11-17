package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	if len1 == 0 {
		if len2%2 == 0 {
			return float64(nums2[(len2-1)/2]+nums2[(len2-1)/2+1]) / 2.0
		} else {
			return float64(nums2[(len2-1)/2])
		}
	}
	total := len1 + len2
	//fmt.Println(total)
	remainLen := (total + 1) / 2
	l1, l2, r1, r2 := 0, 0, len1, len2
	for {
		if l1 >= len1 {
			if total%2 == 0 {
				return float64(nums2[total/2-len1-1]+nums2[total/2-len1]) / 2.0
			} else {
				return float64(nums2[total/2-len1])
			}
		}
		if remainLen == 1 {
			var rst1, rst2 int
			if nums1[l1] < nums2[l2] {
				rst1 = nums1[l1]
				if l1+1 < len1 {
					if nums2[l2] < nums1[l1+1] {
						rst2 = nums2[l2]
					} else {
						rst2 = nums1[l1+1]
					}
				} else {
					rst2 = nums2[l2]
				}
			} else {
				rst1 = nums2[l2]
				if l2+1 < len2 {
					if nums2[l2+1] < nums1[l1] {
						rst2 = nums2[l2+1]
					} else {
						rst2 = nums1[l1]
					}
				} else {
					rst2 = nums1[l1]
				}
			}
			if total%2 == 0 {
				return float64(rst1+rst2) / 2.0
			} else {
				return float64(rst1)
			}
		}
		mid := remainLen / 2
		r1 = mid + l1 - 1
		r2 = mid + l2 - 1
		if mid > len1-1 {
			r1 = len1 - 1
		}
		if nums1[r1] <= nums2[r2] {
			l1 = r1 + 1
			remainLen -= mid
		} else {
			l2 = r2 + 1
			remainLen -= mid
		}
	}
	//return 0.0
}

func main() {
	//nums1 := []int{1, 7, 9, 15, 24, 28}
	//nums2 := []int{2, 4, 8, 17, 23, 32, 44}
	nums1 := []int{100001}
	nums2 := []int{100000}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
