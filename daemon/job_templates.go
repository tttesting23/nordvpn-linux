package daemon

import (
	"encoding/hex"
	"net/http"

	"github.com/NordSecurity/nordvpn-linux/core"
	"github.com/NordSecurity/nordvpn-linux/internal"
)

func JobTemplates(cdn core.CDN) func() {
	return func() {
		getTemplate := func(isObfuscated bool) {
			var digest string
			filepath := internal.OvpnTemplatePath
			if isObfuscated {
				filepath = internal.OvpnObfsTemplatePath
			}
			if internal.FileExists(filepath) {
				if hash, err := internal.FileSha256(filepath); err == nil {
					digest = hex.EncodeToString(hash)
				}
			}

			headers, _, err := cdn.ConfigTemplate(isObfuscated, http.MethodHead)
			if err != nil {
				return
			}

			if digest != headers.Get(core.HeaderDigest) {
				_, body, err := cdn.ConfigTemplate(isObfuscated, http.MethodGet)
				if err != nil {
					return
				}
				err = internal.FileWrite(filepath, body, internal.PermUserRW)
				if err != nil {
					return
				}
			}
		}
		go getTemplate(false)
		go getTemplate(true)
	}
}
