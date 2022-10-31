package factory

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/umma-chain/core-umma/testutil/sample"
	factorysimulation "github.com/umma-chain/core-umma/x/factory/simulation"
	"github.com/umma-chain/core-umma/x/factory/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = factorysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateDenom = "op_weight_msg_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateDenom int = 100

	opWeightMsgUpdateDenom = "op_weight_msg_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateDenom int = 100

	opWeightMsgDeleteDenom = "op_weight_msg_denom"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteDenom int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	factoryGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		DenomList: []types.Denom{
			{
				Owner: sample.AccAddress(),
				Denom: "0",
			},
			{
				Owner: sample.AccAddress(),
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&factoryGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateDenom int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateDenom, &weightMsgCreateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgCreateDenom = defaultWeightMsgCreateDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateDenom,
		factorysimulation.SimulateMsgCreateDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateDenom int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateDenom, &weightMsgUpdateDenom, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateDenom = defaultWeightMsgUpdateDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateDenom,
		factorysimulation.SimulateMsgUpdateDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteDenom int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteDenom, &weightMsgDeleteDenom, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteDenom = defaultWeightMsgDeleteDenom
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteDenom,
		factorysimulation.SimulateMsgDeleteDenom(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
