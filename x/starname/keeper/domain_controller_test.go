package keeper

import (
	"errors"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/umma-chain/umma-core/x/configuration"
	"github.com/umma-chain/umma-core/x/starname/types"
)

func TestDomain_requireDomain(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		k, ctx, _ := NewTestKeeper(t, true)
		ds := k.DomainStore(ctx)
		ds.Create(&types.Domain{
			Name:  "test",
			Admin: AliceKey,
			Type:  types.OpenDomain,
		})
		ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
		err := ctrl.requireDomain()
		if err != nil {
			t.Fatalf("got error: %s", err)
		}
	})
	t.Run("does not exist", func(t *testing.T) {
		k, ctx, _ := NewTestKeeper(t, true)
		ds := k.DomainStore(ctx)
		ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
		err := ctrl.requireDomain()
		if !errors.Is(err, types.ErrDomainDoesNotExist) {
			t.Fatalf("want: %s, got: %s", types.ErrAccountDoesNotExist, err)
		}
	})
}

func TestDomain_domainExpired(t *testing.T) {
	t.Run("domain expired", func(t *testing.T) {
		k, ctx, _ := NewTestKeeper(t, true)
		ds := k.DomainStore(ctx)
		ds.Create(&types.Domain{
			Name:       "test",
			Admin:      AliceKey,
			Type:       types.OpenDomain,
			ValidUntil: 0,
		})
		ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
		err := ctrl.expired()
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
	})
	t.Run("domain not expired", func(t *testing.T) {
		k, ctx, _ := NewTestKeeper(t, true)
		ds := k.DomainStore(ctx)
		now := time.Now()
		ds.Create(&types.Domain{
			Name:       "test",
			Admin:      AliceKey,
			ValidUntil: now.Unix() + 10000,
		})
		ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
		err := ctrl.expired()
		if !errors.Is(err, types.ErrDomainNotExpired) {
			t.Fatalf("expected error: %s, got: %s", types.ErrDomainNotExpired, err)
		}
	})
	t.Run("domain does not exist", func(t *testing.T) {
		k, ctx, _ := NewTestKeeper(t, true)
		store := k.DomainStore(ctx)
		ctrl := NewDomainController(ctx, "test").WithDomains(&store)
		assert.Panics(t, func() { _ = ctrl.expired() }, "domain does not exists")
	})
}

func TestDomain_gracePeriodFinished(t *testing.T) {
	cases := map[string]SubTest{
		"grace period finished": {
			BeforeTestBlockTime: 1,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 1 * time.Second,
				})
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 0,
				})
			},
			TestBlockTime: 10,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				conf := k.ConfigurationKeeper.GetConfiguration(ctx)
				ctrl := NewDomainController(ctx, "test").WithDomains(&ds).WithConfiguration(conf)
				err := ctrl.gracePeriodFinished()
				if err != nil {
					t.Fatal("validation failed: grace period has not expired")
				}
			},
		},
		"grace period not finished": {
			BeforeTestBlockTime: 1,
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					DomainGracePeriod: 15 * time.Second,
				})
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 1,
				})
			},
			TestBlockTime: 3,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				conf := k.ConfigurationKeeper.GetConfiguration(ctx)
				ctrl := NewDomainController(ctx, "test").WithDomains(&ds).WithConfiguration(conf)
				err := ctrl.gracePeriodFinished()
				if !errors.Is(err, types.ErrDomainGracePeriodNotFinished) {
					t.Fatalf("expected error: %s, got: %s", types.ErrDomainGracePeriodNotFinished, err)
				}
			},
		},
	}
	RunTests(t, cases)
}

func TestDomain_ownedBy(t *testing.T) {
	cases := map[string]SubTest{
		"success": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 0,
				})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
				err := ctrl.isAdmin(AliceKey)
				if err != nil {
					t.Fatalf("got error: %s", err)
				}
			},
		},
		"unauthorized": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 0,
				})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
				err := ctrl.isAdmin(BobKey)
				if !errors.Is(err, types.ErrUnauthorized) {
					t.Fatalf("want err: %s, got: %s", types.ErrUnauthorized, err)
				}
			},
		},
	}
	RunTests(t, cases)
}

func TestDomain_notExpired(t *testing.T) {
	cases := map[string]SubTest{
		"success": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 2,
				})
			},
			TestBlockTime: 1,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
				err := ctrl.notExpired()
				if err != nil {
					t.Fatalf("got error: %s", err)
				}
			},
		},
		"expired": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 1,
				})
			},
			TestBlockTime: 2,
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
				err := ctrl.notExpired()
				if !errors.Is(err, types.ErrDomainExpired) {
					t.Fatalf("want err: %s, got: %s", types.ErrDomainExpired, err)
				}
			},
		},
	}
	RunTests(t, cases)
}

func TestDomain_type(t *testing.T) {
	cases := map[string]SubTest{
		"saved": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:  "test",
					Admin: AliceKey,
					Type:  types.ClosedDomain,
				})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
				err := ctrl.dType(types.ClosedDomain)
				if err != nil {
					t.Fatalf("got error: %s", err)
				}
			},
		},
		"fail want type close domain": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:  "test",
					Admin: AliceKey,
					Type:  types.ClosedDomain,
				})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
				err := ctrl.dType(types.OpenDomain)
				if !errors.Is(err, types.ErrInvalidDomainType) {
					t.Fatalf("want err: %s, got: %s", types.ErrInvalidDomainType, err)
				}
			},
		},
		"fail want open domain": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:  "test",
					Admin: AliceKey,
					Type:  types.OpenDomain,
				})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				ds := k.DomainStore(ctx)
				ctrl := NewDomainController(ctx, "test").WithDomains(&ds)
				err := ctrl.dType(types.ClosedDomain)
				if !errors.Is(err, types.ErrInvalidDomainType) {
					t.Fatalf("want err: %s, got: %s", types.ErrInvalidDomainType, err)
				}
			},
		},
	}
	RunTests(t, cases)
}

