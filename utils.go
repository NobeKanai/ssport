package ssport

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	encryptKey = "xfvdmyirsg"
)

func getRunID(token string, lengths int64) (string, error) {
	rsp, err := client.Get(fmt.Sprintf("http://client3.aipao.me/api/%s/QM_Runs/SRS?S1=40.62828&S2=120.79108&S3=%d", token, lengths))
	if err != nil {
		return "", err
	}

	decoder := json.NewDecoder(rsp.Body)

	result := struct {
		Success bool `json:"Success"`
		Data    struct {
			RunID string `json:"RunId"`
		} `json:"Data"`
	}{}
	if err := decoder.Decode(&result); err != nil {
		return "", ErrFailedRunID
	}

	return result.Data.RunID, nil
}

// recordExists check if today's running record is existed
func recordExists(r *RecordInfo) bool {
	date := r.LastResultDate
	if now := time.Now(); now.Year() != date.Year() || now.Month() != date.Month() || now.Day() != date.Day() {
		return false
	}

	return true
}

// run will directly insert a running record which limited by sr(schoolRun)
func run(token, runID string, sr *schoolRun) error {
	// get random speed
	maxSpeed := sr.MaxSpeed - 0.5
	minSpeed := sr.MinSpeed + 0.3

	speed := (maxSpeed-minSpeed)*rand.Float64() + minSpeed

	// get random distance
	distance := sr.Lengths + int64(rand.Intn(5))

	// get costTime
	costTime := int64(math.Round(float64(distance) / speed))

	// get step
	step := 1700 + rand.Intn(500)

	url := fmt.Sprintf("http://client3.aipao.me/api/%s/QM_Runs/ES?S1=%s&S4=%s&S5=%s&S6=A0A2A1A3A0&S7=1&S8=%s&S9=%s",
		token,
		runID,
		encrypt(costTime),
		encrypt(distance),
		encryptKey,
		encrypt(int64(step)),
	)

	rsp, err := client.Get(url)
	if err != nil {
		return err
	}

	decoder := json.NewDecoder(rsp.Body)
	res := struct {
		Success bool
	}{}
	if err := decoder.Decode(&res); err != nil || !res.Success {
		return ErrFailedRun
	}

	return nil
}

func encrypt(num int64) string {
	numbers := make([]int, 0, 10)

	for num != 0 {
		numbers = append(numbers, int(num%10))
		num /= 10
	}

	result := make([]byte, len(numbers))
	for i, n := range numbers {
		result[i] = encryptKey[n]
	}

	// reversing
	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
	}

	return string(result)
}
