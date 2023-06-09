package keeper

import (
	"errors"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umma-chain/umma-core/pkg/utils"
	"github.com/umma-chain/umma-core/x/configuration"
	"github.com/umma-chain/umma-core/x/starname/types"
)

func Test_Closed_handleMsgDomainDelete(t *testing.T) {
	cases := map[string]SubTest{
		"success only admin can delete before grace period": {
			BeforeTestBlockTime: 1,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 10 * time.Second,
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 2,
					Type:       types.ClosedDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
				NewAccountExecutor(ctx, types.Account{
					Domain: "test",
					Name:   utils.StrPtr("1"),
					Owner:  BobKey,
				}).WithAccounts(&accounts).Create()
				NewAccountExecutor(ctx, types.Account{
					Domain: "test",
					Name:   utils.StrPtr("2"),
					Owner:  BobKey,
				}).WithAccounts(&accounts).Create()
			},
			TestBlockTime: 3,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  BobKey.String(),
				}.ToInternal())
				if !errors.Is(err, types.ErrUnauthorized) {
					t.Fatalf("unexpected error: %s", err)
				}
				_, err = deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  AliceKey.String(),
				}.ToInternal())
				if err != nil {
					t.Fatalf("deleteDomain() got error: %s", err)
				}
			},
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				if err := k.DomainStore(ctx).Read((&types.Domain{Name: "test"}).PrimaryKey(), new(types.Domain)); err == nil {
					t.Fatalf("deleteDomain() domain should not exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("1")}).PrimaryKey(), new(types.Account)); err == nil {
					t.Fatalf("deleteDomain() account 1 should not exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("2")}).PrimaryKey(), new(types.Account)); err == nil {
					t.Fatalf("deleteDomain() account 2 should not exist")
				}
			},
		},
		"success anyone can after grace period": {
			BeforeTestBlockTime: 1,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 10 * time.Second,
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 2,
					Type:       types.ClosedDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
				NewAccountExecutor(ctx, types.Account{
					Domain: "test",
					Name:   utils.StrPtr("1"),
					Owner:  BobKey,
				}).WithAccounts(&accounts).Create()
				NewAccountExecutor(ctx, types.Account{
					Domain: "test",
					Name:   utils.StrPtr("2"),
					Owner:  BobKey,
				}).WithAccounts(&accounts).Create()
			},
			TestBlockTime: 1000,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  CharlieKey.String(),
				}.ToInternal())
				if err != nil {
					t.Fatalf("deleteDomain() got error: %s", err)
				}
			},
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				if err := k.DomainStore(ctx).Read((&types.Domain{Name: "test"}).PrimaryKey(), new(types.Domain)); err == nil {
					t.Fatalf("deleteDomain() domain should not exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("1")}).PrimaryKey(), new(types.Account)); err == nil {
					t.Fatalf("deleteDomain() account 1 should not exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("2")}).PrimaryKey(), new(types.Account)); err == nil {
					t.Fatalf("deleteDomain() account 2 should not exist")
				}
			},
		},
	}
	RunTests(t, cases)
}

