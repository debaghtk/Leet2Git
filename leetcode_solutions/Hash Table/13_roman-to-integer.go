# Roman to Integer
# Difficulty: Easy
# Language: golang
# Topic: Hash Table
# Tags: Hash Table, Math, String
# Link: https://leetcode.com/problems/roman-to-integer/

func romanToInt(s string) int {
    romans := map[string]int{}
    romans["I"]=1
    romans["V"]=5
    romans["X"]=10
    romans["L"]=50
    romans["C"]=100
    romans["D"]=500
    romans["M"]=1000

    romans["IV"]=-1
    romans["IX"]=-1
    romans["XL"]=-10
    romans["XC"]=-10
    romans["CD"]=-100
    romans["CM"]=-100
    answer:=0
    length:=len(s)
    for i:=0;i<len(s);i++{
        if i+1<length   {
            c,ok:=romans[string(s[i])+string(s[i+1])]
            if ok {
                answer=answer+c
                continue
            }
        }
        answer=answer+romans[string(s[i])]
    }
    return answer
}