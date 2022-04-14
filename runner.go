package conductor

import (
	"context"
	"os"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/strangelove-ventures/lens/byop"
	lens "github.com/strangelove-ventures/lens/client"
	"github.com/volumefi/conductor/client/cronchain"
	cronchaintypes "github.com/volumefi/conductor/types/cronchain"
	terratypes "github.com/volumefi/conductor/types/terra"

	chain "github.com/volumefi/conductor/client"
)

func Start() {
	panic("main loop goes here")
	// var c terra.Client
	// modules := append(lens.ModuleBasics[:], byop.NewModule(
	// 	"testing",
	// 	(*types.MsgExecuteContract)(nil),
	// ))

	// modules := append(lens.ModuleBasics[:])
	// lensc, err := chain.NewChainClient(
	// 	&lens.ChainClientConfig{
	// 		Key:            "matija",
	// 		ChainID:        "columbus-5",
	// 		RPCAddr:        "http://127.0.0.1:22222",
	// 		KeyringBackend: "file",
	// 		KeyDirectory:   "/home/vizualni/.terra/",
	// 		AccountPrefix:  "terra",
	// 		Modules:        modules,
	// 		Debug:          true,
	// 		GasAdjustment:  1.1,
	// 		GasPrices:      "0.2056735uusd,0luna",
	// 		SignModeStr:    "direct",
	// 	},
	// 	"doesn-tmatter",
	// 	os.Stdin,
	// 	os.Stdout,
	// )
	// if err != nil {
	// 	panic(err)
	// }
	// c.LensClient = lensc
	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	// c.ExecuteSmartContract(ctx)
}

func Start2() {
	// TODO: this is a temporary playing ground to test things in real life
	var c cronchain.Client
	// registering types into the interface registry for codec
	byopmodule := byop.Module{
		ModuleName: "runner",
		MsgsImplementations: []byop.RegisterImplementation{
			{
				Iface: (*sdk.Msg)(nil),
				Msgs: []proto.Message{
					(*terratypes.MsgExecuteContract)(nil),
					(*cronchaintypes.MsgAddMessagesSignatures)(nil),
				},
			},
			{
				Iface: (*cronchaintypes.Signable)(nil),
				Msgs: []proto.Message{
					&cronchaintypes.SignSmartContractExecute{},
				},
			},
		},
	}
	modules := append(lens.ModuleBasics[:], byopmodule)

	lensc, err := chain.NewChainClient(
		&lens.ChainClientConfig{
			Key:            "matija",
			ChainID:        "cronchain",
			RPCAddr:        "http://127.0.0.1:26657",
			KeyringBackend: "test",
			KeyDirectory:   "/home/vizualni/.cronchain/",
			AccountPrefix:  "cosmos",
			Modules:        modules,
			Debug:          true,
			GasAdjustment:  1.1,
			SignModeStr:    "direct",
		},
		"doesn-tmatter",
		os.Stdin,
		os.Stdout,
	)
	if err != nil {
		panic(err)
	}
	c.L = lensc
	r := relayer{
		cronchain: c,
	}

	err = r.signMessagesForExecution(context.Background(), "a")
	if err != nil {
		panic(err)
	}

}

func loop() {
	// get messages for execution
	// get messages for signing
	// get messages for attestation
}