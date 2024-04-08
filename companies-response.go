package main

type CompaniesResponse struct {
	Data                 []Company
	CurrentSortField     string
	CurrentSortDirection string
	CurrentSkip          int
	PreviousSkip         int
	NextSkip             int
	CompanyTableHeader   TableHeader
	ContactTableHeader   TableHeader
	CountryTableHeader   TableHeader
}
