package router

import (
	"github.com/zc2638/swag"
	"github.com/zc2638/swag/option"
)

func CreateSwagger() *swag.API {
	api := swag.New(
		option.Title("CVBuilder API Doc"),
	// 	option.Security("petstore_auth", "read:pets"),
	// 	option.SecurityScheme("petstore_auth",
	// 		option.OAuth2Security("accessCode", "http://example.com/oauth/authorize", "http://example.com/oauth/token"),
	// 		option.OAuth2Scope("write:pets", "modify pets in your account"),
	// 		option.OAuth2Scope("read:pets", "read your pets"),
	// 	),
	)
	return api
}
