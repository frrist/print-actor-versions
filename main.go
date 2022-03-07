package main

import (
	"fmt"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"
	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"
	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"
	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"
)

func main() {
	fmt.Println("paych 7", builtin7.PaymentChannelActorCodeID.String())
	fmt.Println("paych 6", builtin6.PaymentChannelActorCodeID.String())
	fmt.Println("paych 5", builtin5.PaymentChannelActorCodeID.String())
	fmt.Println("paych 4", builtin4.PaymentChannelActorCodeID.String())
	fmt.Println("paych 3", builtin3.PaymentChannelActorCodeID.String())
	fmt.Println("paych 2", builtin2.PaymentChannelActorCodeID.String())

	fmt.Println("miner 7", builtin7.StorageMinerActorCodeID.String())
	fmt.Println("miner 6", builtin6.StorageMinerActorCodeID.String())
	fmt.Println("miner 5", builtin5.StorageMinerActorCodeID.String())
	fmt.Println("miner 4", builtin4.StorageMinerActorCodeID.String())
	fmt.Println("miner 3", builtin3.StorageMinerActorCodeID.String())
	fmt.Println("miner 2", builtin2.StorageMinerActorCodeID.String())

}
