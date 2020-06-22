package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Config:  aws.Config{Region: aws.String("eu-west-1")},
		Profile: "tf_demo",
	})

	if err != nil {
		log.Fatal(err)
	}

	svc := ec2.New(sess)
	result, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(result)

}
