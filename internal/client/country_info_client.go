package client

import (
	"context"
	"soap-client-in-go/internal/service/generated"
	"soap-client-in-go/internal/utils/constants"

	"github.com/hooklift/gowsdl/soap"
)

type CountryInfoClient struct {
	soapClient *soap.Client
}

// NewCountryInfoClient crea una nuova istanza di CountryInfoClient.
// Parametri:
// - url: L'URL del servizio SOAP.
// Restituisce:
// - Una nuova istanza di CountryInfoClient.
func NewCountryInfoClient(url string) *CountryInfoClient {
	return &CountryInfoClient{
		soapClient: soap.NewClient(url),
	}
}

// makeSoapCall esegue una chiamata SOAP al servizio specificato usando il client SOAP sottostante.
// Parametri:
// - ctx: il contesto della chiamata.
// - endpointURL: l'URL dell'endpoint del servizio.
// - request: la struttura di richiesta SOAP.
// - response: la struttura di risposta SOAP.
// Restituisce:
// - Un eventuale errore che si è verificato durante la chiamata SOAP.
func (cic *CountryInfoClient) makeSoapCall(ctx context.Context, endpointURL string, request, response interface{}) error {
	return cic.soapClient.CallContext(ctx, endpointURL, request, response)
}

// makeSoapCallWithoutContext esegue una chiamata SOAP al servizio specificato
// usando il client SOAP sottostante.
//
// Parametri:
// - endpointURL: l'URL dell'endpoint del servizio.
// - request: la struttura di richiesta SOAP.
// - response: la struttura di risposta SOAP.
//
// Restituisce:
// - Un eventuale errore che si è verificato durante la chiamata SOAP.
func (cic *CountryInfoClient) makeSoapCallWithoutContext(endpointURL string, request, response interface{}) error {
	return cic.soapClient.Call(endpointURL, request, response)
}

// ListOfContinentsByName esegue una chiamata SOAP al servizio ListOfContinentsByName utilizzando il client del servizio CountryInfo.
// Parametri:
// - ctx: il contesto della chiamata.
// - request: la struttura di richiesta ListOfContinentsByName.
// Restituisce:
// - la struttura di risposta ListOfContinentsByNameResponse.
// - un eventuale errore che si è verificato durante la chiamata SOAP.
func (cic *CountryInfoClient) ListOfContinentsByName(ctx context.Context, request *generated.ListOfContinentsByName) (*generated.ListOfContinentsByNameResponse, error) {
	response := new(generated.ListOfContinentsByNameResponse)
	err := cic.makeSoapCall(ctx, constants.ListOfContinentsByNameURL, request, response)
	return response, err
}

// ListOfContinentsByNameWithoutContext esegue una chiamata SOAP al servizio ListOfContinentsByName
// senza l'utilizzo di un contesto. Restituisce la struttura di risposta ListOfContinentsByNameResponse
// e un eventuale errore che si è verificato durante la chiamata SOAP.
// Parametri:
// - request: la struttura di richiesta ListOfContinentsByName.
func (cic *CountryInfoClient) ListOfContinentsByNameWithoutContext(request *generated.ListOfContinentsByName) (*generated.ListOfContinentsByNameResponse, error) {
	response := new(generated.ListOfContinentsByNameResponse)
	err := cic.makeSoapCallWithoutContext(constants.ListOfContinentsByNameURL, request, response)
	return response, err
}
