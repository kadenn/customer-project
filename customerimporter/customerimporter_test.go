package customerimporter

import (
	"strings"
	"testing"
)

func TestGetSortedEmailDomainsFromNoCsv(t *testing.T) {
	csvPath := ""

	_, err := GetSortedEmailDomainsFromCsv(csvPath)
	if !strings.Contains(err.Error(), "no such file or directory") {
		t.Errorf("wrong_csv_path GetSortedEmailDomainsFromCsv | got: %s, expected: no such file or directory", err.Error())
	}
}

func TestGetSortedEmailDomainsFromBrokenCustomersCsv(t *testing.T) {
	csvPath := "../data/broken_customers.csv"

	_, err := GetSortedEmailDomainsFromCsv(csvPath)
	if !strings.Contains(err.Error(), "wrong number of fields") {
		t.Errorf("broken_customers.csv GetSortedEmailDomainsFromCsv | got: %s, expected: wrong number of fields", err.Error())
	}
}

func TestGetSortedEmailDomainsFromEmptyCustomersCsv(t *testing.T) {
	csvPath := "../data/empty_customers.csv"

	sortedEmailDomains, err := GetSortedEmailDomainsFromCsv(csvPath)
	if err != nil {
		t.Errorf("empty_customers.csv GetSortedEmailDomainsFromCsv | error: %v", err)
	}

	sortedEmailDomainsLenght := len(sortedEmailDomains)
	if sortedEmailDomainsLenght != 0 {
		t.Errorf("empty_customers.csv sortedEmailDomainsLenght | got: %d, expected: 0", sortedEmailDomainsLenght)
	}
}

func TestGetSortedEmailDomainsFromCustomersCsv(t *testing.T) {
	csvPath := "../data/customers.csv"

	sortedEmailDomains, err := GetSortedEmailDomainsFromCsv(csvPath)
	if err != nil {
		t.Errorf("customers.csv GetSortedEmailDomainsFromCsv | error: %v", err)
	}

	sortedEmailDomainsLenght := len(sortedEmailDomains)
	if sortedEmailDomainsLenght != 500 {
		t.Errorf("customers.csv sortedEmailDomainsLenght | got: %d, expected: 500", sortedEmailDomainsLenght)
	}

	firstSortedEmailDomainName := sortedEmailDomains[0].Name
	if firstSortedEmailDomainName != "123-reg.co.uk" {
		t.Errorf("customers.csv firstSortedEmailDomainName | got: %s, expected: 123-reg.co.uk", firstSortedEmailDomainName)
	}

	firstSortedEmailDomainNumberOfCustomers := sortedEmailDomains[0].NumberOfCustomers
	if firstSortedEmailDomainNumberOfCustomers != 8 {
		t.Errorf("customers.csv firstSortedEmailDomainNumberOfCustomers | got: %d, expected: 8", firstSortedEmailDomainNumberOfCustomers)
	}

	secondSortedEmailDomainName := sortedEmailDomains[1].Name
	if firstSortedEmailDomainName > secondSortedEmailDomainName {
		t.Errorf("customers.csv sortedEmailDomainsSortingComparison | email domains are not sorted properly")
	}
}

func TestGetSortedEmailDomainsFromMoreCustomersCsv(t *testing.T) {
	csvPath := "../data/more_customers.csv"

	sortedEmailDomains, err := GetSortedEmailDomainsFromCsv(csvPath)
	if err != nil {
		t.Errorf("more_customers.csv GetSortedEmailDomainsFromCsv | error: %v", err)
	}

	sortedEmailDomainsLenght := len(sortedEmailDomains)
	if sortedEmailDomainsLenght <= 500 {
		t.Errorf("more_customers.csv sortedEmailDomainsLenght | got: %d, expected: >500", sortedEmailDomainsLenght)
	}
}

func TestReadRowsFromNoCsv(t *testing.T) {
	csvPath := ""

	_, err := readRowsFromCsv(csvPath)
	if !strings.Contains(err.Error(), "no such file or directory") {
		t.Errorf("wrong_csv_path readRowsFromCsv | got: %s, expected: no such file or directory", err.Error())
	}
}

func TestReadRowsFromBrokenCustomersCsv(t *testing.T) {
	csvPath := "../data/broken_customers.csv"

	_, err := readRowsFromCsv(csvPath)
	if !strings.Contains(err.Error(), "wrong number of fields") {
		t.Errorf("broken_customers.csv readRowsFromCsv | got: %s, expected: wrong number of fields", err.Error())
	}
}

