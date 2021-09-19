package main

import (
	"log"
	"sync"
	"time"

	"github.com/kadenn/customer-project/customeradder"
	"github.com/kadenn/customer-project/customerimporter"
)

func main() {
	start := time.Now()
	defer func() {
		log.Println("Execution Time: ", time.Since(start))
	}()

	file := "./data/more_customers.csv"

	// The code below adds 5x1000 new customers to the given csv file concurrently using Goroutines.
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go customeradder.AddNewCustomersToCsv(file, 1000, &wg)
		time.Sleep(1 * time.Second) // Slows down starting new Goroutines to be more polite to the API.
	}
	wg.Wait()

	// The code below gets sorted email domains from the given csv file and prints the results.
	sortedEmailDomains, err := customerimporter.GetSortedEmailDomainsFromCsv(file)
	if err != nil {
		log.Fatal(err)
	}

	totalNumberOfCustomers := 0
	for _, emailDomain := range sortedEmailDomains {
		totalNumberOfCustomers += emailDomain.NumberOfCustomers
	}

	log.Printf("Total number of email domains: %d", len(sortedEmailDomains))
	log.Printf("Total number of valid customers: %d", totalNumberOfCustomers)

}
