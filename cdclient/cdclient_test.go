package cdclient

import (
	"encoding/json"
	"testing"
)

const jsonStr = `[
    {
        "id": 1,
        "name": "CountryA",
        "defaultCurrency": "CurrencyA",
        "documents": [
            {
                "id": 1,
                "code": "DocCodeA",
                "description": "Document A",
                "required": true
            }
        ],
        "options": {
            "withdrawalMaximum": "1000"
        }
    },
	{
        "id": 2,
        "name": "CountryB",
        "defaultCurrency": "CurrencyB",
        "documents": [
            {
                "id": 2,
                "code": "DocCodeB",
                "description": "Document B",
                "required": false
            }
        ],
        "options": {}
    }
]`

func TestUnmarshalCountry(t *testing.T) {
	var countries []Country
	err := json.Unmarshal([]byte(jsonStr), &countries)
	if err != nil {
		t.Errorf("Unmarshal failed: %s", err)
	}

	if len(countries) != 2 {
		t.Fatal("Not all objects unmarshalled")
	}

	if countries[0].Id != 1 ||
		countries[0].Name != "CountryA" ||
		countries[0].DefaultCurrency != "CurrencyA" {
		t.Errorf("Unmarshalling did not produce expected results for root object")
	}

	if countries[0].Documents[0].Id != 1 ||
		countries[0].Documents[0].Code != "DocCodeA" ||
		countries[0].Documents[0].Description != "Document A" {
		t.Errorf("Unmarshalling did not produce expected results for Document object")
	}

	if countries[0].Options.WithdrawalMaximum != "1000" {
		t.Errorf("Unmarshalling did not produce expected results for Options object")
	}
}
