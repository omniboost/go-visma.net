package vismanet_test

import (
	"context"
	"log"
	"net/url"
	"os"
	"testing"

	vismanet "github.com/omniboost/go-visma.net"
	"golang.org/x/oauth2"
)

var (
	client *vismanet.Client
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	companyID := os.Getenv("COMPANY_ID")
	applicationType := os.Getenv("APPLICATION_TYPE")
	tokenURL := os.Getenv("TOKEN_URL")
	debug := os.Getenv("DEBUG")

	oauthConfig := vismanet.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
		// Expiry:       time.Now().AddDate(0, 0, 1),
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background(), token)

	client = vismanet.NewClient(httpClient)
	if companyID != "" {
		client.SetCompanyID(companyID)
	}
	if applicationType != "" {
		client.SetApplicationType(applicationType)
	}
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	m.Run()
}
