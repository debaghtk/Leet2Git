# Employee Importance
# Difficulty: Medium
# Language: golang
# Link: https://leetcode.com/problems/employee-importance/

/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
    candidate:= Employee{}
    for _,v := range employees{
        if v.Id == id {
            candidate = *v
            break
        }
    }
    
    subImportance:=0
    
    for _,v:=range candidate.Subordinates{
        subImportance += getImportance(employees, v)
    }
    
    return candidate.Importance + subImportance
}