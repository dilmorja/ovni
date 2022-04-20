// Copyright (c) 2022, Daniel M. Jaén
// All rights reserved.

package tagger

import (
	"strings"
	"fmt"
)

type Field struct {
	Name string
	Content string
}

func (f *Field) String() string {
	return fmt.Sprintf(`%s="%s"`, f.Name, f.Content)
}

type Tag interface {
	Name string
	Fields []*Field
	Content string
}

func (t *Tag) AddField(name, content string) {
	t.Fields = append(t.Fields, &Field{Name: name, Content: content})
}

func (t *Tag) String() string {
	var fields []string
	for _, f := range t.Fields {
		fields = append(fields, f.String())
	}
	return fmt.Sprintf(`<%s %s>%s</%s>`, t.Name, strings.Join(fields, " "), t.Content, t.Name)
}

type Tagger struct {
	Tags []*Tag
}

func New() *Tagger {
	return &Tagger{
		Tags: make([]*Tag, 0),
	}
}
