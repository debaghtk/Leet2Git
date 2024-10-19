# Map Sum Pairs
# Difficulty: Medium
# Language: golang
# Topic: Trie
# Tags: Hash Table, String, Design, Trie
# Link: https://leetcode.com/problems/map-sum-pairs/

type MapSum struct {
    val int
    character rune
    children []*MapSum
}


func Constructor() MapSum {
    return MapSum{}
}


func (this *MapSum) Insert(key string, val int)  {
        for _,v:= range key {
        if len(this.children) == 0 {
            newTrie:=Constructor()
            newTrie.character=v
            this.children = append(this.children, &newTrie)
            this = this.children[0]
            continue
        }
        matchfound:=false
        for k,c:= range this.children{
            if c.character==v{
                this = this.children[k]
                matchfound=true
                break
            }
        }
        if matchfound {
            continue
        }
        newTrie:=Constructor()
        newTrie.character=v
        this.children = append(this.children, &newTrie)
        this = this.children[len(this.children)-1] 
    }
    this.val=val
}


func (this *MapSum) Sum(prefix string) int {
    result:=0    
    for _,v:= range prefix{
        matchfound:=false
        for k,c:= range this.children{
            if c.character == v{
                this=this.children[k]
                matchfound=true
                break
            }
        }
        if matchfound {
            continue
        }
        return 0
    }
    
    result = result + this.val
    if len(this.children)!=0{
        for _,c:= range this.children{
            result = result + computeSum(c)
        }
    }
    return result
}

func computeSum (this *MapSum) int{
    result := 0
    result = result + this.val
    if len(this.children)!=0{
        for _,c:= range this.children{
            result = result + computeSum(c)
        }
    }
    return result
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */