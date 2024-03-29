package vismanet_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCustomerInvoiceGet(t *testing.T) {
	req := client.NewCustomerInvoiceGet()
	req.PathParams().InvoiceNumber = "2000144"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
