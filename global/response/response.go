package response

import (
	"fmt"
	"fsc/global"
)

func OkWithMessage(message string)  {
	fmt.Printf("「FSC」成功：%s \n", message)
}

func FailWithMessage(message string, err error)  {
	if global.FSC_CONFIG.FSC.Debug {
		if err != nil {
			fmt.Printf("「FSC」失败：%s ，%s。", message, fmt.Sprintf("%s", err))
			return
		}
	}
	fmt.Printf("「FSC」失败：%s \n", message)
}