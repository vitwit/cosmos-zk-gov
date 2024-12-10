package zkgov

import (
	"context"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"github.com/vitwit/cosmos-zk-gov/x/zkgov/client/cli"
	"github.com/vitwit/cosmos-zk-gov/x/zkgov/keeper"

	"cosmossdk.io/core/appmodule"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	zktypes "github.com/vitwit/cosmos-zk-gov/x/zkgov/types"
)

var (
	_ module.AppModuleBasic      = AppModule{}
	_ module.AppModuleSimulation = AppModule{}
	_ module.HasServices         = AppModule{}

	_ appmodule.AppModule = AppModule{}
)

// ConsensusVersion defines the current x/params module consensus version.
const ConsensusVersion = 1

// AppModuleBasic defines the basic application module used by the params module.
type AppModuleBasic struct {
	cdc codec.Codec
}

// Name returns the params module's name.
func (AppModuleBasic) Name() string {
	return zktypes.ModuleName
}

// RegisterLegacyAminoCodec registers the params module's types on the given LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	zktypes.RegisterLegacyAminoCodec(cdc)
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the params module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *gwruntime.ServeMux) {
	if err := zktypes.RegisterQueryHandlerClient(context.Background(), mux, zktypes.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

func (ab AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

func (am AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	zktypes.RegisterInterfaces(registry)
}

// AppModule implements an application module for the distribution module.
type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(cdc codec.Codec, k keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         k,
	}
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// GenerateGenesisState performs a no-op.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {}

// RegisterServices registers a gRPC query service to respond to the
// module-specific gRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	zktypes.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	zktypes.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterStoreDecoder doesn't register any type.
func (AppModule) RegisterStoreDecoder(sdr simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(_ module.SimulationState) []simtypes.WeightedOperation {
	return nil
}

// ConsensusVersion implements AppModule/ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return ConsensusVersion }
