https://leetcode.com/problems/two-sum/

Algorithm:
- Iterate through the slice
- Iterate through the slice again, without the value of i
- add nums[i] + nums[j] and check if it is equal to target
- Check if the numbers are already in the out
- if yes, then continue, else add the elements to out.
- Once the slice is over, return the out