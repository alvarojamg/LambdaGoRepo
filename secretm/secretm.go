package secretm

import (
	"encoding/json"
	"fmt"
	"main/awsgo"
	"main/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(nameSecretet string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println(" > Ask secret" + nameSecretet)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	result, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecretet),
	})

	if err != nil {
		fmt.Println(err.Error())

		return secretData, err
	}

	json.Unmarshal([]byte(*result.SecretString), &secretData)

	fmt.Println("> Secret loading ok" + nameSecretet)

	return secretData, nil

}
