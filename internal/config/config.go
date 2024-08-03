package config

import (
	"os"
)

// Config è la struttura che contiene le configurazioni dell'applicazione.
type Config struct {
	CountryInfoServiceURL string
}

// NewConfig crea una nuova istanza di Config.
// Legge le variabili di ambiente e fornisce i valori di default se non sono settate.
func NewConfig() *Config {
	return &Config{
		CountryInfoServiceURL: getEnv("COUNTRY_INFO_SERVICE_URL", "http://www.oorsprong.org/websamples.countryinfo/CountryInfoService.wso?WSDL"),
	}
}

// getEnv legge una variabile di ambiente e restituisce un valore di default se non è settata.
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
