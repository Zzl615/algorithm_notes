# -*- coding: utf-8 -*-
# @Author: Noaghzil
# @Date:   2023-04-15 16:12:30
# @Description: 二叉搜索树
import pytest
from abc import ABCMeta, abstractmethod

class Node:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None
        
class BSTree(metaclass=ABCMeta):
    @abstractmethod
    def insert(self, value):
        pass
    
    @abstractmethod
    def search(self, value):
        pass
    
    @abstractmethod
    def delete(self, value):
        pass

class BSTImpl(BSTree):
    def __init__(self):
        self.root = None
        
    def insert(self, value):
        if self.root is None:
            self.root = Node(value)
        else:
            self._insert(value, self.root)
            
    def _insert(self, value, node):
        if value < node.value:
            if node.left is None:
                node.left = Node(value)
            else:
                self._insert(value, node.left)
        else:
            if node.right is None:
                node.right = Node(value)
            else:
                self._insert(value, node.right)
                
    def search(self, value):
        return self._search(value, self.root)
    
    def _search(self, value, node):
        if node is None:
            return False
        elif node.value == value:
            return True
        elif value < node.value:
            return self._search(value, node.left)
        else:
            return self._search(value, node.right)
        
    def delete(self, value):
        self.root = self._delete(value, self.root)
        
    def _delete(self, value, node):
        if node is None:
            return None
        elif value < node.value:
            node.left = self._delete(value, node.left)
        elif value > node.value:
            node.right = self._delete(value, node.right)
        else:
            if node.left is None:
                return node.right
            elif node.right is None:
                return node.left
            else:
                min_node = self._find_min(node.right)
                node.value = min_node.value
                node.right = self._delete(min_node.value, node.right)
        return node
                
    def _find_min(self, node):
        while node.left is not None:
            node = node.left
        return node

    def inorder(self):
        return self._inorder(self.root)

    def _inorder(self, node):
        if node is None:
            return []
        return self._inorder(node.left) + [node.value] + self._inorder(node.right)
        
    def is_empty(self):
        return self.root is None


@pytest.fixture
def bst():
    # 构建一棵空树
    bst = BSTImpl()

    # 测试插入操作
    bst.insert(3)
    bst.insert(1)
    bst.insert(5)
    bst.insert(4)
    bst.insert(6)
    yield bst
    print("teardown bstree")

def test_search_bstree(bst):
    # 测试查找操作
    assert bst.search(3) is True
    assert bst.search(2) is False


def test_delete_bstree(bst):
    # 测试删除操作
    bst.delete(3)
    assert bst.search(3) is False
    assert bst.search(4) is True


def test_traverse_operation_bstree(bst):
    # 测试遍历操作
    assert bst.inorder() == [1, 3, 4, 5, 6] 

def test_delete_bstree(bst):
    # 测试删除所有节点后的情况
    bst.delete(1)
    bst.delete(3)
    bst.delete(4)
    bst.delete(5)
    bst.delete(6)
    assert bst.is_empty() is True
