# -*- coding=utf-8 -*-
class Solution(object):
    
    ways_num = 0
    nums = []
    nums_len = 0

    def findTargetSumWays(self, nums, S):
        """
        :type nums: List[int]
        :type S: int
        :rtype: int
        """
        self.nums = nums
        self.nums_len = len(nums)
        way = []
        select = []
        self.traceSumWay(select, way, S)
        return self.ways_num

    def traceSumWay(self, select, way, S):
        sum_way = sum(way)
        if sum_way == S:
            print(select)
            self.ways_num += 1
            return 
        elif sum_way > S:
            return 
        else:
            for i in range(len(self.nums)):
                if i in select:
                    continue
                select.append(i)
                way.append(self.nums[i])
                self.traceSumWay(select, way, S)
                way.pop()
                select.pop()

if __name__ == "__main__":
    res = Solution().findTargetSumWays([1,1,1,1,1],3)
    print(res)