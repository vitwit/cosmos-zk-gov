package zkgov

// import (
// 	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
// 	zkgovv1beta1 "github.com/vitwit/cosmos-zk-gov/api/sdk/zkgov/v1beta1"
// )

// func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
// 	return &autocliv1.ModuleOptions{
// 		Tx: &autocliv1.ServiceCommandDescriptor{
// 			Service:              zkgovv1beta1.Msg_ServiceDesc.ServiceName,
// 			EnhanceCustomCommand: false, // use custom commands only until v0.51
// 			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
// 				{
// 					RpcMethod: "RegisterUser",
// 					Use:       "register-user [commitment] [proposal_id] --from sender ",
// 					Short:     "register user",
// 					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
// 						{ProtoField: "commitment"},
// 						{ProtoField: "proposal_id"},
// 					},
// 				},
// 			},
// 		},
// 		Query: &autocliv1.ServiceCommandDescriptor{
// 			Service: zkgovv1beta1.Query_ServiceDesc.ServiceName,
// 			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
// 				{
// 					RpcMethod: "CommitmentMerkleProof",
// 					Use:       "get-commitment-merkleproof [commitment] [proposal_id]",
// 					Short:     "Query the user",
// 					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
// 						{ProtoField: "commitment"},
// 						{ProtoField: "proposal_id"},
// 					},
// 				},
// 			},
// 		},
// 	}
// }
