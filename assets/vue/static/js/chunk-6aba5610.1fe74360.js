(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-6aba5610","chunk-c7b676d4"],{"089b":function(t,e,a){},"4b59":function(t,e,a){},"7d3c":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-card",{attrs:{shadow:"always"}},[a("div",{attrs:{slot:"header"},slot:"header"},[a("span",{staticStyle:{"padding-left":"100px"}},[t._v("劫持【 "),a("span",{staticStyle:{"font-size":"30px",color:"#3de1ad","font-weight":"900"}},[t._v(t._s(t.getStoreNamespace))]),t._v(" 】通道内 · 实时过境数据")]),a("el-button",{staticStyle:{float:"right",padding:"3px 0","margin-left":"10px"},attrs:{type:"text"},on:{click:function(e){return t.getServerData()}}},[t._v(" 劫持数据 ")]),a("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(e){return t.copyCode()}}},[t._v("一键复制")])],1),a("div",[a("pre",{attrs:{id:"CodeStyle"}},[a("code",{attrs:{id:"CodeContent"}},[t._v(t._s(t.codeJsonContent))])])])])},s=[],l=(a("e9c4"),a("6cac")),o={name:"cCodeFormat",data:function(){return{codeJsonContent:"点击 <劫持数据> 可以看到推送进GoMessage的原始数据...\n\n每次 <劫持数据> 都可以实时抓取到最新的一条数据...\n\n此处对数据进行了格式化与对齐，您可以把数据 <一键复制> 到其它地方使用...",codeUpdateTime:null,arr2:[]}},computed:{getStoreNamespace:function(){return this.$store.getters.getNamespace}},methods:{copyCode:function(){var t=document.getElementById("CodeContent");window.getSelection().selectAllChildren(t),document.execCommand("Copy"),this.$message("复制成功...")},getServerData:function(){var t=this;Object(l["n"])(this.$store.getters.getNamespace,null).then((function(e){var a=e.data.result["request_json"],n=e.data.result["request_time"];null===a||0===a.length?t.$message({showClose:!1,message:"没有数据进入GoMessage服务，无法向您展示数据。"}):(t.codeJsonContent=JSON.stringify(a,null,2),t.codeUpdateTime=t.dateFormat(new Date(n)))})).catch((function(t){console.log(t)}))}}},i=o,c=(a("db83"),a("2877")),r=Object(c["a"])(i,n,s,!1,null,"579de54e",null);e["default"]=r.exports},"8ca6":function(t,e,a){"use strict";a("cc9e")},b01b:function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",[a("el-row",{staticStyle:{margin:"20px 0"},attrs:{gutter:20}},[a("el-col",{attrs:{span:"12"}},[a("DataFormat",{staticClass:"shadow"}),a("br"),a("DataMap",{staticClass:"shadow"})],1),a("el-col",{attrs:{span:12}},[a("CTemplate",{staticClass:"shadow"})],1)],1)],1)},s=[],l=a("7d3c"),o=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-card",{attrs:{shadow:"always"}},[a("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[a("span",{staticStyle:{"padding-left":"50px"}},[t._v("自定义变量映射")]),a("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:t.post_data}},[t._v("保存变量")])],1),a("el-table",{staticStyle:{width:"100%"},attrs:{border:!0,data:t.configList}},[a("el-table-column",{attrs:{label:"Key（自定义变量名）",prop:"key"}}),a("el-table-column",{attrs:{label:"Value（原始数据中的索引）",prop:"value"}}),a("el-table-column",{attrs:{fixed:"right",label:"操作",width:"100"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-button",{attrs:{size:"small",type:"text"},nativeOn:{click:function(a){return a.preventDefault(),t.deleteRow(e.$index,t.configList)}}},[t._v("移除 ")])]}}])})],1),a("br"),t._l(t.mapList2,(function(e,n){return a("div",{key:n},[a("el-row",{attrs:{gutter:20}},[a("el-col",{attrs:{span:12}},[a("div",{staticClass:"DataMapGridContent"},[a("el-input",{attrs:{placeholder:"纯字符串，不能有符号"},model:{value:e.mapKey,callback:function(a){t.$set(e,"mapKey",a)},expression:"list.mapKey"}},[a("template",{slot:"prepend"},[t._v(" Key: ")])],2)],1)]),a("el-col",{attrs:{span:12}},[a("div",{staticClass:"DataMapGridContent"},[a("el-input",{attrs:{placeholder:"Json索引路径"},model:{value:e.mapValue,callback:function(a){t.$set(e,"mapValue",a)},expression:"list.mapValue"}},[a("template",{slot:"prepend"},[t._v(" Value: ")])],2)],1)])],1)],1)})),a("br"),a("el-button",{attrs:{size:"mini",type:"primary"},on:{click:t.addTableData}},[t._v("新增变量")])],2)},i=[],c=(a("a434"),a("ac1f"),a("1276"),a("d3b7"),a("159b"),a("6cac")),r={name:"cConfigMap",data:function(){return{mapList2:[{mapKey:"",mapValue:""}],configList:[{key:"{{ .N1 }} ",value:"alerts.#.labels.alertname"},{key:"{{ .N2 }}",value:"alerts.#.labels.severity"},{key:"{{ .N3 }}",value:"alerts.#.labels.hostname"},{key:"{{ .N4 }}",value:"alerts.#.labels.ping"},{key:"{{ .N5 }}",value:"alerts.#.annotations.description"},{key:"{{ .N6 }}",value:"status"},{key:"{{ .N7 }}",value:"alerts.#.startsAt"},{key:"{{ .N8 }}",value:"alerts.#.endsAt"}]}},computed:{},methods:{addTableData:function(){var t=this.mapList2[0].mapKey,e=this.mapList2[0].mapValue;if(0===t.length||0===e.length)this.$message.error("输入框不能为空...");else{var a={key:"{{ ."+t+" }}",value:e};this.configList.push(a),this.mapList2[0].mapKey="",this.mapList2[0].mapValue="",this.$message.success("添加成功..."),this.post_data()}},deleteRow:function(t,e){e.splice(t,1),this.post_data()},post_data:function(){for(var t=this,e=[],a=0;a<this.configList.length;a++){var n=this.configList[a].key,s=n.split(" ")[1].split(".")[1],l=this.configList[a].value,o={};o[s]=l,e.push(o)}Object(c["k"])(this.$store.getters.getNamespace,{key_value_list:e}).then((function(e){console.log(e.data),t.$message.success("数据库更新成功...")})).catch((function(t){console.log(t)}))},pullMapData:function(){var t=this;Object(c["o"])(this.$store.getters.getNamespace,null).then((function(e){var a=[];0===e.data.result.length?console.log("数据库里没有数据"):(e.data.result.forEach((function(t){var e={key:"{{ ."+t["key"]+" }}",value:t["value"]};a.push(e)})),t.configList=a)})).catch((function(t){console.log(t)}))}},created:function(){this.pullMapData()}},p=r,u=(a("fc08"),a("2877")),d=Object(u["a"])(p,o,i,!1,null,"dde7b8ac",null),m=d.exports,f=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-card",{attrs:{shadow:"always"}},[a("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[a("el-switch",{staticStyle:{float:"left",padding:"3px 0"},attrs:{"inactive-text":"聚合发送"},on:{change:t.pushTemplateData},model:{value:t.template.template_is_merge,callback:function(e){t.$set(t.template,"template_is_merge",e)},expression:"template.template_is_merge"}}),a("el-tooltip",{staticClass:"item",staticStyle:{float:"left","margin-left":"30px","padding-top":"3px"},attrs:{content:'跳转到新页面：查看"消息模板"编写教程',effect:"dark",placement:"bottom"}},[a("el-link",{attrs:{underline:!1,type:"primary"},on:{click:t.newTagPage}},[a("span",[a("i",{staticClass:"el-icon-info"})])])],1),a("span",{staticStyle:{"margin-right":"70px"}},[t._v("自定义消息模板")]),a("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(e){return t.pushTemplateData()}}},[t._v("保存模板")])],1),a("div",[a("el-input",{attrs:{autosize:{minRows:10,maxRows:200},placeholder:"请输入Golang语法的模板内容",resize:"none",type:"textarea"},model:{value:t.template.template_content,callback:function(e){t.$set(t.template,"template_content",e)},expression:"template.template_content"}})],1)])},h=[],g={name:"cTemplate",data:function(){return{template:{template_is_merge:!1,template_content:"{{ if eq .N6 \"firing\" }}\n\n## <font color='#FF0000'>【报警中】服务器{{ .N3 }}</font>\n\n{{ else if eq .N6 \"resolved\" }}\n\n## <font color='#20B2AA'>【已恢复】服务器{{ .N3 }}</font>\n\n{{ else }}\n\n## 标题：信息通知\n\n{{ end }}\n\n========================\n\n**告警规则**：{{ .N1 }}\n\n**告警级别**：{{ .N2 }}\n\n**主机名称**：{{ .N3 }} \n\n**主机地址**：{{ .N4 }}\n\n**告警详情**：{{ .N5 }}\n\n**告警状态**：{{ .N6 }}\n\n**触发时间**：{{ .N7 }}\n\n**发送时间**：{{ .N8 }}\n\n**规则详情**：[Prometheus控制台](https://www.baidu.com)\n\n**报警详情**：[Alertmanager控制台](https://www.baidu.com)\n"}}},components:{},methods:{newTagPage:function(){var t="https://github.com/gomessage/gomessage#gomessage";window.open(t)},pushTemplateData:function(){var t=this;Object(c["j"])(this.$store.getters.getNamespace,this.template).then((function(e){console.log(e.data.result),t.$message.success("数据库更新成功...")})).catch((function(t){console.log(t)}))},pullTemplateData:function(){var t=this;Object(c["f"])(this.$store.getters.getNamespace,null).then((function(e){var a=e.data.result[0]["template_content"],n=e.data.result[0]["template_is_merge"];void 0===a||null===a||""===a?console.log("数据库里没有数据"):(t.template.template_content=a,t.template.template_is_merge=n)})).catch((function(t){console.log(t)}))}},created:function(){this.$store.commit("updateStepsActive",2),this.pullTemplateData()}},v=g,b=(a("ce60"),Object(u["a"])(v,f,h,!1,null,"04ca15f5",null)),y=b.exports,_={name:"ViewRequestData",components:{DataFormat:l["default"],DataMap:m,CTemplate:y},created:function(){this.$store.commit("updateStepsActive",1)}},w=_,k=(a("8ca6"),Object(u["a"])(w,n,s,!1,null,"081eaa92",null));e["default"]=k.exports},b671:function(t,e,a){},cc9e:function(t,e,a){},ce60:function(t,e,a){"use strict";a("4b59")},db83:function(t,e,a){"use strict";a("089b")},fc08:function(t,e,a){"use strict";a("b671")}}]);
//# sourceMappingURL=chunk-6aba5610.1fe74360.js.map