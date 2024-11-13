package handler

import "rest/domain"

type baseService interface {
	Info() (domain.Base, error)
}

type userService interface{}
