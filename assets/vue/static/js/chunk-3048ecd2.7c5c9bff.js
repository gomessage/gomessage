(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-3048ecd2"],{1360:function(e,t,a){},"1da9":function(e,t,a){},"23d6":function(e,t,a){"use strict";a("799b")},"74f9":function(e,t,a){"use strict";a("9a58")},"799b":function(e,t,a){},"7dbb":function(e,t,a){},"859a":function(e,t,a){"use strict";a.r(t);var s=function(){var e=this,t=e._self._c;return t("div",[t("el-row",{directives:[{name:"loading",rawName:"v-loading",value:e.dialogVisible,expression:"dialogVisible"}],staticStyle:{padding:"20px 0","margin-left":"0","margin-right":"0"},attrs:{gutter:20,"element-loading-text":'启用【渲染功能】会把过境数据"渲染为人类可读"的信息；若不开启则只会将数据"原封不动"的送达至目标客户端！',"element-loading-spinner":"el-icon-info","element-loading-background":"rgba(0, 0, 0, 0.95)"}},[t("el-col",{attrs:{span:12}},[t("DataFormat",{staticClass:"shadow"}),t("br"),!1===e.configMapType?[t("DataMap",{staticClass:"shadow"})]:[t("DataMap2",{staticClass:"shadow"})]],2),t("el-col",{attrs:{span:12}},[t("CTemplate",{staticClass:"shadow"})],1)],1)],1)},l=[],n=function(){var e=this,t=e._self._c;return t("el-card",{attrs:{shadow:"always"}},[t("div",{attrs:{slot:"header"},slot:"header"},[t("span",{staticStyle:{"padding-left":"100px"}},[e._v("劫持【 "),t("span",{staticStyle:{"font-size":"18px",color:"#3de1ad","font-weight":"900"}},[e._v(e._s(e.getStoreNamespace))]),e._v(" 】通道中的最新过境数据")]),t("el-button",{staticStyle:{float:"right",padding:"3px 0","margin-left":"10px"},attrs:{type:"text"},on:{click:function(t){return e.getServerData()}}},[e._v(" 劫持数据 ")]),t("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(t){return e.copyCode()}}},[e._v("一键复制")])],1),t("div",[t("pre",{attrs:{id:"CodeStyle"}},[t("code",{attrs:{id:"CodeContent"}},[e._v(e._s(e.codeJsonContent))])])])])},i=[],o=a("6cac"),r={name:"cCodeFormat",data(){return{codeJsonContent:"点击 <劫持数据> 可以看到推送进GoMessage的原始数据...\n\n每次 <劫持数据> 都可以实时抓取到最新的一条数据...\n\n此处对数据进行了格式化与对齐，您可以把数据 <一键复制> 到其它地方使用...",codeUpdateTime:null,arr2:[]}},computed:{getStoreNamespace:function(){return this.$store.getters.getNamespace}},methods:{copyCode:function(){const e=document.getElementById("CodeContent");window.getSelection().selectAllChildren(e),document.execCommand("Copy"),this.$message("复制成功...")},getServerData:function(){Object(o["s"])(this.$store.getters.getNamespace,null).then(e=>{const t=e.data.result["request_json"],a=e.data.result["request_time"];null===t||0===t.length?this.$message({showClose:!1,message:"没有数据进入GoMessage服务，无法向您展示数据。"}):(this.codeJsonContent=JSON.stringify(t,null,2),this.codeUpdateTime=this.dateFormat(new Date(a)))}).catch((function(e){console.log(e)}))}}},c=r,p=(a("a9da"),a("2877")),u=Object(p["a"])(c,n,i,!1,null,"5e781ea1",null),d=u.exports,m=function(){var e=this,t=e._self._c;return t("el-card",{attrs:{shadow:"always"}},[t("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[t("span",{staticStyle:{"padding-left":"50px"}},[e._v("自定义变量映射")]),t("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:e.PushVarData}},[e._v("保存变量")])],1),t("el-table",{staticStyle:{width:"100%"},attrs:{border:!0,data:e.configList}},[t("el-table-column",{attrs:{label:"Key（自定义变量名）",prop:"key"}}),t("el-table-column",{attrs:{label:"Value（原始数据中的索引）",prop:"value"}}),t("el-table-column",{attrs:{fixed:"right",label:"操作",width:"100"},scopedSlots:e._u([{key:"default",fn:function(a){return[t("el-button",{attrs:{size:"small",type:"text"},nativeOn:{click:function(t){return t.preventDefault(),e.deleteRow(a.$index,e.configList)}}},[e._v("移除 ")])]}}])})],1),t("br"),t("el-form",{ref:"newMap",attrs:{model:e.newMap,rules:e.newMapRules}},[t("el-row",{attrs:{gutter:20}},[t("el-col",{attrs:{span:12}},[t("div",{staticClass:"DataMapGridContent"},[t("el-form-item",{attrs:{prop:"mapKey"}},[t("el-input",{attrs:{placeholder:"纯字符串，不能有符号"},model:{value:e.newMap.mapKey,callback:function(t){e.$set(e.newMap,"mapKey",t)},expression:"newMap.mapKey"}},[t("template",{slot:"prepend"},[e._v("Key:")])],2)],1)],1)]),t("el-col",{attrs:{span:12}},[t("div",{staticClass:"DataMapGridContent"},[t("el-form-item",{attrs:{prop:"mapValue"}},[t("el-input",{attrs:{placeholder:"Json索引路径"},model:{value:e.newMap.mapValue,callback:function(t){e.$set(e.newMap,"mapValue",t)},expression:"newMap.mapValue"}},[t("template",{slot:"prepend"},[e._v("Value:")])],2)],1)],1)])],1)],1),t("br"),t("el-button",{attrs:{size:"mini",type:"primary"},on:{click:e.addTableData}},[e._v("新增变量")])],1)},h=[],g=(a("14d9"),a("61f7")),f={name:"cConfigMap",data(){return{newMap:{mapKey:"",mapValue:""},newMapRules:{mapKey:[{required:!0,message:"不能为空",trigger:"blur"},{validator:g["b"],trigger:"blur"},{validator:g["c"],trigger:"blur"}],mapValue:[{required:!0,message:"不能为空",trigger:"blur"},{validator:g["b"],trigger:"blur"}]},configList:[{key:"{{ .N1 }} ",value:"alerts.#.labels.alertname"},{key:"{{ .N2 }}",value:"alerts.#.labels.severity"},{key:"{{ .N3 }}",value:"alerts.#.labels.instance"},{key:"{{ .N4 }}",value:"alerts.#.annotations.description"},{key:"{{ .N5 }}",value:"status"},{key:"{{ .N6 }}",value:"alerts.#.startsAt"},{key:"{{ .N7 }}",value:"alerts.#.endsAt"}]}},computed:{},methods:{addTableData:function(){this.$refs["newMap"].validate(e=>{if(e){let e=this.newMap.mapKey,t=this.newMap.mapValue;if(0===e.length||0===t.length)this.$message.error("输入框不能为空...");else{let a={key:"{{ ."+e+" }}",value:t};this.configList.push(a),this.newMap.mapKey="",this.newMap.mapValue="",this.$message.success("添加成功..."),this.PushVarData()}}})},deleteRow(e,t){t.splice(e,1),this.PushVarData()},PushVarData:function(e=!0){const t=[];for(let a=0;a<this.configList.length;a++){let e=this.configList[a].key,s=e.split(" ")[1].split(".")[1],l=this.configList[a].value,n={};n[s]=l,t.push(n)}Object(o["n"])(this.$store.getters.getNamespace,{key_value_list:t}).then(t=>{console.log(t.data),e&&this.$message.success("数据库更新成功...")}).catch(e=>{console.log(e)})},PullVarData:function(){Object(o["t"])(this.$store.getters.getNamespace,null).then(e=>{if(0===e.data.result.length)this.PushVarData(!1);else{let t=[];e.data.result.forEach(e=>{let a={key:"{{ ."+e["key"]+" }}",value:e["value"]};t.push(a)}),this.configList=t}}).catch(e=>{console.log(e)})}},created(){this.PullVarData()}},v=f,y=(a("23d6"),Object(p["a"])(v,m,h,!1,null,"7ca0a878",null)),b=y.exports,w=function(){var e=this,t=e._self._c;return t("el-card",{attrs:{shadow:"always"}},[t("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[t("span",{staticStyle:{"padding-left":"50px"}},[e._v("分解过境数据")]),t("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:e.getServerData}},[e._v("刷新")])],1),t("el-table",{staticStyle:{width:"100%"},attrs:{border:!0,data:e.configList}},[t("el-table-column",{attrs:{label:"Key",prop:"key"}}),t("el-table-column",{attrs:{label:"Value",prop:"value"},scopedSlots:e._u([{key:"default",fn:function(a){return[t("div",{staticStyle:{"white-space":"nowrap",overflow:"hidden","text-overflow":"ellipsis"}},[e._v(e._s(a.row.value))])]}}])})],1)],1)},_=[],k={name:"cConfigMap2",data(){return{configList:[{key:"{{ .N1 }} ",value:"alerts.#.labels.alertname"},{key:"{{ .N2 }}",value:"alerts.#.labels.severity"},{key:"{{ .N3 }}",value:"alerts.#.labels.instance"},{key:"{{ .N4 }}",value:"alerts.#.annotations.description"},{key:"{{ .N5 }}",value:"status"},{key:"{{ .N6 }}",value:"alerts.#.startsAt"},{key:"{{ .N7 }}",value:"alerts.#.endsAt"}]}},computed:{},methods:{getServerData:function(){Object(o["r"])(this.$store.getters.getNamespace,null).then(e=>{const t=e.data.result["key_value_map"],a=e.data.result["request_time"];if(null===t||0===t.length)this.$message({showClose:!1,message:"没有数据进入GoMessage服务，无法向您展示数据。"});else{let e=[],s=Object.keys(t);s.forEach(a=>{let s={key:"{{ ."+a+" }}",value:t[a]};e.push(s)}),this.configList=e,this.codeUpdateTime=this.dateFormat(new Date(a))}}).catch((function(e){console.log(e)}))},pullMapData:function(){Object(o["t"])(this.$store.getters.getNamespace,null).then(e=>{let t=[];0===e.data.result.length?console.log("数据库里没有数据"):(e.data.result.forEach(e=>{let a={key:"{{ ."+e["key"]+" }}",value:e["value"]};t.push(a)}),this.configList=t)}).catch(e=>{console.log(e)})}},created(){this.getServerData()}},N=k,x=(a("8f43"),Object(p["a"])(N,w,_,!1,null,"54066553",null)),C=x.exports,D=function(){var e=this,t=e._self._c;return t("el-card",{attrs:{shadow:"always"}},[t("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[t("el-switch",{staticStyle:{float:"left",padding:"3px 0"},attrs:{"inactive-text":"聚合发送"},on:{change:e.PushTemplateData},model:{value:e.template.template_is_merge,callback:function(t){e.$set(e.template,"template_is_merge",t)},expression:"template.template_is_merge"}}),t("el-tooltip",{staticClass:"item",staticStyle:{float:"left","margin-left":"30px","padding-top":"3px"},attrs:{content:'跳转到新页面：查看"消息模板"编写教程',effect:"dark",placement:"bottom"}},[t("el-link",{attrs:{underline:!1,type:"primary"},on:{click:e.newTagPage}},[t("span",[t("i",{staticClass:"el-icon-info"})])])],1),t("span",{staticStyle:{"margin-right":"70px"}},[e._v("消息渲染模板")]),t("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(t){return e.PushTemplateData()}}},[e._v("保存模板")])],1),t("div",[t("el-input",{attrs:{autosize:{minRows:10,maxRows:200},placeholder:"请输入Golang语法的模板内容",resize:"none",type:"textarea"},model:{value:e.template.template_content,callback:function(t){e.$set(e.template,"template_content",t)},expression:"template.template_content"}})],1)])},M=[],S={name:"cTemplate",data(){return{template:{template_is_merge:!1,template_content:"{{ if eq .N5 \"firing\" }}\n\n## <font color='#FF0000'>【报警中】服务器{{ .N3 }}</font>\n\n{{ else if eq .N5 \"resolved\" }}\n\n## <font color='#20B2AA'>【已恢复】服务器{{ .N3 }}</font>\n\n{{ else }}\n\n## 标题：信息通知\n\n{{ end }}\n\n====================\n\n**告警规则**：{{ .N1 }}\n\n**告警级别**：{{ .N2 }}\n\n**主机名称**：{{ .N3 }} \n\n**告警详情**：{{ .N4 }}\n\n**告警状态**：{{ .N5 }}\n\n**触发时间**：{{ .N6 }}\n\n**发送时间**：{{ .N7 }}\n\n**规则详情**：[Prometheus控制台](https://www.baidu.com)\n\n**报警详情**：[Alertmanager控制台](https://www.baidu.com)\n"}}},components:{},methods:{newTagPage:function(){let e="https://github.com/gomessage/gomessage#gomessage";window.open(e)},PushTemplateData:function(e=!0){Object(o["m"])(this.$store.getters.getNamespace,this.template).then(t=>{console.log(t.data.result),e&&this.$message.success("数据库更新成功...")}).catch(e=>{console.log(e)})},PullTemplateData:function(){Object(o["g"])(this.$store.getters.getNamespace,null).then(e=>{if(0===e.data.result.length)this.PushTemplateData(!1);else{let t=e.data.result[0]["template_content"],a=e.data.result[0]["template_is_merge"];this.template.template_content=t,this.template.template_is_merge=a}}).catch(e=>{console.log(e)})}},created(){this.$store.commit("updateStepsActive",2),this.PullTemplateData()}},$=S,V=(a("74f9"),Object(p["a"])($,D,M,!1,null,"57ef3e70",null)),T=V.exports,O={name:"ViewRequestData",data(){return{thisRenders:!0,dialogVisible:!1,configMapType:!1}},components:{DataFormat:d,DataMap:b,CTemplate:T,DataMap2:C},computed:{getThisRenders:function(){return this.$store.getters.getNamespaceInfo["is_renders"]}},methods:{getNamespaceRenders:function(){let e=this.$store.getters.getNamespaceInfo;Object(o["f"])(e.id,null).then(e=>{this.thisRenders=e.data.result.is_renders,this.$store.commit("updateNamespaceInfo",e.data.result)})},updateNamespaceRenders:function(){let e=this.$store.getters.getNamespaceInfo;e["is_renders"]=this.thisRenders,Object(o["q"])(e.id,e).then(e=>{this.thisRenders=e.data.result["is_renders"],this.$store.commit("updateNamespaceInfo",e.data.result)})}},watch:{thisRenders:{immediate:!0,handler(e,t){console.log(e,t),this.dialogVisible=!e}}},created(){this.getNamespaceRenders(),this.$store.commit("updateStepsActive",1)}},j=O,P=(a("dee5"),Object(p["a"])(j,s,l,!1,null,"29c74960",null));t["default"]=P.exports},"8f43":function(e,t,a){"use strict";a("7dbb")},"9a58":function(e,t,a){},a9da:function(e,t,a){"use strict";a("1360")},dee5:function(e,t,a){"use strict";a("1da9")}}]);
//# sourceMappingURL=chunk-3048ecd2.7c5c9bff.js.map