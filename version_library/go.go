package version_library

import "runtime"

// GetVersion returns the Go version
func GetVersion() string {
	return runtime.Version()
}
