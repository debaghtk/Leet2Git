# Minimum Genetic Mutation
# Difficulty: Medium
# Language: golang
# Topic: Hash Table
# Tags: Hash Table, String, Breadth-First Search
# Link: https://leetcode.com/problems/minimum-genetic-mutation/

import "slices"

func minMutation(startGene string, endGene string, bank []string) int {
	if slices.Contains(bank, endGene) == false {
		return -1
	}
	distance := map[string]int{}
	for _, v := range bank {
		distance[v] = -1
	}
    distance[startGene]=0
	queue := []string{}
	queue = append(queue, startGene)

	for len(queue) > 0 {
		gene := queue[0]
		queue = queue[1:]

		if gene == endGene {
			// do something
            return distance[gene]
		}

		for _, v := range bank {
			if distanceOf1(gene, v) && distance[v] == -1 {
				queue = append(queue, v)
				distance[v] = distance[gene] + 1
			}
		}

	}

	return -1
}

func distanceOf1(gene, gene2 string) bool {
	diff := 0
	for i, _ := range gene {
		if gene2[i] != gene[i] {
			diff++
		}
	}
	if diff == 1 {
		return true
	}
	return false
}