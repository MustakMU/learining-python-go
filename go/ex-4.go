package main

import (
	"sort"
)

/**
Write a  program to sort a list of tuples using Lambda.
Original list of tuples:
[('English', 88), ('Science', 90), ('Maths', 97), ('Social sciences', 82)]
Sorting the List of Tuples:
[('Social sciences', 82), ('English', 88), ('Science', 90), ('Maths', 97)]

**/
type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	marks := map[string]int{
		"English":         99,
		"Science":         90,
		"Maths":           97,
		"Social sciences": 82,
	}

	p := make(PairList, len(marks))

	i := 0
	for k, v := range marks {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)
	//p is sorted

	for _, k := range p {
		println(k.Key, k.Value)
	}

}
