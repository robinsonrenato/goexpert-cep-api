package main

import (
	"flag"
	"log"
	"time"

	"github.com/robinsonrenato/goexpert-cep-api/internal/dto"
	"github.com/robinsonrenato/goexpert-cep-api/internal/usecase"
	"github.com/robinsonrenato/goexpert-cep-api/pkg/formatter"
)

func main() {

	cepArg := flag.String("cep", "00000000", "Digite o CEP para retornar o Endereco")
	flag.Parse()

	apiCepUsecase := usecase.NewApiCepUsecase()
	viaCepUsecase := usecase.NewViaCepUsecase()

	addressChannel := make(chan *dto.AddressApi)
	addressErrChannel := make(chan *dto.AddressApiError)

	go func() {
		address, err := apiCepUsecase.Execute(*cepArg)
		if err != nil {
			addressErrChannel <- err
			return
		}
		addressChannel <- address
	}()

	go func() {
		address, err := viaCepUsecase.Execute(*cepArg)
		if err != nil {
			addressErrChannel <- err
			return
		}
		addressChannel <- address
	}()

	select {
	case address := <-addressChannel:
		log.Println(formatter.JSON(address))

	case err := <-addressErrChannel:
		log.Println(formatter.JSON(err))

	case <-time.After(time.Second):
		log.Println("timeout exceeded")
	}
}
