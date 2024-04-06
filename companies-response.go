package main

type CompaniesResponse struct {
	Data                 []Company
	CurrentSortField     string
	CurrentSortDirection string
	PageIndex            int
	CompanyTableHeader   TableHeader
	ContactTableHeader   TableHeader
	CountryTableHeader   TableHeader
}
