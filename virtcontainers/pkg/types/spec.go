package types

import (
	spec "github.com/opencontainers/runtime-spec/specs-go"
)

// CompatOCIProcess is a structure inheriting from spec.Process defined
// in runtime-spec/specs-go package. The goal is to be compatible with
// both v1.0.0-rc4 and v1.0.0-rc5 since the latter introduced a change
// about the type of the Capabilities field.
// Refer to: https://github.com/opencontainers/runtime-spec/commit/37391fb
type CompatOCIProcess struct {
	spec.Process
	Capabilities interface{} `json:"capabilities,omitempty" platform:"linux"`
}

// CompatOCISpec is a structure inheriting from spec.Spec defined
// in runtime-spec/specs-go package. It relies on the CompatOCIProcess
// structure declared above, in order to be compatible with both
// v1.0.0-rc4 and v1.0.0-rc5.
// Refer to: https://github.com/opencontainers/runtime-spec/commit/37391fb
type CompatOCISpec struct {
	spec.Spec
	Process *CompatOCIProcess `json:"process,omitempty"`
}
