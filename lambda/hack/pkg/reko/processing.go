package reko

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"log"
)

type Reko struct {
	Rekog *rekognition.Rekognition
}

func (r *Reko) DetectText(bucket, name string) (*rekognition.DetectTextOutput, error) {
	input := &rekognition.DetectTextInput{
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: aws.String(bucket),
				Name:   aws.String(name),
			},
		},
	}
	result, err := r.Rekog.DetectText(input)
	if err != nil {
		log.Printf("detect lable processing failed %v ", err)
		return nil, err
	}

	return result, nil
}
