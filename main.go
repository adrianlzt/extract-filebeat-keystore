package main

import (
	"fmt"

	"github.com/elastic/elastic-agent-libs/keystore"
)

func main() {
	ks, err := keystore.NewFileKeystoreWithPassword("filebeat.keystore", keystore.NewSecureString([]byte("")))
	if err != nil {
		panic(err)
	}

	keys, err := ks.(*keystore.FileKeystore).List()
	if err != nil {
		panic(err)
	}

	for _, k := range keys {
		secureText, err := ks.Retrieve(k)
		if err != nil {
			panic(err)
		}
		text, err := secureText.Get()
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v: %v\n", k, string(text))
	}

	// Example writing and reading a value

	// ksw, err := keystore.AsWritableKeystore(ks)
	// if err != nil {
	// 	panic(err)
	// }

	// err = ksw.Store("test", []byte("test"))
	// if err != nil {
	// 	panic(err)
	// }

	// text, err := ks.Retrieve("test")
	// if err != nil {
	// 	panic(err)
	// }

	// s, err := text.Get()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(s))
}
