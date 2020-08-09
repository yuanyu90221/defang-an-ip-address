package defang_ip

import "strings"

func defangIPaddr(address string) string {
	ret := ""
	splitedAddr := strings.Split(address, ".")
	ret = strings.Join(splitedAddr, "[.]")
	return ret
}
