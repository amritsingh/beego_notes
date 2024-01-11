package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"golang.org/x/oauth2"
)

type OauthUser struct {
	Id           uint64
	UserId       uint64
	Provider     string // Like Facebook, Google, Twitter etc.
	Uuid         string
	Name         string
	Email        string
	AccessToken  string
	ExpiresAt    time.Time
	TokenType    string
	RefreshToken string
	ProfilePic   string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *OauthUser) TableName() string {
	return "oauth_users"
}

func SaveFBOauthDetails(token *oauth2.Token) *OauthUser {
	// Get the details of the user
	resp, _ := http.Get("https://graph.facebook.com/me?access_token=" +
		url.QueryEscape(token.AccessToken) +
		"&fields=email,name")

	defer resp.Body.Close()
	me := map[string]interface{}{}

	err := json.NewDecoder(resp.Body).Decode(&me)
	if err != nil {
		fmt.Println("json decode error", "error", err)
	}

	uuid := me["id"].(string)
	email := me["email"].(string)
	name := me["name"].(string)
	provider := "Facebook"

	// Get Profile pic
	// TODO
	profilePic := getFBProfilePic(uuid, token.AccessToken)

	o := orm.NewOrm()
	var oauthUser OauthUser
	o.QueryTable(new(OauthUser)).Filter("email", email).Filter("provider", provider).One(&oauthUser)

	if oauthUser.Id != 0 {
		// Update the oauth user
		o.Update(&OauthUser{Name: name, ProfilePic: profilePic,
			AccessToken: token.AccessToken,
			ExpiresAt:   token.Expiry, TokenType: token.TokenType,
			RefreshToken: token.RefreshToken,
			UpdatedAt:    time.Now()})
	} else {
		// Create new entry
		oauthUser = OauthUser{Provider: provider, Email: email, Uuid: uuid,
			UserId: 0, // There is no user yet
			Name:   name, ProfilePic: profilePic,
			AccessToken: token.AccessToken,
			ExpiresAt:   token.Expiry, TokenType: token.TokenType,
			RefreshToken: token.RefreshToken,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now()}
		o.Insert(&oauthUser)
	}
	return &oauthUser
}

func (oauthUser *OauthUser) UpdateUserID(user *User) {
	o := orm.NewOrm()
	oauthUser.UserId = user.Id
	o.Update(oauthUser)
}

func getFBProfilePic(uuid string, accessToken string) string {
	profilePicApi := "https://graph.facebook.com/" + uuid +
		"?fields=picture.width(720).height(720)&access_token=" + accessToken

	resp, _ := http.Get(profilePicApi)
	defer resp.Body.Close()
	picResp := map[string]interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&picResp); err != nil {
		fmt.Println("json decode error", err)
	}
	return picResp["picture"].(map[string]interface{})["data"].(map[string]interface{})["url"].(string)
}
