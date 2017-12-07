package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["gofamily/controllers:TopicController"] = append(beego.GlobalControllerRouter["gofamily/controllers:TopicController"],
		beego.ControllerComments{
			Method: "Publish",
			Router: `/`,
			AllowHTTPMethods: []string{"post","options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gofamily/controllers:TopicController"] = append(beego.GlobalControllerRouter["gofamily/controllers:TopicController"],
		beego.ControllerComments{
			Method: "GetTopicCommentList",
			Router: `/:tid/comments`,
			AllowHTTPMethods: []string{"get","options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gofamily/controllers:UserController"] = append(beego.GlobalControllerRouter["gofamily/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserDetail",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gofamily/controllers:UserController"] = append(beego.GlobalControllerRouter["gofamily/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserDynamic",
			Router: `/:uid/dynamic`,
			AllowHTTPMethods: []string{"get","options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gofamily/controllers:UserController"] = append(beego.GlobalControllerRouter["gofamily/controllers:UserController"],
		beego.ControllerComments{
			Method: "PublishTopic",
			Router: `/:uid/publish/topic`,
			AllowHTTPMethods: []string{"post","options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gofamily/controllers:UserController"] = append(beego.GlobalControllerRouter["gofamily/controllers:UserController"],
		beego.ControllerComments{
			Method: "PublishTopicComment",
			Router: `/:uid/topic/:tid/comment`,
			AllowHTTPMethods: []string{"post","options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gofamily/controllers:UserController"] = append(beego.GlobalControllerRouter["gofamily/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetUserContacts",
			Router: `/contacts`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gofamily/controllers:UserController"] = append(beego.GlobalControllerRouter["gofamily/controllers:UserController"],
		beego.ControllerComments{
			Method: "LoginUser",
			Router: `/login`,
			AllowHTTPMethods: []string{"post","options"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["gofamily/controllers:UserController"] = append(beego.GlobalControllerRouter["gofamily/controllers:UserController"],
		beego.ControllerComments{
			Method: "RegisterUser",
			Router: `/register`,
			AllowHTTPMethods: []string{"post","options"},
			MethodParams: param.Make(),
			Params: nil})

}
