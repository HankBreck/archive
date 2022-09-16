package cda

import (
	"math/rand"

	"archive/testutil/sample"
	cdasimulation "archive/x/cda/simulation"
	"archive/x/cda/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = cdasimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateCDA = "op_weight_msg_create_cda"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateCDA int = 100

	opWeightMsgApproveCda = "op_weight_msg_approve_cda"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveCda int = 100

	opWeightMsgFinalizeCda = "op_weight_msg_finalize_cda"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFinalizeCda int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	cdaGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&cdaGenesis)
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

	var weightMsgCreateCDA int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateCDA, &weightMsgCreateCDA, nil,
		func(_ *rand.Rand) {
			weightMsgCreateCDA = defaultWeightMsgCreateCDA
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateCDA,
		cdasimulation.SimulateMsgCreateCDA(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveCda int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveCda, &weightMsgApproveCda, nil,
		func(_ *rand.Rand) {
			weightMsgApproveCda = defaultWeightMsgApproveCda
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveCda,
		cdasimulation.SimulateMsgApproveCda(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFinalizeCda int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFinalizeCda, &weightMsgFinalizeCda, nil,
		func(_ *rand.Rand) {
			weightMsgFinalizeCda = defaultWeightMsgFinalizeCda
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFinalizeCda,
		cdasimulation.SimulateMsgFinalizeCda(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
