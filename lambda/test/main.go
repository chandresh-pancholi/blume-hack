package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	rekognition2 "github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

type Handler struct {
	s3 *s3.S3
	rekog *rekognition2.Rekognition
}

func (h Handler) Handle(ctx context.Context, e events.S3Event) {

}


func main() {
	s := session.Must(session.NewSession(aws.NewConfig()))

	s3Session := s3.New(s)

	reko := rekognition2.New(s)

	h := Handler{
		s3: s3Session,
		rekog: reko,
	}

	lambda.Start(h.Handle)
}



//func parseEvent(e events.S3Event) (string, string, error) {
//	if len(e.Records) == 0 {
//		return "", "", errors.New("no records found")
//	}
//
//	s3 := e.Records[0].S3
//
//	return s3.Bucket.Name, s3.Object.Key, nil
//}
