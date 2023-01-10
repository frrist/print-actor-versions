package main

import (
	"fmt"
	"reflect"

	builtin9 "github.com/filecoin-project/go-state-types/builtin"
)

func main() {
	// Print | Method Name | Method ID |
	// Account Methods
	fmt.Println("\nAccount Methods")
	for i := 0; i < reflect.TypeOf(builtin9.MethodsAccount).NumField(); i++ {
		fmt.Printf("| %s | %v |\n", reflect.TypeOf(builtin9.MethodsAccount).Field(i).Name, reflect.ValueOf(builtin9.MethodsAccount).Field(i).Interface())
	}

	// Init Methods
	fmt.Println("\nInit Methods")
	for i := 0; i < reflect.TypeOf(builtin9.MethodsInit).NumField(); i++ {
		fmt.Printf("| %s | %v |\n", reflect.TypeOf(builtin9.MethodsInit).Field(i).Name, reflect.ValueOf(builtin9.MethodsInit).Field(i).Interface())
	}

	// Cron Methods
	fmt.Println("\nCron Methods")
	for i := 0; i < reflect.TypeOf(builtin9.MethodsCron).NumField(); i++ {
		fmt.Printf("| %s | %v |\n", reflect.TypeOf(builtin9.MethodsCron).Field(i).Name, reflect.ValueOf(builtin9.MethodsCron).Field(i).Interface())
	}

	// Multisig Methods
	fmt.Println("\nMultisig Methods")
	for i := 0; i < reflect.TypeOf(builtin9.MethodsMultisig).NumField(); i++ {
		fmt.Printf("| %s | %v |\n", reflect.TypeOf(builtin9.MethodsMultisig).Field(i).Name, reflect.ValueOf(builtin9.MethodsMultisig).Field(i).Interface())
	}

	// Payment Channel Methods
	fmt.Println("\nPayment Channel Methods")
	for i := 0; i < reflect.TypeOf(builtin9.MethodsPaych).NumField(); i++ {
		fmt.Printf("| %s | %v |\n", reflect.TypeOf(builtin9.MethodsPaych).Field(i).Name, reflect.ValueOf(builtin9.MethodsPaych).Field(i).Interface())
	}

	// Power Methods
	fmt.Println("\nPower Methods")
	for i := 0; i < reflect.TypeOf(builtin9.MethodsPower).NumField(); i++ {
		fmt.Printf("| %s | %v |\n", reflect.TypeOf(builtin9.MethodsPower).Field(i).Name, reflect.ValueOf(builtin9.MethodsPower).Field(i).Interface())
	}

	// Reward Methods
	fmt.Println("\nReward Methods")
	for i := 0; i < reflect.TypeOf(builtin9.MethodsReward).NumField(); i++ {
		fmt.Printf("| %s | %v |\n", reflect.TypeOf(builtin9.MethodsReward).Field(i).Name, reflect.ValueOf(builtin9.MethodsReward).Field(i).Interface())
	}

	// Verifreg Methods
	fmt.Println("\nVerifreg Methods")
	for i := 0; i < reflect.TypeOf(builtin9.MethodsVerifiedRegistry ).NumField(); i++ {
		fmt.Printf("| %s | %v |\n", reflect.TypeOf(builtin9.MethodsVerifiedRegistry ).Field(i).Name, reflect.ValueOf(builtin9.MethodsVerifiedRegistry ).Field(i).Interface())
	}
}