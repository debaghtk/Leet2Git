# Reverse String
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/reverse-string/

func reverseString(s []byte)  {
    for i:=0;i<len(s)-1-i;i++ { 
        s[i], s[len(s)-1-i]=s[len(s)-1-i],s[i]
    }
}