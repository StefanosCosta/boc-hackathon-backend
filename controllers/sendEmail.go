package controllers

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendEmail(subject, body, recipient, sender string) {
	// Initialize a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1"), // Replace with your desired AWS region
	})
	if err != nil {
		log.Fatal(err)
	}

	// Create a new SES session
	svc := ses.New(sess)

	// Specify the email parameters
	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(recipient),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data: aws.String(body),
				},
			},
			Subject: &ses.Content{
				Data: aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	// Send the email
	result, err := svc.SendEmail(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully. Message ID:", *result.MessageId)
}
