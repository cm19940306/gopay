package gopay

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAliPayClient_AliPayTradePay(t *testing.T) {
	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl("https://api.iguiyu.com/gy_pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("subject", "条码支付")
	body.Set("scene", "bar_code")
	body.Set("auth_code", "284680185023768373")
	body.Set("out_trade_no", "GYWX201901301040355706100458")
	body.Set("total_amount", "100.00")
	body.Set("timeout_express", "2m")
	//条码支付
	aliRsp, err := client.AliPayTradePay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestAliPayClient_AliPayTradeCreate(t *testing.T) {
	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl("https://api.iguiyu.com/gy_pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("subject", "创建订单")
	body.Set("buyer_id", "2088802095984694")
	body.Set("out_trade_no", "GZ201901301040355706100469")
	body.Set("total_amount", "0.01")
	//创建订单
	aliRsp, err := client.AliPayTradeCreate(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
	fmt.Println("aliRsp.TradeNo:", aliRsp.AliPayTradeCreateResponse.TradeNo)
}

func TestAliPayClient_AliPayTradeClose(t *testing.T) {
	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl("https://api.iguiyu.com/gy_pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("out_trade_no", "GYWX201901301040355706100459")
	//条码支付
	aliRsp, err := client.AliPayTradeClose(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestAliPayClient_AliPayTradeCancel(t *testing.T) {
	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl("https://api.iguiyu.com/gy_pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("out_trade_no", "GYWX201901301040355706100457")
	//撤销支付订单
	aliRsp, err := client.AliPayTradeCancel(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestAliPayClient_AliPayTradeQuery(t *testing.T) {
	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2")
	//请求参数
	body := make(BodyMap)
	body.Set("out_trade_no", "GYWX201901301040355706100457")

	//查询订单
	aliRsp, err := client.AliPayTradeQuery(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestAliPayClient_AliPayTradeWapPay(t *testing.T) {
	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		//SetReturnUrl("https://www.gopay.ink").
		SetNotifyUrl("https://api.iguiyu.com/gy_pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("subject", "手机网站测试支付")
	body.Set("out_trade_no", "GZ201901301040355706100477")
	body.Set("quit_url", "https://www.gopay.ink")
	body.Set("total_amount", "100.00")
	body.Set("product_code", "QUICK_WAP_WAY")
	//手机网站支付请求
	payUrl, err := client.AliPayTradeWapPay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("payUrl:", payUrl)
}

func TestAliPayClient_AliPayTradeAppPay(t *testing.T) {
	//aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1wn1sU/8Q0rYLlZ6sq3enrPZw2ptp6FecHR2bBFLjJ+sKzepROd0bKddgj+Mr1ffr3Ej78mLdWV8IzLfpXUi945DkrQcOUWLY0MHhYVG2jSs/qzFfpzmtut2Cl2TozYpE84zom9ei06u2AXLMBkU6VpznZl+R4qIgnUfByt3Ix5b3h4Cl6gzXMAB1hJrrrCkq+WvWb3Fy0vmk/DUbJEz8i8mQPff2gsHBE1nMPvHVAMw1GMk9ImB4PxucVek4ZbUzVqxZXphaAgUXFK2FSFU+Q+q1SPvHbUsjtIyL+cLA6H/6ybFF9Ffp27Y14AHPw29+243/SpMisbGcj2KD+evBwIDAQAB"
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl("https://api.iguiyu.com/gy_pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("subject", "测试APP支付")
	body.Set("out_trade_no", "GZ201901301040355706100469")
	body.Set("total_amount", "1.00")
	//手机APP支付参数请求
	payParam, err := client.AliPayTradeAppPay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("payParam:", payParam)
}

func TestAliPayClient_AliPayTradePagePay(t *testing.T) {
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl("https://api.iguiyu.com/gy_pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("subject", "网站测试支付")
	body.Set("out_trade_no", "GZ201901301040355706100469")
	body.Set("total_amount", "88.88")
	body.Set("product_code", "FAST_INSTANT_TRADE_PAY")

	//电脑网站支付请求
	payUrl, err := client.AliPayTradePagePay(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("payUrl:", payUrl)
}

func TestAliPayClient_AliPayTradeRefund(t *testing.T) {
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, true)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2")
	//请求参数
	body := make(BodyMap)
	body.Set("out_trade_no", "GZ201907301420334577")
	body.Set("refund_amount", "5")
	body.Set("refund_reason", "退款")
	//发起退款请求
	aliRsp, err := client.AliPayTradeRefund(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestAliPayClient_AliPayTradeFastPayRefundQuery(t *testing.T) {
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, true)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2")
	//请求参数
	body := make(BodyMap)
	body.Set("out_trade_no", "GZ201907301420334577")
	body.Set("out_request_no", "GZ201907301420334577")
	//发起退款查询请求
	aliRsp, err := client.AliPayTradeFastPayRefundQuery(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestAliPayClient_AliPayTradeOrderSettle(t *testing.T) {
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, true)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2")
	//请求参数
	body := make(BodyMap)
	body.Set("out_request_no", "201907301518083384")
	body.Set("trade_no", "2019072522001484690549776067")

	type OpenApiRoyaltyDetailInfoPojo struct {
		RoyaltyType  string `json:"royalty_type"`
		TransOut     string `json:"trans_out"`
		TransOutType string `json:"trans_out_type"`
		TransInType  string `json:"trans_in_type"`
		TransIn      string `json:"trans_in"`
		Amount       string `json:"amount"`
		Desc         string `json:"desc"`
	}
	var listParams []OpenApiRoyaltyDetailInfoPojo
	listParams = append(listParams, OpenApiRoyaltyDetailInfoPojo{"transfer", "2088802095984694", "userId", "userId", "2088102363632794", "0.01", "分账给2088102363632794"})
	body.Set("royalty_parameters", listParams)

	//发起交易结算接口
	aliRsp, err := client.AliPayTradeOrderSettle(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
}

func TestAliPayClient_AliPayTradePrecreate(t *testing.T) {
	privateKey := "MIIEogIBAAKCAQEAy+CRzKw4krA2RzCDTqg5KJg92XkOY0RN3pW4sYInPqnGtHV7YDHu5nMuxY6un+dLfo91OFOEg+RI+WTOPoM4xJtsOaJwQ1lpjycoeLq1OyetGW5Q8wO+iLWJASaMQM/t/aXR/JHaguycJyqlHSlxANvKKs/tOHx9AhW3LqumaCwz71CDF/+70scYuZG/7wxSjmrbRBswxd1Sz9KHdcdjqT8pmieyPqnM24EKBexHDmQ0ySXvLJJy6eu1dJsPIz+ivX6HEfDXmSmJ71AZVqZyCI1MhK813R5E7XCv5NOtskTe3y8uiIhgGpZSdB77DOyPLcmVayzFVLAQ3AOBDmsY6wIDAQABAoIBAHjsNq31zAw9FcR9orQJlPVd7vlJEt6Pybvmg8hNESfanO+16rpwg2kOEkS8zxgqoJ1tSzJgXu23fgzl3Go5fHcoVDWPAhUAOFre9+M7onh2nPXDd6Hbq6v8OEmFapSaf2b9biHnBHq5Chk08v/r74l501w3PVVOiPqulJrK1oVb+0/YmCvVFpGatBcNaefKUEcA+vekWPL7Yl46k6XeUvRfTwomCD6jpYLUhsAKqZiQJhMGoaLglZvkokQMF/4G78K7FbbVLMM1+JDh8zJ/DDVdY2vHREUcCGhl4mCVQtkzIbpxG++vFg7/g/fDI+PquG22hFILTDdtt2g2fV/4wmkCgYEA6goRQYSiM03y8Tt/M4u1Mm7OWYCksqAsU7rzQllHekIN3WjD41Xrjv6uklsX3sTG1syo7Jr9PGE1xQgjDEIyO8h/3lDQyLyycYnyUPGNNMX8ZjmGwcM51DQ/QfIrY/CXjnnW+MVpmNclAva3L33KXCWjw20VsROV1EA8LCL94BUCgYEA3wH4ANpzo7NqXf+2WlPPMuyRrF0QPIRGlFBNtaKFy0mvoclkREPmK7+N4NIGtMf5JNODS5HkFRgmU4YNdupA2I8lIYpD+TsIobZxGUKUkYzRZYZ1m1ttL69YYvCVz9Xosw/VoQ+RrW0scS5yUKqFMIUOV2R/Imi//c5TdKx6VP8CgYAnJ1ADugC4vI2sNdvt7618pnT3HEJxb8J6r4gKzYzbszlGlURQQAuMfKcP7RVtO1ZYkRyhmLxM4aZxNA9I+boVrlFWDAchzg+8VuunBwIslgLHx0/4EoUWLzd1/OGtco6oU1HXhI9J9pRGjqfO1iiIifN/ujwqx7AFNknayG/YkQKBgD6yNgA/ak12rovYzXKdp14Axn+39k2dPp6J6R8MnyLlB3yruwW6NSbNhtzTD1GZ+wCQepQvYvlPPc8zm+t3tl1r+Rtx3ORf5XBZc3iPkGdPOLubTssrrAnA+U9vph61W+OjqwLJ9sHUNK9pSHhHSIS4k6ycM2YAHyIC9NGTgB0PAoGAJjwd1DgMaQldtWnuXjvohPOo8cQudxXYcs6zVRbx6vtjKe2v7e+eK1SSVrR5qFV9AqxDfGwq8THenRa0LC3vNNplqostuehLhkWCKE7Y75vXMR7N6KU1kdoVWgN4BhXSwuRxmHMQfSY7q3HG3rDGz7mzXo1FVMr/uE4iDGm0IXY="
	//初始化支付宝客户端
	//    appId：应用ID
	//    privateKey：应用秘钥
	//    isProd：是否是正式环境
	client := NewAliPayClient("2016091200494382", privateKey, false)
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType("RSA2").
		SetNotifyUrl("https://api.iguiyu.com/gy_pay/alipay/notify")
	//请求参数
	body := make(BodyMap)
	body.Set("subject", "预创建创建订单")
	body.Set("out_trade_no", "GZ201901301040355703")
	body.Set("total_amount", "100")
	//创建订单
	aliRsp, err := client.AliPayTradePrecreate(body)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("aliRsp:", *aliRsp)
	fmt.Println("aliRsp.QrCode:", aliRsp.AlipayTradePrecreateResponse.QrCode)
	fmt.Println("aliRsp.OutTradeNo:", aliRsp.AlipayTradePrecreateResponse.OutTradeNo)
}

//==================================================//
func TestFormatPrivateKey(t *testing.T) {
	privateKey := "MIIEowIBAAKCAQEAxorj0SisU93LxawVWybZzDYfq8bxzw9VIOKDSW1M5ombl/MvZTT9IuBSGvqUxXFSWqY05e+CBlquFrOHTzKttmzeZJEKHpGAkW+4TsX6N4ZCHfz3Sz5MBAzBqQvscUD0FTZ+VlQbQ1HjQiKLbHaNY+fFVn3q2XL6dNtWwFhad592M1lCbzEjZ7yqFcxJ0jIDBh1lWwemHu52iW3YkvLD5lE8IXRiJedhgPFXCFAkX5HLcc0/jicRTarQ9gRNpHC6B87T2SDjKmgxqcws4CVdYJbztU1KLrsbMALFMOp24x8xsLgR88XeraQyUWe6V3lt1OYEOd9XLHWRx33bRHSwIDAQABAoIBAHjrGfjG1r11Nae8OH19WeRfikZqMdcztVsD2YWcxdsaL+MJPvJapVjaWecIehcN/2QqGcl4Zy5Lh/9Xc68uZFHYWFHTa+BWKYFqE0wWk1/Bqv7slAgFdvJ4enHkSypmrsFEoQkezEPh2ZDrzRJP2ajg/XTB14h72EHXXCxlIyP6q9ldlHyYSc+KOdC3WYOg1FoFXsliZHVKGZUxo4jck+xwdTGRAIKYpdLjpw7DqAWS6N25cx4XK9GuBYoV7AkIM0kpdjDDXAciG2aws6ef4kumWuW/JSbdrFWGLdiAN8GVpBx6+9WTeDKerS9KyDLNf6rsz9Hm3LNWOYCMLlrEFiECgYEA/XB6pZGrbmSebn2lZO+WFs3LYCXoCUy4ePouLYZO7lKsHNeTYRxCc6cbEmap0hpuMCYVPJkqK+nL/CDwBad1QChN5rPVFl2rLtLu0owvoAuTVjWjYNPgDfWb3spXh0p+lZ9oi53kZd4iDGe/jQJzAcpUa3Yj2me6XFFTD+8FNCkCgYEAyIxrkCo2oqGg1aJ7xc9+aBcpsrVg/upptdTgiMGSVZ3XuixufZV36lzJOdmCBoFGKwmLgKkStJSOm3SHUqdEKQBbHI95aG3HgnAMRXOtkn1exExfpAmCnAAnAx8RONorOTjrMrW0BO0/NII7O7NkLg/ocahr/lXEylsC8dLlMCgYBDiwiEu6/Oee5nUAEWR2veo/YBp9iRMeswAqzv4Q2EInBQN3vFs7xaCj0CyG2V2wlmt5+NSNyeW27LwRN2zkxHTvaD94VgspH+pqSTZF0E8FDR9vWVxqG91qk11QNCwS2/Pn6kRu4p3+t/Ft9L+00fOwcIpLGlcWOPWvUiF/dxEQKBgQCNEEhwpWC80FejLaFGKIdPjEtmSrKpXBV0rfTF+LkizuUBJ3/9zQNRyeGxnnuRj+nlvO1e3sWACySHRu4G53MvR8qqVr13ecfuuA0nOvPojuq4THKrlzVsUqGelXBrlEdiFFJMY7axfvBzoYIyqq+aoTxFjJ6Z/czFOZyp6tnpxQKBgClvDZ9pUc+WH28fDnDPd36bC6HmBq/fkxo92RJey1aRFSoCtKNW5Eaqem8iDD+WAVYak2Vg7xUHkhwIEyVVfHIxZBXc0X1w3jNFjG1/Fyul4hLjqCH2QI8gOjHXAcDZe+MJa8b33ZTiiilUu5A0N8+Xz8qpMQ84cXODHJcPMPYb"

	pKey := FormatPrivateKey(privateKey)
	fmt.Println(pKey)
}

func TestFormatAliPayPublicKey(t *testing.T) {
	aliPayPublicKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAp8gueNlkbiDidz6FBQEBpqoRgH8h7JtsPtYW0nzAqy1MME4mFnDSMfSKlreUomS3a55gmBopL1eF4/Km/dEnaL5tCY9+24SKn1D4iyls+lvz/ZjvUjVwxoUYBh8kkcxMZSDeDz8//o+9qZTrICVP2a4sBB8T0XmU4gxfw8FsmtoomBH1nLk3AO7wgRN2a3+SRSAmxrhIGDmF1lljSlhY32eJpJ2TZQKaWNW+7yDBU/0Wt3kQVY84vr14yYagnSCiIfqyVFqePayRtmVJDr5qvSXr51tdqs2zKZCu+26X7JAF4BSsaq4gmY5DmDTm4TohCnBduI1+bPGD+igVmtl05wIDAQAB"

	pKey := FormatAliPayPublicKey(aliPayPublicKey)
	fmt.Println(pKey)
}

func TestAlipaySystemOauthToken(t *testing.T) {
	privateKey := "MIIEpAIBAAKCAQEA0N6CzO4T6NEgWrlW+mrZYMZSNAQ/vBeaCzo1qT+VYNBbklYWxwyUT4JptswSkupdSZGzKVX8rgTBnIbJXjaL9OuUS4KpkOVtqnW9uaPlei7kh2vHkyzbJS85HG8r4XwrrNGY/o4kp/ur51X5dSNPR+Bd5L/0/SUxztsMS8w2onkKk+VNh95yNQx3ZX0XXDhJGmURHFhuWf3bd0XYGoAsLfuEVj95dKOVUodP6RduGJYd/23CBWQfsapA2SxB/X8/ol/7pCgI9UxYkSm88dFZF5BygBv4Alc77ZP2G05VLRtyaEgvzooXzb7N9QEPWsMD4e7sjuuyhUWXakP78BN7tQIDAQABAoIBAGSqeVwPeqUA9ZkELrlueeo8ay4KxjT/Iw06NtalgHcytI+o0j2JAIn5jr29SU5piiCFClZnzOqJ/E5WYAL6QyX9zHFAPj1Jdjtrl887AMSMiTQDbCKwwcYAJoRfyERatioKsrQtCdm0YDuUojAknndD2b0ADzFaldDeneSCPeq7kHVpAylhLIbzGmE7rGWy1gAnJ5ljza/QMQQkVX7MP852IFEwl+8dYeMK4snfKH4SBWbGgLSkfgemEL3mbHaocl2encFHQYjjFSWabUedOfmf0s81+KytsylKLxzZePAKCpbw7cBzf09wXO7YfFPWP3ORSZtEb6UsJAZBD7Kr1ZkCgYEA9DFdgInepyyFgg4w0ljg7HzRT+2U8xGV/rzO6Yc3HerXKWlhUQlAxq+nLW//MZ4RKRilkQmp9876S2jchATXroU+D620q7RmwN80EPhJC3HuAD/vv13rXJpRLhOElL4eMfbllqN/p7Xu0i4P8qrN/WmyICBN4HQxwiAPacOejNMCgYEA2vfoQpcpfDh+pfFvh1kpZLqEsqNdap9Su/dBmtv0oXDST78Azh4WXskoF0g9MQIkUZDjM3Aaa3qWeVI7p8FW8pk1vC1ktD/3x5K19d/eTJYjw4zaxLSDu/8KXkk8+NjpbBJNiUmidFk8y9n15o98AmLmRx5xYeHosH2Y+W6D4FcCgYEA0F7rbYq3/qkEB7V6PZPv4nyj03NlHYENIEEWwrlJ8/J3xXORwieGKcSrgB9IgJtkA1Bvv3Kiob3xYLXLLbBAJ0C1c4WDXN1krQIJmM107xo5v+bO+tn2w58+1HV9NbuBYptO0BsWDBM4BgueOjE78xjlEXrfWT9tUjwvSs3NFlcCgYBakX2ExYdINQnVU1aBEhT379nhFYcXCc0rp+p/xgiaGdMzXtpTZshz7hnh/2Qn+C3Teu0uomibGfpRNf0stuBiuZIrQk9L6sVuy7TtfOoynUsvn/wArnVqdePw+bP5baamp69gYI+MNjjaTE3UTnbJeeSrd+EHzmvzTA/Q47AyQwKBgQC50qZ0jV1WsNfyvv3YyGmlLiyLFEPZ6aDY+gntgeXJ3fCORy90LOk+4FfkaWdIEGH21+dIx7g7oEa9rTt2foqZnCmWWC3vRlDYsdpRYXz4UOKbHsBocWvSCSOmQ8XgwoDve1MhAz7yOlrHIzX816Ld7QHQdI/d/Z4j1LrMyIWP7A=="
	rsp, err := AlipaySystemOauthToken("2019071565794663", privateKey, "", "06e8961891d647c0ac99bb1cebe7SE69")
	if err != nil {
		fmt.Println("gopay.AlipaySystemOauthToken:", err)
		return
	}
	fmt.Println("rsp:", *rsp)
	//fmt.Println("rsp:", rsp.ErrorResponse.Code)
	//fmt.Println("rsp:", rsp.ErrorResponse.Msg)
	//fmt.Println("rsp:", rsp.ErrorResponse.SubCode)
	//fmt.Println("rsp:", rsp.ErrorResponse.SubMsg)
}

func TestDecryptAliPayOpenDataToStruct(t *testing.T) {
	data := "MkvuiIZsGOC8S038cu/JIpoRKnF+ZFjoIRGf5d/K4+ctYjCtb/eEkwgrdB5TeH/93bxff1Ylb+SE+UGStlpvcg=="
	key := "TDftre9FpItr46e9BVNJcw=="
	rsp := new(PhoneNumberResponse)
	err := DecryptAliPayOpenDataToStruct(data, key, rsp)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("rsp.Code:", rsp.Code)
	fmt.Println("rsp.Msg:", rsp.Msg)
	fmt.Println("rsp.SubCode:", rsp.SubCode)
	fmt.Println("rsp.SubMsg:", rsp.SubMsg)
	fmt.Println("rsp.Mobile:", rsp.Mobile)
}

//==================================================//
type List struct {
	BillList []fundBillListInfo `json:"bill_list"`
}

func TestJsonToString(t *testing.T) {

	list := new(List)
	infos := make([]fundBillListInfo, 0)

	infos = append(infos, fundBillListInfo{Amount: "1.0.0", FundChannel: "iguiyu"})
	infos = append(infos, fundBillListInfo{Amount: "2.0.2", FundChannel: "Jerry"})

	list.BillList = infos

	bs, err := json.Marshal(list)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("string:", string(bs))
}

type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestAliPayParams(t *testing.T) {
	bodyMap := make(BodyMap)

	//people := new(People)
	//people.Name = "Jerry"
	//people.Age = 18
	people := make(map[string]interface{})
	people["name"] = "jerry"
	people["age"] = 18
	bodyMap.Set("people", people)

	fmt.Println("result:", bodyMap.Get("people"))
}
