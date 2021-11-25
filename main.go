package main

import "gotag/parse"

type User struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	About 		string `json:"about" `
	FavLanguage string `json:"fav_language" must:"typescript,go,python"`
}

func main() {
	user := User{
		Name:        "Nick",
		Age:         19,
		FavLanguage: "go",
	}

	if err := parse.Must(&user); err != nil {
		panic(err)
	}
}