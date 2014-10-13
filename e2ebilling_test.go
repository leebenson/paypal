package paypal

import (
	"net/http"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBilling(t *testing.T) {
	withContext(func(client *Client) {
		Convey("With the billing plan endpoint", t, func() {
			Convey("Creating a billing plan with valid data should be successful", func() {
				plan := &Plan{
					Name:        "T-Shirt of the Month Club Plan",
					Description: "Template creation.",
					Type:        PlanTypeFixed,
					PaymentDefinitions: []PaymentDefinition{
						PaymentDefinition{
							Name:              "Regular Payments",
							Type:              PaymentDefinitionTypeRegular,
							Frequency:         "MONTH",
							FrequencyInterval: "2",
							Amount: &Currency{
								Value:    "100",
								Currency: "USD",
							},
							Cycles: "12",
							ChargeModels: []ChargeModels{
								ChargeModels{
									Type: ChargeModelsTypeShipping,
									Amount: &Currency{
										Value:    "10",
										Currency: "USD",
									},
								},
								ChargeModels{
									Type: ChargeModelsTypeTax,
									Amount: &Currency{
										Value:    "12",
										Currency: "USD",
									},
								},
							},
						},
					},
					MerchantPreferences: &MerchantPreferences{
						SetupFee: &Currency{
							Value:    "1",
							Currency: "USD",
						},
						ReturnURL:               "http://www.return.com",
						CancelURL:               "http://www.cancel.com",
						AutoBillAmount:          "YES",
						InitialFailAmountAction: "CONTINUE",
						MaxFailAttempts:         "0",
					},
				}

				newPlan, err, resp := client.CreateBillingPlan(plan)

				So(err, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, http.StatusCreated)
				So(newPlan.ID, ShouldNotEqual, "")
				So(newPlan.Name, ShouldEqual, plan.Name)
				So(newPlan.Description, ShouldEqual, plan.Description)

				Convey("Updating the billing plan should be successful", func() {
					err, resp := client.UpdateBillingPlan(newPlan.ID, &PatchPlan{
						State: PlanStateActive,
					})

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)

					Convey("Retrieving the updated plan should be successful", func() {
						plan, err, resp := client.GetBillingPlan(newPlan.ID)

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusOK)
						So(plan.State, ShouldEqual, PlanStateActive)
					})

					Convey("Creating a billing agreement with valid data should be successful", func() {
						startDate := DatetimeRFC3339{time.Now().Add(24 * time.Hour)}
						agreement := &Agreement{
							Name:        "T-Shirt of the Month Club Agreement",
							Description: "Agreement for T-Shirt of the Month Club Plan",
							StartDate:   &startDate,
							Plan: &PatchPlan{
								ID: newPlan.ID,
							},
							Payer: &AgreementPayer{
								PaymentMethod: PaymentMethodPaypal,
							},
							ShippingAddress: &Address{
								Line1:       "111 First Street",
								City:        "Saratoga",
								State:       "CA",
								PostalCode:  "95070",
								CountryCode: "US",
							},
						}

						newAgreement, err, resp := client.CreateAgreement(agreement)

						So(err, ShouldBeNil)
						So(resp.StatusCode, ShouldEqual, http.StatusCreated)
						So(newAgreement.Name, ShouldEqual, agreement.Name)
						So(newAgreement.Description, ShouldEqual, agreement.Description)

						SkipConvey("Executing the new agreement should be successful", func() {
							_, err, resp := client.ExecuteAgreement("123")

							So(err, ShouldBeNil)
							So(resp.StatusCode, ShouldEqual, http.StatusCreated)
						})
					})

				})

				Convey("Listing billing plans shold return valid data", func() {
					plans, err, resp := client.ListBillingPlans(nil)

					So(err, ShouldBeNil)
					So(resp.StatusCode, ShouldEqual, http.StatusOK)
					So(len(plans), ShouldBeGreaterThan, 0)
				})

			})

		})
	})

}
