package paypal

import (
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestVault(t *testing.T) {
	withContext(func(client *Client) {
		Convey("With the vault endpoint", t, func() {

			Convey("Storing a credit card with valid data should be successful", func() {
				creditCard := CreditCard{
					PayerID:     "user12345",
					Type:        CreditCardTypeVisa,
					Number:      "4417119669820331",
					ExpireMonth: "11",
					ExpireYear:  "2018",
					FirstName:   "Betsy",
					LastName:    "Buyer",
				}

				newCreditCard, err, resp := client.StoreCreditCard(&creditCard)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusCreated)
				So(newCreditCard.ID, ShouldNotBeNil)
				So(newCreditCard.Number, ShouldEqual, "xxxxxxxxxxxx0331")

				Convey("Retrieving the stored credit card should be successful", func() {
					creditCard, err, resp := client.GetStoredCreditCard(newCreditCard.ID)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
					So(creditCard, ShouldResemble, newCreditCard)

					// Skip this test for now as it seems the endpoint does not behave as specified in documentation
					SkipConvey("Updating the stored credit card should updates the data", func() {
						newCreditCard.FirstName = "Carol"
						creditCard, err, resp := client.UpdateStoredCreditCard(newCreditCard.ID, newCreditCard)

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusOK)
						So(creditCard.FirstName, ShouldEqual, newCreditCard.FirstName)
					})

					Convey("Deleting the stored credit card should be successful", func() {
						err, resp := client.DeleteStoredCreditCard(newCreditCard.ID)

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusNoContent)

						Convey("Retrieving the deleted credit card should not return any data", func() {
							creditCard, err, resp := client.GetStoredCreditCard(newCreditCard.ID)

							So(resp.StatusCode, ShouldEqual, http.StatusNotFound)
							So(err, ShouldNotBeNil)
							So(creditCard, ShouldEqual, nil)
						})
					})
				})

			})
		})
	})
}
