package simulation

import (
	"math/rand"

	"archive/x/identity/keeper"
	"archive/x/identity/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgAddIdentityMember(
	ak types.AccountKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAddIdentityMember{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AddIdentityMember simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AddIdentityMember simulation not implemented"), nil, nil
	}
}
