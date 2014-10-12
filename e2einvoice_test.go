package paypal

import (
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInvoice(t *testing.T) {
	withContext(func(client *Client) {
		Convey("With the invoice endpoint", t, func() {
			Convey("Creating a invoice with valid data should be successful", func() {
				invoice := &Invoice{

					MerchantInfo: &MerchantInfo{
						Email:        "lee-facilitator@fundary.com",
						FirstName:    "Dennis",
						LastName:     "Doctor",
						BusinessName: "Medical Professionals, LLC",
						Phone: &Phone{
							CountryCode:    "001",
							NationalNumber: "5032141716",
						},
						Address: &Address{
							Line1:       "1234 Main St.",
							City:        "Portland",
							State:       "OR",
							PostalCode:  "97217",
							CountryCode: "US",
						},
					},

					BillingInfo: []BillingInfo{
						BillingInfo{
							Email: "example@example.com",
						},
					},

					Items: []InvoiceItem{
						InvoiceItem{
							Name:     "Sutures",
							Quantity: 100,
							UnitPrice: &Currency{
								Currency: "USD",
								Value:    "5.00",
							},
						},
					},

					Note: "Medical Invoice 16 Jul, 2013 PST",

					PaymentTerm: &PaymentTerm{
						TermType: PaymentTermTypeNet45,
					},

					ShippingInfo: &ShippingInfo{
						FirstName:    "Sally",
						LastName:     "Patient",
						BusinessName: "Not applicable",
						Address: &Address{
							Line1:       "1234 Broad St.",
							City:        "Portland",
							State:       "OR",
							PostalCode:  "97216",
							CountryCode: "US",
						},
					},
				}

				newInvoice, err, resp := client.CreateInvoice(invoice)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusCreated)
				So(newInvoice.ID, ShouldNotEqual, "")
				So(newInvoice.Status, ShouldEqual, InvoiceStatusDraft)
				So(newInvoice.Number, ShouldNotEqual, "")
				So(newInvoice.Items[0].UnitPrice.Value, ShouldEqual, "5.00")
				So(newInvoice.PaymentTerm.TermType, ShouldEqual, PaymentTermTypeNet45)

				Convey("Sending the new invoice to the payer should be successful", func() {
					err, resp := client.SendInvoice(newInvoice.ID)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusAccepted)
				})

				// Currently paypal returns 500 when updating invoice
				SkipConvey("Updating the new invoice should be successful", func() {
					newInvoice.Items[0].UnitPrice.Value = "250.00"
					newInvoice.PaymentTerm.TermType = PaymentTermTypeNoDueDate

					updatedInvoice, err, resp := client.UpdateInvoice(newInvoice.ID, newInvoice)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
					So(updatedInvoice.Items[0].UnitPrice.Value, ShouldEqual, "250.00")
					So(updatedInvoice.PaymentTerm.TermType, ShouldEqual, PaymentTermTypeNoDueDate)

					Convey("Retrieving the updated invoice should be successful", func() {
						newInvoice, err, resp := client.GetInvoice(newInvoice.ID)

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusOK)
						So(newInvoice.Items[0].UnitPrice.Value, ShouldEqual, "250.00")
						So(newInvoice.ID, ShouldEqual, updatedInvoice.ID)
					})
				})

				Convey("Listing invoices should be successful", func() {
					invoices, err, resp := client.ListInvoices(nil)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
					So(len(invoices), ShouldBeGreaterThan, 0)
				})

				Convey("Searching for invoices should be successful", func() {
					invoices, err, resp := client.SearchInvoices(&InvoiceSearch{
						Email: "example@example.com",
					})

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
					So(len(invoices), ShouldBeGreaterThan, 0)
				})

				// Require invoices with status "SENT"
				SkipConvey("Sending an invoice reminder should be successful", func() {
					notification := &Notification{
						Subject:        "Past due",
						Note:           "Please pay soon",
						SendToMerchant: true,
					}

					err, resp := client.SendInvoiceReminder(newInvoice.ID, notification)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusAccepted)
				})

				Convey("Retrieving QR code for the new invoice should be successful", func() {
					img, err, resp := client.GetInvoiceQRCode(newInvoice.ID, 150, 150)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
					So(img, ShouldNotEqual, "")
				})

				// Require status SENT
				SkipConvey("Recording a payment for the new invoice should be successful", func() {
					err, resp := client.RecordInvoicePayment(newInvoice.ID, PaymentDetailMethodCash, &Datetime{time.Now()}, "Cash received.")

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
				})

				// Require status PAID
				SkipConvey("Recording a refund for the new invoice should be successful", func() {
					err, resp := client.RecordInvoiceRefund(newInvoice.ID, &Datetime{time.Now()}, "Refund provided by cash.")

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
				})

				// Require status SENT
				SkipConvey("Cancelling the new invoice should be successful", func() {
					err, resp := client.CancelInvoice(newInvoice.ID, &Notification{
						Subject:        "Past due",
						Note:           "Canceling invoice",
						SendToMerchant: true,
						SendToPayer:    true,
					})

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusAccepted)
				})

				Convey("Deleting the new invoice should be successful", func() {
					err, resp := client.DeleteInvoice(newInvoice.ID)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusNoContent)
				})
			})

		})
	})

}
