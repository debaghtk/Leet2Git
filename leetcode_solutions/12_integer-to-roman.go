# Integer to Roman
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/integer-to-roman/

func intToRoman(num int) string {
   answer:=""
   for num>0{
    if (int(num/1000)>0){
        answer=answer+"M"
        num=num-1000
        continue
    }
    if (int(num/900)>0){
        answer=answer+"CM"
        num=num-900
        continue
    }
    if (int(num/500)>0){
        answer=answer+"D"
        num=num-500
        continue
    }
    if (int(num/400)>0){
        answer=answer+"CD"
        num=num-400
        continue
    }
    if (int(num/100)>0){
        answer=answer+"C"
        num=num-100
        continue
    }
    if (int(num/90)>0){
        answer=answer+"XC"
        num=num-90
        continue
    }
    if (int(num/50)>0){
        answer=answer+"L"
        num=num-50
        continue
    }
    if (int(num/40)>0){
        answer=answer+"XL"
        num=num-40
        continue
    }
    if (int(num/10)>0){
        answer=answer+"X"
        num=num-10
        continue
    }
    if (int(num/9)>0){
        answer=answer+"IX"
        num=num-9
        continue
    }
    if (int(num/5)>0){
        answer=answer+"V"
        num=num-5
        continue
    }
    if (int(num/4)>0){
        answer=answer+"IV"
        num=num-4
        continue
    }
    if (int(num/1)>0){
        answer=answer+"I"
        num=num-1
        continue
    }
   }
   return answer
}