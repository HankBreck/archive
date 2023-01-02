package simulation

import (
	"math/rand"

	"archive/x/identity/keeper"
	"archive/x/identity/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgRenounceIdentity(
	ak types.AccountKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgRenounceIdentity{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the RenounceIdentity simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "RenounceIdentity simulation not implemented"), nil, nil
	}
}
