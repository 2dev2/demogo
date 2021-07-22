package demogo

import "embed"

//go:embed web/build
var EmbededFiles embed.FS
