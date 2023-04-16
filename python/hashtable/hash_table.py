# -*- coding: utf-8 -*-
# @Author: Noaghzil
# @Date:   2023-04-15 15:58:07
# @Description: 散列表

import abc

class HashTableBase(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def insert(self, key, value):
        pass

    @abc.abstractmethod
    def find(self, key):
        pass

    @abc.abstractmethod
    def delete(self, key):
        pass
    
class HashTable(HashTableBase):
    def __init__(self, size):
        self.size = size
        self.table = [[] for _ in range(self.size)]
    
    def hash_func(self, key):
        return hash(key) % self.size
    
    def insert(self, key, value):
        hash_value = self.hash_func(key)
        self.table[hash_value].append((key, value))
    
    def find(self, key):
        hash_value = self.hash_func(key)
        for k, v in self.table[hash_value]:
            if k == key:
                return v
        return None
    
    def delete(self, key):
        hash_value = self.hash_func(key)
        for i, (k, v) in enumerate(self.table[hash_value]):
            # print(i, k, v)
            # import pdb; pdb.set_trace()
            if k == key:
                del self.table[hash_value][i]
                return True
        return False


if __name__ == "__main__":
    hash_table = HashTable(10)
    for i in range(100):
        hash_table.insert("key" + str(i), "value" + str(i))
    print(hash_table.find("key1"))
    print(hash_table.delete("key8"))
    print(hash_table.delete("key108"))