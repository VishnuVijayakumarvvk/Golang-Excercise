package main

import "fmt"

func main() {
	var languages = make(map[string]string)
	languages["js"] = "Javascript"
	languages["php"] = "phpphp"
	languages["py"] = "python"
	languages["ty"] = "tygor"
	fmt.Println("The map is : ", languages)
	delete(languages, "ty")
	fmt.Println("The map is : ", languages)
	fmt.Println("The map  value is : ", languages["js"])

	for key, value := range languages {
		fmt.Printf("The key is %v and the value is %v\n", key, value)
	}
}
