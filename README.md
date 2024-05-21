# webapi


## go

```bash
go mod init webapi
go get github.com/labstack/echo/v4
go get github.com/aws/aws-lambda-go/lambda
go get github.com/awslabs/aws-lambda-go-api-proxy/echo
```

## serverlessframework

```bash
npm install -g serverless
serverless create --template aws-go-mod
serverless plugin install -n serverless-go-plugin
```
