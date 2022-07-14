package keeper_test

import (
	"testing"

	testkeeper "github.com/Metapolis21/PolisChain/testutil/keeper"
	"github.com/Metapolis21/PolisChain/x/polischain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PolischainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
