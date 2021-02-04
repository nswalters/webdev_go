package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name        string
		Weather     string
		WeekWeather map[string]string
		IsCool      bool
	}{"John Smith", "Cloudy", map[string]string{"Monday": "Rain", "Tuesday": "Partly Sunny"}, false}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
