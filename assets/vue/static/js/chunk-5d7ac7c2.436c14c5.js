(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-5d7ac7c2","chunk-d4fab21a"],{"07c5":function(e,t,a){},"3ca5":function(e,t,a){"use strict";a("c2ac")},"46d9":function(e,t,a){},"7d3c":function(e,t,a){"use strict";a.r(t);var n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("el-card",{attrs:{shadow:"always"}},[a("div",{attrs:{slot:"header"},slot:"header"},[a("span",{staticStyle:{"padding-left":"100px"}},[e._v("劫持【 "),a("span",{staticStyle:{"font-size":"18px",color:"#3de1ad","font-weight":"900"}},[e._v(e._s(e.getStoreNamespace))]),e._v(" 】通道中的最新过境数据")]),a("el-button",{staticStyle:{float:"right",padding:"3px 0","margin-left":"10px"},attrs:{type:"text"},on:{click:function(t){return e.getServerData()}}},[e._v(" 劫持数据 ")]),a("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(t){return e.copyCode()}}},[e._v("一键复制")])],1),a("div",[a("pre",{attrs:{id:"CodeStyle"}},[a("code",{attrs:{id:"CodeContent"}},[e._v(e._s(e.codeJsonContent))])])])])},s=[],l=(a("e9c4"),a("6cac")),o={name:"cCodeFormat",data:function(){return{codeJsonContent:"点击 <劫持数据> 可以看到推送进GoMessage的原始数据...\n\n每次 <劫持数据> 都可以实时抓取到最新的一条数据...\n\n此处对数据进行了格式化与对齐，您可以把数据 <一键复制> 到其它地方使用...",codeUpdateTime:null,arr2:[]}},computed:{getStoreNamespace:function(){return this.$store.getters.getNamespace}},methods:{copyCode:function(){var e=document.getElementById("CodeContent");window.getSelection().selectAllChildren(e),document.execCommand("Copy"),this.$message("复制成功...")},getServerData:function(){var e=this;Object(l["p"])(this.$store.getters.getNamespace,null).then((function(t){var a=t.data.result["request_json"],n=t.data.result["request_time"];null===a||0===a.length?e.$message({showClose:!1,message:"没有数据进入GoMessage服务，无法向您展示数据。"}):(e.codeJsonContent=JSON.stringify(a,null,2),e.codeUpdateTime=e.dateFormat(new Date(n)))})).catch((function(e){console.log(e)}))}}},i=o,r=(a("fa18"),a("2877")),c=Object(r["a"])(i,n,s,!1,null,"5e781ea1",null);t["default"]=c.exports},8382:function(e,t,a){},a525:function(e,t,a){"use strict";a("f220")},b64b:function(e,t,a){var n=a("23e7"),s=a("7b0b"),l=a("df75"),o=a("d039"),i=o((function(){l(1)}));n({target:"Object",stat:!0,forced:i},{keys:function(e){return l(s(e))}})},bc05:function(e,t,a){"use strict";a("07c5")},c2ac:function(e,t,a){},e00d:function(e,t,a){"use strict";a("46d9")},f220:function(e,t,a){},f744:function(e,t,a){"use strict";a.r(t);var n=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",[a("el-row",{directives:[{name:"loading",rawName:"v-loading",value:e.dialogVisible,expression:"dialogVisible"}],staticStyle:{padding:"20px 0","margin-left":"0","margin-right":"0"},attrs:{gutter:20,"element-loading-text":'启用【渲染功能】会把过境数据"渲染为人类可读"的信息；若不开启则只会将数据"原封不动"的送达至目标客户端！',"element-loading-spinner":"el-icon-info","element-loading-background":"rgba(0, 0, 0, 0.95)"}},[a("el-col",{attrs:{span:12}},[a("DataFormat",{staticClass:"shadow"}),a("br"),!1===e.configMapType?[a("DataMap",{staticClass:"shadow"})]:[a("DataMap2",{staticClass:"shadow"})]],2),a("el-col",{attrs:{span:12}},[a("CTemplate",{staticClass:"shadow"})],1)],1)],1)},s=[],l=a("7d3c"),o=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("el-card",{attrs:{shadow:"always"}},[a("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[a("span",{staticStyle:{"padding-left":"50px"}},[e._v("自定义变量映射")]),a("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:e.post_data}},[e._v("保存变量")])],1),a("el-table",{staticStyle:{width:"100%"},attrs:{border:!0,data:e.configList}},[a("el-table-column",{attrs:{label:"Key（自定义变量名）",prop:"key"}}),a("el-table-column",{attrs:{label:"Value（原始数据中的索引）",prop:"value"}}),a("el-table-column",{attrs:{fixed:"right",label:"操作",width:"100"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("el-button",{attrs:{size:"small",type:"text"},nativeOn:{click:function(a){return a.preventDefault(),e.deleteRow(t.$index,e.configList)}}},[e._v("移除 ")])]}}])})],1),a("br"),a("el-form",{ref:"newMap",attrs:{model:e.newMap,rules:e.newMapRules}},[a("el-row",{attrs:{gutter:20}},[a("el-col",{attrs:{span:12}},[a("div",{staticClass:"DataMapGridContent"},[a("el-form-item",{attrs:{prop:"mapKey"}},[a("el-input",{attrs:{placeholder:"纯字符串，不能有符号"},model:{value:e.newMap.mapKey,callback:function(t){e.$set(e.newMap,"mapKey",t)},expression:"newMap.mapKey"}},[a("template",{slot:"prepend"},[e._v("Key:")])],2)],1)],1)]),a("el-col",{attrs:{span:12}},[a("div",{staticClass:"DataMapGridContent"},[a("el-form-item",{attrs:{prop:"mapValue"}},[a("el-input",{attrs:{placeholder:"Json索引路径"},model:{value:e.newMap.mapValue,callback:function(t){e.$set(e.newMap,"mapValue",t)},expression:"newMap.mapValue"}},[a("template",{slot:"prepend"},[e._v("Value:")])],2)],1)],1)])],1)],1),a("br"),a("el-button",{attrs:{size:"mini",type:"primary"},on:{click:e.addTableData}},[e._v("新增变量")])],1)},i=[],r=(a("a434"),a("ac1f"),a("1276"),a("d3b7"),a("159b"),a("6cac")),c=a("61f7"),u={name:"cConfigMap",data:function(){return{newMap:{mapKey:"",mapValue:""},newMapRules:{mapKey:[{required:!0,message:"不能为空",trigger:"blur"},{validator:c["b"],trigger:"blur"},{validator:c["c"],trigger:"blur"}],mapValue:[{required:!0,message:"不能为空",trigger:"blur"},{validator:c["b"],trigger:"blur"}]},configList:[{key:"{{ .N1 }} ",value:"alerts.#.labels.alertname"},{key:"{{ .N2 }}",value:"alerts.#.labels.severity"},{key:"{{ .N3 }}",value:"alerts.#.labels.hostname"},{key:"{{ .N4 }}",value:"alerts.#.labels.ping"},{key:"{{ .N5 }}",value:"alerts.#.annotations.description"},{key:"{{ .N6 }}",value:"status"},{key:"{{ .N7 }}",value:"alerts.#.startsAt"},{key:"{{ .N8 }}",value:"alerts.#.endsAt"}]}},computed:{},methods:{addTableData:function(){var e=this;this.$refs["newMap"].validate((function(t){if(t){var a=e.newMap.mapKey,n=e.newMap.mapValue;if(0===a.length||0===n.length)e.$message.error("输入框不能为空...");else{var s={key:"{{ ."+a+" }}",value:n};e.configList.push(s),e.newMap.mapKey="",e.newMap.mapValue="",e.$message.success("添加成功..."),e.post_data()}}}))},deleteRow:function(e,t){t.splice(e,1),this.post_data()},post_data:function(){for(var e=this,t=!(arguments.length>0&&void 0!==arguments[0])||arguments[0],a=[],n=0;n<this.configList.length;n++){var s=this.configList[n].key,l=s.split(" ")[1].split(".")[1],o=this.configList[n].value,i={};i[l]=o,a.push(i)}Object(r["l"])(this.$store.getters.getNamespace,{key_value_list:a}).then((function(a){console.log(a.data),t&&e.$message.success("数据库更新成功...")})).catch((function(e){console.log(e)}))},pullMapData:function(){var e=this;Object(r["q"])(this.$store.getters.getNamespace,null).then((function(t){var a=[];0===t.data.result.length?(console.log("数据库里没有数据"),e.post_data(!1)):(t.data.result.forEach((function(e){var t={key:"{{ ."+e["key"]+" }}",value:e["value"]};a.push(t)})),e.configList=a)})).catch((function(e){console.log(e)}))}},created:function(){this.pullMapData()}},p=u,d=(a("a525"),a("2877")),m=Object(d["a"])(p,o,i,!1,null,"cd8ab430",null),f=m.exports,g=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("el-card",{attrs:{shadow:"always"}},[a("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[a("span",{staticStyle:{"padding-left":"50px"}},[e._v("分解过境数据")]),a("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:e.getServerData}},[e._v("刷新")])],1),a("el-table",{staticStyle:{width:"100%"},attrs:{border:!0,data:e.configList}},[a("el-table-column",{attrs:{label:"Key",prop:"key"}}),a("el-table-column",{attrs:{label:"Value",prop:"value"},scopedSlots:e._u([{key:"default",fn:function(t){return[a("div",{staticStyle:{"white-space":"nowrap",overflow:"hidden","text-overflow":"ellipsis"}},[e._v(e._s(t.row.value))])]}}])})],1)],1)},h=[],v=(a("b64b"),{name:"cConfigMap2",data:function(){return{configList:[{key:"{{ .N1 }} ",value:"alerts.#.labels.alertname"},{key:"{{ .N2 }}",value:"alerts.#.labels.severity"},{key:"{{ .N3 }}",value:"alerts.#.labels.hostname"},{key:"{{ .N4 }}",value:"alerts.#.labels.ping"},{key:"{{ .N5 }}",value:"alerts.#.annotations.description"},{key:"{{ .N6 }}",value:"status"},{key:"{{ .N7 }}",value:"alerts.#.startsAt"},{key:"{{ .N8 }}",value:"alerts.#.endsAt"}]}},computed:{},methods:{getServerData:function(){var e=this;Object(r["o"])(this.$store.getters.getNamespace,null).then((function(t){var a=t.data.result["key_value_map"],n=t.data.result["request_time"];if(null===a||0===a.length)e.$message({showClose:!1,message:"没有数据进入GoMessage服务，无法向您展示数据。"});else{var s=[],l=Object.keys(a);l.forEach((function(e){var t={key:"{{ ."+e+" }}",value:a[e]};s.push(t)})),e.configList=s,e.codeUpdateTime=e.dateFormat(new Date(n))}})).catch((function(e){console.log(e)}))},pullMapData:function(){var e=this;Object(r["q"])(this.$store.getters.getNamespace,null).then((function(t){var a=[];0===t.data.result.length?console.log("数据库里没有数据"):(t.data.result.forEach((function(e){var t={key:"{{ ."+e["key"]+" }}",value:e["value"]};a.push(t)})),e.configList=a)})).catch((function(e){console.log(e)}))}},created:function(){this.getServerData()}}),b=v,y=(a("3ca5"),Object(d["a"])(b,g,h,!1,null,"083c0292",null)),w=y.exports,_=function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("el-card",{attrs:{shadow:"always"}},[a("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[a("el-switch",{staticStyle:{float:"left",padding:"3px 0"},attrs:{"inactive-text":"聚合发送"},on:{change:e.pushTemplateData},model:{value:e.template.template_is_merge,callback:function(t){e.$set(e.template,"template_is_merge",t)},expression:"template.template_is_merge"}}),a("el-tooltip",{staticClass:"item",staticStyle:{float:"left","margin-left":"30px","padding-top":"3px"},attrs:{content:'跳转到新页面：查看"消息模板"编写教程',effect:"dark",placement:"bottom"}},[a("el-link",{attrs:{underline:!1,type:"primary"},on:{click:e.newTagPage}},[a("span",[a("i",{staticClass:"el-icon-info"})])])],1),a("span",{staticStyle:{"margin-right":"70px"}},[e._v("消息渲染模板")]),a("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(t){return e.pushTemplateData()}}},[e._v("保存模板")])],1),a("div",[a("el-input",{attrs:{autosize:{minRows:10,maxRows:200},placeholder:"请输入Golang语法的模板内容",resize:"none",type:"textarea"},model:{value:e.template.template_content,callback:function(t){e.$set(e.template,"template_content",t)},expression:"template.template_content"}})],1)])},k=[],N={name:"cTemplate",data:function(){return{template:{template_is_merge:!1,template_content:"{{ if eq .N6 \"firing\" }}\n\n## <font color='#FF0000'>【报警中】服务器{{ .N3 }}</font>\n\n{{ else if eq .N6 \"resolved\" }}\n\n## <font color='#20B2AA'>【已恢复】服务器{{ .N3 }}</font>\n\n{{ else }}\n\n## 标题：信息通知\n\n{{ end }}\n\n====================\n\n**告警规则**：{{ .N1 }}\n\n**告警级别**：{{ .N2 }}\n\n**主机名称**：{{ .N3 }} \n\n**主机地址**：{{ .N4 }}\n\n**告警详情**：{{ .N5 }}\n\n**告警状态**：{{ .N6 }}\n\n**触发时间**：{{ .N7 }}\n\n**发送时间**：{{ .N8 }}\n\n**规则详情**：[Prometheus控制台](https://www.baidu.com)\n\n**报警详情**：[Alertmanager控制台](https://www.baidu.com)\n"}}},components:{},methods:{newTagPage:function(){var e="https://github.com/gomessage/gomessage#gomessage";window.open(e)},pushTemplateData:function(){var e=this,t=!(arguments.length>0&&void 0!==arguments[0])||arguments[0];Object(r["k"])(this.$store.getters.getNamespace,this.template).then((function(a){console.log(a.data.result),t&&e.$message.success("数据库更新成功...")})).catch((function(e){console.log(e)}))},pullTemplateData:function(){var e=this;Object(r["g"])(this.$store.getters.getNamespace,null).then((function(t){var a=t.data.result[0]["template_content"],n=t.data.result[0]["template_is_merge"];void 0===a||null===a||""===a?(console.log("数据库里没有数据"),e.pushTemplateData(!1)):(e.template.template_content=a,e.template.template_is_merge=n)})).catch((function(e){console.log(e)}))}},created:function(){this.$store.commit("updateStepsActive",2),this.pullTemplateData()}},x=N,C=(a("e00d"),Object(d["a"])(x,_,k,!1,null,"819f13a6",null)),$=C.exports,M={name:"ViewRequestData",data:function(){return{thisRenders:!0,dialogVisible:!1,configMapType:!1}},components:{DataFormat:l["default"],DataMap:f,CTemplate:$,DataMap2:w},computed:{getThisRenders:function(){return this.$store.getters.getNamespaceInfo["is_renders"]}},methods:{getNamespaceRenders:function(){var e=this,t=this.$store.getters.getNamespaceInfo;Object(r["f"])(t.id,null).then((function(t){e.thisRenders=t.data.result.is_renders,e.$store.commit("updateNamespaceInfo",t.data.result)}))},updateNamespaceRenders:function(){var e=this,t=this.$store.getters.getNamespaceInfo;t["is_renders"]=this.thisRenders,Object(r["n"])(t.id,t).then((function(t){e.thisRenders=t.data.result["is_renders"],e.$store.commit("updateNamespaceInfo",t.data.result)}))}},watch:{thisRenders:{immediate:!0,handler:function(e,t){console.log(e,t),this.dialogVisible=!e}}},created:function(){this.getNamespaceRenders(),this.$store.commit("updateStepsActive",1)}},S=M,D=(a("bc05"),Object(d["a"])(S,n,s,!1,null,"66538b6a",null));t["default"]=D.exports},fa18:function(e,t,a){"use strict";a("8382")}}]);
//# sourceMappingURL=chunk-5d7ac7c2.436c14c5.js.map