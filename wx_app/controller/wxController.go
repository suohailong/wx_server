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

func VerifyWx(c *gin.Context){
	signature := c.Query("signature")
	echostr  := c.Query("echostr")
	nonce:= c.Query("nonce")
	flysnowRegexp:=regexp.MustCompile(`xtamp=(\d+)`)
	params := flysnowRegexp.FindStringSubmatch(echostr)
	
	fmt.Println("你好啊，怎么回事，怎么这里会超出index呢")
	fmt.Println(echostr)
	fmt.Println(params)

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
		c.String(http.StatusOK,"true")
	}else{
		c.String(http.StatusOK,"false")
	}
	
	
	

}