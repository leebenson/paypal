package paypal

import (
	"fmt"
	"net/http"
	"time"
)

// https://developer.paypal.com/webapps/developer/docs/api/#invoicing

// CreateInvoice creates an invoice in draft state. After an invoice is created
// with items array, it can be sent.
func (c *Client) CreateInvoice(i *Invoice) (*Invoice, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/invoicing/invoices", c.APIBase), i)
	if err != nil {
		return nil, err, nil
	}

	v := &Invoice{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// SendInvoice sends an invoice to the payer. An invoice cannot be sent unless it
// includes the item array
func (c *Client) SendInvoice(invoiceID string) (error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/invoicing/invoices/%s/send", c.APIBase, invoiceID), nil)
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return err, resp
	}

	return nil, resp
}

// UpdateInvoice updates an invoic
func (c *Client) UpdateInvoice(i *Invoice) (*Invoice, error, *http.Response) {
	req, err := NewRequest("PUT", fmt.Sprintf("%s/invoicing/invoices/%s", c.APIBase, i.ID), i)
	if err != nil {
		return nil, err, nil
	}

	v := &Invoice{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// GetInvoice returns the specific invoice
func (c *Client) GetInvoice(invoiceID string) (*Invoice, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/invoicing/invoices/%s", c.APIBase, invoiceID), nil)
	if err != nil {
		return nil, err, nil
	}

	v := &Invoice{}

	resp, err := c.SendWithAuth(req, v)
	if err != nil {
		return nil, err, resp
	}

	return v, nil, resp
}

// ListInvoices returns invoices that belong to the merchant who makes the call
func (c *Client) ListInvoices(filter map[string]string) ([]Invoice, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/invoicing/invoices", c.APIBase), nil)
	if err != nil {
		return nil, err, nil
	}

	if filter != nil {
		q := req.URL.Query()

		for k, v := range filter {
			q.Set(k, v)
		}

		req.URL.RawQuery = q.Encode()
	}

	var v struct {
		Invoices []Invoice `json:"invoices"`
	}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err, resp
	}

	return v.Invoices, nil, resp
}

// SearchInvoices returns invoices that match the specificed criteria
func (c *Client) SearchInvoices(s *Search) ([]Invoice, error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/invoicing/search", c.APIBase), s)
	if err != nil {
		return nil, err, nil
	}

	var v struct {
		Invoices []Invoice `json:"invoices"`
	}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return nil, err, resp
	}

	return v.Invoices, nil, resp
}

// SendInvoiceReminder sends a reminder that a payment is due for an existing invoice.
func (c *Client) SendInvoiceReminder(invoiceID string, n *Notification) (error, *http.Response) {
	// Do not pass in send_to_payer param
	if n != nil {
		n.SendToPayer = false
	}
	req, err := NewRequest("POST", fmt.Sprintf("%s/invoicing/invoices/%s/remind", c.APIBase, invoiceID), n)
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return err, resp
	}

	return nil, resp
}

// CancelInvoice cancels an invoice and (optionally) notifies the payer of the cancellation.
func (c *Client) CancelInvoice(invoiceID string, n *Notification) (error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/invoicing/invoices/%s/cancel", c.APIBase, invoiceID), n)
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return err, resp
	}

	return nil, resp
}

// DeleteInvoice deletes a draft invoice. Note that this call works for invoices in the draft state only.
// For invoices that have already been sent, it can be cancelled instead.. Once a draft invoice is
// deleted, it can no longer be used or retrieved, but its invoice number can be reuse.
func (c *Client) DeleteInvoice(invoiceID string) (error, *http.Response) {
	req, err := NewRequest("DELETE", fmt.Sprintf("%s/invoicing/invoices/%s", c.APIBase, invoiceID), nil)
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return err, resp
	}

	return nil, resp
}

// GetInvoiceQRCode returns a QR code as PNG image, in base-64 encoded format. Before getting a QR code,
// an invoice must be created. It is recommended that to specify qrinvoice@paypal.com as the recipient
// email address in the billing_info object. (Use a customer email address only if you want the invoice
// to be emailed.). After the invoice has been created, it must be sent. This step is necessary to move
// the invoice from a draft state to a payable state. As stated above, if you specify
// qrinvoice@paypal.com as the recipient email address, the invoice will not be emailed.
func (c *Client) GetInvoiceQRCode(invoiceID string, width, height int) (string, error, *http.Response) {
	req, err := NewRequest("GET", fmt.Sprintf("%s/invoicing/invoices/%s/qr-code?width=%s&height%s", c.APIBase, invoiceID, width, height), nil)
	if err != nil {
		return "", err, nil
	}

	var v struct {
		Image string `json:"image"`
	}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return "", err, resp
	}

	return v.Image, nil, resp
}

// RecordInvoicePayment marks an invoice as paid
func (c *Client) RecordInvoicePayment(invoiceID string, method PaymentDetailMethod, date *time.Time, note string) (error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/invoicing/invoices/%s/record-payment", c.APIBase, invoiceID), &struct {
		Method PaymentDetailMethod `json:"method"`
		Date   *time.Time          `json:"date"`
		Note   string              `json:"note"`
	}{
		Method: method,
		Date:   date,
		Note:   note,
	})
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return err, resp
	}

	return nil, resp
}

// RecordInvoiceRefund marks an invoice as refunded
func (c *Client) RecordInvoiceRefund(invoiceID string, date *time.Time, note string) (error, *http.Response) {
	req, err := NewRequest("POST", fmt.Sprintf("%s/invoicing/invoices/%s/record-refund", c.APIBase, invoiceID), &struct {
		Date *time.Time `json:"date"`
		Note string     `json:"note"`
	}{
		Date: date,
		Note: note,
	})
	if err != nil {
		return err, nil
	}

	v := &struct{}{}

	resp, err := c.SendWithAuth(req, &v)
	if err != nil {
		return err, resp
	}

	return nil, resp
}
