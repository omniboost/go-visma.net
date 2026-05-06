package vismanet_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerGetByCD(t *testing.T) {
	req := client.NewCustomerGetByCD()
	req.PathParams().CustomerCD = "M131848"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
