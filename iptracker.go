package main

import (
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"flag"
	"os"
)

var DEBUG = false

func main() {

	var ipAddress = flag.String("i", "NULL", "Specify ip address")
	flag.BoolVar(&DEBUG, "d", false, "Debugging")
	flag.Parse()

	if *ipAddress == "NULL" {
		flag.Usage()
		os.Exit(1)
	}

	url := fmt.Sprintf("http://www.geoplugin.net/json.gp?ip=%s", *ipAddress)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Couldn't make http connection\n")
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatalf("Couldn't read body\n")
	}

	var objmap map[string]*json.RawMessage
	err3 := json.Unmarshal(body, &objmap)
	if err3 != nil {
		log.Fatalf("Error unmarshalling: %s\n", err3)
		log.Fatalf("%+v", objmap)
	} 

	// get the contry and country code
	var cc string
	var cn string
	err = json.Unmarshal(*objmap["geoplugin_countryCode"], &cc)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
	err = json.Unmarshal(*objmap["geoplugin_countryName"], &cn)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	if cc != "US" {	
		fmt.Printf("IP: %s not in the US\n", *ipAddress)
		fmt.Printf("Results...\n")
		fmt.Printf(string(body))
	} else {
		if DEBUG == true {
			fmt.Printf("IP: %s is in the US\n", *ipAddress)
			fmt.Printf("Results...\n")
			fmt.Printf(string(body))
		}
	}
}