func Test_Open_handleMsgDomainDelete(t *testing.T) {
	cases := map[string]SubTest{
		"success anyone can after grace period": {
			BeforeTestBlockTime: 1,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 10 * time.Second,
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 2,
					Type:       types.OpenDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
				NewAccountExecutor(ctx, types.Account{
					Domain: "test",
					Name:   utils.StrPtr("1"),
					Owner:  BobKey,
				}).WithAccounts(&accounts).Create()
				NewAccountExecutor(ctx, types.Account{
					Domain: "test",
					Name:   utils.StrPtr("2"),
					Owner:  BobKey,
				}).WithAccounts(&accounts).Create()
			},
			TestBlockTime: 1000,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  CharlieKey.String(),
				}.ToInternal())
				if err != nil {
					t.Fatalf("deleteDomain() got error: %s", err)
				}
			},
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				if err := k.DomainStore(ctx).Read((&types.Domain{Name: "test"}).PrimaryKey(), new(types.Domain)); err == nil {
					t.Fatalf("deleteDomain() domain should not exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("1")}).PrimaryKey(), new(types.Account)); err == nil {
					t.Fatalf("deleteDomain() account 1 should not exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("2")}).PrimaryKey(), new(types.Account)); err == nil {
					t.Fatalf("deleteDomain() account 2 should not exist")
				}
			},
		},
		"domain cannot be deleted before grace period even by admin": {
			BeforeTestBlockTime: 1,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 10 * time.Second,
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 2,
					Type:       types.OpenDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
				NewAccountExecutor(ctx, types.Account{
					Domain: "test",
					Name:   utils.StrPtr("1"),
					Owner:  BobKey,
				}).WithAccounts(&accounts).Create()
				NewAccountExecutor(ctx, types.Account{
					Domain: "test",
					Name:   utils.StrPtr("2"),
					Owner:  BobKey,
				}).WithAccounts(&accounts).Create()
			},
			TestBlockTime: 3,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  CharlieKey.String(),
				}.ToInternal())
				if !errors.Is(err, types.ErrDomainGracePeriodNotFinished) {
					t.Fatalf("unexpected error: %s", err)
				}
				_, err = deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  AliceKey.String(),
				}.ToInternal())
				if !errors.Is(err, types.ErrDomainGracePeriodNotFinished) {
					t.Fatalf("unexpected error: %s", err)
				}
			},
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				if err := k.DomainStore(ctx).Read((&types.Domain{Name: "test"}).PrimaryKey(), new(types.Domain)); err != nil {
					t.Fatalf("deleteDomain() domain should exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("1")}).PrimaryKey(), new(types.Account)); err != nil {
					t.Fatalf("deleteDomain() account 1 should exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("2")}).PrimaryKey(), new(types.Account)); err != nil {
					t.Fatalf("deleteDomain() account 2 should exist")
				}
			},
		},
	}
	RunTests(t, cases)
}

