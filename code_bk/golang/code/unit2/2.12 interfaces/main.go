package main

import (
	"fmt"
	"strings"
)

type Printer interface {
	Print() string
}

type PDF struct {
	Content string
}

func (p PDF) Print() string {
	return p.Content
}

type Excel struct {
	Cells []string
}

// []string{"A", "B"} => A,B
func (e Excel) Print() string {
	return strings.Join(e.Cells, ",")
}

func main() {
	docs := []Printer{
		PDF{Content: "Hello PDF"},
		Excel{Cells: []string{"A", "B"}},
	}

	for _, doc := range docs {
		fmt.Println(doc.Print())
	}

	// interface{}
	docs2 := []interface{}{
		PDF{Content: "Hello PDF"},
		Excel{Cells: []string{"A", "B"}},
	}

	fmt.Println(docs2)

	// interface{} => any
	docs3 := []any{
		PDF{Content: "Hello PDF"},
		Excel{Cells: []string{"A", "B"}},
	}

	fmt.Println(docs3)
}
