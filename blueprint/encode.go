package blueprint

import (
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"io"
)

func Encode(in IsBlueprint, out io.Writer) error {
	if _, err := out.Write([]byte{'0'}); err != nil {
		return err
	}

	b64Writer := base64.NewEncoder(base64.StdEncoding, out)
	defer b64Writer.Close()

	zlibWriter, err := zlib.NewWriterLevel(b64Writer, zlib.BestCompression)
	if err != nil {
		return err
	}
	defer zlibWriter.Close()

	return json.NewEncoder(zlibWriter).Encode(BookEntry{Entry: in})
}
