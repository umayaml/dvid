/*
	This file supports the keyspace for the keyvalue data type.
*/

package keyvalue

import (
	"fmt"

	"github.com/janelia-flyem/dvid/storage"
)

const (
	// keyUnknown should never be used and is a check for corrupt or incorrectly set keys
	keyUnknown storage.TKeyClass = iota

	// the byte id for a standard key of a keyvalue
	keyStandard = 177
)

// NewTKey returns the "key" key component.
func NewTKey(key string) (storage.TKey, error) {
	// Make sure the key has no embedded 0 values
	for i := 0; i < len(key); i++ {
		if key[i] == 0 {
			return nil, fmt.Errorf("key cannot have embedded 0 value")
		}
	}
	return storage.NewTKey(keyStandard, append([]byte(key), 0)), nil
}

// DecodeTKey returns the string key used for this keyvalue.
func DecodeTKey(tk storage.TKey) (string, error) {
	ibytes, err := tk.ClassBytes(keyStandard)
	if err != nil {
		return "", err
	}
	sz := len(ibytes) - 1
	if sz <= 0 {
		return "", fmt.Errorf("empty key")
	}
	if ibytes[sz] != 0 {
		return "", fmt.Errorf("expected 0 byte ending key of keyvalue key, got %d", ibytes[sz])
	}
	return string(ibytes[:sz]), nil
}
