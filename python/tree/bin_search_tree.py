# -*- coding: utf-8 -*-
# @Author: Noaghzil
# @Date:   2023-04-15 19:24:00
# @Description: 二叉树的层序遍历

class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
        
def bin_tree_level_order(root):
    if not root:
        return []
    
    result = []
    queue = [root]
    
    while queue:
        level_size = len(queue)
        current_level = []
        
        for i in range(level_size):
            node = queue.pop(0)
            current_level.append(node.val)
            
            if node.left:
                queue.append(node.left)
                
            if node.right:
                queue.append(node.right)
                
        result.append(current_level)
        
    return result


def test_bin_tree_order():
    # 构建一棵二叉树
    #         3
    #        / \
    #       9   20
    #          /  \
    #         15   7
    root = TreeNode(3)
    root.left = TreeNode(9)
    root.right = TreeNode(20)
    root.right.left = TreeNode(15)
    root.right.right = TreeNode(7)

    # 预期的层序遍历结果
    expected_output = [[3], [9, 20], [15, 7]]

    # 调用层序遍历函数得到实际的结果
    assert bin_tree_level_order(root) == expected_output
    print(f"expected_output: {expected_output}")
    print("test_bin_tree_order passed")


def test_none_tree_order():
    # 构建一棵空树
    root = None

    # 预期的层序遍历结果为空列表
    expected_output = []

    # 调用层序遍历函数得到实际的结果
    assert bin_tree_level_order(root) == expected_output
    print(f"expected_output: {expected_output}")
    print("test_none_tree_order passed")

if __name__ == "__main__":
    test_bin_tree_order()
    test_none_tree_order()