package main

import (
	"context"
	"log"
	"os"
	"soap-client-in-go/internal/client"
	"soap-client-in-go/internal/config"
	"soap-client-in-go/internal/service"
	"time"
)

// main is the entry point of the program. It initializes the CountryInfoService, performs two SOAP API calls
// using different contexts, and prints the list of continents returned by each API call.
// The first API call is made without a context, and the second API call is made with a timeout context of 5 seconds.
// The program also logs debug information before and after each API call.
func main() {
	countryInfoService := initializeCountryInfoService()
	countryInfoService.ListOfContinentByNameWithoutContext()
	// Creare un contesto con timeout di 5 secondi
	log.Println("•••••••••••••••••••••••••••••••••••••••••••••••••")
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	countryInfoService.ListOfContinentsByName(ctx)
}

// init is called implicitly at program startup. It sets the log flags, log output,
// and the log prefix.
// It sets the log flags to 0, so that no additional information is included in the log messages.
// It sets the log output to os.Stdout, so that log messages are written to the standard output.
// It sets the log prefix to "••••••• ✈️✈️ ", which is added at the beginning of each log message.
func init() {
	log.SetFlags(4)
	log.SetOutput(os.Stdout)
	log.SetPrefix("✈️ ")
}

// initializeCountryInfoService creates a new instance of the CountryInfoService by initializing all its dependencies.
// It retrieves the configuration from NewConfig, creates a CountryInfoClient using NewCountryInfoClient,
// and finally initializes the CountryInfoService using NewCountryInfoService.
// The initialized CountryInfoService is returned.
func initializeCountryInfoService() *service.CountryInfoService {
	configuration := config.NewConfig()
	soapClient := client.NewCountryInfoClient(configuration.CountryInfoServiceURL)
	countryInfoService := service.NewCountryInfoService(soapClient)
	return countryInfoService
}
