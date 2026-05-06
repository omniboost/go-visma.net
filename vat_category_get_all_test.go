package vismanet_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestVATCategoryGetAll(t *testing.T) {
	req := client.NewVATCategoryGetAll()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
