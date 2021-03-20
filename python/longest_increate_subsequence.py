#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   longest_increate_subsequence.py
@Time    :   2021/02/10 11:32:28
@Author  :   Zzl
@Contact :   noaghzil@gmail.com
@Desc    :   一个未排序的整数数组，找到最长递增子序列的个数。
'''

# here put the import lib

# nums list is [1,3,5,4,7]，[1,1,1,2,2,2,3,3,3], [1,1,2,2,3]
# dp[{"mlen": 1, "num":1}] 
# 如果这些序列比 length[j] 长，那么我们就知道我们有count[i] 个长度为 length 的序列。如果这些序列的长度与 length[j] 相等，那么我们就知道现在有 count[i] 个额外的序列（即 count[j]+=count[i]）


class Solution1():
    
    def findNumberOfLIS(self, nums: list) -> int:
        dp = self.getLISDP(nums)
        print(dp)
        return self.findNumber(dp)
    
    def getLISDP(self, nums:list) -> dict:
        dp = {}
        for i, num in enumerate(nums):
            mxlis = 1
            mxnum = 1  
            for j in range(i):
                sub = nums[j]
                if sub < num:
                    if mxlis < dp[j]["mlen"] + 1:
                        mxlis = dp[j]["mlen"] +1
                        mxnum = dp[j]["num"]
                    elif mxlis == dp[j]["mlen"] + 1:
                        mxnum = mxnum + dp[j]["num"]
            dp[i] = {"mlen": mxlis, "num": mxnum}
        return dp
    
    def findNumber(self, dp:dict) -> int:
        max_len = 0
        max_len_num = 0
        for v in dp.values():
            if max_len > v["mlen"]:
                continue
            elif max_len < v["mlen"]:
                max_len = v["mlen"]
                max_len_num = v["num"]
            else:
                max_len_num += v["num"]
        return max_len_num


class Solution2(object):
    def findNumberOfLIS(self, nums):
        N = len(nums)
        if N <= 1: return N
        lengths = [1] * N #lengths[i] = longest ending in nums[i]
        counts = [1] * N #count[i] = number of longest ending in nums[i]

        for j, num in enumerate(nums):
            for i in range(j):
                if nums[i] < nums[j]:
                    if lengths[i] >= lengths[j]:
                        lengths[j] = 1 + lengths[i]
                        counts[j] = counts[i]
                    elif lengths[i] + 1 == lengths[j]:
                        counts[j] += counts[i]
        print(lengths,counts)
        longest = max(lengths)
        return sum(c for i, c in enumerate(counts) if lengths[i] == longest)



if __name__ == "__main__":
    pass
