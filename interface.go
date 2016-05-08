package lordcommander

import "bytes"

// Commander is a testable interface built to capture the complexities of
// shelling out in Go.
type Commander interface {
	// Run executes the command with the same arguments as exec.Cmd.
	Run(command string, args ...string) error

	// Stdout returns standard output of the process as captured by the
	// underlying bytes.Buffer.
	Stdout() *bytes.Buffer

	// Stderr returns standard error of the process as captured by the
	// underlying bytes.Buffer.
	Stderr() *bytes.Buffer
}
