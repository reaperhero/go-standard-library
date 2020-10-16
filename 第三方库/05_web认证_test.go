package third

import (
	//"testing"
	//"github.com/abbot/go-http-auth"
	//"github.com/bradrydzewski/go.auth"
)
// 参考地址：https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/14.4.md

//三个方面的认证
//
//HTTP Basic和 HTTP Digest认证
//第三方集成认证：QQ、微博、豆瓣、OPENID、google、GitHub、facebook和twitter等
//自定义的用户登录、注册、登出，一般都是基于session、cookie认证



//HTTP Basic 认证 （HTTP Digest认证同理）
//func Secret(user, realm string) string {
//	if user == "john" {
//		// password is "hello"
//		return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1"
//	}
//	return ""
//}
//
//func (this *MainController) Prepare() {
//	a := auth.NewBasicAuthenticator("example.com", Secret)
//	if username := a.CheckAuth(this.Ctx.Request); username == "" {
//		a.RequireAuth(this.Ctx.ResponseWriter, this.Ctx.Request)
//	}
//}
//
//func (this *MainController) Get() {
//	this.Data["Username"] = "astaxie"
//	this.Data["Email"] = "astaxie@gmail.com"
//	this.TplNames = "index.tpl"
//}




//oauth和oauth2的认证
//package controllers
//
//import (
//	"github.com/astaxie/beego"
//	"github.com/bradrydzewski/go.auth"
//)
//
//const (
//	githubClientKey = "a0864ea791ce7e7bd0df"
//	githubSecretKey = "a0ec09a647a688a64a28f6190b5a0d2705df56ca"
//)
//
//type GithubController struct {
//	beego.Controller
//}
//
//func (this *GithubController) Get() {
//	// set the auth parameters
//	auth.Config.CookieSecret = []byte("7H9xiimk2QdTdYI7rDddfJeV")
//	auth.Config.LoginSuccessRedirect = "/mainpage"
//	auth.Config.CookieSecure = false
//
//	githubHandler := auth.Github(githubClientKey, githubSecretKey)
//
//	githubHandler.ServeHTTP(this.Ctx.ResponseWriter, this.Ctx.Request)
//}




//自定义认证