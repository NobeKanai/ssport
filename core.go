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

// FetchInfo get all information needs by user
func FetchInfo(imeicode string) (*AuthInfo, *UserInfo, error) {
	authInfo, err := getAuthInfo(imeicode)
	if err != nil {
		return nil, nil, err
	}

	user, err := getUserInfoWithToken(authInfo.Token)
	if err != nil {
		return nil, nil, err
	}

	return authInfo, user, nil
}

// Run insert a record
func Run(auth *AuthInfo, user *UserInfo) error {
	// check for duplicate submissions
	if recordExists(auth.Token, auth.UserID) {
		return ErrDuplicated
	}

	// fetch run id
	runID, err := getRunID(auth.Token, user.SchoolRun.Lengths)
	if err != nil {
		return fmt.Errorf("error when get run id: %w", err)
	}

	// run it
	if err := run(auth.Token, runID, user.SchoolRun); err != nil {
		return err
	}

	return nil
}
