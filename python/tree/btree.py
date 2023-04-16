# -*- coding: utf-8 -*-
# @Author: Noaghzil
# @Date:   2023-04-16 12:40:49
# @Description: B树

from abc import ABCMeta, abstractmethod
import pytest


class BTreeNode(metaclass=ABCMeta):
    @abstractmethod
    def is_full(self):
        pass

    @abstractmethod
    def insert(self, key):
        pass

    @abstractmethod
    def search(self, key):
        pass


class BTree(BTreeNode):
    def __init__(self, order=3):
        self.order = order
        self.root = BTree.Node(order)

    def is_full(self):
        return self.root.is_full()

    def insert(self, key):
        if self.is_full():
            new_root = BTree.Node(self.order)
            new_root.children.append(self.root)
            new_root.split_child(0)
            self.root = new_root
        self.root.insert(key)

    def search(self, key):
        return self.root.search(key)

    class Node(BTreeNode):
        def __init__(self, order):
            self.order = order
            self.keys = []
            self.children = []

        def is_full(self):
            return len(self.keys) == 2 * self.order - 1

        def insert(self, key):
            if self.is_leaf():
                i = len(self.keys) - 1
                while i >= 0 and key < self.keys[i]:
                    self.keys[i + 1] = self.keys[i]
                    i -= 1
                self.keys[i + 1] = key
            else:
                i = len(self.keys) - 1
                while i >= 0 and key < self.keys[i]:
                    i -= 1
                i += 1
                if self.children[i].is_full():
                    self.split_child(i)
                    if key > self.keys[i]:
                        i += 1
                self.children[i].insert(key)

        def split_child(self, i):
            z = BTree.Node(self.order)
            y = self.children[i]
            z.keys = y.keys[self.order:]
            self.keys.insert(i, y.keys[self.order - 1])
            self.children.insert(i + 1, z)
            y.keys = y.keys[:self.order - 1]
            z.children = y.children[self.order:]
            y.children = y.children[:self.order]

        def is_leaf(self):
            return len(self.children) == 0

        def search(self, key):
            i = 0
            while i < len(self.keys) and key > self.keys[i]:
                i += 1
            if i < len(self.keys) and key == self.keys[i]:
                return True
            elif self.is_leaf():
                return False
            else:
                return self.children[i].search(key)


def test_b_tree():
    b_tree = BTree(order=3)
    b_tree.insert(4)
    b_tree.insert(2)
    b_tree.insert(1)
    b_tree.insert(5)
    b_tree.insert(3)
    assert b_tree.search(4) is True
    assert b_tree.search(6) is False
    assert b_tree.search(3) is True
    assert b_tree.search(2) is True
    assert b_tree.search(1) is True
    assert b_tree.search(5) is True


if __name__ == '__main__':
    pytest.main()
