package app

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
)

const (
	// DefaultUmmaInstanceCost is initially set the same as in wasmd
	DefaultUmmaInstanceCost uint64 = 60_000
	// DefaultUmmaCompileCost set to a large number for testing
	DefaultUmmaCompileCost uint64 = 3
)

// UmmaGasRegisterConfig is defaults plus a custom compile amount
func UmmaGasRegisterConfig() wasmkeeper.WasmGasRegisterConfig {
	gasConfig := wasmkeeper.DefaultGasRegisterConfig()
	gasConfig.InstanceCost = DefaultUmmaInstanceCost
	gasConfig.CompileCost = DefaultUmmaCompileCost

	return gasConfig
}

func NewUmmaWasmGasRegister() wasmkeeper.WasmGasRegister {
	return wasmkeeper.NewWasmGasRegister(UmmaGasRegisterConfig())
}