func Test_handleMsgDomainDelete(t *testing.T) {
	cases := map[string]SubTest{
		"fail domain does not exist": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				GetConfigSetter(k.ConfigurationKeeper).SetConfig(ctx, configuration.Config{})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "this does not exist",
					Owner:  BobKey.String(),
				}.ToInternal())
				if !errors.Is(err, types.ErrDomainDoesNotExist) {
					t.Fatalf("deleteDomain() expected error: %s, got: %s", types.ErrDomainDoesNotExist, err)
				}
			},
		},
		"fail domain admin does not match msg owner": {
			BeforeTestBlockTime: 0,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 1000000000000000,
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Admin:      BobKey,
					ValidUntil: 0,
					Type:       types.ClosedDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
			},
			TestBlockTime: 1,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  AliceKey.String(),
				}.ToInternal())
				if !errors.Is(err, types.ErrUnauthorized) {
					t.Fatalf("deleteDomain() expected error: %s, got: %s", types.ErrUnauthorized, err)
				}
			},
			AfterTest: nil,
		},
		"fail domain grace period not over": {
			BeforeTestBlockTime: 0,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 5,
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Admin:      BobKey,
					ValidUntil: 3,
					Type:       types.ClosedDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
			},
			TestBlockTime: 3,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  AliceKey.String(),
				}.ToInternal())
				if !errors.Is(err, types.ErrUnauthorized) {
					t.Fatalf("deleteDomain() expected error: %s, got: %s", types.ErrUnauthorized, err)
				}
			},
			AfterTest: nil,
		},
		"success domain grace period over": {
			BeforeTestBlockTime: 0,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 5,
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Admin:      BobKey,
					ValidUntil: 4,
					Type:       types.ClosedDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
			},
			TestBlockTime: 10,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  AliceKey.String(),
				}.ToInternal())
				if err != nil {
					t.Fatalf("deleteDomain() got error: %s", err)
				}
			},
			AfterTest: nil,
		},
		"success owner can delete one of the domains after one expires and deleted": {
			BeforeTestBlockTime: 1589826438,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 1,
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test1",
					Admin:      BobKey,
					ValidUntil: 1589826439,
					Type:       types.ClosedDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test2",
					Admin:      BobKey,
					ValidUntil: 1589828251,
					Type:       types.ClosedDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
			},
			TestBlockTime: 1589826441,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				// another user can delete expired domain
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test1",
					Owner:  AliceKey.String(),
				}.ToInternal())
				if err != nil {
					t.Fatalf("deleteDomain() got error: %s", err)
				}
				_, err = deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test2",
					Owner:  AliceKey.String(),
				}.ToInternal())
				if !errors.Is(err, types.ErrUnauthorized) {
					t.Fatalf("deleteDomain() expected error: %s, got: %s", types.ErrUnauthorized, err)
				}
				_, err = deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test2",
					Owner:  BobKey.String(),
				}.ToInternal())
				if err != nil {
					t.Fatalf("deleteDomain() got error: %s", err)
				}
			},
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {

			},
		},
		"success owner can delete their domain before grace period": {
			BeforeTestBlockTime: 0,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 1000000000000000, // unexpired domain
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx) // set domain
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 0,
					Type:       types.ClosedDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
			},
			TestBlockTime: 4,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  AliceKey.String(),
				}.ToInternal())
				if err != nil {
					t.Fatalf("deleteDomain() got error: %s", err)
				}
			},
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				if err := k.DomainStore(ctx).Read((&types.Domain{Name: "test"}).PrimaryKey(), new(types.Domain)); err == nil {
					t.Fatalf("deleteDomain() domain should not exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("1")}).PrimaryKey(), new(types.Account)); err == nil {
					t.Fatalf("deleteDomain() account 1 should not exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("2")}).PrimaryKey(), new(types.Account)); err == nil {
					t.Fatalf("deleteDomain() account 2 should not exist")
				}
			},
		},
		"success claim expired": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 1,
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				// set domain
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 0,
					Type:       types.ClosedDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
				// add two accounts
				NewAccountExecutor(ctx, types.Account{
					Domain: "test",
					Name:   utils.StrPtr("1"),
					Owner:  BobKey,
				}).WithAccounts(&accounts).Create()
				// add two accounts
				NewAccountExecutor(ctx, types.Account{
					Domain: "test",
					Name:   utils.StrPtr("2"),
					Owner:  BobKey,
				}).WithAccounts(&accounts).Create()
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := deleteDomain(ctx, k, types.MsgDeleteDomain{
					Domain: "test",
					Owner:  BobKey.String(),
				}.ToInternal())
				if err != nil {
					t.Fatalf("deleteDomain() got error: %s", err)
				}
			},
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				if err := k.DomainStore(ctx).Read((&types.Domain{Name: "test"}).PrimaryKey(), new(types.Domain)); err == nil {
					t.Fatalf("deleteDomain() domain should not exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("1")}).PrimaryKey(), new(types.Account)); err == nil {
					t.Fatalf("deleteDomain() account 1 should not exist")
				}
				if err := k.AccountStore(ctx).Read((&types.Account{Domain: "test", Name: utils.StrPtr("2")}).PrimaryKey(), new(types.Account)); err == nil {
					t.Fatalf("deleteDomain() account 2 should not exist")
				}
			},
		},
	}
	RunTests(t, cases)
}

