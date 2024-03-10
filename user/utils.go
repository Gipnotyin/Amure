package user

import (
	"fmt"
	"strings"
)

const (
	Login      = "login"
	Name       = "name"
	LastName   = "last_name"
	Email      = "email"
	Phone      = "phone"
	SortByDesc = "DESC"
	SortByAsc  = "ASC"
)

func IsCorrectSort(sort string) error {
	sort = strings.ToUpper(sort)

	switch sort {
	case SortByAsc, SortByDesc:
		return nil
	default:
		return fmt.Errorf("correct sort are ASC or DESC")
	}
}

func IsCorrectField(field string) error {
	field = strings.ToLower(field)

	switch field {
	case Login, Email, Phone, LastName, Name:
		return nil
	default:
		return fmt.Errorf("this is no correct field")
	}
}
