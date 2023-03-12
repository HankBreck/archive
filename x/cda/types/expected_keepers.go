package types

import (
	"github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IdentityKeeper interface {
	HasMember(ctx sdk.Context, certificateId uint64, member sdk.AccAddress) (bool, error)
	HasCertificate(ctx sdk.Context, id uint64) bool
}

type WasmKeeper interface {
	Instantiate(
		ctx sdk.Context,
		codeID uint64,
		creator, admin sdk.AccAddress,
		initMsg []byte,
		label string,
		deposit sdk.Coins,
	) (sdk.AccAddress, []byte, error)
	ClearContractAdmin(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress) error
	SetContractInfoExtension(ctx sdk.Context, contract sdk.AccAddress, extra types.ContractInfoExtension) error
}
