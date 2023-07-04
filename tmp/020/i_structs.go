package main

import (
	"fmt"
	"reflect"
	"time"
)

func structs() {
	type address struct {
		zipCode int
		city    string
		country string
	}
	berlinCentral := address{zipCode: 10557, city: "Berlin"}
	fmt.Println(berlinCentral)

	hamburgFishmarket := address{22767, "Hamburg", "DE"}
	fmt.Println(hamburgFishmarket)
}

func structs2() {
	type address struct {
		zipCode int
		city    string
		country string
	}
	type contact struct {
		fullName    string
		homeAddress address
	}
	me := contact{"Thomas Auinger", address{97234, "Reichenberg", "DE"}}
	fmt.Println(me)
}

func anonStruct() {
	logData := struct {
		level     string
		message   string
		timestamp int64
	}{
		"INFO",
		"The application is starting up",
		time.Now().UnixMilli(),
	}
	fmt.Println(logData)
}

func taggedStruct() {
	type logData struct {
		level           string `json:"lvl" xml:"the-level"`
		message         string `json:"msg"`
		ignoreThisField int    `json:"-"`
	}
}

func printTags() {
	type User struct {
		username string `json:"user-name"`
		roles    []string
		disabled bool
	}
	u := User{username: "john.doe", disabled: true}

	fmt.Println("Fields of struct:")
	t := reflect.TypeOf(u)
	for n := 0; n < t.NumField(); n++ {
		structField := t.Field(n)
		fmt.Printf("#%d %s\n", n, structField.Name)
		fmt.Printf("   value of `json` tag: %s\n", structField.Tag.Get("json"))
	}
}
