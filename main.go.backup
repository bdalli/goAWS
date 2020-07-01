package main

import (
	"fmt"
	//"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	sess, err := session.NewSessionWithOptions(session.Options{
		// Specify profile to load for the session's config
		Config:  aws.Config{Region: aws.String("eu-west-1")},
		Profile: "dev",
	})

	if err != nil {
		log.Fatal(err)
	}

	input := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String("i-060798e03cf3aa650"),
		},
	}

	//	input := &ec2.DescribeInstancesInput{}

	svc := ec2.New(sess)
	result, err := svc.DescribeInstances(input)
	if err != nil {
		log.Fatal(err)
		return
	}
	// results is a struct with a list of results or next_token
	//fmt.Println(len(result.Reservations), result.Reservations)

	for _, reservation := range result.Reservations {
		for i, instance := range reservation.Instances {
			fmt.Println(i, *instance)
		}
	}

}
