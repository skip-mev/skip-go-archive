package skip

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tendermint/tendermint/crypto/secp256k1"
)

func SignBundle(bundle [][]byte, privateKeyBytes []byte) ([]string, []byte) {
	// Append list of bytes to a single byte slice
	var bundleBytes []byte
	for _, tx := range bundle {
		bundleBytes = append(bundleBytes, tx...)
	}

	// Create private key object to sign
	privKey := secp256k1.PrivKey(privateKeyBytes)

	// Sign digest of bundleBytes, digest is created by
	// hashing the bundleBytes within the Sign method
	bundleSignature, err := privKey.Sign(bundleBytes)

	// Check for errors
	if err != nil {
		panic(err)
	}

	// Create b64 encoded bundle
	base64EncodedBundle := []string{}
	for _, tx := range bundle {
		base64EncodedBundle = append(base64EncodedBundle, base64.StdEncoding.EncodeToString(tx))
	}

	return base64EncodedBundle, bundleSignature
}

func SendBundle(b64EncodedSignedBundle []string, bundleSignature []byte, publicKey string, rpcURL string, desiredHeight string, sync bool) {
	// Send signed bundle to RPC
	var method string
	if sync {
		method = "broadcast_bundle_sync"
	} else {
		method = "broadcast_bundle_async"
	}

	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params": []interface{}{
			b64EncodedSignedBundle,
			desiredHeight,
			publicKey,
			bundleSignature,
		},
		"id": 1,
	}

	json_data, err := json.Marshal(data)

	if err != nil {
		panic(err)
	}

	response, err := http.Post(rpcURL, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
		panic(err)
	}

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	fmt.Println(result)
}

func SignAndSendBundle(bundle [][]byte, privateKeyBytes []byte, publicKey string, rpcURL string, desiredHeight string, sync bool) {
	b64EncodedSignedBundle, bundleSignature := SignBundle(bundle, privateKeyBytes)
	SendBundle(b64EncodedSignedBundle, bundleSignature, publicKey, rpcURL, desiredHeight, sync)
}
