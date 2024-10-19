# Design HashSet
# Difficulty: Easy
# Language: golang
# Topic: Linked List
# Tags: Array, Hash Table, Linked List, Design, Hash Function
# Link: https://leetcode.com/problems/design-hashset/

type MyHashSet struct {
    Set [1001][1001]int
}


func Constructor() MyHashSet {
    obj:=MyHashSet{}
    obj.Set[0][0]=-1
    return obj
}


func (this *MyHashSet) Add(key int)  {
    i:=key%1001
    j:=key/1001
    this.Set[i][j] = key
}


func (this *MyHashSet) Remove(key int)  {
    i:=key%1001
    j:=key/1001
    this.Set[i][j] = -1
}


func (this *MyHashSet) Contains(key int) bool {
    i:=key%1001
    j:=key/1001
    if this.Set[i][j] == key{
        return true
    }
    return false
}


/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */