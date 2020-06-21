package main

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/opeonikute/panda"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context) (bool, error) {
	goPanda := panda.GoPanda{
		Config: panda.Settings{
			MgDomain:       os.Getenv("MG_DOMAIN"),
			MgKey:          os.Getenv("MG_API_KEY"),
			MailRecipients: os.Getenv("MAIL_RECIPIENT"),
			CdCloudName:    os.Getenv("CD_CLOUD_NAME"),
			CdUploadPreset: os.Getenv("CD_UPLOAD_PRESET"),
			MongoURL:       os.Getenv("MONGO_URL"),
			MongoDB:        os.Getenv("MONGO_DATABASE"),
		},
	}
	result := goPanda.Run(0)
	if result != true {
		return false, errors.New("Email was not sent. Image function returned false")
	}
	return true, nil
}
