package awsTools

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

const (
	AWS_REGION = "us-east-1"
)

type AWSServices struct {
	Session *session.Session
	SSM *ssm.SSM
}

// AWS Services constructor
func NewAWSService() (*AWSServices, error) {
	// Initialization of session and services
	session, err := StartNewAWSSession()
	if err != nil {
		return nil, err
	}
	// Initialize our ssm
	smmSvc := StartNewSSMService(session)

	return &AWSServices{
		Session: session,
		SSM: smmSvc,
	}, nil
}

func StartNewAWSSession() (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(AWS_REGION)},
		SharedConfigState: session.SharedConfigEnable,
	})

	return sess, err
}
