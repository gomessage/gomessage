package send

import (
	"bytes"
	"fmt"
	"html/template"
)

// CompleteMessage 模板渲染（这仅仅只是把模板中内嵌的变量渲染成消息实体）
func CompleteMessage(thisTemplate string, dataList []map[string]string) []string {
	var templateList []string

	for _, data := range dataList {

		//创建一个字节缓冲器
		var buf bytes.Buffer

		//以 K8sMessageTemplate 为蓝本，新建一个模板，新起一个名字叫tmplNewName
		tmpl, err := template.New("tmplNewName").Parse(thisTemplate)
		if err != nil {
			fmt.Println(err)
		}
		//实例化属性
		//tmplData := data
		//fmt.Println(data)

		//渲染tmpl模板，渲染的过程中把tmplData的值填充到模板之内，最后把渲染后的结果存入到buf这个字节缓冲器中
		if err = tmpl.Execute(&buf, data); err != nil {
			panic(err)
		}

		//格式化字节缓冲器中的内容为String串，然后返回给调用方
		//return buf.String()
		templateList = append(templateList, buf.String())
	}
	return templateList
}
