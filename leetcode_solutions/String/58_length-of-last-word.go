# Length of Last Word
# Difficulty: Easy
# Language: golang
# Topic: String
# Tags: String
# Link: https://leetcode.com/problems/length-of-last-word/

func lengthOfLastWord(s string) int {
	answer := 0
	counter := 0
	for i, v := range s {
		if v != ' ' {
			counter++
			continue
		}
		if v == ' ' && i == 0 {
			continue
		}
		if v == ' ' && s[i-1] != ' ' {
			answer = counter
			counter = 0
			continue
		}
		if v == ' ' {
			continue
		}
	}

    if counter!=0{
        return counter
    }
	return answer
}