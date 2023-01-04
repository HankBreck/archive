package identity

import (
	"math/rand"

	"archive/testutil/sample"
	identitysimulation "archive/x/identity/simulation"
	"archive/x/identity/types"

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
	_ = identitysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRegisterIssuer = "op_weight_msg_register_issuer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegisterIssuer int = 100

	opWeightMsgAcceptIdentity = "op_weight_msg_accept_identity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAcceptIdentity int = 100

	opWeightMsgRejectIdentity = "op_weight_msg_reject_identity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRejectIdentity int = 100

	opWeightMsgRevokeIdentity = "op_weight_msg_revoke_identity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRevokeIdentity int = 100

	opWeightMsgRenounceIdentity = "op_weight_msg_renounce_identity"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRenounceIdentity int = 100

	opWeightMsgAddIdentityMember = "op_weight_msg_add_identity_member"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddIdentityMember int = 100

	opWeightMsgUpdateOperators = "op_weight_msg_update_operators"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateOperators int = 100

	opWeightMsgUpdateMembers = "op_weight_msg_update_members"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateMembers int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	identityGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&identityGenesis)
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

	var weightMsgRegisterIssuer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegisterIssuer, &weightMsgRegisterIssuer, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterIssuer = defaultWeightMsgRegisterIssuer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegisterIssuer,
		identitysimulation.SimulateMsgRegisterIssuer(am.accountKeeper, am.keeper),
	))

	var weightMsgAcceptIdentity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAcceptIdentity, &weightMsgAcceptIdentity, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptIdentity = defaultWeightMsgAcceptIdentity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAcceptIdentity,
		identitysimulation.SimulateMsgAcceptIdentity(am.accountKeeper, am.keeper),
	))

	var weightMsgRejectIdentity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRejectIdentity, &weightMsgRejectIdentity, nil,
		func(_ *rand.Rand) {
			weightMsgRejectIdentity = defaultWeightMsgRejectIdentity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRejectIdentity,
		identitysimulation.SimulateMsgRejectIdentity(am.accountKeeper, am.keeper),
	))

	var weightMsgRevokeIdentity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRevokeIdentity, &weightMsgRevokeIdentity, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeIdentity = defaultWeightMsgRevokeIdentity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRevokeIdentity,
		identitysimulation.SimulateMsgRevokeIdentity(am.accountKeeper, am.keeper),
	))

	var weightMsgRenounceIdentity int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRenounceIdentity, &weightMsgRenounceIdentity, nil,
		func(_ *rand.Rand) {
			weightMsgRenounceIdentity = defaultWeightMsgRenounceIdentity
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRenounceIdentity,
		identitysimulation.SimulateMsgRenounceIdentity(am.accountKeeper, am.keeper),
	))

	var weightMsgAddIdentityMember int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddIdentityMember, &weightMsgAddIdentityMember, nil,
		func(_ *rand.Rand) {
			weightMsgAddIdentityMember = defaultWeightMsgAddIdentityMember
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddIdentityMember,
		identitysimulation.SimulateMsgAddIdentityMember(am.accountKeeper, am.keeper),
	))

	var weightMsgUpdateOperators int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateOperators, &weightMsgUpdateOperators, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateOperators = defaultWeightMsgUpdateOperators
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateOperators,
		identitysimulation.SimulateMsgUpdateOperators(am.accountKeeper, am.keeper),
	))

	var weightMsgUpdateMembers int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateMembers, &weightMsgUpdateMembers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateMembers = defaultWeightMsgUpdateMembers
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateMembers,
		identitysimulation.SimulateMsgUpdateMembers(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
