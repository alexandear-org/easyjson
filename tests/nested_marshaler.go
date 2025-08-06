package tests

import (
	"github.com/alexandear-org/easyjson"
	"github.com/alexandear-org/easyjson/jlexer"
	"github.com/alexandear-org/easyjson/jwriter"
)

//easyjson:json
type NestedMarshaler struct {
	Value easyjson.MarshalerUnmarshaler
	Value2 int
}

type StructWithMarshaler struct {
	Value int
}

func (s *StructWithMarshaler) UnmarshalEasyJSON(w *jlexer.Lexer) {
	s.Value = w.Int()
}

func (s *StructWithMarshaler) MarshalEasyJSON(w *jwriter.Writer) {
	w.Int(s.Value)
}
