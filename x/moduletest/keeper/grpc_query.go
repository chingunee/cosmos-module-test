package keeper

import (
	"module-test/x/moduletest/types"
)

var _ types.QueryServer = Keeper{}
