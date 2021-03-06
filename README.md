# Sunshine Sport Api

This tool provides several APIs for Sunshine Sport(阳光体育平台).

## Document

```golang
package main

import (
	"fmt"
	"log"

	"github.com/NobeKanai/ssport"
)

func main() {
	info, err := ssport.FetchInfo("your imeicode") // fetch information by imeicode
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println("user name is", info.UserInfo.User.NickName)
	fmt.Println("user id is", info.AuthInfo.UserID)

	fmt.Println("last run date is", info.RecordInfo.LastResultDate)
	fmt.Println("total race number is", info.RecordInfo.RaceNums)

	// start inserting a record
	if err := ssport.Run(info); err != nil {
		log.Fatal(err)
	}
}
```
