package keeper_test

import (
	"encoding/hex"
	"math/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/umma-chain/umma-core/x/escrow/keeper"
	"github.com/umma-chain/umma-core/x/escrow/test"
	"github.com/umma-chain/umma-core/x/escrow/types"
)

type InvariantsSuite struct {
	BaseKeeperSuite
	escrows []types.Escrow
}

func (s *InvariantsSuite) saveObject(obj *types.TestObject, escrowId string) {
	// Escrows are already created, so all your objects are belong to us
	obj.Owner = s.keeper.GetEscrowAddress(escrowId)
	// Save the object to the store
	s.Assert().Nil(s.store.Create(obj), "Error while saving a generated object")
}

func (s *InvariantsSuite) addEscrow(escrow types.Escrow) {
	s.escrows = append(s.escrows, escrow)
	s.keeper.SaveEscrow(s.ctx, escrow)
	s.keeper.ImportNextID(s.ctx, s.generator.GetNextId())
}

func (s *InvariantsSuite) deleteAllEscrows() {
	for _, e := range s.escrows {
		test.DeleteEscrow(s.ctx, s.storeKey, e.Id)
	}
	s.escrows = nil
}

func (s *InvariantsSuite) reinitEscrows() {
	s.deleteAllEscrows()

	// Add 100 valid escrow
	for i := 0; i < 100; i++ {
		escrow, obj := s.generator.NewRandomTestEscrow()
		s.addEscrow(escrow)
		s.saveObject(obj, escrow.Id)
	}
	// Add 20 valid expired escrows
	for i := 0; i < 20; i++ {
		escrow, obj := s.generator.NewRandomTestEscrow()
		escrow.State = types.EscrowState_Expired
		escrow.Deadline = s.generator.NowAfter(0) - (rand.Uint64() % 5)
		s.addEscrow(escrow)
		s.saveObject(obj, escrow.Id)
	}

	//Invariant should hold
	s.expects(keeper.StateInvariant(s.keeper), false, "Valid escrows")
}

func (s *InvariantsSuite) SetupTest() {
	s.Setup(nil, true)
	s.reinitEscrows()
}

func TestInvariants(t *testing.T) {
	suite.Run(t, new(InvariantsSuite))

}

func (s *InvariantsSuite) expects(invariant sdk.Invariant, shouldBeBroken bool, testName string) {
	errorMsg, broken := invariant(s.ctx)

	if broken != shouldBeBroken {
		errorStr := "no "
		if shouldBeBroken {
			errorStr = ""
		}
		s.Fail("Expected " + errorStr + "error but got " + errorMsg + " for test " + testName)
	}
}

