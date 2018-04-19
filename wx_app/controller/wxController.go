package controller


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"crypto/sha1"
	"strings"
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
	for _,v := range arr{
		tmpStr = tmpStr+v;
	}
	fmt.Println(tmpStr)
	h := sha1.New()
	io.WriteString(h, tmpStr)


	result := fmt.Sprintf("%x",h.Sum(nil))
	fmt.Println(result)

	

	if strings.Compare(signature,result)==0{
		c.String(http.StatusOK,"true")
	}else{
		c.String(http.StatusOK,"false")
	}
	
	
	

}