func TestHandleMsgRegisterDomain(t *testing.T) {
	testCases := map[string]SubTest{
		"success": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				configSetter := GetConfigSetter(k.ConfigurationKeeper)
				// set config
				configSetter.SetConfig(ctx, configuration.Config{
					Configurer:      AliceKey.String(),
					ValidDomainName: "^(.*?)?",
				})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := registerDomain(ctx, k, types.MsgRegisterDomain{
					Name:       "domain-closed",
					DomainType: types.ClosedDomain,
					Admin:      BobKey.String(),
				}.ToInternal())
				if err != nil {
					t.Fatalf("registerDomain() with close domain, got error: %s", err)
				}
				_, err = registerDomain(ctx, k, types.MsgRegisterDomain{
					Name:       "domain-open",
					Admin:      AliceKey.String(),
					DomainType: types.OpenDomain,
					Broker:     "",
				}.ToInternal())
				if err != nil {
					t.Fatalf("registerDomain() with open domain, got error: %s", err)
				}
			},
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				// TODO do reflect.DeepEqual checks on expected results vs results returned
				if err := k.DomainStore(ctx).Read((&types.Domain{Name: "domain-closed"}).PrimaryKey(), new(types.Domain)); err != nil {
					t.Fatalf("registerDomain() could not find 'domain-closed'")
				}
				if err := k.DomainStore(ctx).Read((&types.Domain{Name: "domain-open"}).PrimaryKey(), new(types.Domain)); err != nil {
					t.Fatalf("registerDomain() could not find 'domain-open'")
				}
			},
		},
		"fail domain name exists": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				GetConfigSetter(k.ConfigurationKeeper).SetConfig(ctx, configuration.Config{})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)

				NewDomainExecutor(ctx, types.Domain{
					Name:       "exists",
					Admin:      BobKey,
					ValidUntil: 0,
					Type:       types.ClosedDomain,
					Broker:     nil,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := registerDomain(ctx, k, types.MsgRegisterDomain{
					Name:       "exists",
					Admin:      AliceKey.String(),
					DomainType: types.ClosedDomain,
				}.ToInternal())
				if !errors.Is(err, types.ErrDomainAlreadyExists) {
					t.Fatalf("registerDomain() expected: %s got: %s", types.ErrDomainAlreadyExists, err)
				}
			},
			AfterTest: nil,
		},
		"fail domain does not match valid domain regexp": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				// get set config function
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				// set configs with a domain regexp that matches nothing
				setConfig(ctx, configuration.Config{
					ValidDomainName:     "$^",
					DomainRenewalPeriod: 0,
				})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := registerDomain(ctx, k, types.MsgRegisterDomain{
					Name:       "invalid-name",
					Admin:      "",
					DomainType: types.OpenDomain,
					Broker:     "",
				}.ToInternal())
				if !errors.Is(err, types.ErrInvalidDomainName) {
					t.Fatalf("registerDomain() expected error: %s, got: %s", types.ErrInvalidDomainName, err)
				}
			},
			// TODO ADD AFTER TEST
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {

			},
		},
	}
	// run all test cases
	RunTests(t, testCases)
}

func Test_handlerDomainRenew(t *testing.T) {
	cases := map[string]SubTest{
		"domain not found": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				GetConfigSetter(k.ConfigurationKeeper).SetConfig(ctx, configuration.Config{})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := renewDomain(ctx, k, types.MsgRenewDomain{Domain: "does not exist"}.ToInternal())
				if !errors.Is(err, types.ErrDomainDoesNotExist) {
					t.Fatalf("renewDomain() expected error: %s, got: %s", types.ErrDomainDoesNotExist, err)
				}
			},
			AfterTest: nil,
		},
		"success": {
			BeforeTestBlockTime: 1000,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				// add config
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainRenewalCountMax: 2,
					DomainRenewalPeriod:   1 * time.Second,
					DomainGracePeriod:     10 * time.Second,
				})
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				// add domain
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					ValidUntil: 1000,
					Admin:      BobKey,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := renewDomain(ctx, k, types.MsgRenewDomain{Domain: "test"}.ToInternal())
				if err != nil {
					t.Fatalf("renewDomain() got error: %s", err)
				}
			},
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				// get domain
				domain := new(types.Domain)
				_ = k.DomainStore(ctx).Read((&types.Domain{Name: "test"}).PrimaryKey(), domain)
				if domain.ValidUntil != 1001 {
					t.Fatalf("renewDomain() expected 1001, got: %d", domain.ValidUntil)
				}
			},
		},
	}
	// run tests
	RunTests(t, cases)
}

