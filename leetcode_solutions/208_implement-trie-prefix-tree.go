# Implement Trie (Prefix Tree)
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/implement-trie-prefix-tree/

type Trie struct {
    character rune
    children []*Trie
    terminal bool
}


func Constructor() Trie {
    return Trie{}
}


func (this *Trie) Insert(word string)  {
    //length := len(word)
    for _,v:= range word {
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
    this.terminal=true
}


func (this *Trie) Search(word string) bool {
    for _,v:= range word{
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
        return false
    }
    return this.terminal
}


func (this *Trie) StartsWith(prefix string) bool {
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
        return false
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */