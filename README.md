# fenv

This is a Go library to parse environment variables in a file.

## Install

````bash
go get github.com/lbrun25/fenv
````

## Example

Can parse any environment variable in a file

````yml
server:
  hostname: ${HOSTNAME}
  port: ${PORT}
````

and returns []byte

````yml
server:
  hostname: localhost
  port: 8080
````

## Usage

````go
package main

import (
	"bytes"
	"log"
	
	"github.com/lbrun25/fenv"
	"github.com/spf13/viper"
)

func main() {
	f := fenv.Mock()
	
	// Parse the file
	newContent, err := f.Parse("./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Use parsed configuration with viper for example
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err = viper.ReadConfig(bytes.NewBuffer(newContent))
	if err != nil {
		log.Fatal(err)
	}
}
````

	
