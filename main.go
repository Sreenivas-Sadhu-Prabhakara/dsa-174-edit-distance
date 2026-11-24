package main

import "fmt"

// EditDistance returns the minimum number of single-character insertions,
// deletions and substitutions turning a into b (Levenshtein), with two
// rolling rows.
func EditDistance(a, b string) int {
	prev := make([]int, len(b)+1)
	cur := make([]int, len(b)+1)
	for j := range prev {
		prev[j] = j // empty a: j insertions build b[:j]
	}
	for i := 1; i <= len(a); i++ {
		cur[0] = i // empty b: i deletions erase a[:i]
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				cur[j] = prev[j-1]
				continue
			}
			best := prev[j-1] // substitute
			if prev[j] < best {
				best = prev[j] // delete from a
			}
			if cur[j-1] < best {
				best = cur[j-1] // insert into a
			}
			cur[j] = best + 1
		}
		prev, cur = cur, prev
	}
	return prev[len(b)]
}

func main() {
	fmt.Println(EditDistance("horse", "ros"))
}
