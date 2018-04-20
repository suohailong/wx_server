package controller


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"crypto/sha1"
	"strings"
	"fmt"
	"io"
	"regexp"
)

func verifyWxSignature(signature,url,nonce string) bool{
	flysnowRegexp:=regexp.MustCompile(`xtamp=(\d+)`)
	fmt.Printf("我是url%q\n",url)
	params := flysnowRegexp.FindStringSubmatch(url)
	arr  := []string{"suohailong",params[1],nonce}
	sort.Strings(arr)
	tmpStr:=""
	for _,v := range arr{
		tmpStr = tmpStr+v;
	}
	fmt.Println(tmpStr)
	h := sha1.New()
	io.WriteString(h, tmpStr)

	result := fmt.Sprintf("%x",h.Sum(nil))
	if strings.Compare(signature,result)==0{
		return true
	}else{
		return false
	}
}

func VerifyWx(c *gin.Context){

	if r:=recover();r!=nil{
		fmt.Printf("runTime error caught:%v",r)
	}
	echostr := c.Query("echostr")

	c.String(http.StatusOK,echostr)
}