func TestReadRowsFromEmptyCustomerCsv(t *testing.T) {
	csvPath := "../data/empty_customers.csv"

	rows, err := readRowsFromCsv(csvPath)
	if err != nil {
		t.Errorf("empty_customers.csv readRowsFromCsv | error: %v", err)
	}

	rowsLenght := len(rows)
	if rowsLenght != 1 {
		t.Errorf("empty_customers.csv rowsLenght | got: %d, expected: 1", rowsLenght)
	}
}

func TestReadRowsFromCustomersCsv(t *testing.T) {
	csvPath := "../data/customers.csv"

	rows, err := readRowsFromCsv(csvPath)
	if err != nil {
		t.Errorf("customers.csv readRowsFromCsv | error: %v", err)
	}

	rowsLenght := len(rows)
	if rowsLenght != 3003 {
		t.Errorf("customers.csv rowsLenght | got: %d, expected: 3003", rowsLenght)
	}

	firstRowFirstName := rows[1][0]
	if firstRowFirstName != "Mildred" {
		t.Errorf("customers.csv firstRowFirstName | got: %s, expected: Mildred", firstRowFirstName)
	}

	firstRowLastName := rows[1][1]
	if firstRowLastName != "Hernandez" {
		t.Errorf("customers.csv firstRowLastName | got: %s, expected: Hernandez", firstRowLastName)
	}

	firstRowEmail := rows[1][2]
	if firstRowEmail != "mhernandez0@github.io" {
		t.Errorf("customers.csv firstRowEmail | got: %s, expected: mhernandez0@github.io", firstRowEmail)
	}

	firstRowGender := rows[1][3]
	if firstRowGender != "Female" {
		t.Errorf("customers.csv firstRowGender | got: %s, expected: Female", firstRowGender)
	}

	firstRowIpAddress := rows[1][4]
	if firstRowIpAddress != "38.194.51.128" {
		t.Errorf("customers.csv firstRowIpAddress | got: %s, expected: 38.194.51.128", firstRowIpAddress)
	}
}

func TestReadRowsFromMoreCustomersCsv(t *testing.T) {
	csvPath := "../data/more_customers.csv"

	rows, err := readRowsFromCsv(csvPath)
	if err != nil {
		t.Errorf("more_customers.csv readRowsFromCsv | error: %v", err)
	}

	rowsLenght := len(rows)
	if rowsLenght < 500 {
		t.Errorf("more_customers.csv rowsLenght | got: %d, expected: >500", rowsLenght)
	}
}

func TestGetEmailDomains(t *testing.T) {
	mockRows := [][]string{
		{"first_name", "last_name", "email", "gender", "ip_address"},
		{"Mildred", "Hernandez", "mhernandez0@github.io", "Female", "38.194.51.128"},
		{"Bonnie", "Ortiz", "bortiz1@cyberchimps.com", "Female", "197.54.209.129"},
		{"Dennis", "Henry", "dhenry2@hubpages.com", "Male", "155.75.186.217"},
		{"Justin", "Hansen", "jhansen3@360.cn", "Male", "251.166.224.119"},
		{"Carlos", "Garcia", "cgarcia4@statcounter.com", "Male", "57.171.52.110"},
		{"Ernest", "Reid", "ereid5@rediff.com", "Male", "243.219.170.46"},
		{"Gary", "Henderson", "ghenderson6@acquirethisname.com", "Male", "30.97.220.14"},
		{"Dennis", "Henderson", "dhenderson7@chicagotribune.com", "Male", "27.122.100.11"},
		{"Norma", "Allen", "nallen8@cnet.com", "Female", "168.67.162.1"},
		{"Lillian", "Lawrence", "llawrence9@blogtalkradio.com", "Female", "190.106.124.105"},
	}

	emailDomainsList := getEmailDomains(mockRows)

	emailDomainsListLenght := len(emailDomainsList)
	if emailDomainsListLenght != 10 {
		t.Errorf("emailDomainsListLenght | got: %d, expected: 10", emailDomainsListLenght)
	}

	emailDomainsListFirst := emailDomainsList[0]
	if emailDomainsListFirst != "github.io" {
		t.Errorf("emailDomainsListFirst | got: %s, expected: github.io", emailDomainsListFirst)
	}

	emailDomainsListSecond := emailDomainsList[1]
	if emailDomainsListSecond != "cyberchimps.com" {
		t.Errorf("emailDomainsListSecond | got: %s, expected: cyberchimps.com", emailDomainsListSecond)
	}

	emailDomainsListThird := emailDomainsList[2]
	if emailDomainsListThird != "hubpages.com" {
		t.Errorf("emailDomainsListThird | got: %s, expected: hubpages.com", emailDomainsListThird)
	}
}

