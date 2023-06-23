package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/smartwalle/alipay/v3"
	"github.com/smartwalle/xid"
)

var client *alipay.Client

const (
	kAppId      = "9021000122693643"
	kPrivateKey = "MIIEowIBAAKCAQEAjasPxFS+AqqrzU4CEsGWSgPYHtSt3jCO+Wr0OpsOnmJC8fxbse8BaQ9B2CZmmwTrCU7+hOhUJE7KZrAoCSGrUY+D4VzL0cGKCoghHhLA+Bj22g9d5iVksWL1c6bFHNHUViZRnQKPYSfRzOIbiLvXXbdoqEmvR9F4jubTM9yKqq6ROr9O7UOuZGArwNJj5wVuDHqRvYtVrmqhdCDC2N1nkLlPOWIF3dW/SMMxW/V5brwuDz83UwU4OVrU+Zt9S321evDdozT12fykDs3Hz0I5QkwU7vmOu/krhLsOnw1We2XEv5NHVQFyVcWai6nI0Kg5YsC3bRTCRj6E8rJduTEfLQIDAQABAoIBAGBiPbvpm8zcqT5pEGgKZWG989AmCpTcnKl85uToka2YVMumUgxg+6iqxLV2iuB1HCOdJvTtzDFXmcT7nIF58sDuz18Ib3MqJey8aSu4vqkM8wWw5sWqMzOsjCo/EGRnLfve3i/gigQjv02RmAITKpeMLRwx5ZsLHOa8jH0AUyeDyDrW6AKJ3A8cPKR5GhZn1EcieD0XM6eP1bZ4+ULA9I4nf2hYlBAYQ3c/Prf685hAfuetFPWoQaHKXKKiLmsuPtGTPtDAig1Zjlf4O6nV6cLo76NZ7yCy4PyhlC0iZOGdtbE9ldcHDj9UPKPujWoXtUHk/L8RvHuo1GI+M4hkp9ECgYEA+qI9qVMi0wd9zrpVs+yUZ9LhEtt0BzqyH/bdQJEIJA1A+mI+TpCj31/bKpi9qxceVKWz1+Ks56F27EFOCXK/WYEUSd3oXBpLwEub1XQoT/WuBlU+Kl3agnZVdLaRplGh9XdJRIa2Rkaoh70VmOXU2xXiWxGOobT8vGHlPfZ5Sn8CgYEAkLOQrFKgAPPKevf8URYatA8qe4J8z6RRv9nqVYt0ZdpB8HsL8a1CTDgrNo8TdB02yJPNNlB2IOePCVNf41aAKAdDBofqUFDzuNSiY7QJ7e8EJgBD7r6cUdW5uGoEwnGiDc0miLdMuflfpsw3z5hNkr64YhNfwP+laLBLp9F2CFMCgYEA8oW7lPVVZU6Cdl/oPkpW8IAmOtEP3U72vedrAOdWMUyB68bU88ESUTu4H+fc4IlmWBJfjC9TkG+1W7d1g779WeaFc1S/WyA23MRM7qePuo/I7Wfe17ApmMRYvK9NENs9FQFwbFbOWgJVYcOd+m2rIrRybBz5H5C6Xq973BF8HJECgYAQ9zzLN5TTB8KB0Hmg+DqBShXtDUuc/OMo31/+T6+CEakRjRvFZk8TqLJIX1YkkxhNRCcXay+ug/fdXA0uFoJILzHInj7208b/sM6paf6QCyotWnKcl3S0k56G5MdSFQivHZKbDQLN+1Mft2oDyROvVPVBaowbkN3P/HuiXYi9SQKBgCgJc+DBw3S0DN8jRiaNdAb0p5fprWiEeyTEsjAy07b/pJb5Sy8T1PzhhczXKXO/4bb6/ThaP0EAjSu0qht7lyepuhKjM1snUMeUATO00K4qaGTZyJcj2Vsn971TYlyllz1qphHRV8EvOG4jHrRhWSjxxqiE8y6ftiXjpvnTt4LF"
	kServerPort = "9999"
	// TODO 设置回调地址域名
	kServerDomain = "https://efb5-222-249-139-67.ngrok-free.app"
)

const (
	newGatewayAddr = "https://openapi-sandbox.dl.alipaydev.com/gateway.do"
)

func main() {
	var err error

	if client, err = alipay.New(kAppId, kPrivateKey, false, alipay.WithSandboxGateway(newGatewayAddr)); err != nil {
		log.Println("初始化支付宝失败", err)
		return
	}

	// 加载证书
	if err = client.LoadAppCertPublicKeyFromFile("appPublicCert.cer"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	if err = client.LoadAliPayRootCertFromFile("alipayRootCert.cer"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}
	if err = client.LoadAlipayCertPublicKeyFromFile("alipayPublicCert.cer"); err != nil {
		log.Println("加载证书发生错误", err)
		return
	}

	if err = client.SetEncryptKey("jd+GhHDawCRYkCuNjzv7rQ=="); err != nil {
		log.Println("加载内容加密密钥发生错误", err)
		return
	}

	http.HandleFunc("/alipay/pay", pay)
	http.HandleFunc("/alipay/return_url", returnURL)
	http.HandleFunc("/alipay/notify", notify)

	http.ListenAndServe(":"+kServerPort, nil)
}

func pay(writer http.ResponseWriter, request *http.Request) {
	var tradeNo = fmt.Sprintf("%d", xid.Next())

	var p = alipay.TradePagePay{}
	p.NotifyURL = kServerDomain + "/alipay/notify"
	p.ReturnURL = kServerDomain + "/alipay/return_url"
	p.Subject = "支付测试:" + tradeNo
	p.OutTradeNo = tradeNo
	p.TotalAmount = "10.00"
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		log.Printf("[WARN] trade page pay error: %v", err)
	}
	http.Redirect(writer, request, url.String(), http.StatusTemporaryRedirect)
}

func returnURL(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	if err := client.VerifySign(request.Form); err != nil {
		log.Println("回调验证签名发生错误", err)
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte("回调验证签名发生错误"))
		return
	}

	log.Println("回调验证签名通过")

	var outTradeNo = request.Form.Get("out_trade_no")
	var p = alipay.TradeQuery{}
	p.OutTradeNo = outTradeNo

	rsp, err := client.TradeQuery(p)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("验证订单 %s 信息发生错误: %s", outTradeNo, err.Error())))
		return
	}

	if rsp.IsFailure() {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(fmt.Sprintf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg)))
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(fmt.Sprintf("订单 %s 支付成功", outTradeNo)))
}

func notify(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	var notification, err = client.DecodeNotification(request.Form)
	if err != nil {
		log.Println("解析异步通知发生错误", err)
		return
	}

	log.Println("解析异步通知成功:", notification.NotifyId)

	var p = alipay.NewPayload("alipay.trade.query")
	p.Set("out_trade_no", notification.OutTradeNo)

	var rsp *alipay.TradeQueryRsp
	if err = client.Request(p, &rsp); err != nil {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s \n", notification.OutTradeNo, err.Error())
		return
	}
	if rsp.IsFailure() {
		log.Printf("异步通知验证订单 %s 信息发生错误: %s-%s \n", notification.OutTradeNo, rsp.Msg, rsp.SubMsg)
		return
	}

	log.Printf("订单 %s 支付成功 \n", notification.OutTradeNo)

	client.ACKNotification(writer)
}
