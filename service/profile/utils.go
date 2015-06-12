package profile

import (
	"github.com/pritunl/pritunl-client-electron/service/utils"
	"path/filepath"
	"runtime"
)

func getOpenvpnPath() (pth string) {
	switch runtime.GOOS {
	case "windows":
		pth = filepath.Join(utils.GetRootDir(), "openvpn", "openvpn.exe")
	case "linux":
		pth = "openvpn"
	default:
		panic("profile: Not implemented")
	}

	return
}