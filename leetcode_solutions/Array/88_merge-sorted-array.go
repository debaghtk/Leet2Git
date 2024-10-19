# Merge Sorted Array
# Difficulty: Easy
# Language: golang
# Topic: Array
# Tags: Array, Two Pointers, Sorting
# Link: https://leetcode.com/problems/merge-sorted-array/

func merge(nums1 []int, m int, nums2 []int, n int)  {
    
    for i,j:=0,0;  i<= m && j<=n; i,j=i,j{
        if i==m {
            copy(nums1[i:], nums2[j:])
            break
        }
        
        if j==n {
            break
        }
        
        if nums1[i] > nums2[j]{
            copy(nums1[i+1:], nums1[i:])
            nums1[i]=nums2[j]
            j++
            m++
            continue
        }
        
        if nums1[i] <= nums2[j]{
            i++
        }
    
    }
    
}
