package reqPkg

import (
	"fmt"
	"regexp"
)

func IsValidEndPoint(rgx, endpoint string) bool {
	regx, err := regexp.Compile(fmt.Sprintf("^%s$", rgx))
	if err != nil {
		return false
	}
	return regx.Match([]byte(endpoint))
}
