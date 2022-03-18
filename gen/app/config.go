package app

import "github.com/bitwormhole/starter/application"

// ExportConfigForCdepApp ...
func ExportConfigForCdepApp(cb application.ConfigBuilder) error {
	return autoGenConfig(cb)
}
