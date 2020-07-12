package main

import (
	"context"
	"html/template"
	"io/ioutil"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"github.com/koheiiwamura/golang-statik-gin-sample/functions/oauth/controller"

	// 生成した statik パッケージをimportする
	_ "github.com/koheiiwamura/golang-statik-gin-sample/functions/oauth/statik"
	"github.com/rakyll/statik/fs"
)

func initTemplates() *template.Template {
	// zipdata["default"]の値("/templates/" 配下のデータ)を取得したhttp.FileSystemを生成。
	statikFs, err := fs.New()
	if err != nil {
		panic(err)
	}

	// "/templates/authorize.html"のデータを取得
	f, err := statikFs.Open("/authorize.html")
	if err != nil {
		panic(err)
	}

	// "authorize.html"を読み込む
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	// template.Template を生成
	t, err := template.New("authorize.html").Parse(string(b))
	if err != nil {
		panic(err)
	}

	return t
}

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()
	// ginのHTMLレンダラーに、テンプレートを関連付ける。
	r.SetHTMLTemplate(initTemplates())
	oauthEngine := r.Group("/oauth")
	{
		oauthEngine.GET("/authorize", controller.Authorize)
		oauthEngine.POST("/approve", controller.Approve)

	}
	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