func (s *InvariantsSuite) TestStateInvariant() {
	invariant := keeper.AllInvariants(s.keeper)

	// Test invalid expired escrows
	escrow, obj := s.generator.NewRandomTestEscrow()
	// Escrow is not actually expired
	escrow.State = types.EscrowState_Expired
	s.addEscrow(escrow)
	s.saveObject(obj, escrow.Id)

	s.expects(invariant, true, "Escrow with expired state but with future deadline")
	s.reinitEscrows()

	escrow, obj = s.generator.NewRandomTestEscrow()
	// Escrow should be expired
	escrow.Deadline = s.generator.NowAfter(0) - 1
	s.addEscrow(escrow)
	s.saveObject(obj, escrow.Id)

	s.expects(invariant, true, "Escrow with open state but with expired deadline")
	s.reinitEscrows()

	// Escrow should be just expired
	escrow.Deadline = s.generator.NowAfter(0)
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with open state but with just expired deadline")
	s.reinitEscrows()

	// Test completed and refunded escrows
	escrow, obj = s.generator.NewRandomTestEscrow()
	escrow.State = types.EscrowState_Completed
	s.addEscrow(escrow)
	s.saveObject(obj, escrow.Id)

	s.expects(invariant, true, "Escrow with completed state in store")
	s.reinitEscrows()

	escrow.State = types.EscrowState_Refunded
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with refunded state in store")
	s.reinitEscrows()

	//Test object not in crud store
	escrow, obj = s.generator.NewRandomTestEscrow()
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with non existing object")
	s.reinitEscrows()

	// Test object not owned by module / not owned by module in store
	escrow, obj = s.generator.NewRandomTestEscrow()
	obj.Owner = s.generator.NewAccAddress()
	s.addEscrow(escrow)
	s.Assert().Nil(s.store.Create(obj))

	s.expects(invariant, true, "Escrow with object not owned by module")
	s.reinitEscrows()

	// Object is valid but not his version in store
	obj.Owner = s.keeper.GetEscrowAddress(escrow.Id)
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with object not owned by module in store version")
	s.reinitEscrows()

	// Test malformed id
	escrow, obj = s.generator.NewRandomTestEscrow()
	escrow.Id = "0123456789"
	s.addEscrow(escrow)
	s.saveObject(obj, escrow.Id)
	s.expects(invariant, true, "Escrow with invalid ID: not enough characters")
	s.reinitEscrows()

	escrow.Id = "1234567890123456789a"
	s.addEscrow(escrow)
	s.expects(invariant, true, "Escrow with invalid ID: too much characters")
	s.reinitEscrows()

	// Test invalid price
	escrow, obj = s.generator.NewRandomTestEscrow()
	price := sdk.NewCoin(test.Denom, sdk.NewInt(1))
	price.Amount = price.Amount.SubRaw(10)
	escrow.Price = sdk.Coins{price}
	s.addEscrow(escrow)
	s.saveObject(obj, escrow.Id)

	s.expects(invariant, true, "Escrow with negative price")
	s.reinitEscrows()

	// Invalid price denomination
	escrow.Price = sdk.NewCoins(sdk.NewCoin("abc", sdk.NewInt(1)))
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with invalid price denomination")
	s.reinitEscrows()

	escrow.Price = sdk.NewCoins(sdk.NewCoin(test.Denom, sdk.NewInt(1)), sdk.NewCoin("abc", sdk.NewInt(1)))
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with invalid price denomination with multiple coins")
	s.reinitEscrows()

	// Empty price
	escrow.Price = sdk.Coins{}
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with empty price")
	s.reinitEscrows()

	// Test id on the future
	escrow, obj = s.generator.NewRandomTestEscrow()
	escrow.Id = hex.EncodeToString(sdk.Uint64ToBigEndian(s.generator.GetNextId() + 500))
	s.addEscrow(escrow)
	s.saveObject(obj, escrow.Id)

	s.expects(invariant, true, "Escrow with invalid ID: this ID will be generated for future escrows")
	s.reinitEscrows()

	// Test invalid seller address
	escrow, obj = s.generator.NewRandomTestEscrow()
	escrow.Seller = "star15d4e5d4e54d5e4d"
	s.addEscrow(escrow)
	s.saveObject(obj, escrow.Id)

	s.expects(invariant, true, "Escrow with invalid seller: invalid bech32 address")
	s.reinitEscrows()

	escrow.Seller = "cosmos1veah54c394fk9y9ed95dka8npz6m77ncg5luhs"
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with invalid seller: invalid prefix in the bech32 address")
	s.reinitEscrows()

	// Test invalid broker address
	escrow.Seller = s.generator.NewAccAddress().String()
	escrow.BrokerAddress = "star15d4e5d4e54d5e4d"
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with invalid broker: invalid bech32 address")
	s.reinitEscrows()

	escrow.BrokerAddress = "cosmos1veah54c394fk9y9ed95dka8npz6m77ncg5luhs"
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with invalid broker: invalid prefix in the bech32 address")
	s.reinitEscrows()

	// Test invalid broker commission
	escrow, obj = s.generator.NewRandomTestEscrow()
	escrow.BrokerCommission = sdk.NewDec(-1)
	s.addEscrow(escrow)
	s.saveObject(obj, escrow.Id)

	s.expects(invariant, true, "Escrow with invalid commission: negative commission")
	s.reinitEscrows()

	escrow.BrokerCommission = sdk.NewDec(2)
	s.addEscrow(escrow)

	s.expects(invariant, true, "Escrow with invalid commission: commission over 100%")
	s.reinitEscrows()

	// Deadline is exceeding the maximum allowed duration for an escrow
	escrow, obj = s.generator.NewRandomTestEscrow()
	escrow.Deadline = s.generator.NowAfter(1 + uint64(s.keeper.GetMaximumEscrowDuration(s.ctx).Seconds()))
	s.addEscrow(escrow)
	s.saveObject(obj, escrow.Id)

	s.expects(invariant, true, "Escrow with invalid deadline: exceeds the maximum allowed duration for an escrow")
	s.reinitEscrows()
}
