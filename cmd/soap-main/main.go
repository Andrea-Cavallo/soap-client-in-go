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

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
	log.SetPrefix("••••••• ✈️✈️ ")
}
func main() {
	countryInfoService := initializeCountryInfoService()
	//countryInfoService.ListOfContinentByNameWithoutContext()
	// Creare un contesto con timeout di 5 secondi
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	countryInfoService.ListOfContinentsByName(ctx)
}

func initializeCountryInfoService() *service.CountryInfoService {
	configuration := config.NewConfig()
	soapClient := client.NewCountryInfoClient(configuration.CountryInfoServiceURL)
	countryInfoService := service.NewCountryInfoService(soapClient)
	return countryInfoService
}
