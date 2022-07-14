package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Metapolis21/PolisChain/testutil/keeper"
	"github.com/Metapolis21/PolisChain/x/polischain/keeper"
	"github.com/Metapolis21/PolisChain/x/polischain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PolischainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
