package service

import (
	"context"
	"log"
	"soap-client-in-go/internal/client"
	"soap-client-in-go/internal/service/generated"
	"time"
)

type CountryInfoService struct {
	client *client.CountryInfoClient
}

// NewCountryInfoService crea una nuova istanza di CountryInfoService.
// Parametri:
// - client: Un'istanza di CountryInfoClient.
// Restituisce:
// - Una nuova istanza di CountryInfoService.
func NewCountryInfoService(client *client.CountryInfoClient) *CountryInfoService {
	return &CountryInfoService{
		client: client,
	}
}

// ListOfContinentsByName recupera una lista di continenti tramite una chiamata SOAP
// utilizzando un contesto. Se il contesto viene annullato prima della chiamata,
// la funzione restituisce immediatamente senza eseguire la chiamata SOAP.
// La lista dei continenti viene stampata sulla console. Non restituisce alcun valore.
func (cis *CountryInfoService) ListOfContinentsByName(ctx context.Context) {
	request := &generated.ListOfContinentsByName{}
	time.Sleep(6 * time.Second)
	select {
	case <-ctx.Done():
		log.Println("Contesto annullato prima della chiamata SOAP:", ctx.Err())
		return
	default:
		response, _ := cis.client.ListOfContinentsByName(ctx, request)
		for _, continent := range response.ListOfContinentsByNameResult.TContinent {
			log.Printf("Codice Continente: %s, Nome Continente: %s\n", continent.SCode, continent.SName)
		}
	}

}

// ListOfContinentByNameWithoutContext recupera una lista di continenti tramite una chiamata SOAP
// senza l'utilizzo di un contesto.
// La lista dei continenti viene stampata sulla console.
// Non restituisce alcun valore.
func (cis *CountryInfoService) ListOfContinentByNameWithoutContext() {
	request := &generated.ListOfContinentsByName{}
	response, _ := cis.client.ListOfContinentsByNameWithoutContext(request)
	for _, continent := range response.ListOfContinentsByNameResult.TContinent {
		log.Printf("Codice Continente: %s, Nome Continente: %s\n", continent.SCode, continent.SName)
	}

}
