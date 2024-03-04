package vismanet_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestContextGet(t *testing.T) {
	req := client.NewContextGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
