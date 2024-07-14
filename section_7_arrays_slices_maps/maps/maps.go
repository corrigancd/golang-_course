package main

import (
	"fmt"
)

func main() {
	// creating a map
	websites := map[string]string{"google": "https://google.com", "aws": "https://aws.com"}
	fmt.Println(websites)
	fmt.Println(websites["aws"])

	// adding new items to map
	websites["facebook"] = "https://facebook.com"
	fmt.Println(websites)

	// deleting items from map
	delete(websites, "google")
	fmt.Println(websites)

	// iterating over map

}
