package types_test

import (
	"encoding/hex"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/umma-chain/umma-core/x/escrow/test"
	"github.com/umma-chain/umma-core/x/escrow/types"
)

type GenesisTestSuite struct {
	suite.Suite
	generator *test.EscrowGenerator
}

func (s *GenesisTestSuite) SetupTest() {
	s.generator = test.NewEscrowGenerator(200)
	test.SetConfig()
}

func TestGenesisTestSuite(t *testing.T) {
	suite.Run(t, new(GenesisTestSuite))
}

func (s *GenesisTestSuite) TestValidate() {

	testCases := []struct {
		name          string
		mutateGenesis func(state *types.GenesisState)
	}{
		{
			name:          "valid genesis",
			mutateGenesis: func(state *types.GenesisState) {},
		},
		{
			name: "valid genesis with expired escrows",
			mutateGenesis: func(state *types.GenesisState) {
				for i := 0; i < 20; i++ {
					escrow, _ := s.generator.NewRandomTestEscrow()
					escrow.State = types.EscrowState_Expired
					escrow.Deadline = s.generator.NowAfter(0) - 10
					state.Escrows = append(state.Escrows, escrow)
				}
			},
		},
		{
			name: "invalid genesis: block time in the future",
			mutateGenesis: func(state *types.GenesisState) {
				state.LastBlockTime = uint64(time.Now().Unix() + 100)
			},
		},
		{
			name: "invalid genesis: Escrow with expired state but with future deadline",
			mutateGenesis: func(state *types.GenesisState) {
				// Test invalid expired escrows
				escrow, _ := s.generator.NewRandomTestEscrow()
				// Escrow is not actually expired
				escrow.State = types.EscrowState_Expired
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with open state but with expired deadline",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				// Escrow should be expired
				escrow.Deadline = s.generator.NowAfter(0) - 1
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with completed state in store",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.State = types.EscrowState_Completed
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with refunded state in store",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.State = types.EscrowState_Refunded
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "Invalid genesis: Escrow with invalid ID: not enough characters",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.Id = "0123456789"
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with invalid ID: too much characters",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.Id = "1234567890123456789a"
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with negative price",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.Price = sdk.Coins{sdk.Coin{Denom: test.Denom, Amount: sdk.NewInt(-10)}}
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with empty price",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.Price = sdk.Coins{}
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with invalid ID: this ID will be generated for future escrows",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.Id = hex.EncodeToString(sdk.Uint64ToBigEndian(s.generator.GetNextId() + 500))
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with invalid seller",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.Seller = "star15555f8e8f7e4"
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with invalid broker",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.BrokerAddress = "star1455e5f4e5fe"
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with commission too low",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.BrokerCommission = sdk.NewDec(-1)
				state.Escrows = append(state.Escrows, escrow)
			},
		},
		{
			name: "invalid genesis: Escrow with commission too high",
			mutateGenesis: func(state *types.GenesisState) {
				escrow, _ := s.generator.NewRandomTestEscrow()
				escrow.BrokerCommission = sdk.NewDec(2)
				state.Escrows = append(state.Escrows, escrow)
			},
		},
	}

	for _, tc := range testCases {
		baseGenesis := s.generator.NewEscrowGenesis(100)
		tc.mutateGenesis(baseGenesis)
		// This has to be done for all tests
		baseGenesis.NextEscrowId = s.generator.GetNextId()

		validate := func(*testing.T) error {
			return types.ValidateGenesis(*baseGenesis)
		}

		test.EvaluateTest(s.T(), tc.name, validate)
	}
}
