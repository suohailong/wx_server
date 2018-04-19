package controller


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"crypto/sha1"
	"bytes"
	"fmt"
	"io"
)

func VerifyWx(c *gin.Context){
	signature := c.Query("signature")
	tamp  := c.Query("tamp")
	nonce:= c.Query("nonce")

	arr  := []string{"suohailong",tamp,nonce}
	sort.Strings(arr)
	
	tmpStr:=""
	for i,v := range arr{
		i=i;
		tmpStr = tmpStr+v;
	}
	fmt.Println(tmpStr)
	h := sha1.New()
	io.WriteString(h, tmpStr)
	Result:=h.Sum(nil)

	fmt.Println("% x", Result)
	fmt.Println("% x", []byte(signature))

	

	if bytes.Compare([]byte(signature),[]byte(signature))==0{
		c.String(http.StatusOK,"true")
	}else{
		c.String(http.StatusOK,"false")
	}
	
	
	

}