package main

import (
	"os"
	"os/exec"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"github.com/joho/godotenv"
)

type CdkGoLambdaStackProps struct {
	awscdk.StackProps
}

func NewCdkGoLambdaStack(scope constructs.Construct, id string, props *CdkGoLambdaStackProps) (awscdk.Stack, error) {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Gol のビルド
	buildCmd := exec.Command("go", "build", "-tags", " lambda.norpc", "-o", "bin/bootstrap", "handler/main.go")
	buildCmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
	if err := buildCmd.Run(); err != nil {
		return nil, err
	}

	// Lambda関数の定義
	awslambda.NewFunction(stack, jsii.String("golang-lambda"), &awslambda.FunctionProps{
		FunctionName: jsii.String("golang-lambda-function"),
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Code:         awslambda.Code_FromAsset(jsii.String("bin"), nil),
		Handler:      jsii.String("main"),
	})

	return stack, nil
}

func main() {
	defer jsii.Close()

	// envファイルの読み込み
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	app := awscdk.NewApp(nil)

	_, err := NewCdkGoLambdaStack(app, "CdkGoLambdaStack", &CdkGoLambdaStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	if err != nil {
		panic(err)
	}

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("AWS_ACCOUNT")),
		Region:  jsii.String(os.Getenv("AWS_REGION")),
	}
}
