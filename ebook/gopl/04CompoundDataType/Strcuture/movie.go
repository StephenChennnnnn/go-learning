package Strcuture

import (
	"encoding/json"
	"fmt"
	"log"
)

type Moive struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Moive{
	{Title: "Casablanca", Year: 1942, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	// ...
}

func main() {
	// Marshal
	{
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)
	}

	// Unmarshal
	{
		data, err := json.MarshalIndent(movies, "", "    ")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s", err)
		}
		fmt.Printf("%s\n", data)

		var titles []struct{ Title string }
		err = json.Unmarshal(data, &titles)
		if err != nil {
			log.Fatalf("JSON unmarshaling failed: %s", err)
		}
		fmt.Println(titles)
		// "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
	}
}
