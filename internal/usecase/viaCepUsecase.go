package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/robinsonrenato/goexpert-cep-api/internal/dto"
	"github.com/robinsonrenato/goexpert-cep-api/pkg/formatter"
	"github.com/robinsonrenato/goexpert-cep-api/pkg/validator"
)

const VIA_CEP_SERVICE_NAME string = "Via Cep Service"

type viaCepOutput struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        bool   `json:"erro"`
}

type viaCepUsecase struct {
	url string
}

func NewViaCepUsecase() *viaCepUsecase {
	return &viaCepUsecase{
		url: "https://viacep.com.br",
	}
}

func (uc *viaCepUsecase) Execute(rawCep string) (*dto.AddressApi, *dto.AddressApiError) {
	cep := formatter.SanitalizeCep(rawCep)

	if !validator.Cep(cep) {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, "CEP incorreto")
	}

	cep = formatter.Cep(cep)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	endpoint := fmt.Sprintf("%s/ws/%s/json", uc.url, cep)
	req, err := http.NewRequestWithContext(ctx, "GET", endpoint, nil)
	if err != nil {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, err.Error())
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, err.Error())
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, "Erro ao obter os dados do CEP")
	}

	var response viaCepOutput

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, err.Error())
	}

	if response.Erro {
		return nil, dto.NewAddressApiError(VIA_CEP_SERVICE_NAME, "CEP nao Encontrado")
	}

	address := dto.NewAddressApi(VIA_CEP_SERVICE_NAME, response.Logradouro, response.Bairro, response.Localidade, response.Uf)
	return address, nil
}
