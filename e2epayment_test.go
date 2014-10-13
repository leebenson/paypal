package paypal

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPayment(t *testing.T) {
	withContext(func(client *Client) {
		Convey("With the payments endpoint", t, func() {
			Convey("Creating a payment with valid data should be successful", func() {
				billingAddress := Address{
					Line1:       "111 First Street",
					City:        "Saratoga",
					State:       "CA",
					PostalCode:  "95070",
					CountryCode: "US",
				}
				creditCard := CreditCard{
					Number:         "4417119669820331",
					Type:           "visa",
					ExpireMonth:    "11",
					ExpireYear:     "2018",
					CVV2:           "874",
					FirstName:      "Betsy",
					LastName:       "Buyer",
					BillingAddress: &billingAddress,
				}
				payer := Payer{
					PaymentMethod: PaymentMethodCreditCard,
					FundingInstruments: []FundingInstrument{
						FundingInstrument{
							CreditCard: &creditCard,
						},
					},
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
				payment := Payment{
					Intent:       PaymentIntentSale,
					Payer:        &payer,
					Transactions: []Transaction{transaction},
				}
				newPaymentResp, err, resp := client.CreatePayment(payment)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusCreated)
				So(newPaymentResp.Intent, ShouldEqual, PaymentIntentSale)
				So(newPaymentResp.ID, ShouldNotBeNil)

				// This requires manual test as the payer needs to approve inside Paypal
				SkipConvey("Execute the newly created payment should be successful", func() {
					_, err, resp := client.ExecutePayment(newPaymentResp.ID, "123", nil)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
				})

				Convey("Fetching the newly created payment should return valid results", func() {
					payment, err, resp := client.GetPayment(newPaymentResp.ID)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
					So(payment.ID, ShouldEqual, newPaymentResp.ID)
					So(payment.Intent, ShouldEqual, PaymentIntentSale)
					So(payment.Payer.PaymentMethod, ShouldEqual, PaymentMethodCreditCard)
					So(payment.Transactions[0].RelatedResources[0], ShouldNotBeNil)

					Convey("With the sale endpoints", func() {
						Convey("Fetching an existing sale should return valid data", func() {
							sale, err, resp := client.GetSale(newPaymentResp.Transactions[0].RelatedResources[0].Sale.ID)

							So(err, ShouldBeNil)
							So(resp.StatusCode, ShouldEqual, http.StatusOK)
							So(sale.ID, ShouldEqual, newPaymentResp.Transactions[0].RelatedResources[0].Sale.ID)

							// Cannot test refund as it require that a payment is approved and completed
							SkipConvey("A partial refund for an existing sale should be successful", func() {
								amount := Amount{
									Total:    "2.34",
									Currency: "USD",
								}

								refund, err, resp := client.RefundSale(sale.ID, &amount)

								So(err, ShouldBeNil)
								So(resp.StatusCode, ShouldEqual, http.StatusOK)

								refundedSale, err, resp := client.GetSale(newPaymentResp.Transactions[0].RelatedResources[0].Sale.ID)

								So(err, ShouldBeNil)
								So(resp.StatusCode, ShouldEqual, http.StatusOK)
								So(refund.Amount.Total, ShouldEqual, amount.Total)
								So(refund.ParentPayment, ShouldEqual, payment.ID)
								So(refundedSale.State, ShouldEqual, SaleStatePartiallyRefunded)
								So(refundedSale.Amount.Total, ShouldEqual, "5.13")

								Convey("Retrieving the new refund should return valid results", func() {
									newRefund, err, resp := client.GetRefund(refund.ID)

									So(err, ShouldBeNil)
									So(resp.StatusCode, ShouldEqual, http.StatusOK)
									So(newRefund.ID, ShouldEqual, refund.ID)
									So(newRefund.Amount, ShouldResemble, refund.Amount)

								})
							})
						})

					})
				})

				Convey("List payments should include the newly created payment", func() {
					payments, err, resp := client.ListPayments(map[string]string{
						"count":   "10",
						"sort_by": "create_time",
					})

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
					So(len(payments), ShouldBeGreaterThan, 0)
					So(payments[0].ID, ShouldEqual, newPaymentResp.ID)
				})

				Convey("Authorize a new payment should be successful", func() {
					payment.Intent = PaymentIntentAuthorize

					authorizedPayment, err, resp := client.CreatePayment(payment)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusCreated)
					So(authorizedPayment.Intent, ShouldEqual, PaymentIntentAuthorize)
					So(authorizedPayment.Transactions[0].RelatedResources[0].Authorization.ID, ShouldNotEqual, "")

					authID := authorizedPayment.Transactions[0].RelatedResources[0].Authorization.ID

					Convey("Looking up the payment's authorization should return valid data", func() {
						authorization, err, resp := client.GetAuthorization(authID)

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusOK)
						So(authorization.Amount.Total, ShouldEqual, "7.47")
					})

					Convey("Capturing the authorization should be successful", func() {
						capture, err, resp := client.CaptureAuthorization(authID, &Amount{
							Currency: "USD",
							Total:    "4.54",
						}, true)

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusOK)
						So(capture.Amount.Total, ShouldEqual, "4.54")
						So(capture.IsFinalCapture, ShouldEqual, true)
						So(capture.State, ShouldEqual, CaptureStatePending)

						Convey("Retrieving the new capture should returns valid data", func() {
							newCapture, err, resp := client.GetCapture(capture.ID)

							So(err, ShouldBeNil)
							So(resp.StatusCode, ShouldEqual, http.StatusOK)
							So(newCapture.Amount.Total, ShouldEqual, capture.Amount.Total)
							So(newCapture.State, ShouldEqual, CaptureStateAuthorized)
						})

						// Cannot refund
						SkipConvey("Refunding the new capture should returns a valid refund object", func() {
							refund, err, resp := client.RefundCapture(capture.ID, &Amount{
								Currency: "USD",
								Total:    "4.54",
							})

							So(err, ShouldBeNil)
							So(resp.StatusCode, ShouldEqual, http.StatusOK)
							So(refund.ID, ShouldNotEqual, "")
							So(refund.State, ShouldEqual, RefundStatePending)
							So(refund.ParentPayment, ShouldEqual, authorizedPayment.ID)
						})
					})

					Convey("Voiding an authorization should be successful", func() {
						voidedAuthorization, err, resp := client.VoidAuthorization(authID)

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusOK)
						So(voidedAuthorization.ID, ShouldEqual, authID)
						So(voidedAuthorization.State, ShouldEqual, AuthorizationStateVoided)
					})

					// TODO: Add test for reauthorize payments
				})
			})

		})
	})

}
