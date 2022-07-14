package polischain_test

import (
	"testing"

	keepertest "github.com/Metapolis21/PolisChain/testutil/keeper"
	"github.com/Metapolis21/PolisChain/testutil/nullify"
	"github.com/Metapolis21/PolisChain/x/polischain"
	"github.com/Metapolis21/PolisChain/x/polischain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PolischainKeeper(t)
	polischain.InitGenesis(ctx, *k, genesisState)
	got := polischain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
