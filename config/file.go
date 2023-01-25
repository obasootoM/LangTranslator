package config

import (
	"encoding/base64"
	"fmt"
	_"log"
	_"os"
	_"strings"
)





func ReadFile(file string) (string,error) {
	 bytess, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		fmt.Printf("failed to decode %s\n",err)
	}else{
		fmt.Printf("byte begin with %q\n",bytess[0:4])
	}
	fmt.Println(len(bytess))
	return file, err
}