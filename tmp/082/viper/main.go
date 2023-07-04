package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("app.some-key", 123)
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	foo()
}

func foo() {
	n := viper.GetInt("app.some-key")
	fmt.Println(n)
}