func TestCountEmailDomains(t *testing.T) {
	mockEmailDomainsList := []string{
		"github.io",
		"cyberchimps.com",
		"github.io",
		"hubpages.com",
		"360.cn",
		"github.io",
		"statcounter.com",
		"rediff.com",
		"acquirethisname.com",
		"chicagotribune.com",
		"cnet.com",
		"blogtalkradio.com",
		"tinyurl.com",
		"tinyurl.com",
		"tinyurl.com",
		"tinyurl.com",
		"tinyurl.com",
		"tinyurl.com",
		"tinyurl.com",
	}

	emailDomainsMap := countEmailDomains(mockEmailDomainsList)

	emailDomainsMapLenght := len(emailDomainsMap)
	if emailDomainsMapLenght != 11 {
		t.Errorf("emailDomainsMapLenght | got: %d, expected: 11", emailDomainsMapLenght)
	}

	emailDomainsMapGithubCount := emailDomainsMap["github.io"]
	if emailDomainsMapGithubCount != 3 {
		t.Errorf("emailDomainsMapGithubCount | got: %d, expected: 3", emailDomainsMapGithubCount)
	}

	emailDomainsMapTinyurlCount := emailDomainsMap["tinyurl.com"]
	if emailDomainsMapTinyurlCount != 7 {
		t.Errorf("emailDomainsMapTinyurlCount | got: %d, expected: 7", emailDomainsMapTinyurlCount)
	}

	emailDomainsMapCnetCount := emailDomainsMap["cnet.com"]
	if emailDomainsMapCnetCount != 1 {
		t.Errorf("emailDomainsMapCnetCount | got: %d, expected: 1", emailDomainsMapCnetCount)
	}
}

func TestSortEmailDomains(t *testing.T) {
	mockEmailDomainsMap := map[string]int{
		"gerlach.com":         22,
		"schroeder.info":      3,
		"jacobi.com":          15,
		"romaguera.org":       4,
		"lakin.com":           24,
		"rempel.org":          3,
		"brakus.info":         3,
		"hud.gov":             6,
		"will.biz":            6,
		"123-reg.co.uk":       8,
		"163.com":             6,
		"1688.com":            3,
		"1und1.de":            5,
		"pouros.com":          20,
		"marks.com":           26,
		"stamm.com":           22,
		"rolfson.biz":         1,
		"herzog.info":         2,
		"douglas.net":         3,
		"alexa.com":           6,
		"larkin.biz":          2,
		"aboutads.info":       2,
		"accuweather.com":     6,
		"acquirethisname.com": 6,
		"addthis.com":         10,
	}

	sortedEmailDomains := sortEmailDomains(mockEmailDomainsMap)

	sortedEmailDomainsLenght := len(sortedEmailDomains)
	if sortedEmailDomainsLenght != 25 {
		t.Errorf("sortedEmailDomainsLenght | got: %d, expected: 25", sortedEmailDomainsLenght)
	}

	sortedEmailDomains123 := sortedEmailDomains[0]
	if sortedEmailDomains123.Name != "123-reg.co.uk" {
		t.Errorf("sortedEmailDomainsFirst.Name | got: %s, expected: 123-reg.co.uk", sortedEmailDomains123.Name)
	}
	if sortedEmailDomains123.NumberOfCustomers != 8 {
		t.Errorf("sortedEmailDomainsFirst.NumberOfCustomers | got: %d, expected: 8", sortedEmailDomains123.NumberOfCustomers)
	}

	sortedEmailDomainsAccuweather := sortedEmailDomains[5]
	if sortedEmailDomainsAccuweather.Name != "accuweather.com" {
		t.Errorf("sortedEmailDomainsSecond.Name | got: %s, expected: accuweather.com", sortedEmailDomainsAccuweather.Name)
	}
	if sortedEmailDomainsAccuweather.NumberOfCustomers != 6 {
		t.Errorf("sortedEmailDomainsSecond.NumberOfCustomers | got: %d, expected: 6", sortedEmailDomainsAccuweather.NumberOfCustomers)
	}

	sortedEmailDomainsGerlach := sortedEmailDomains[10]
	if sortedEmailDomainsGerlach.Name != "douglas.net" {
		t.Errorf("sortedEmailDomainsThird.Name | got: %s, expected: douglas.net", sortedEmailDomainsGerlach.Name)
	}
	if sortedEmailDomainsGerlach.NumberOfCustomers != 3 {
		t.Errorf("sortedEmailDomainsThird.NumberOfCustomers | got: %d, expected: 3", sortedEmailDomainsGerlach.NumberOfCustomers)
	}

	if sortedEmailDomains123.Name > sortedEmailDomainsAccuweather.Name || sortedEmailDomains123.Name > sortedEmailDomainsGerlach.Name || sortedEmailDomainsAccuweather.Name > sortedEmailDomainsGerlach.Name {
		t.Errorf("sortedEmailDomainsSortingComparison | email domains are not sorted properly")
	}
}
