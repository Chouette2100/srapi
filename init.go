package srapi

import (
	"fmt"
	"log"
	"runtime"
)

// 初期化処理
func init() {
	// GoのバージョンとOS/Archを取得
	goVersion := runtime.Version()
	goOS := runtime.GOOS
	goArch := runtime.GOARCH

	var major, minor, patch int
	fmt.Sscanf(Version, "%1d%3d%2d", &major, &minor, &patch)

	useragent = fmt.Sprintf("srapi/%d.%d.%d (%s; Go/%s; %s/%s)",
		major, minor, patch, mailAddress, goVersion, goOS, goArch)

	log.Printf("User-Agent: %s\n", useragent)
}
