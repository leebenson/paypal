package paypal

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOrder(t *testing.T) {
	withContext(func(client *Client) {
		Convey("With the orders endpoint", t, func() {
			Convey("Creating a order with valid data should be successful", func() {
				payer := Payer{
					PaymentMethod: PaymentMethodPaypal,
				}
				amountDetail := Details{
					Subtotal: "7.41",
					Tax:      "0.03",
					Shipping: "0.03",
				}
				amount := Amount{
					Total:    "7.47",
					Currency: "USD",
					Details:  &amountDetail,
				}
				transaction := Transaction{
					Amount:      &amount,
					Description: "This is the payment transaction description.",
				}
				order := Payment{
					Intent:       PaymentIntentOrder,
					Payer:        &payer,
					Transactions: []Transaction{transaction},
					RedirectURLs: RedirectURLs{
						ReturnURL: "http://www.return.com",
						CancelURL: "http://www.cancel.com",
					},
				}
				newOrder, err, resp := client.CreatePayment(order)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusCreated)
				So(newOrder.Intent, ShouldEqual, PaymentIntentOrder)
				So(newOrder.ID, ShouldNotBeNil)

				// Require user's approval
				SkipConvey("Retrieving an order should returns valid data", func() {
					order, err, resp := client.GetOrder(newOrder.ID)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
					So(order.ID, ShouldNotEqual, "")
					So(order.State, ShouldEqual, OrderStatePending)
					So(order.PendingReason, ShouldEqual, PendingReasonOrder)

					Convey("Authorizing the order should return a valid authorization object", func() {
						authorization, err, resp := client.AuthorizeOrder(newOrder.ID, &Amount{
							Currency: "USD",
							Total:    "4.54",
						})

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusOK)
						So(authorization.ID, ShouldNotEqual, "")
						So(authorization.State, ShouldEqual, AuthorizationStateAuthorized)
					})

					Convey("Capturing the order should return a valid capture object", func() {
						capture, err, resp := client.CaptureOrder(newOrder.ID, &Amount{
							Currency: "USD",
							Total:    "4.54",
						}, true)

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusOK)
						So(capture.Amount.Total, ShouldEqual, "4.54")
						So(capture.IsFinalCapture, ShouldEqual, true)
						So(capture.State, ShouldEqual, CaptureStatePending)
					})

					Convey("Voiding the order should be successful", func() {
						voidedOrder, err, resp := client.VoidOrder(newOrder.ID)

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusOK)
						So(voidedOrder.State, ShouldEqual, "voided")
					})
				})
			})

		})
	})

}
