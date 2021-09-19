// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain. Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type EmailDomain struct {
	Name              string
	NumberOfCustomers int
}

func GetSortedEmailDomainsFromCsv(path string) ([]EmailDomain, error) {
	rows, err := readRowsFromCsv(path)
	if err != nil {
		log.Println("An error occurred while loading the CSV file.")
		return nil, err
	}
	emailDomainsList := getEmailDomains(rows)
	emailDomainsMap := countEmailDomains(emailDomainsList)
	sortedEmailDomains := sortEmailDomains(emailDomainsMap)

	return sortedEmailDomains, nil
}

func readRowsFromCsv(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	rows := [][]string{}
	reader := csv.NewReader(file)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		rows = append(rows, row)
	}

	return rows, nil
}

func getEmailDomains(rows [][]string) []string {
	var emailDomains []string

	for _, row := range rows[1:] {
		email := row[2]

		s := strings.Split(email, "@")
		if len(s) == 2 {
			emailDomains = append(emailDomains, s[1])
		} else {
			log.Printf("invalid email: %s", email)
		}
	}

	return emailDomains
}

func countEmailDomains(emailDomains []string) map[string]int {
	emailDomainsMap := make(map[string]int)

	for _, emailDomain := range emailDomains {
		_, exist := emailDomainsMap[emailDomain]
		if exist {
			emailDomainsMap[emailDomain] += 1
		} else {
			emailDomainsMap[emailDomain] = 1
		}
	}

	return emailDomainsMap
}

func sortEmailDomains(emailDomainsMap map[string]int) []EmailDomain {
	keys := make([]string, 0, len(emailDomainsMap))
	for key := range emailDomainsMap {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sortedEmailDomains []EmailDomain
	for _, key := range keys {
		sortedEmailDomains = append(sortedEmailDomains, EmailDomain{key, emailDomainsMap[key]})
	}

	return sortedEmailDomains
}
