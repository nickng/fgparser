package parser

import "fmt"

// pos is a single character Position.
type pos struct {
	line   int
	column int
}

func (p pos) Line() int   { return p.line }
func (p pos) Column() int { return p.column }
func (p pos) String() string {
	return fmt.Sprintf("%d:%d", p.Line(), p.Column())
}

// posRange is a character range Position.
type posRange struct {
	line        int
	columnStart int
	columnEnd   int
}

func (p posRange) Line() int   { return p.line }
func (p posRange) Column() int { return p.columnStart } // start column
func (p posRange) String() string {
	return fmt.Sprintf("%d:%d-%d", p.Line(), p.columnStart, p.columnEnd)
}
