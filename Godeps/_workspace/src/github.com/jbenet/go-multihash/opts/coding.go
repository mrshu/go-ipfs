package opts

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	base58 "github.com/ipfs/go-ipfs/Godeps/_workspace/src/github.com/jbenet/go-base58"
	mh "gx/ipfs/QmYf7ng2hG5XBtJA3tN34DQ2GUN5HNksEw1rLDkmr6vGku/go-multihash"
)

func Decode(encoding, digest string) (mh.Multihash, error) {
	switch encoding {
	case "raw":
		return mh.Cast([]byte(digest))
	case "hex":
		return hex.DecodeString(digest)
	case "base58":
		return base58.Decode(digest), nil
	case "base64":
		return base64.StdEncoding.DecodeString(digest)
	default:
		return nil, fmt.Errorf("unknown encoding: %s", encoding)
	}
}

func Encode(encoding string, hash mh.Multihash) (string, error) {
	switch encoding {
	case "raw":
		return string(hash), nil
	case "hex":
		return hex.EncodeToString(hash), nil
	case "base58":
		return base58.Encode(hash), nil
	case "base64":
		return base64.StdEncoding.EncodeToString(hash), nil
	default:
		return "", fmt.Errorf("unknown encoding: %s", encoding)
	}
}
