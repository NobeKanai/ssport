package ssport

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
)

var (
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	// ErrDuplicated indicates today's sport record already exists
	ErrDuplicated = errors.New("duplicated sport record")
	// ErrInvalidIMEICode indicates an error occured when get auth information
	ErrInvalidIMEICode = errors.New("imeicode is invalid")
	// ErrInvalidToken indicates token is invalid
	ErrInvalidToken = errors.New("invalid token")
	// ErrFailedRunID indicates an error when get run id
	ErrFailedRunID = errors.New("failed to get run id")
	// ErrFailedRun indicates error when insert a record
	ErrFailedRun = errors.New("failed to insert running record to server")
)

// Info contains all information about user
type Info struct {
	AuthInfo   *AuthInfo
	UserInfo   *UserInfo
	RecordInfo *RecordInfo
}

// FetchInfo get all information needs by user
func FetchInfo(imeicode string) (*Info, error) {
	authInfo, err := createAuth(imeicode)
	if err != nil {
		return nil, err
	}

	user, err := createUser(authInfo.Token)
	if err != nil {
		return nil, err
	}

	record, err := createRecord(authInfo.Token, authInfo.UserID)
	if err != nil {
		return nil, err
	}

	info := &Info{
		AuthInfo:   authInfo,
		UserInfo:   user,
		RecordInfo: record,
	}

	return info, nil
}

// Run insert a record
func Run(info *Info) error {
	// check for duplicate submissions
	if recordExists(info.RecordInfo) {
		return ErrDuplicated
	}

	// fetch run id
	runID, err := getRunID(info.AuthInfo.Token, info.UserInfo.SchoolRun.Lengths)
	if err != nil {
		return fmt.Errorf("error when get run id: %w", err)
	}

	// run it
	if err := run(info.AuthInfo.Token, runID, info.UserInfo.SchoolRun); err != nil {
		return err
	}

	return nil
}