func Test_transferDomain(t *testing.T) {
	cases := map[string]SubTest{
		"domain does not exist": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {

			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := transferDomain(ctx, k, types.MsgTransferDomain{
					Domain:   "does not exist",
					Owner:    "",
					NewAdmin: "",
				}.ToInternal())
				if !errors.Is(err, types.ErrDomainDoesNotExist) {
					t.Fatalf("transferDomain() expected error: %s, got error: %s", types.ErrDomainDoesNotExist, err)
				}
			},
			AfterTest: nil,
		},
		/*
			"domain type open": {
				BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
					NewDomainExecutor(ctx, types.Domain{
						Name:  "test",
						Type:  types.OpenDomain,
						Admin: AliceKey,
					}).Create()
				},
				Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
					_, err := transferDomain(ctx, k, &types.MsgTransferDomain{
						Domain:   "test",
						Owner:    nil,
						NewAdmin: nil,
					})
					if !errors.Is(err, types.ErrInvalidDomainType) {
						t.Fatalf("transferDomain() expected error: %s, got error: %s", types.ErrInvalidDomainType, err)
					}
				},
				AfterTest: nil,
			},
		*/
		"domain type closed": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:  "test",
					Type:  types.ClosedDomain,
					Admin: AliceKey,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := transferDomain(ctx, k, types.MsgTransferDomain{
					Domain:   "test",
					Owner:    "",
					NewAdmin: "",
				}.ToInternal())
				if !errors.Is(err, types.ErrUnauthorized) {
					t.Fatalf("transferDomain() expected error: %s, got error: %s", types.ErrUnauthorized, err)
				}
			},
			AfterTest: nil,
		},
		"domain has expired": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:  "test",
					Type:  types.ClosedDomain,
					Admin: BobKey,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := transferDomain(ctx, k, types.MsgTransferDomain{
					Domain:   "test",
					Owner:    BobKey.String(),
					NewAdmin: "",
				}.ToInternal())
				if !errors.Is(err, types.ErrDomainExpired) {
					t.Fatalf("transferDomain() expected error: %s, got error: %s", types.ErrDomainExpired, err)
				}
			},
			AfterTest: nil,
		},
		"msg signer is not domain admin": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Type:       types.ClosedDomain,
					ValidUntil: utils.TimeToSeconds(ctx.BlockTime().Add(1000 * time.Hour)),
					Admin:      AliceKey,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := transferDomain(ctx, k, types.MsgTransferDomain{
					Domain:   "test",
					Owner:    BobKey.String(),
					NewAdmin: "",
				}.ToInternal())
				if !errors.Is(err, types.ErrUnauthorized) {
					t.Fatalf("transferDomain() expected error: %s, got error: %s", types.ErrUnauthorized, err)
				}
			},
			AfterTest: nil,
		},
		"success": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				domains := k.DomainStore(ctx)
				accounts := k.AccountStore(ctx)
				// create domain
				NewDomainExecutor(ctx, types.Domain{
					Name:       "test",
					Type:       types.ClosedDomain,
					ValidUntil: utils.TimeToSeconds(ctx.BlockTime().Add(1000 * time.Hour)),
					Admin:      AliceKey,
				}).WithDomains(&domains).WithAccounts(&accounts).Create()
				// add account 1
				NewAccountExecutor(ctx, types.Account{
					Domain:     "test",
					Name:       utils.StrPtr("1"),
					Owner:      AliceKey,
					ValidUntil: 0,
					Resources: []*types.Resource{{
						URI:      "test",
						Resource: "test",
					}},
					Certificates: [][]byte{[]byte("cert")},
					Broker:       nil,
				}).WithAccounts(&accounts).Create()
				// add account 2
				NewAccountExecutor(ctx, types.Account{
					Domain:     "test",
					Name:       utils.StrPtr("2"),
					Owner:      AliceKey,
					ValidUntil: 0,
					Resources: []*types.Resource{{
						URI:      "test",
						Resource: "test",
					}},
					Certificates: [][]byte{[]byte("cert")},
					Broker:       nil,
				}).WithAccounts(&accounts).Create()
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				_, err := transferDomain(ctx, k, types.MsgTransferDomain{
					Domain:       "test",
					Owner:        AliceKey.String(),
					NewAdmin:     BobKey.String(),
					TransferFlag: types.TransferOwned,
				}.ToInternal())
				if err != nil {
					t.Fatalf("transferDomain() got error: %s", err)
				}
			},
			AfterTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {

			},
		},
	}

	RunTests(t, cases)
}
