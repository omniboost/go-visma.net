package vismanet_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestDepartmentGetAll(t *testing.T) {
	req := client.NewDepartmentGetAll()
	// req.QueryParams().PeriodID = "202103"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
