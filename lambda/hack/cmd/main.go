package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
	"hack/pkg/kafka"
	"hack/pkg/lambdas"
)

func main() {
	s := session.Must(session.NewSession(aws.NewConfig()))
	h := lambdas.Handler{
		S3:       s3.New(s),
		Rekog:    rekognition.New(s),
		Producer: kafka.Producer(),
	}

	h.Successes()
	h.Errors()

	lambda.Start(h.Handle)
}
