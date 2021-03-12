package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	empJson := `{
		"id": 11,
		"name": "Irshad",
		"department": "IT",
		"designation": "Product Manager",
		"address": {
			"city": "Mumbai",
			"state": "Maharashtra",
			"country": "India"
		}
	}`

	var result map[string]interface{}

	json.Unmarshal([]byte(empJson), &result)

	address := result["address"].(map[string]interface{})

	fmt.Println("Id :", result["id"],
		"\nName :", result["name"],
		"\nDepartment :", result["department"],
		"\nDesignation :", result["designation"],
		"\nAddress :", address["city"], address["state"], address["country"])
}
