package blueprint

import (
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

func Decode[T Book | Blueprint](in io.Reader, v *T) error {
	// read the first letter and check if it's a 0
	var versionBit [1]byte
	if n, err := in.Read(versionBit[:]); err != nil && err != io.EOF {
		return err
	} else if n == 0 {
		return errors.New("blueprint string is empty")
	}

	if ver := versionBit[0]; ver != '0' {
		return fmt.Errorf("invalid version byte: expected '0', got '%c'", ver)
	}

	rc, err := zlib.NewReader(base64.NewDecoder(base64.StdEncoding, in))
	if err != nil {
		return err
	}
	defer rc.Close()

	var be BookEntry
	if err := json.NewDecoder(rc).Decode(&be); err != nil {
		return err
	}

	if b, ok := be.Entry.(T); ok {
		*v = b
		return nil
	}

	return fmt.Errorf("unsupported type for BookEntry: %T", v)
}
