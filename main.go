package main

import (
	"fmt"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"
	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"
	// Attempts to get v8 and v9 to work
	// TODO: Figure out why this doesn't work
	builtin8 "github.com/filecoin-project/specs-actors/v8/actors/builtin"
	builtin9 "github.com/filecoin-project/go-state-types/builtin"
	migration9 "github.com/filecoin-project/go-state-types/builtin/v9/migration"
	miner9 "github.com/filecoin-project/go-state-types/builtin/v9/miner"
)

func main() {
	// the Miner Actor
	fmt.Println("Miner Actor", builtin9.MethodsMiner)
	fmt.Println("| 2 | ", builtin2.StorageMinerActorCodeID.String(), " |")
	fmt.Println("| 3 | ", builtin3.StorageMinerActorCodeID.String(), " |")
	fmt.Println("| 4 | ", builtin4.StorageMinerActorCodeID.String(), " |")
	fmt.Println("| 5 | ", builtin5.StorageMinerActorCodeID.String(), " |")
	fmt.Println("| 6 | ", builtin6.StorageMinerActorCodeID.String(), " |")
	fmt.Println("| 7 | ", builtin7.StorageMinerActorCodeID.String(), " |")
	fmt.Println("| 8 | ", builtin8.StorageMinerActorCodeID.String(), " |")
	// Get v9's CodeID
	v9_info := migration9.migratedCodeCID()
	fmt.Println("| 9 | ", v9_info, " |")

	// the Account Actor
	fmt.Println("\nAccount Actor", builtin9.MethodsAccount)
	fmt.Println("| 2 | ", builtin2.AccountActorCodeID.String(), " |")
	fmt.Println("| 3 | ", builtin3.AccountActorCodeID.String(), " |")
	fmt.Println("| 4 | ", builtin4.AccountActorCodeID.String(), " |")
	fmt.Println("| 5 | ", builtin5.AccountActorCodeID.String(), " |")
	fmt.Println("| 6 | ", builtin6.AccountActorCodeID.String(), " |")
	fmt.Println("| 7 | ", builtin7.AccountActorCodeID.String(), " |")
	fmt.Println("| 8 | ", builtin8.AccountActorCodeID.String(), " |")
	// fmt.Println("account 9", builtin9.AccountActorCodeID.String(), "\n")

	// the Init Actor
	fmt.Println("\nInit Actor", builtin9.MethodsInit)
	fmt.Println("| 2 | ", builtin2.InitActorCodeID.String(), " |")
	fmt.Println("| 3 | ", builtin3.InitActorCodeID.String(), " |")
	fmt.Println("| 4 | ", builtin4.InitActorCodeID.String(), " |")
	fmt.Println("| 5 | ", builtin5.InitActorCodeID.String(), " |")
	fmt.Println("| 6 | ", builtin6.InitActorCodeID.String(), " |")
	fmt.Println("| 7 | ", builtin7.InitActorCodeID.String(), " |")
	fmt.Println("| 8 | ", builtin8.InitActorCodeID.String(), " |")

	// the Cron Actor
	fmt.Println("\nCron Actor", builtin9.MethodsCron)
	fmt.Println("| 2 | ", builtin2.CronActorCodeID.String(), " |")
	fmt.Println("| 3 | ", builtin3.CronActorCodeID.String(), " |")
	fmt.Println("| 4 | ", builtin4.CronActorCodeID.String(), " |")
	fmt.Println("| 5 | ", builtin5.CronActorCodeID.String(), " |")
	fmt.Println("| 6 | ", builtin6.CronActorCodeID.String(), " |")
	fmt.Println("| 7 | ", builtin7.CronActorCodeID.String(), " |")
	fmt.Println("| 8 | ", builtin8.CronActorCodeID.String(), " |")

	// the Market Actor
	fmt.Println("\nMarket Actor", builtin9.MethodsMarket)
	fmt.Println("| 2 | ", builtin2.StorageMarketActorCodeID.String(), " |")
	fmt.Println("| 3 | ", builtin3.StorageMarketActorCodeID.String(), " |")
	fmt.Println("| 4 | ", builtin4.StorageMarketActorCodeID.String(), " |")
	fmt.Println("| 5 | ", builtin5.StorageMarketActorCodeID.String(), " |")
	fmt.Println("| 6 | ", builtin6.StorageMarketActorCodeID.String(), " |")
	fmt.Println("| 7 | ", builtin7.StorageMarketActorCodeID.String(), " |")
	fmt.Println("| 8 | ", builtin8.StorageMarketActorCodeID.String(), " |")
	// print 9 separately with a newline at the end
	fmt.Println("| 9 | ", builtin9.StorageMarketActorAddr.String(), " |\n")

	// the Multisig Actor
	fmt.Println("\nMultisig Actor", builtin9.MethodsMultisig)
	fmt.Println("| 2 | ", builtin2.MultisigActorCodeID.String(), " |")
	fmt.Println("| 3 | ", builtin3.MultisigActorCodeID.String(), " |")
	fmt.Println("| 4 | ", builtin4.MultisigActorCodeID.String(), " |")
	fmt.Println("| 5 | ", builtin5.MultisigActorCodeID.String(), " |")
	fmt.Println("| 6 | ", builtin6.MultisigActorCodeID.String(), " |")
	fmt.Println("| 7 | ", builtin7.MultisigActorCodeID.String(), " |")
	fmt.Println("| 8 | ", builtin8.MultisigActorCodeID.String(), " |")
	// print 9 separately with a newline at the end
	// fmt.Println("| 9 | ", builtin9.MultisigActorCodeID.String(), " |\n")

	// the Payment Channel Actor
	fmt.Println("\nPayment Channel Actor", builtin9.MethodsPaych)
	fmt.Println("| 2 | ", builtin2.PaymentChannelActorCodeID.String(), " |")
	fmt.Println("| 3 | ", builtin3.PaymentChannelActorCodeID.String(), " |")
	fmt.Println("| 4 | ", builtin4.PaymentChannelActorCodeID.String(), " |")
	fmt.Println("| 5 | ", builtin5.PaymentChannelActorCodeID.String(), " |")
	fmt.Println("| 6 | ", builtin6.PaymentChannelActorCodeID.String(), " |")
	fmt.Println("| 7 | ", builtin7.PaymentChannelActorCodeID.String(), " |")
	fmt.Println("| 8 | ", builtin8.PaymentChannelActorCodeID.String(), " |")
	// print 9 separately with a newline at the end
	// fmt.Println("| 9 | ", builtin9.PaymentChannelActorCodeID.String(), " |\n")

	// The Power Actor
	fmt.Println("\nPower Actor", builtin9.MethodsPower)
	fmt.Println("| 2 | ", builtin2.StoragePowerActorCodeID.String(), " |")
	fmt.Println("| 3 | ", builtin3.StoragePowerActorCodeID.String(), " |")
	fmt.Println("| 4 | ", builtin4.StoragePowerActorCodeID.String(), " |")
	fmt.Println("| 5 | ", builtin5.StoragePowerActorCodeID.String(), " |")
	fmt.Println("| 6 | ", builtin6.StoragePowerActorCodeID.String(), " |")
	fmt.Println("| 7 | ", builtin7.StoragePowerActorCodeID.String(), " |")
	fmt.Println("| 8 | ", builtin8.StoragePowerActorCodeID.String(), " |")
	// print 9 separately with a newline at the end
	// fmt.Println("| 9 | ", builtin9.StoragePowerActorCodeID.String(), " |\n")

	// The Reward Actor
	fmt.Println("\nReward Actor", builtin9.MethodsReward)
	fmt.Println("| 2 | ", builtin2.RewardActorCodeID.String(), " |")
	fmt.Println("| 3 | ", builtin3.RewardActorCodeID.String(), " |")
	fmt.Println("| 4 | ", builtin4.RewardActorCodeID.String(), " |")
	fmt.Println("| 5 | ", builtin5.RewardActorCodeID.String(), " |")
	fmt.Println("| 6 | ", builtin6.RewardActorCodeID.String(), " |")
	fmt.Println("| 7 | ", builtin7.RewardActorCodeID.String(), " |")
	fmt.Println("| 8 | ", builtin8.RewardActorCodeID.String(), " |")
	// print 9 separately with a newline at the end
	// fmt.Println("| 9 | ", builtin9.RewardActorCodeID.String(), " |\n")

	// The System Actor
	fmt.Println("\nSystem Actor")
	fmt.Println("| 2 | ", builtin2.SystemActorCodeID.String(), " |")
	fmt.Println("| 3 | ", builtin3.SystemActorCodeID.String(), " |")
	fmt.Println("| 4 | ", builtin4.SystemActorCodeID.String(), " |")
	fmt.Println("| 5 | ", builtin5.SystemActorCodeID.String(), " |")
	fmt.Println("| 6 | ", builtin6.SystemActorCodeID.String(), " |")
	fmt.Println("| 7 | ", builtin7.SystemActorCodeID.String(), " |")
	fmt.Println("| 8 | ", builtin8.SystemActorCodeID.String(), " |")
	// print 9 separately with a newline at the end
	// fmt.Println("| 9 | ", builtin9.SystemActorCodeID.String(), " |\n")
	
}
