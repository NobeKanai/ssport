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
	auth, user, err := ssport.FetchInfo("your imeicode")
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Println("token is", auth.Token)
	fmt.Println("user id is", auth.UserID)

	if err := ssport.Run(auth, user); err != nil {
		log.Fatal(err)
	}
}
```
