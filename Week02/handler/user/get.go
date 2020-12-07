package user

import "Week02/service"

// get user handler
func Get()  {
	// 假设请求进来了
	// 获取参数
	userID := 1

	srv := service.NewService()
	u, err := srv.GetUser(userID)
	if err != nil {
		// logger ...
	}

	// 省略返回的封装了。包括code之类的
	//rsp := respone {
	//}
	//return SendResponse(rsp)
}