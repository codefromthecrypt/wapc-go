package metering

// #include <wasmer.h>
// __attribute__((weak))
// extern uint64_t metering_delegate_alt(enum wasmer_parser_operator_t op);
import "C"
import "unsafe"

//export meteringFn
func meteringFn(op C.wasmer_parser_operator_t) C.uint64_t {
	if op >= C.I32Load && op <= C.I64TruncSatF64U {
		return 1
	}
	return 0 // no value means no cost
}

// CFunctionPointer returns the C pointer to meteringFn.
func CFunctionPointer() unsafe.Pointer {
	return unsafe.Pointer(C.meteringFn)
}
