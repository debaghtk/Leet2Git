# Reverse Bits
# Difficulty: Easy
# Language: golang
# Link: https://leetcode.com/problems/reverse-bits/

func reverseBits(num uint32) uint32 {
    ans:=[]int{}
    for num>0 {
        ans = append(ans, int(num%2))
        num=num/2
    }
    result:=0
    for i,v:=range ans{
        if v==0{
            continue
        }
        if v==1{
            result += int(math.Pow(2, float64(31-i)))
        }
    }
    return uint32(result)
}