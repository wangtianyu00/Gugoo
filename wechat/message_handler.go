package wechat

import (
	"fmt"
	"net/http"

	"github.com/chanxuehong/wechat/corp"
	"github.com/chanxuehong/wechat/corp/message/request"
	"github.com/chanxuehong/wechat/corp/message/response"
)

// 文本消息的 Handler
func TextMessageHandler(w http.ResponseWriter, r *corp.Request) {
	text := request.GetText(r.MixedMsg) // 可以省略, 直接从 r.MixedMsg 取值
	resp := response.NewText(text.FromUserName, text.ToUserName, text.CreateTime, text.Content)
	corp.WriteResponse(w, r, resp)
}

// 位置消息的 Handler
func LocationMessageHandler(w http.ResponseWriter, r *corp.Request) {
	location := request.GetLocation(r.MixedMsg) // 可以省略, 直接从 r.MixedMsg 取值
	distance := EarthDistance(LocationX, LocationY, location.LocationX, location.LocationY)
	state := "未打卡"
	if distance < 100 {
		state = "打卡成功"
	} else {
		state = "打卡失败"
	}
	resp := response.NewText(location.FromUserName, location.ToUserName, location.CreateTime, "x:"+fmt.Sprint(location.LocationX)+"	y:"+fmt.Sprint(location.LocationY)+"	scale:"+fmt.Sprint(location.Scale)+location.Label+"	距离："+fmt.Sprint(distance)+"	打卡状态："+fmt.Sprint(state))
	fmt.Println(location)
	fmt.Println(resp)
	corp.WriteResponse(w, r, resp)
}