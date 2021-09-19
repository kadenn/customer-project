// Package customeradder adds new customers to a given csv file from fakerapi.
// API Docs: https://fakerapi.it/en
package customeradder

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
)

type Customer struct {
	UUID       string `json:"uuid"`
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	IP         string `json:"ip"`
	MacAddress string `json:"macAddress"`
	Website    string `json:"website"`
	Image      string `json:"image"`
}

type ApiResponse struct {
	Status string     `json:"status"`
	Code   int        `json:"code"`
	Total  int        `json:"total"`
	Data   []Customer `json:"data"`
}

func AddNewCustomersToCsv(path string, number int, wg *sync.WaitGroup) {
	log.Printf("Adding %d new customers to the dataset...", number)
	defer wg.Done()

	customers, err := getCustomersFromApi(number)
	if err != nil {
		log.Printf("An error occurred while getting customers from API: %s", err)
	}

	addErr := addCustomersToCsv(path, customers)
	if addErr != nil {
		log.Printf("An error occurred while adding customers to CSV: %s", err)
	}

}

func getCustomersFromApi(number int) ([]Customer, error) {

	if number < 1 || number > 1000 {
		err := errors.New("the number of customers must be an integer between 0 and 1000")
		return nil, err
	}

	url := fmt.Sprintf("https://fakerapi.it/api/v1/users?_quantity=%d", number)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	apiResponse := ApiResponse{}
	jsonErr := json.Unmarshal(body, &apiResponse)
	if jsonErr != nil {
		return nil, jsonErr
	}

	return apiResponse.Data, nil
}

func addCustomersToCsv(path string, customers []Customer) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}

	defer file.Close()

	var newCustomers [][]string

	for _, customer := range customers {
		// customers.csv columns: first_name, last_name, email, gender, ip_address
		genders := [3]string{"Male", "Female", "Nonbinary"}
		newCustomer := []string{customer.Firstname, customer.Lastname, customer.Email, genders[rand.Intn(3)], customer.IP}
		newCustomers = append(newCustomers, newCustomer)
	}

	writer := csv.NewWriter(file)

	writeErr := writer.WriteAll(newCustomers)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
