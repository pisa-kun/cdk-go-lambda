# cdk-go-lambda project

IaSをgolangで記載して、golangのlambdaをデプロイするサンプル

### 初期化

> mkdir cdk_go_lambda  
> cd cdk_go_lambda

プロジェクト初期化

> cdk init app --language=go

必要なCDKライブラリのインポート

> go mod tidy

**※ソースコード記載後にこのコマンドを実行する**

~~環境変数にアカウントIDとリージョン情報を登録すること。~~  
.envファイルを作成して下記のkeyに値を書き込む
> AWS_REGION=XXX  
> AWS_ACCOUNT=XXX

### Lambda関数の作成


#### 参考
- cdk deployまで

https://qiita.com/yudai2929/items/1ed87643e38002f57ed6

https://zenn.dev/kasega0/articles/31a8b8f9911217

- godotenvの使い方について

https://zenn.dev/a_ichi1/articles/c9f3870350c5e2

#### エラー時のリファレンス

> Another CLI (PID=2044) is currently synthing to cdk.out. Invoke the CLI in sequence, or use '--output' to synth into different directories.

事象: 環境変数に値をセットした後で発生  
対応: cdk.out をディレクトリごと削除

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests
