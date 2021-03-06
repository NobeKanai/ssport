package ssport

import (
	"encoding/json"
	"fmt"
)

func createUser(token string) (*UserInfo, error) {
	url := fmt.Sprintf("http://client3.aipao.me/api/%s/QM_Users/GS", token)

	rsp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(rsp.Body)

	result := struct {
		Success bool
		Data    *UserInfo
	}{}
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}

	if !result.Success {
		return nil, ErrInvalidToken
	}

	return result.Data, nil
}

// UserInfo infomation
type UserInfo struct {
	User      *user      `json:"User"`
	SchoolRun *schoolRun `json:"SchoolRun"`
}

type schoolRun struct {
	Sex         string  `json:"Sex"`
	SchoolID    string  `json:"SchoolId"`
	SchoolName  string  `json:"SchoolName"`
	MinSpeed    float64 `json:"MinSpeed"`
	MaxSpeed    float64 `json:"MaxSpeed"`
	Lengths     int64   `json:"Lengths"`
	IsNeedPhoto string  `json:"IsNeedPhoto"`
	IsShowAd    int64   `json:"IsShowAd"`
}

type user struct {
	UserID           int64       `json:"UserID"`
	NickName         string      `json:"NickName"`
	UserName         string      `json:"UserName"`
	Sex              string      `json:"Sex"`
	Province         interface{} `json:"Province"`
	City             interface{} `json:"City"`
	Country          interface{} `json:"Country"`
	HeadImgURL       string      `json:"HeadImgUrl"`
	Mobile           interface{} `json:"Mobile"`
	MobileVerifyCode interface{} `json:"MobileVerifyCode"`
	IsMoblileVerify  string      `json:"IsMoblileVerify"`
	Weights          float64     `json:"Weights"`
	Bmi              float64     `json:"BMI"`
	Heights          float64     `json:"Heights"`
	Birthday         string      `json:"Birthday"`
	OldYears         int64       `json:"OldYears"`
	IsInfoOk         string      `json:"IsInfoOk"`
	WXNickName       interface{} `json:"WXNickName"`
	WXSex            interface{} `json:"WXSex"`
	IsStationOpen    string      `json:"IsStationOpen"`
	IsBgMusic        string      `json:"IsBgMusic"`
	IsReciveMsg      string      `json:"IsReciveMsg"`
	IsSchoolMode     string      `json:"IsSchoolMode"`
	LevelLengh       int64       `json:"Level_Lengh"`
	LevelLenghDate   string      `json:"Level_Lengh_Date"`
	DaysStart        int64       `json:"Days_Start"`
	DaysStartDate    string      `json:"Days_Start_Date"`
}
