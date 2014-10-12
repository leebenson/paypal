package paypal

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNotifications(t *testing.T) {
	withContext(func(client *Client) {
		Convey("With the notifications endpoint", t, func() {
			Convey("Listing event types should return valid data", func() {
				eventTypes, err, resp := client.ListWebhookEventTypes()

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
				So(len(eventTypes), ShouldEqual, 6)
				So(eventTypes[0].Name, ShouldEqual, "PAYMENT.AUTHORIZATION.CREATED")
				So(eventTypes[0].Description, ShouldEqual, "A payment authorization was created")
			})

			// Skipping as endpoint requires a valid webhook URL
			SkipConvey("Creating a webhook with valid data should be successful", func() {
				w := &Webhook{
					URL: "http://www.yeowza.com/paypal_webhook",
					EventTypes: []EventType{
						EventType{Name: "PAYMENT.AUTHORIZATION.CREATED"},
						EventType{Name: "PAYMENT.AUTHORIZATION.VOIDED"},
					},
				}

				newWebhook, err, resp := client.CreateWebhook(w)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusCreated)
				So(newWebhook.URL, ShouldEqual, w.URL)
				So(len(newWebhook.EventTypes), ShouldEqual, len(w.EventTypes))
				So(newWebhook.EventTypes[0].Name, ShouldEqual, w.EventTypes[0].Name)
				So(newWebhook.EventTypes[0].Description, ShouldEqual, "A payment authorization was created")
				So(newWebhook.ID, ShouldNotBeNil)

				Convey("Retrieving the newly created webhook should be successful", func() {
					w, err, resp := client.GetWebhook(newWebhook.ID)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusCreated)
					So(newWebhook.URL, ShouldEqual, w.URL)
				})
			})

			Convey("Listing all webhooks should return valid data", func() {
				_, err, resp := client.ListWebhooks()

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusOK)
			})
		})
	})
}
