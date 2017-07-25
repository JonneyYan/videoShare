package api

import (
	"fmt"
	"json"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/JonneyYan/videoShare/server/wxpay"
)

type PayResponse struct {
	Code      string
	Msg       string
	Sign      string
	Timestamp string
	Noncestr  string
	Appid     string
	PrepayID  string `json:"prepay_id"`
}

func HandlePay(w http.ResponseWriter, r *http.Request) {
	parent := r.Form.Get("parent")
	openID := r.Form.Get("openid")

	payAPI := wxpay.NewAPI(appID, "", appkey, "")
	params := map[string]string{
		appid: appID,
	}
	pm := payAPI.NewMap()
	pm.BasicCheckSet()
	pm.SetDeviceInfo("WEB")
	pm.SetBody("百仕汇和-代理注册费")
	pm.SetAttach(parent)

	date := time.Now().Format("20060102150405")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randNumber := r.Intn(100) + 100
	tradeNo := fmt.Sprintf("%s%d", date, randNumber)
	pm.SetOutTradeNo(tradeNo)

	pm.SetTotalFee("48.00")

	clientIPAddress := strings.Split(r.RemoteAddr, ":")[0]
	pm.SetSpbillCreateIP(clientIPAddress)

	pm.SetNotifyUrl("http://shop.baishihuihe.com/wxpay/paycallback")
	pm.SetTradeType("JSAPI")
	pm.SetOpenID(openID)
	result, err := payAPI.UnifiedOrder(pm)

	var resp = PayResponse

	if result["return_code"] == "FAIL" {
		resp.Code = result["return_code"]
		resp.Msg = result["return_msg"]
	} else if result["result_code"] == "FAIL" {
		resp.Code = result["err_code"]
		resp.Msg = result["err_code_des"]
	} else {
		resp.Code = "0"
		resp.Msg = "成功"
		resp.Sign = result["sign"]
		resp.Timestamp = result["timestamp"]
		resp.Noncestr = result["noncestr"]
		resp.Appid = result["appid"]
		resp.PrepayID = result["prepayid"]
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Fprint(w, string(respJSON))
}
