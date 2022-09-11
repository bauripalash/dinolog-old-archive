package lib

import (
	"bytes"
	"compress/gzip"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func compressData(data string) []byte {
	var buf bytes.Buffer

	gw := gzip.NewWriter(&buf)

	if _, err := gw.Write([]byte(data)); err != nil {
		log.Fatal(err)
	}

	err := gw.Close()

	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}

func NewResponse(raw_res Response, compress bool) []byte {
	extra_options := 0
	output := []byte(raw_res.Data)
	if compress {
		extra_options = 1
		output = compressData(raw_res.Data)
	}

	resp := []byte(fmt.Sprintf("D~%d~%d~%d~X~%d\r\n", raw_res.Status, len(output), raw_res.NumPost, extra_options))

	if extra_options == 1 {
		resp = append(resp, []byte("gzip=1")...)
	}

	resp = append(resp, output...)

	return resp

}
