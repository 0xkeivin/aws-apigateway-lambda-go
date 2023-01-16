# aws-apigateway-lambda-go
Purpose of this repo is to demo the usage of AWS API Gateway and Lambda with Go.
# Setup
Terraform is used to provision the required AWS resources:
- terraform v1.3.7

```bash
export AWS_ACCESS_KEY_ID="xxx"
export AWS_SECRET_ACCESS_KEY="yyy"
```

# objectives
1. [x] - return a status 400 if the "authentication" header is malformed or missing in the following lambda code 
2. [x] - return a status 403 if "token" is invalid or empty
3. [x] - Fix Query user notes
4. [x] - Limit user notes to 10