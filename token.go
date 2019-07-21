package parser

// Position is a location in a source file.
type Position interface {
	Line() int
	Column() int
	String() string
}

// Token is a generic token
type Token interface {
	String() string
	Pos() Position
}
