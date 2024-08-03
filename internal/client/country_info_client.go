package client

import (
	"context"
	"soap-client-in-go/internal/service/generated"
	"soap-client-in-go/internal/utils/constants"

	"github.com/hooklift/gowsdl/soap"
)

// CountryInfoClient è il client per interagire con il servizio SOAP di informazioni sui paesi.
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

func (cic *CountryInfoClient) makeSoapCall(ctx context.Context, endpointURL string, request, response interface{}) error {
	return cic.soapClient.CallContext(ctx, endpointURL, request, response)
}

func (cic *CountryInfoClient) makeSoapCallWithoutContext(endpointURL string, request, response interface{}) error {
	return cic.soapClient.Call(endpointURL, request, response)
}

func (cic *CountryInfoClient) ListOfContinentsByName(ctx context.Context, request *generated.ListOfContinentsByName) (*generated.ListOfContinentsByNameResponse, error) {
	response := new(generated.ListOfContinentsByNameResponse)
	err := cic.makeSoapCall(ctx, constants.ListOfContinentsByNameURL, request, response)
	return response, err
}

func (cic *CountryInfoClient) ListOfContinentsByNameWithoutContext(request *generated.ListOfContinentsByName) (*generated.ListOfContinentsByNameResponse, error) {
	response := new(generated.ListOfContinentsByNameResponse)
	err := cic.makeSoapCallWithoutContext(constants.ListOfContinentsByNameURL, request, response)
	return response, err
}
