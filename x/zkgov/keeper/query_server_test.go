package keeper_test

// import (
// 	"fmt"

// 	"github.com/vishal-kanna/zk/zk-gov/x/zkgov/types"
// )

// func (s *TestSuite) TestPlanByID() {

// 	alice, err := s.addressCodec.BytesToString([]byte("alice"))
// 	s.Require().NoError(err)

// 	_, err = s.msgSrvr.RegisterUser(s.ctx, &types.RegisterUserRequest{
// 		Sender: alice,
// 	})
// 	s.Require().NoError(err)

// 	_, err = s.msgSrvr.RegisterUser(s.ctx, &types.RegisterUserRequest{
// 		Sender: alice,
// 	})
// 	s.Require().NoError(err)
// 	bz, _ := s.addressCodec.StringToBytes("cosmos15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqjwl8sq")
// 	fmt.Println(">>>>>>>>.", bz, ">>>>>>>>>>>>>")
// 	res, err := s.queryClient.GetUser(s.ctx, &types.QueryUserRequset{Userid: 1})
// 	fmt.Println("the result is>>>>>>>>>>", res)
// 	s.Require().Equal(res.Ust.Userid, uint64(1))
// 	s.Require().NoError(err)

// 	res, err = s.queryClient.GetUser(s.ctx, &types.QueryUserRequset{Userid: 2})
// 	fmt.Println("the result is>>>>>>>>>>", res)
// 	s.Require().Equal(res.Ust.Userid, uint64(2))
// 	s.Require().NoError(err)
// }
