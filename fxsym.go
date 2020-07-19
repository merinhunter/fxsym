package fxsym

import (
	"errors"
	"fmt"
	"strings"
)

const (
	SNone = iota
	SProg
	SFunc
	SType
	SVar
	SConst
)

type sType int

type Env map[string]*Sym

type StkEnv []Env

func (envs *StkEnv) PushEnv() {
	env := Env{}
	*envs = append(*envs, env)
}

func (envs *StkEnv) PopEnv() {
	eS := *envs
	if len(eS) == 1 {
		panic("cannot pop builtin")
	}
	*envs = eS[:len(eS)-1]
}

func (envs *StkEnv) NewSym(name string, sType int) (s *Sym, err error) {
	eS := *envs
	e := eS[len(eS)-1]

	for i := len(eS) - 1; i >= 0; i-- {
		if _, ok := eS[i][name]; ok {
			return nil, errors.New("already declared sym")
		}
	}

	s = &Sym{name: name, sType: sType}
	e[name] = s

	return s, nil
}

func (envs *StkEnv) GetSym(name string) (s *Sym) {
	eS := *envs
	for i := len(eS) - 1; i >= 0; i-- {
		if s, ok := eS[i][name]; ok {
			return s
		}
	}
	return nil
}

type Sym struct {
	name  string
	sType int
	tType *Type

	tokKind int
	depth   int

	file string
	line int

	symContent interface{}
}

func (s *Sym) Name() string {
	return s.name
}

func (s *Sym) SymType() string {
	return sType(s.sType).String()
}

func (s *Sym) SetType(tp int) {
	s.tType = &Type{tp}
}

func (s *Sym) SetDepth(depth int) {
	s.depth = depth
}

func (s *Sym) AddSymType(sType int) {
	s.sType = sType
}

func (s *Sym) AddTokKind(tokKind int) {
	s.tokKind = tokKind
}

func (s *Sym) AddPlace(file string, line int) {
	s.file = file
	s.line = line
}

func (s *Sym) AddContent(content interface{}) {
	s.symContent = content
}

func (s *Sym) Content() interface{} {
	return s.symContent
}

func (s *Sym) String() string {
	if s == nil {
		return "nil"
	}

	tabs := strings.Repeat("\t", s.depth)
	return fmt.Sprintf("%s%p SYM[%s](%s)", tabs, s, s.SymType(), s.name)
}

type Type struct {
	id int
}
