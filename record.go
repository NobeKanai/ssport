package ssport

import (
	"encoding/json"
	"fmt"
	"time"
)

const dateLayout = "2006年01月02日"

// RecordInfo contains race num and last result date
type RecordInfo struct {
	RaceNums       int64
	LastResultDate time.Time
}

func createRecord(token string, userID int64) (*RecordInfo, error) {
	var r RecordInfo

	url := fmt.Sprintf("https://client4.aipao.me/api/%s/QM_Runs/getResultsofValidByUser?UserId=%d&pageIndex=1&pageSize=1", token, userID)

	rsp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(rsp.Body)

	result := struct {
		// Success   bool
		RaceNums  int64
		ListValue []struct {
			ResultDate string
		}
	}{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	r.RaceNums = result.RaceNums
	if len(result.ListValue) == 0 {
		return &r, nil
	}

	date, err := time.Parse(dateLayout, result.ListValue[0].ResultDate)
	if err != nil {
		return nil, err
	}

	r.LastResultDate = date
	return &r, nil
}
