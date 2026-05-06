package vismanet_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerCreditNoteGetAll(t *testing.T) {
	req := client.NewCustomerCreditNoteGetAll()
	req.QueryParams().ExternalReference = "202331913"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
