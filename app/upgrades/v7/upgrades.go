package v7

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/persistenceOne/persistenceCore/v8/app/upgrades"
)

func CreateUpgradeHandler(args upgrades.UpgradeHandlerArgs) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		ctx.Logger().Info("start to run module migrations...")

		return args.ModuleManager.RunMigrations(ctx, args.Configurator, vm)
	}
}
