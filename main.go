package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
)

func main() {
	stackName := flag.String("stack-name", "", "Name of the stack to read resources from")
	flag.Parse()

	fmt.Println("Stack name:", *stackName)

	if *stackName == "" {
		fmt.Println("Please provide a CloudFormation stack name with --stack-name")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	client := cloudformation.NewFromConfig(cfg)

	input := &cloudformation.DescribeStackResourcesInput{
		StackName: stackName,
	}

	resp, err := client.DescribeStackResources(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to describe stack resources, %v", err)
	}

	fmt.Printf("Resources in stack '%s':\n", *stackName)
	for _, resource := range resp.StackResources {
		fmt.Printf("Logical ID: %s, Type: %s, Physical ID: %s\n", *resource.LogicalResourceId, *resource.ResourceType, *resource.PhysicalResourceId)
	}
}