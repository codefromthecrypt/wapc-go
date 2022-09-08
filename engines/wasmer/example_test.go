package wasmer_test

import (
	"github.com/wapc/wapc-go/engines/wasmer"
	"github.com/wapc/wapc-go/engines/wasmer/testdata/metering"
)

// This example shows how to configure an engine-specific option, in this case
// wasmer.WithMetering.
func Example_WithMetering() {
	// This configures the maximum gas usable as well a function that decides
	// how much gas each operation takes.
	engOpt := wasmer.WithMetering(uint64(800000000), metering.CFunctionPointer())
	// You can now pass this option to the engine, which enables metering.
	_ = wasmer.Engine(engOpt)
}
