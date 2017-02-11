package tukilastic

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/edoardo849/apex-aws-signer"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
)

func newClient(url string, sniff bool, awsRegion string) (*elastic.Client, error) {
	ses := session.New(&aws.Config{Region: aws.String(awsRegion)})

	transport := signer.NewTransport(ses, elasticsearchservice.ServiceName)

	httpClient := &http.Client{
		Transport: transport,
	}

	client, err := elastic.NewClient(
		elastic.SetURL(url),
		elastic.SetSniff(sniff),
		elastic.SetHttpClient(httpClient),
	)

	return client, err
}
