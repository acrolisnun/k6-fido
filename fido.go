package k6fido

import (
	"encoding/json"
	"fmt"
)

type SendUafResponse struct {
	UafResponse string `json:"uafResponse"`
	Context     string `json:"context"`
}

func (k6fido *K6Fido) GenerateRegistrationResponse(aaid string, uafRequest string,
	trustedFacetId string, overriddenSignature string, signatureSignData string,
	privKey string, pubKey string) (string, error) {
	fidoRegistrationUafRequest := NewFidoRegistrationReturnUafRequest(uafRequest)

	fidoRegistrationResponse := FidoRegistrationResponse{
		facetId:          trustedFacetId,
		returnUafRequest: *fidoRegistrationUafRequest,
	}

	sendUafResponse, _ := fidoRegistrationResponse.Build(aaid, overriddenSignature, signatureSignData, privKey, pubKey)

	fidoRegistrationResponseString, err := json.Marshal(sendUafResponse)
	if err != nil {
		return "", fmt.Errorf("Failed to unmarshall ufa response: %s", err)
	}

	return string(fidoRegistrationResponseString), nil
}

func (k6fido *K6Fido) GenerateAuthenticationResponse(aaid string, uafRequest string,
	trustedFacetId string, overriddenSignature string, signatureSignData string,
	privKey string, pubKey string, username string) (string, error) {

	// fmt.Println(">> Generating Authentication Response.")

	fidoAuthenticationUafRequest := NewFidoAuthenticationReturnUafRequest(uafRequest)

	// fmt.Print("Fido Authentication UAF Request:")
	// fmt.Println(fidoAuthenticationUafRequest)

	fidoAuthenticationResponse := FidoAuthenticationResponse{
		facetId:          trustedFacetId,
		returnUafRequest: *fidoAuthenticationUafRequest,
		username:         username,
	}

	sendUafResponse, _ := fidoAuthenticationResponse.Build(aaid, overriddenSignature, signatureSignData, privKey, pubKey)

	fidoRegistrationResponseString, err := json.Marshal(sendUafResponse)
	if err != nil {
		return "", fmt.Errorf("Failed to marshall send ufa response: %s", err)
	}
	// fmt.Println(">> Generated Authentication Response.")
	return string(fidoRegistrationResponseString), nil
}
