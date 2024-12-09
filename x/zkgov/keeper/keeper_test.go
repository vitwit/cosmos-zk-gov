package keeper_test

import (
	"fmt"
	"testing"

	addresstypes "cosmossdk.io/core/address"
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	cmttime "github.com/cometbft/cometbft/types/time"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/stretchr/testify/suite"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/keeper"
	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
)

type TestSuite struct {
	suite.Suite

	ctx          sdk.Context
	addrs        []sdk.AccAddress
	baseApp      *baseapp.BaseApp
	encCfg       moduletestutil.TestEncodingConfig
	queryClient  types.QueryClient
	msgSrvr      types.MsgServer
	keeper       keeper.Keeper
	addressCodec addresstypes.Codec

	addresses []string
}

func (s *TestSuite) SetupTest() {
	key := storetypes.NewKVStoreKey(types.ModuleName)
	storeService := runtime.NewKVStoreService(key)
	testCtx := testutil.DefaultContextWithDB(s.T(), key, storetypes.NewTransientStoreKey("transient_test"))
	s.ctx = testCtx.Ctx.WithBlockHeader(cmtproto.Header{Time: cmttime.Now()})
	s.encCfg = moduletestutil.MakeTestEncodingConfig(zkgov.AppModuleBasic{})

	s.addressCodec = address.NewBech32Codec("cosmos")
	// ctrl := gomock.NewController(s.T())
	s.baseApp = baseapp.NewBaseApp(
		"zk-gov",
		log.NewNopLogger(),
		testCtx.DB,
		s.encCfg.TxConfig.TxDecoder(),
	)
	s.baseApp.SetCMS(testCtx.CMS)
	s.baseApp.SetInterfaceRegistry(s.encCfg.InterfaceRegistry)

	s.addrs = simtestutil.CreateIncrementalAccounts(7)

	s.keeper = keeper.NewKeeper(s.encCfg.Codec, storeService)

	queryHelper := baseapp.NewQueryServerTestHelper(s.ctx, s.encCfg.InterfaceRegistry)
	types.RegisterQueryServer(queryHelper, s.keeper)
	queryClient := types.NewQueryClient(queryHelper)
	s.queryClient = queryClient

	s.msgSrvr = keeper.NewMsgServerImpl(s.keeper)

	alice, err := s.addressCodec.BytesToString([]byte("alice"))
	s.Require().NoError(err)

	bob, err := s.addressCodec.BytesToString([]byte("bob"))
	s.Require().NoError(err)

	_, err = s.addressCodec.StringToBytes(bob)
	s.Require().NoError(err)

	charlie, err := s.addressCodec.BytesToString([]byte("charlie"))
	s.Require().NoError(err)

	s.addresses = []string{alice, bob, charlie}
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestGetUserInfo() {
	alice := "cosmos1ee0gzn3f3f0h2l4djk9rfvwmr4h0z8xlqvdc3k"

	bob := "cosmos1rns9w3kgpsc5rtcleagaul79hvjxeu4h26jyfd"

	_, err := s.msgSrvr.CreateProposal(s.ctx, &types.MsgCreateProposal{
		Title:                "dummy proposal",
		Description:          "This is the dummy proposal",
		RegistrationDeadline: nil,
		VotingDeadline:       nil,
		Sender:               alice,
	})
	s.Require().NoError(err)

	_, err = s.msgSrvr.RegisterUser(s.ctx, &types.MsgRegisterUser{
		Sender:     alice,
		Commitment: "0a8b34dc58d41b24c4a3e961cd78b45221b9eac800bb2e173133e3496381f898",
		ProposalId: 1,
	})
	s.Require().NoError(err)

	_, err = s.msgSrvr.RegisterUser(s.ctx, &types.MsgRegisterUser{
		Sender:     bob,
		Commitment: "0a8b34dc58d41b24c4a3e961cd78b45221b9eac800bb2e173133e3496381f898",
		ProposalId: 1,
	})
	s.Require().NoError(err)

	res, err := s.queryClient.ProposalAllInfo(s.ctx, &types.QueryProposalAllInfoRequest{ProposalId: 1})
	fmt.Println("the result is>>>>>>>>>>", res.Votes)

	s.Require().NoError(err)
}
