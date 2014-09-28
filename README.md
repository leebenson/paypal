# Payment REST API Go client

[![Coverage Status](https://coveralls.io/repos/fundary/paypal/badge.png)](https://coveralls.io/r/fundary/paypal) [![GoDoc](https://godoc.org/github.com/fundary/paypal?status.svg)](https://godoc.org/github.com/fundary/paypal)

This is a client for the Paypal REST API ([https://developer.paypal.com/webapps/developer/docs/api/](https://developer.paypal.com/webapps/developer/docs/api/)

## Goals

- [x] Tests where feasible (some actions requires manual testing)
- [ ] Concurrency safety by utilizing `PayPal-Request-Id`

## Usage

```bash
go get github.com/fundary/paypal
```

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
- [x ] [Payments - Sale transactions](https://developer.paypal.com/webapps/developer/docs/api/#sale-transactions)
- [ ] [Payments - Refunds](https://developer.paypal.com/webapps/developer/docs/api/#refunds)
- [ ] [Payments - Authorizations](https://developer.paypal.com/webapps/developer/docs/api/#authorizations)
- [ ] [Payments - Captures](https://developer.paypal.com/webapps/developer/docs/api/#billing-plans-and-agreements)
- [ ] [Payments - Billing Plans and Agreements](https://developer.paypal.com/webapps/developer/docs/api/#billing-plans-and-agreements)
- [ ] [Payments - Order](https://developer.paypal.com/webapps/developer/docs/api/#orders)
- [ ] [Vault](https://developer.paypal.com/webapps/developer/docs/api/#vault)
- [ ] [Identity](https://developer.paypal.com/webapps/developer/docs/api/#identity)
- [ ] [Invoicing](https://developer.paypal.com/webapps/developer/docs/api/#invoicing)
- [ ] [Payment Experience](https://developer.paypal.com/webapps/developer/docs/api/#payment-experience)
