package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"log"
	"net/http"
)

var apigw = apigateway.New(session.Must(session.NewSessionWithOptions(session.Options{
	Profile:           "my_profile", // TODO 書き換え
	SharedConfigState: session.SharedConfigEnable,
},
)))

func main() {

	params := &apigateway.TestInvokeMethodInput{
		RestApiId:           aws.String("1234123412"), // TODO 書き換え
		ResourceId:          aws.String("3gapai"),     // TODO 書き換え
		HttpMethod:          aws.String(http.MethodGet),
		PathWithQueryString: aws.String("/v1/health"),
	}

	req, resp := apigw.TestInvokeMethodRequest(params)
	if err := req.Send(); err != nil {
		log.Fatalf("test invoke: %v", err)
	}
	fmt.Printf("%d\n%v\n%v\n", *resp.Status, aws.StringValueMap(resp.Headers), *resp.Body)
}
