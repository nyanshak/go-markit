package markit

import(
	"strings"
)

const (
	BaseUrl		= "http://dev.markitondemand.com/Api/v2"
)

func sanitize(body string) string {
	return strings.Replace(strings.TrimSpace(body), "\"", "", -1)
}
