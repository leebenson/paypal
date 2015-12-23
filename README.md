# Payment REST API Go client
 [![GoDoc](https://godoc.org/github.com/leebenson/paypal?status.svg)](https://godoc.org/github.com/leebenson/paypal)

This is a client for the Paypal REST API ([https://developer.paypal.com/webapps/developer/docs/api/](https://developer.paypal.com/webapps/developer/docs/api/)

## Goals

- [x] Automated tests that don't require manual approval in Paypal account
- [ ] Automated tests that require manual approval in a Paypal account (with a different build tag, eg. `PAYPAL_APPROVED_PAYMENT_ID`
- [ ] Concurrency safety by utilizing `PayPal-Request-Id`

## Usage

```bash
go get github.com/leebenson/paypal
```

Import into your app and start using it:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/leebenson/paypal"
)

func main() {
	clientID := os.Getenv("PAYPAL_CLIENTID")
	if clientID == "" {
		panic("PayPal clientID is missing")
	}

	secret := os.Getenv("PAYPAL_SECRET")
	if secret == "" {
		panic("PayPal secret is missing")
	}

	client := paypal.NewClient(clientID, secret, paypal.APIBaseLive)

	payments, err := client.ListPayments(map[string]string{
		"count":   "10",
		"sort_by": "create_time",
	})
	if err != nil {
		log.Fatal("Could not retrieve payments: ", err)
	}

	fmt.Println(payments)
}
```

## Run tests

This library use [Goconvey](http://goconvey.co/) for tests, so to run them, start Goconvey:

```
PAYPAL_TEST_CLIENTID=[Paypal Client ID] PAYPAL_TEST_SECRET=[Paypal Secret] goconvey
```

Or you can just use `go test`

```
PAYPAL_TEST_CLIENTID=[Paypal Client ID] PAYPAL_TEST_SECRET=[Paypal Secret] go test
```

## Roadmap

- [x] [Payments - Payment](https://developer.paypal.com/webapps/developer/docs/api/#payments)
- [x] [Payments - Sale transactions](https://developer.paypal.com/webapps/developer/docs/api/#sale-transactions)
- [x] [Payments - Refunds](https://developer.paypal.com/webapps/developer/docs/api/#refunds)
- [x] [Payments - Authorizations](https://developer.paypal.com/webapps/developer/docs/api/#authorizations)
- [x] [Payments - Captures](https://developer.paypal.com/webapps/developer/docs/api/#billing-plans-and-agreements)
- [ ] [Payments - Billing Plans and Agreements](https://developer.paypal.com/webapps/developer/docs/api/#billing-plans-and-agreements)
- [ ] [Payments - Order](https://developer.paypal.com/webapps/developer/docs/api/#orders)
- [ ] [Vault](https://developer.paypal.com/webapps/developer/docs/api/#vault)
- [ ] [Identity](https://developer.paypal.com/webapps/developer/docs/api/#identity)
- [ ] [Invoicing](https://developer.paypal.com/webapps/developer/docs/api/#invoicing)
- [ ] [Payment Experience](https://developer.paypal.com/webapps/developer/docs/api/#payment-experience)
