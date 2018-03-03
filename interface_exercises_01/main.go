package main

import (
	"fmt"
	"sort"
)

/* type People []string
type ByName []string
func (b ByName) Len() int {return len(b)}
func (b ByName) Swap(i, j int) {b[i],b[j] = b[j],b[i] }
func (b ByName) Less(i,j int) bool {return b[i] < b[j]}
*/

type People struct {
	Name string
	Age  int
}

func PrintNameAge(p []People) {
	for i := 0; i < len(p); i++ {
		fmt.Printf("Name: %s, Age: %d\n", p[i].Name, p[i].Age)
	}
}

type SortByName []People

func (s SortByName) Len() int           { return len(s) }
func (s SortByName) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortByName) Less(i, j int) bool { return s[i].Name < s[j].Name }

type SortByAge []People

func (s SortByAge) Len() int           { return len(s) }
func (s SortByAge) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortByAge) Less(i, j int) bool { return s[i].Age < s[j].Age }

func main() {
	//studyGroup := People{"Zeno","John","Al","Jenny"}

	studyGroup := []People{
		{"Zeno", 34},
		{"John", 4},
		{"Al", 84},
		{"Jenny", 32},
	}

	sort.Sort(SortByAge(studyGroup))
	//fmt.Println(studyGroup)
	PrintNameAge(studyGroup)
	// Use this if studyGroup is a slice of strings
	//sort.Sort(sort.Reverse(studyGroup))

	// Use this if studyGroup is a slice of structs
	//sort.Reverse(SortByAge(studyGroup))

	sort.Sort(SortByName(studyGroup))
	//fmt.Println(studyGroup)
	PrintNameAge(studyGroup)
}
