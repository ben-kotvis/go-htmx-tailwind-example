package main

type CompaniesResponse struct {
	Data               []Company
	CompanyTableHeader TableHeader
	ContactTableHeader TableHeader
	CountryTableHeader TableHeader
}
