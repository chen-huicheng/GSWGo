package main

import (
	"crypto"
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"strings"

	"encoding/hex"
	"strconv"
	"time"
)

const (
	defaultHash = crypto.SHA256
)

var akSk map[string]string

func init() {
	akSk = make(map[string]string)
	akSk["7QFxrZ4aDA1mR5J9"] = "96LwPoiO8HUgnefC"
	akSk["vq2SMlz0IEtYsuyj"] = "efCvq2SMlz0IEtYe"
}

// 使用 Sk 生成 body 签名
func generateSign(key, body string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(body))
	return hex.EncodeToString(mac.Sum(nil))
}

// AK SK 生成 token
func GenerateToken(body, ak, sk string, expires int64) string {
	token := fmt.Sprintf("%s/%d", ak, expires)
	signKey := generateSign(sk, token) // 生成 ak/expires 的签名, 擅自修改 expires 鉴权会不通过
	sign := generateSign(signKey, body)
	return fmt.Sprintf("%s/%s", token, sign)
}

// 判断 token 是否是约定的 ak sk 生成的
func AuthToken(body, token string) bool {
	tokens := strings.Split(token, "/")
	if len(tokens) != 3 { // 约定数量
		return false
	}
	ak := tokens[0] // 约定格式
	if _, ok := akSk[ak]; !ok {
		return false
	}
	sk := akSk[ak]
	expires := tokens[1] // 约定格式
	expiresTime, err := strconv.ParseInt(expires, 10, 64)
	if err != nil {
		return false
	}
	if expiresTime > time.Now().Unix() { // 判断是否过期
		return false
	}
	sign := GenerateToken(body, ak, sk, expiresTime)
	if sign != token {
		return false
	}
	return true
}

func GenerateTokenTest() string {
	expires := time.Now().Unix()
	body := "hello ak sk"
	ak := "vq2SMlz0IEtYsuyj"
	sk := akSk[ak]
	return GenerateToken(body, ak, sk, expires)
}
func main() {
	// 请求者/客户端 使用约定的 ak sk 生成请求数据的 token
	token := GenerateTokenTest()

	// 服务端 使用同样方式 生成 sign 通过比较判断是否一致
	// sign, err := hex.DecodeString(strings.Split(token, "/")[2])
	// fmt.Println(string(sign), err)
	if AuthToken("hello ak sk", token) {
		fmt.Println("true")
	}
}
