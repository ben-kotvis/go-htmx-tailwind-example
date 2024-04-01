package main

import (
	"cmp"
	"net/url"
	"slices"
)

func sort(queryValues *url.Values) *SortReponse {
	nextSortDirection := "descending"
	arrowUpClass := "hide"
	arrowDownClass := "hide"
	slices.SortFunc(data, getSortByName(queryValues))
	if queryValues.Get("direction") == "descending" {
		slices.Reverse(data)
		nextSortDirection = "ascending"
		arrowDownClass = "icon"
	} else {
		arrowUpClass = "icon"
	}
	response := SortReponse{
		Data:              data,
		NextSortDirection: nextSortDirection,
		ArrowUpClass:      arrowUpClass,
		ArrowDownClass:    arrowDownClass,
	}
	return &response
}

func getSortByName(queryValues *url.Values) func(a Company, b Company) int {
	switch sortName := queryValues.Get("sort"); sortName {
	default:
	case "Company":
		return func(a Company, b Company) int {
			return cmp.Compare(a.Company, b.Company)
		}
	case "Contact":
		return func(a Company, b Company) int {
			return cmp.Compare(a.Contact, b.Contact)
		}
	case "Country":
		return func(a Company, b Company) int {
			return cmp.Compare(a.Country, b.Country)
		}
	}
	return func(a Company, b Company) int {
		return cmp.Compare(a.Company, b.Company)
	}
}
