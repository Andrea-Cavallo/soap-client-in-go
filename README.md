# Soap Client implementation with GO.

## Introduzione:

Questo progetto implementa un client `SOAP: Simple Object Access Protocol`
Protocollo standard progettato per permettere la comunicazione tra applicazioni costruite con linguaggi diversi e su piattaforme diverse.
Essendo un protocollo, impone regole integrate che aumentano la complessità e l'overhead, il che può portare a tempi di caricamento più lunghi.
Tuttavia, questi standard offrono sicurezza e ACID che può essere preferibile per scenari aziendali

In questo esempio il client è progettato per interagire con un servizio SOAP che fornisce informazioni sui paesi preso da: http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso

## Come leggere un file .wsdl e generare una service.go:

1) Installare gowsdl:

```bash
go install github.com/hooklift/gowsdl/cmd/gowsdl@latest
```

2) Generare il service.go a partire dal .wsdl:
```bash
PS gowsdl wsdl2go --output client/ wsdl/countryinfo.wsdl
🍀  Reading file ...\soap-client-spd\wsdl\service.wsdl
🍀  Done 👍
```

## Descrizione dei packages principali - client,service:

### Client:

* **CountryInfoClient** è la struttura principale che rappresenta il client per il servizio `SOAP`. Include un'istanza di `soap.Client` fornita dalla libreria `github.com/hooklift/gowsdl/soap`.
* **NewCountryInfoClient** restituisce un nuovo oggetto `CountryInfoClient` con il client `SOAP` inizializzato.

Il client include due metodi principali per effettuare chiamate `SOAP`:

- **makeSoapCall**: Effettua una chiamata `SOAP` utilizzando un contesto (`context.Context`), permettendo l'annullamento della chiamata in base al contesto.
- **makeSoapCallWithoutContext**: Effettua una chiamata `SOAP` senza contesto, adatta per chiamate che non richiedono gestione del contesto.

Il client fornisce metodi specifici per interagire con il servizio `SOAP`:

- **ListOfContinentsByName**: Recupera una lista di continenti per nome utilizzando un contesto.
- **ListOfContinentsByNameWithoutContext**: Recupera una lista di continenti per nome senza utilizzare un contesto.

### Service:

`CountryInfoService` è la struttura principale che rappresenta il servizio per gestire le informazioni sui paesi.
Per creare un'istanza di `CountryInfoService`, si utilizza la funzione `NewCountryInfoService` che richiede un'istanza di `CountryInfoClient` come parametro.

Nel package service sono presenti le logiche di business principali, tra cui due metodi che evidenziano la differenza nell'utilizzo dello stesso metodo con o senza il parametro context.
Questi metodi permettono di comprendere al meglio le potenzialità dell'uso del contesto nelle chiamate SOAP. 
Entrambi i servizi richiamano i metodi esposti dal client e stampano la risposta.


### Main:

* **initializeCountryInfoService** - metodo initialize carica l'url che serve per istanziare l'oggetto soapClient che serve poi per istanziare il countryInfoService, tramite il quale accediamo ai servizi.

```go
func initializeCountryInfoService() *service.CountryInfoService {
    configuration := config.NewConfig()
    soapClient := client.NewCountryInfoClient(configuration.CountryInfoServiceURL)
    countryInfoService := service.NewCountryInfoService(soapClient)
    return countryInfoService
}
```

## Come testare il progetto:

1) **Clonare il repository**:
```sh
git clone <url_del_repository>
```

2) **Configurare le dipendenze**:
```sh
go mod tidy
```

3) **Eseguire il progetto**:

```sh
go run main.go
```

## Ulteriori note:

Punto fondamentale visto in questo progetto è:
L'uso del `context` nei metodi di servizio generati da `gowsdl` è una peculiarità del linguaggio Go.
Il context è una funzionalità potente di Go che permette:

- **Timeout e Cancellazione**: Permette di specificare un timeout o di cancellare una richiesta se non è più necessaria.
- **Propagazione di valori**: Facilita il passaggio di valori tra diverse parti dell'applicazione.
- **Gestione concorrente**: Migliora la gestione delle richieste concorrenti, riducendo il rischio di goroutine orfane.

## Resources utili:

* **W3C WSDL spec:** http://www.w3.org/TR/wsdl
* **W3C SOAP 1.2 spec:** http://www.w3.org/TR/soap/

## Author:
* **Andrea-Cavallo**