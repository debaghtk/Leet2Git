# Design HashMap
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/design-hashmap/

type MyHashMap struct {
    Collection []Kvstore
}

type Kvstore struct{
    Key int
    Value int
}


func Constructor() MyHashMap {
    return MyHashMap{}
}


func (this *MyHashMap) Put(key int, value int)  {
    for i,coll:=range this.Collection{
        if coll.Key == key{
            this.Collection[i].Value = value
            return
        }
    }
    this.Collection=append(this.Collection, Kvstore{Key:key, Value:value})
}


func (this *MyHashMap) Get(key int) int {
    for _,coll:=range this.Collection{
        if coll.Key == key{
            return coll.Value
        }
    }
    return -1
}


func (this *MyHashMap) Remove(key int)  {
    copyColl:=this.Collection
    this.Collection=[]Kvstore{}
    for _,coll:=range copyColl{
        if coll.Key == key{
            continue
        }
        this.Collection=append(this.Collection, Kvstore{Key:coll.Key, Value: coll.Value})
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */