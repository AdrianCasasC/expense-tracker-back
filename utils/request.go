package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
)

func GetListOfFields(c *gin.Context) (map[string]bool, error) {
	// Consume y almacena el cuerpo de la solicitud
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	// Restaura el cuerpo de la solicitud para su posterior uso
	c.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))

	// Parse JSON en un mapa
	rawFields := make(map[string]json.RawMessage)
	if err = json.Unmarshal(bodyBytes, &rawFields); err != nil {
		return nil, err
	}

	// Extraer nombres de los campos incluyendo objetos anidados
	fieldsToUpdate := make(map[string]bool)
	extractFields("", rawFields, fieldsToUpdate)

	return fieldsToUpdate, nil
}

func extractFields(prefix string, data map[string]json.RawMessage, fields map[string]bool) {
	for key, rawValue := range data {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		// Intentamos parsear el valor como un objeto JSON
		var nestedMap map[string]json.RawMessage
		if err := json.Unmarshal(rawValue, &nestedMap); err == nil {
			// Si es un objeto, llamamos recursivamente a la función
			extractFields(fullKey, nestedMap, fields)
		} else {
			// Si no es un objeto, simplemente lo agregamos
			fields[fullKey] = true
		}
	}
}
