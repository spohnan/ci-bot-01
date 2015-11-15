// Package bot is the top level container for ci-bot-01 code
package bot

import (
	_ "github.com/spohnan/ci-bot-01/paas" // Run paas::init()
)

func init() {
	initWebHandlers()
}
