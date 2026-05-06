package vismanet_test

import (
	"context"
	"encoding/json"
	"log"
	"testing"
)

func TestJournalTransactionV2GetAll(t *testing.T) {
	req := client.NewJournalTransactionV2GetAll()
	req.QueryParams().PeriodID = "202104"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
