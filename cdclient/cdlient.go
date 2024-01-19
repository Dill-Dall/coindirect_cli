package cdclient

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Country struct {
	Id              uint        `json:"id"`
	Name            string      `json:"name"`
	DefaultCurrency string      `json:"defaultCurrency"`
	Documents       []Documents `json:"documents"`
	Options         Options     `json:"options"`
}

type Options struct {
	WithdrawalMaximum string `json:"withdrawalMaximum"`
}

type Documents struct {
	Id          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type SortBy string

const (
	Id       SortBy = "id"
	Currency SortBy = "currency"
	Name     SortBy = "name"
)

const baseURL = "https://api.sandbox.coindirect.com/api/"

// Don`t know what this means. Maybe countries not elligble? Though response seems the same
const enabled = false
const offset = 0

// Max countries are 250 according to the api, according to wikipedia, 195.
const max = 300

// Used to map currency to countries.
var currencyMap = make(map[string]string)

// Toggle whether to output the map or the default list.
var isOutputCurrencyMap = false

/*
I imagine the sandbox have a lot of invalid data. Norway by instance have no default currency.
None of those not having a default currency, has required documents. Which is interesting.
*/
func FetchCountries(descending bool, sb SortBy, isCurrencyMap bool) {
	//Needed because cli flags default to false
	var ascending = !descending
	
	isOutputCurrencyMap = isCurrencyMap
	var countries = getCountriesFromApi()
	sortCountries(ascending, sb, &countries)
	output(countries)
}

func getCountriesFromApi() []Country {
	u, err := url.Parse(baseURL + "country")
	if err != nil {
		log.Fatal(err)
	}

	query := u.Query()
	query.Set("enabled", strconv.FormatBool(enabled))
	query.Set("offset", strconv.Itoa(offset))
	query.Set("max", strconv.Itoa(max))
	u.RawQuery = query.Encode()

	fmt.Println("Get from: " + u.String())
	response, err := http.Get(u.String())
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	if response.StatusCode != 200 {
		fmt.Print("Error: " + response.Status)
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject []Country
	err = json.Unmarshal(responseData, &responseObject)
	if err != nil {
		log.Fatalf("Error occurred during unmarshalling: %s", err)
	}

	return responseObject
}

func sortCountries(asc bool, sb SortBy, countries *[]Country) {
	switch sb {
	case Id:
		sort.Slice(*countries, func(i, j int) bool {
			if asc {
				return (*countries)[i].Id < (*countries)[j].Id
			} else {
				return (*countries)[i].Id > (*countries)[j].Id
			}
		})
	case Currency:
		sort.Slice(*countries, func(i, j int) bool {
			if asc {
				return (*countries)[i].DefaultCurrency < (*countries)[j].DefaultCurrency
			} else {
				return (*countries)[i].DefaultCurrency > (*countries)[j].DefaultCurrency
			}
		})
	case Name:
		sort.Slice(*countries, func(i, j int) bool {
			if asc {
				return strings.ToLower((*countries)[i].Name) < strings.ToLower((*countries)[j].Name)
			} else {
				return strings.ToLower((*countries)[i].Name) > strings.ToLower((*countries)[j].Name)
			}
		})
	}
}

func output(countries []Country) {
	fmt.Printf("ID, N:Name, C:Currency, WMax:WithdrawalMaximum\n\n%s", currencyMap[""])
	for _, country := range countries {
		if !isOutputCurrencyMap {
			var withdrawalMaximum = country.Options.WithdrawalMaximum
			if withdrawalMaximum == "" {
				withdrawalMaximum = "0"
			}
			fmt.Printf("ID: %d, N: %s, C: %s, WMax: %s\n", country.Id, country.Name, country.DefaultCurrency, country.Options.WithdrawalMaximum)
			outPutRequiredDocuments(country)
		} else {
			currencyMap[country.DefaultCurrency] += (country.Name + " ")
		}

	}

	if isOutputCurrencyMap {
		outputCurrencyMap()
	}
}

func outputCurrencyMap() {
	for key, val := range currencyMap {
		if key != "" {
			fmt.Printf("Nr: %s, N: %s\n", key, val)
		}
	}
}

func outPutRequiredDocuments(country Country) {
	var requiredDocuments = ""
	for _, doc := range country.Documents {

		if doc.Required {
			requiredDocuments += fmt.Sprintf("	ID: %d, %s - %s\n", doc.Id, doc.Code, doc.Description)
		}
	}
	if requiredDocuments != "" {
		fmt.Printf("Required Documents: {\n%s}\n", requiredDocuments)
	}
}

func ParseSortBy(str string) (SortBy, error) {
	switch str {
	case "id":
		return Id, nil
	case "currency":
		return Currency, nil
	case "name":
		return Name, nil
	default:
		return "", fmt.Errorf("invalid SortBy value: %s", str)
	}
}
