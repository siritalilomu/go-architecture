package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("username:password"))) // dXNlcm5hbWU6cGFzc3dvcmQ=
	// curl -u username:password -v google.com
}
