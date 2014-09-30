package paypal

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var testClient *Client
var done = make(chan bool)

func init() {
	getTestClient()
}

func getTestClient() *Client {
	if testClient == nil {

		clientID := os.Getenv("PAYPAL_TEST_CLIENTID")
		if clientID == "" {
			panic("Test Paypal clientID is missing")
		}

		secret := os.Getenv("PAYPAL_TEST_SECRET")
		if secret == "" {
			panic("Test Paypal secret is missing")
		}

		testClient = NewClient(clientID, secret, APIBaseSandBox)
		close(done)
	}

	return testClient
}

func TestAuth(t *testing.T) {
	client := getTestClient()

	Convey("Requesting an access token should returns token response", t, func() {
		tokenResp, err := client.GetAccessToken()

		So(err, ShouldBeNil)
		So(tokenResp.Token, ShouldNotBeBlank)
		So(tokenResp.AppID, ShouldNotBeBlank)
		So(tokenResp.ExpiresIn, ShouldBeGreaterThan, 0)
	})
}

func withContext(fn func(c *Client)) {
	for {
		_, ok := <-done
		if !ok {
			break
		}
	}
	fn(getTestClient())
}