func TestDomain_validName(t *testing.T) {
	cases := map[string]SubTest{
		"success": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					ValidDomainName: RegexMatchAll,
				})
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 0,
				})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				conf := k.ConfigurationKeeper.GetConfiguration(ctx)
				ctrl := NewDomainController(ctx, "test").WithConfiguration(conf)
				err := ctrl.validName()
				if err != nil {
					t.Fatalf("got error: %s", err)
				}
			},
		},
		"invalid name": {
			BeforeTest: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
				setConfig(ctx, configuration.Config{
					ValidDomainName: RegexMatchNothing,
				})
				ds := k.DomainStore(ctx)
				ds.Create(&types.Domain{
					Name:       "test",
					Admin:      AliceKey,
					ValidUntil: 0,
				})
			},
			Test: func(t *testing.T, k Keeper, ctx sdk.Context, mocks *Mocks) {
				conf := k.ConfigurationKeeper.GetConfiguration(ctx)
				ctrl := NewDomainController(ctx, "test").WithConfiguration(conf)
				err := ctrl.validName()
				if !errors.Is(err, types.ErrInvalidDomainName) {
					t.Fatalf("want err: %s, got: %s", types.ErrInvalidDomainName, err)
				}
			},
		},
	}
	RunTests(t, cases)
}

func TestDomain_Renewable(t *testing.T) {
	k, ctx, _ := NewTestKeeper(t, true)
	ctx = ctx.WithBlockTime(time.Unix(1, 0))
	setConfig := GetConfigSetter(k.ConfigurationKeeper).SetConfig
	setConfig(ctx, configuration.Config{
		DomainGracePeriod:     100 * time.Second,
		DomainRenewalCountMax: 1, // increased by one inside controller
		DomainRenewalPeriod:   10 * time.Second,
	})
	conf := k.ConfigurationKeeper.GetConfiguration(ctx)
	ds := k.DomainStore(ctx)
	ds.Create(&types.Domain{
		Name:       "open",
		Admin:      AliceKey,
		ValidUntil: time.Unix(18, 0).Unix(),
		Type:       types.OpenDomain,
	})
	ds.Create(&types.Domain{
		Name:       "closed",
		Admin:      AliceKey,
		ValidUntil: time.Unix(18, 0).Unix(),
		Type:       types.ClosedDomain,
	})
	ds.Create(&types.Domain{
		Name:       "deadline-exceeded",
		Admin:      AliceKey,
		ValidUntil: time.Unix(10, 0).Unix(),
		Type:       types.ClosedDomain,
	})
	// 120(DomainValidUntil) + 10(DomainRP) = 130 newValidUntil
	t.Run("beyond grace period", func(t *testing.T) {
		d := NewDomainController(ctx.WithBlockTime(time.Unix(241, 0)), "deadline-exceeded").WithDomains(&ds).WithConfiguration(conf)
		err := d.Renewable().Validate()
		if !errors.Is(err, types.ErrRenewalDeadlineExceeded) {
			t.Fatalf("want: %s, got: %s", types.ErrRenewalDeadlineExceeded, err)
		}
	})

	// 18(DomainValidUntil) + 10 (DomainRP) = 28 newValidUntil
	t.Run("open domain", func(t *testing.T) {
		// 7(time) + 2(DomainRCM) * 10(DomainRP) = 27 maxValidUntil
		d := NewDomainController(ctx.WithBlockTime(time.Unix(7, 0)), "open").WithDomains(&ds).WithConfiguration(conf)
		err := d.Renewable().Validate()
		if !errors.Is(err, types.ErrUnauthorized) {
			t.Fatalf("want: %s, got: %s", types.ErrUnauthorized, err)
		}
		// 100(time) + 2(DomainRCM) * 10(DomainRP) = 120 maxValidUntil
		d = NewDomainController(ctx.WithBlockTime(time.Unix(100, 0)), "open").WithDomains(&ds).WithConfiguration(conf)
		if err := d.Renewable().Validate(); err != nil {
			t.Fatalf("got error: %s", err)
		}
	})
	// 18(DomainValidUntil) + 10 (DomainRP) = 28 newValidUntil
	t.Run("closed domain", func(t *testing.T) {
		// 7(time) + 2(DomainRCM) * 10(DomainRP) = 27 maxValidUntil
		d := NewDomainController(ctx.WithBlockTime(time.Unix(7, 0)), "closed").WithDomains(&ds).WithConfiguration(conf)
		err := d.Renewable().Validate()
		if !errors.Is(err, types.ErrUnauthorized) {
			t.Fatalf("want: %s, got: %s", types.ErrUnauthorized, err)
		}
		// 100(time) + 2(DomainRCM) * 10(DomainRP) = 120 maxValidUntil
		d = NewDomainController(ctx.WithBlockTime(time.Unix(100, 0)), "closed").WithDomains(&ds).WithConfiguration(conf)
		if err := d.Renewable().Validate(); err != nil {
			t.Fatalf("got error: %s", err)
		}
	})
}
