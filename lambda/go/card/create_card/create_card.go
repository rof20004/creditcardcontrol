package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	f "github.com/fauna/faunadb-go/faunadb"
	m "github.com/rof20004/creditcardcontrol/card"
)

var faunaDB f.Value

func init() {
	client := f.NewFaunaClient(os.Getenv("FAUNADB_SERVER_SECRET"))

	res, err := client.Query(f.Get(f.Ref("collections/Card")))
	if err != nil {
		panic(err)
	}

	faunaDB = res
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod != "POST" {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       "Método inválido",
		}, nil
	}

	var cards []*m.Card

	if err := faunaDB.At(f.ObjKey("data")).Get(&cards); err != nil {
		log.Println(err)
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       "Erro ao recuperar os dados",
		}, nil
	}

	b, err := json.Marshal(cards)
	if err != nil {
		log.Println(err)
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusMethodNotAllowed,
			Body:       "Erro ao converter dados",
		}, nil
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(b),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
