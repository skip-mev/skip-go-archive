// A script to test skip helper library in go
package main

import (
	"os"

	"github.com/joho/godotenv"

	"skip-go/skip"

	"encoding/base64"
)

func main() {
	// Load .env file
	err := godotenv.Load("local.env")

	if err != nil {
		panic(err)
	}

	privateKey := os.Getenv("PRIVATE_KEY")
	publicKey := os.Getenv("PUBLIC_KEY")
	skipRPCURL := os.Getenv("SKIP_RPC_URL")

	txBytes := []byte("\n\x90\x01\n\x8d\x01\n\x1c/cosmos.bank.v1beta1.MsgSend\x12m\n+juno1zhqrfu9w3sugwykef3rq8t0vlxkz72vwnnptts\x12+juno1ptcltmzllgu0am4c0wmgdlkv5y7r5grsn9h76m\x1a\x11\n\x05junox\x12\x0810000000\x12d\nP\nF\n\x1f/cosmos.crypto.secp256k1.PubKey\x12#\n!\x03H\x14l=[\x1f\xf6bg*,\n\x954\xcc9\x8e\xd2\x0eF\x8dz\x9b\xfdXec\xe7\xbeo\x16\x95\x12\x04\n\x02\x08\x01\x18\x07\x12\x10\n\n\n\x05junox\x12\x010\x10\xa0\x8d\x06\x1a@\x82MzmjC#\xba\xec`\xd0\xde-p\xb6\xba\x1d1\xe5\xdc\r,\x0e59\x88b\x05\x02\xf8]Nf\xd5`\xd0u4V\xfc#\xf2R\xad\xa3\xfe\xaf\x85\xf6\xac\x9a\x8f\x11\xb2\xfaYM#m\xbd\xd4Ozd")

	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKey)

	if err != nil {
		panic(err)
	}

	skip.SignAndSendBundle([][]byte{txBytes}, privateKeyBytes, publicKey, skipRPCURL, "0", true)
}
