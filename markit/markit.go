package markit

import(
	"bytes"
)

const (
	BaseUrl		= "http://dev.markitondemand.com/Api/v2"
)

func sanitize(body []byte) []byte {
	return bytes.Replace(bytes.TrimSpace(body), []byte{'"'}, []byte{}, -1)
}
