package dto

type AddressApi struct {
	ApiName      string `json:"api_name"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
	City         string `json:"cidade"`
	State        string `json:"uf"`
}

type AddressApiError struct {
	ApiName string `json:"api_name"`
	Message string `json:"message"`
}

func NewAddressApi(apiName, street string, neighborhood string, city, state string) *AddressApi {
	return &AddressApi{
		ApiName:      apiName,
		Street:       street,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
	}
}

func NewAddressApiError(apiName string, message string) *AddressApiError {
	return &AddressApiError{
		ApiName: apiName,
		Message: message,
	}
}
