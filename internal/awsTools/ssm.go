package awsTools

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var allowedTypes = map[string] bool {
	ssm.ParameterTypeString: true,
	ssm.ParameterTypeSecureString: true,
	ssm.ParameterTypeStringList: true,
}

func StartNewSSMService(awsSession *session.Session) *ssm.SSM {
	ssmsvc := ssm.New(awsSession, aws.NewConfig().WithRegion(AWS_REGION))

	return ssmsvc
}

func (awsSVC *AWSServices) GetSSMParameter(location string, isDecrypted bool) (string, error) {
	param, err := awsSVC.SSM.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(location),
		WithDecryption: aws.Bool(isDecrypted),
	})

	return *param.Parameter.Value, err
}


//Creates a new parameter in the designated location within the paramName.
//
// Allowed paramTypes are ["String", "SecureString", "StringList"]
func (awsSVC *AWSServices) PutSSMParameter(name, value, paramType string, overWrite bool) error {

	// Validate param type
	if !allowedTypes[paramType] {
		return errors.New("invalid param type entered. Valid: [String, SecureString, StringList]")
	}

	_, err := awsSVC.SSM.PutParameter(&ssm.PutParameterInput{
		Name:  &name,
		Value: &value,
		Type:  &paramType,
		// We'll need to have a merchant id check before we can allow for overwrite requests
		Overwrite: &overWrite,
	})

	return err
}
