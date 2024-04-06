package main

import (
	"cmp"
	"net/url"
	"slices"
	"strconv"
)

func companiesResponse(queryValues *url.Values) *CompaniesResponse {

	slices.SortFunc(data, getSortByName(queryValues))
	if queryValues.Get("direction") == "descending" {
		slices.Reverse(data)
	}
	skip := "0"
	if queryValues.Has("skip") {
		skip = queryValues.Get("skip")
	}
	pageIndex, err := strconv.Atoi(skip)
	if err != nil {
		panic(err)
	}
	return &CompaniesResponse{
		Data:                 data,
		CurrentSortField:     queryValues.Get("sort"),
		CurrentSortDirection: queryValues.Get("direction"),
		PageIndex:            pageIndex,
		CompanyTableHeader:   *getHeader(queryValues, "Company"),
		CountryTableHeader:   *getHeader(queryValues, "Country"),
		ContactTableHeader:   *getHeader(queryValues, "Contact"),
	}
}

func getHeader(queryValues *url.Values, label string) *TableHeader {
	currentSortDirection := queryValues.Get("direction")
	selected := queryValues.Get("sort") == label
	response := TableHeader{
		Label:             label,
		NextSortDirection: getNextDirectionIfSelected(selected, checkForValue(currentSortDirection)),
		Selected:          selected,
		ArrowUpClass:      getArrowClassIfSelected(selected, checkForValue(currentSortDirection), true),
		ArrowDownClass:    getArrowClassIfSelected(selected, checkForValue(currentSortDirection), false),
	}
	return &response
}

func getSortMap() map[bool]string {

	return map[bool]string{
		true:  "ascending",
		false: "descending",
	}
}

func getNextDirectionIfSelected(selected bool, currentDirection bool) string {
	sortMap := getSortMap()
	if selected {
		return sortMap[!currentDirection]
	}
	return sortMap[true]
}

func getArrowClassIfSelected(selected bool, currentDirection bool, up bool) string {
	if selected && currentDirection == up {
		return "icon"
	}
	return "hide"
}

// function to check if a value present in the map
func checkForValue(sortValue string) bool {

	//traverse through the map
	for key, value := range getSortMap() {

		//check if present value is equals to userValue
		if value == sortValue {

			//if same return true
			return key
		}
	}

	//if value not found return false
	return false
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
