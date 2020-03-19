package main

import (
	"extres"
	"fmt"
	"os"
)

// App is the application's global context
var App struct {
	X extres.ExternalResources
}

func main() {
	err := extres.ReadConfig("./config.json", &App.X)
	if err != nil {
		fmt.Printf("Error reading confing file: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Dbuser = %s\n", App.X.Dbuser)
	fmt.Printf("MojoWebAddr = %s\n", App.X.MojoWebAddr)
	fmt.Printf("Timezone = %s\n", App.X.Timezone)
	fmt.Printf("Tester1 = %s\n", App.X.Tester1Name)
	fmt.Printf("Tester2 = %s\n", App.X.Tester2Name)
	fmt.Printf("WREISDbname = %s\n", App.X.WREISDbname)
}
