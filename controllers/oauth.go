package controllers

import (
	"beego_notes/models"
	"context"
	"fmt"
	"net/http"
	"os"

	beego "github.com/beego/beego/v2/server/web"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

type OauthController struct {
	beego.Controller
}

var FACEBOOK = &oauth2.Config{
	// Get your own ClienID and ClientSecret
	ClientID:     os.Getenv("FB_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("FB_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"email"},
	Endpoint:     facebook.Endpoint,
	RedirectURL:  "http://localhost:8080/auth/facebook",
}

func (c *OauthController) FacebookAuth() {
	code := c.GetString("code", "")
	fmt.Println(code)
	token, err := FACEBOOK.Exchange(context.Background(), code)

	if err != nil {
		c.Redirect("/", http.StatusTemporaryRedirect)
	}

	// Create or update user
	user := models.UserUpdateOrCreate(token)

	// Create Login Session
	c.SetSession("user_id", user.Id)

	c.Redirect("/", http.StatusTemporaryRedirect)
}
