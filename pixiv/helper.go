package pixiv

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
)

func withOpts(opts interface{}, values url.Values, caller string) (url.Values, error) {
	// Overwrite opts with values
	if opts != nil {
		q, err := query.Values(opts)
		if err != nil {
			return nil, fmt.Errorf("pixiv: %s: query encode: %w", caller, err)
		}
		for k, v := range values {
			q[k] = v
		}
		return q, nil
	}
	return values, nil
}

func intsToStrings(idns []int) []string {
	ids := make([]string, len(idns))
	for i, x := range idns {
		ids[i] = strconv.Itoa(x)
	}
	return ids
}
