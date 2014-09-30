package paypal

import (
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
					payment, err := client.GetPayment(newPaymentResp.ID)

					So(err, ShouldBeNil)
					So(payment.ID, ShouldEqual, newPaymentResp.ID)
					So(payment.Intent, ShouldEqual, PaymentIntentSale)
					So(payment.Payer.PaymentMethod, ShouldEqual, PaymentMethodCreditCard)
					So(payment.Transactions[0].RelatedResources[0], ShouldNotBeNil)

					Convey("With the sale endpoints", func() {
						Convey("Fetching an existing sale should return valid data", func() {
							sale, err := client.GetSale(newPaymentResp.Transactions[0].RelatedResources[0].Sale.ID)

							So(err, ShouldBeNil)
							So(sale.ID, ShouldEqual, newPaymentResp.Transactions[0].RelatedResources[0].Sale.ID)

							// Cannot test refund as it require that a payment is approved and completed
							// Convey("A partial refund for an existing sale should be successful", func() {
							// 	amount := Amount{
							// 		Total:    "2.34",
							// 		Currency: "USD",
							// 	}
							//
							// 	refund, err := client.RefundSale(sale.ID, &amount)
							// 	So(err, ShouldBeNil)
							//
							// 	refundedSale, err := client.GetSale(newPaymentResp.Transactions[0].RelatedResources[0].Sale.ID)
							//
							// 	So(err, ShouldBeNil)
							// 	So(refund.Amount.Total, ShouldEqual, amount.Total)
							// 	So(refund.ParentPayment, ShouldEqual, payment.ID)
							// 	So(refundedSale.State, ShouldEqual, SaleStatePartiallyRefunded)
							// 	So(refundedSale.Amount.Total, ShouldEqual, "5.13")
							//
							// 	Convey("Retrieving the new refund should return valid results", func() {
							// 		newRefund, err := client.GetRefund(refund.ID)
							//
							// 		So(err, ShouldBeNil)
							// 		So(newRefund.ID, ShouldEqual, refund.ID)
							// 		So(newRefund.Amount, ShouldResemble, refund.Amount)
							//
							// 	})
							// })
						})

					})
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
	})

}
