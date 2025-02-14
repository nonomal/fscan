package Plugins

import (
	"github.com/shadow1ng/fscan/Common"
	"github.com/shadow1ng/fscan/WebScan"
)

// WebPoc 直接执行Web漏洞扫描
func WebPoc(info *Common.HostInfo) error {
	WebScan.WebScan(info)
	return nil
}
