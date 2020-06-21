#! /bin/bash

# Delete lambda function first
aws lambda delete-function --function-name go-panda-lambda

GOOS=linux go build

zip function.zip go-panda-lambda

aws lambda create-function --function-name go-panda-lambda --runtime go1.x \
  --zip-file fileb://function.zip --handler go-panda-lambda \
  --role arn:aws:iam::417282402508:role/lamba-execution-role \
  --timeout 10 \
  --environment=Variables="{MG_DOMAIN=mg.opeonikute.dev,MG_API_KEY=xxxxxxxxxxxx,MAIL_RECIPIENT=test@yahoo.com,CD_UPLOAD_PRESET=test,CD_CLOUD_NAME=sdsdsd,MONGO_DATABASE=go-panda,MONGO_URL=mongo://localhost}"


# Cleanup
rm -f function.zip go-panda 