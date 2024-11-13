package domain

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidArgument = errors.New("invalid argument")
	ErrServerInternal  = errors.New("internal error")
)

type Kind uint8

const (
	KindUnspecified Kind = iota
	// list all kinds of error might occur
)

func (t Kind) String() string {
	switch t {
	// case Kind...:
	// return "mycustomKindOfEntity"
	}
	return "entity"
}

type ErrNotFound struct {
	kind Kind
}

func NewErrNotFound(kind Kind) error {
	return &ErrNotFound{
		kind: kind,
	}
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("%s not found", e.kind)
}
