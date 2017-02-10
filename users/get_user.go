package main

import intercom "gopkg.in/intercom/intercom-go.v2"
import "os"
import (
        "fmt"
)

func main() {
    user_id := os.Args[1]
    ic := intercom.NewClient(os.Getenv("AT"), "")
    user, err := ic.Users.FindByUserID(user_id)
    fmt.Printf("INTERCOM GO \n")
    fmt.Printf("%s \n", user.Email)
    _=err
}
