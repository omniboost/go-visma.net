package vismanet_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSubaccountGetAll(t *testing.T) {
	req := client.NewSubaccountGetAll()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
