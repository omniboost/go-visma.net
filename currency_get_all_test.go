package vismanet_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestCurrencyGetAll(t *testing.T) {
	req := client.NewCurrencyGetAll()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
