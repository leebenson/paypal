package paypal

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPaypal(t *testing.T) {
	clientID := os.Getenv("PAYPAL_TEST_CLIENTID")
	if clientID == "" {
		panic("Test Paypal clientID is missing")
	}

	secret := os.Getenv("PAYPAL_TEST_SECRET")
	if secret == "" {
		panic("Test Paypal secret is missing")
	}

	client := NewClient(clientID, secret, APIBaseSandBox)

	Convey("Requesting an access token should returns token response", t, func() {
		tokenResp, err := client.GetAccessToken()

		So(err, ShouldBeNil)
		So(tokenResp.Token, ShouldNotBeBlank)
		So(tokenResp.AppID, ShouldNotBeBlank)
		So(tokenResp.ExpiresIn, ShouldBeGreaterThan, 0)
	})

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
			newPaymentResp, err := client.CreatePayment(payment)

			So(err, ShouldBeNil)
			So(newPaymentResp.Intent, ShouldEqual, PaymentIntentSale)
			So(newPaymentResp.ID, ShouldNotBeNil)

			// This requires manual test as the payer needs to approve inside Paypal
			// Convey("Execute the newly created payment should be successful", func() {
			// 	resp, err := client.ExecutePayment(newPaymentResp.ID, payer.ID, nil)
			//
			// 	So(err, ShouldBeNil)
			// 	So(resp.ID, ShouldEqual, newPaymentResp.ID)
			// 	So(resp.State, ShouldEqual, PaymentStateApproved)
			// })

			Convey("Fetching the newly created payment should return valid results", func() {
				resp, err := client.GetPayment(newPaymentResp.ID)

				So(err, ShouldBeNil)
				So(resp.ID, ShouldEqual, newPaymentResp.ID)
				So(resp.Intent, ShouldEqual, PaymentIntentSale)
				So(resp.Payer.PaymentMethod, ShouldEqual, PaymentMethodCreditCard)
			})

			Convey("List payments should include the newly created payment", func() {
				payments, err := client.ListPayments(map[string]string{
					"count":   "10",
					"sort_by": "create_time",
				})

				So(err, ShouldBeNil)
				So(len(payments), ShouldBeGreaterThan, 0)
				So(payments[0].ID, ShouldEqual, newPaymentResp.ID)
			})
		})

	})

}