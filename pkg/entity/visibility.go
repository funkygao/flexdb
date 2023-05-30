package entity

import (
	"errors"
	"strconv"
)

type Visibility int16

func (v Visibility) String() string {
	return strconv.Itoa(int(v))
}

const (
	unknownVisibility = Visibility(0)

	PublicVisibility  = Visibility(1)
	PrivateVisibility = Visibility(2)
)

var (
	ErrInvalidVisibility = errors.New("invalid visibility")
)

func ToVisibility(v string) (Visibility, error) {
	visibility, err := strconv.Atoi(v)
	if err != nil {
		return unknownVisibility, err
	}

	vis := Visibility(visibility)
	if vis != PublicVisibility && vis != PrivateVisibility {
		return unknownVisibility, ErrInvalidVisibility
	}

	return vis, nil
}
