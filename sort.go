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
	slices.SortFunc(data,
		func(a, b Company) int {
			return cmp.Compare(a.Company, b.Company)
		})
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
