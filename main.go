package openapi

import (
	"context"

	openapi "github.com/Niceblueman/whatsapp-rapid-api"
)

func main() {
	ctx := context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
		"{classname}Service.{nickname}": {
			"port": "8443",
		},
	})
}
