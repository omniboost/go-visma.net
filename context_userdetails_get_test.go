package vismanet_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestContextUserdetailsGet(t *testing.T) {
	req := client.NewContextUserdetailsGet()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
