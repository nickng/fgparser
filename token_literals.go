package parser

import "fmt"

type TypeLit struct {
	pos posRange
}

func (t TypeLit) Pos() Position  { return t.pos }
func (t TypeLit) String() string { return "Type" }

type StructLit struct {
	pos posRange
}

func (t StructLit) Pos() Position  { return t.pos }
func (t StructLit) String() string { return "Struct" }

type FuncLit struct {
	pos posRange
}

func (t FuncLit) Pos() Position  { return t.pos }
func (t FuncLit) String() string { return "Func" }

type InterfaceLit struct {
	pos posRange
}

func (t InterfaceLit) Pos() Position  { return t.pos }
func (t InterfaceLit) String() string { return "Interface" }

type PackageLit struct {
	pos posRange
}

func (t PackageLit) Pos() Position  { return t.pos }
func (t PackageLit) String() string { return "Package" }

type ReturnLit struct {
	pos posRange
}

func (t ReturnLit) Pos() Position  { return t.pos }
func (t ReturnLit) String() string { return "Return" }

type Ident struct {
	str string
	pos posRange
}

func (t Ident) Pos() Position  { return t.pos }
func (t Ident) String() string { return fmt.Sprintf("Ident %s", t.str) }
