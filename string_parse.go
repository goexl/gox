package gox

import (
	`strings`
)

// ParsePackageName 从域名解析出包名
func ParsePackageName(domain string) (name string) {
	// 包名不能包含特殊字符，比如-、_、@、#等
	domain = strings.ReplaceAll(domain, "-", ".")
	domain = strings.ReplaceAll(domain, "_", ".")
	domain = strings.ReplaceAll(domain, "@", ".")
	domain = strings.ReplaceAll(domain, "#", ".")

	originalHosts := strings.Split(domain, ".")
	hosts := make([]string, 0, len(originalHosts))
	for _, originalHost := range originalHosts {
		host := strings.TrimSpace(originalHost)
		if "" != host {
			hosts = append(hosts, host)
		}
	}

	// 包名必须包含多个单词，如果只有一个单词，默认增加一个com后缀
	if 0 == len(hosts) {
		hosts = append(hosts, "com")
	}

	// 反转，和包名规范一致
	for i, j := 0, len(hosts)-1; i < j; i, j = i+1, j-1 {
		hosts[i], hosts[j] = hosts[j], hosts[i]
	}
	name = strings.Join(hosts, ".")

	return
}
