package sshsig

import (
	"fmt"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/testdata"
)

var (
	testPrivateKeys map[string]interface{}
	testSigners     map[string]ssh.Signer
)

func init() {
	var err error

	n := len(testdata.PEMBytes)
	testPrivateKeys = make(map[string]interface{}, n)
	testSigners = make(map[string]ssh.Signer, n)
	for t, k := range testdata.PEMBytes {
		testPrivateKeys[t], err = ssh.ParseRawPrivateKey(k)
		if err != nil {
			panic(fmt.Sprintf("Unable to parse test key %s: %v", t, err))
		}
		testSigners[t], err = ssh.NewSignerFromKey(testPrivateKeys[t])
		if err != nil {
			panic(fmt.Sprintf("Unable to create signer for test key %s: %v", t, err))
		}
	}
}
