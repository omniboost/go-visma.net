package vismanet_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestContextUserdetailsGet(t *testing.T) {
	req := client.NewContextUserdetailsGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
