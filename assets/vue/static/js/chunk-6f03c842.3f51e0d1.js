(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-6f03c842"],{1360:function(t,e,a){},"23d6":function(t,e,a){"use strict";a("799b")},"61f7":function(t,e,a){"use strict";function n(t,e,a){let n=/^[0-9]+$/,s=n.test(e.toString());s&&a("不能以数字开头"),a()}function s(t,e,a){e.length<3&&a("长度不能小于3个字符"),a()}function l(t,e,a){let n=/^[A-Za-z0-9]+$/,s=n.test(e.toString());s||a("只能是字母或数字，不能包含特殊符号"),a()}a.d(e,"b",(function(){return n})),a.d(e,"a",(function(){return s})),a.d(e,"c",(function(){return l}))},"6cac":function(t,e,a){"use strict";a.d(e,"k",(function(){return s})),a.d(e,"s",(function(){return l})),a.d(e,"t",(function(){return i})),a.d(e,"n",(function(){return r})),a.d(e,"r",(function(){return o})),a.d(e,"g",(function(){return c})),a.d(e,"m",(function(){return u})),a.d(e,"c",(function(){return p})),a.d(e,"j",(function(){return d})),a.d(e,"d",(function(){return m})),a.d(e,"p",(function(){return f})),a.d(e,"o",(function(){return g})),a.d(e,"a",(function(){return h})),a.d(e,"e",(function(){return v})),a.d(e,"l",(function(){return b})),a.d(e,"f",(function(){return y})),a.d(e,"b",(function(){return w})),a.d(e,"q",(function(){return _})),a.d(e,"h",(function(){return k})),a.d(e,"i",(function(){return N}));var n=a("be3b");const s=(t,e)=>n["a"].Post("/go/"+t,e),l=(t,e)=>n["a"].Get("/api/v1/"+t+"/json",e),i=(t,e)=>n["a"].Get("/api/v1/"+t+"/vars",e),r=(t,e)=>n["a"].Post("/api/v1/"+t+"/vars",e),o=(t,e)=>n["a"].Get("/api/v1/"+t+"/flattening",e),c=(t,e)=>n["a"].Get("/api/v1/"+t+"/template",e),u=(t,e)=>n["a"].Post("/api/v1/"+t+"/template",e),p=(t,e)=>n["a"].Get("/api/v1/"+t+"/client",e),d=(t,e)=>n["a"].Post("/api/v1/"+t+"/client",e),m=(t,e,a)=>n["a"].Get("/api/v1/"+t+"/client/"+e,a),f=(t,e,a)=>n["a"].Put("/api/v1/"+t+"/client/"+e,a),g=(t,e,a)=>n["a"].Put("/api/v1/"+t+"/client-info/"+e,a),h=(t,e,a)=>n["a"].Delete("/api/v1/"+t+"/client/"+e,a),v=t=>n["a"].Get("/api/v1/namespace",t),b=t=>n["a"].Post("/api/v1/namespace",t),y=(t,e)=>n["a"].Get("/api/v1/namespace/"+t,e),w=t=>n["a"].Delete("/api/v1/namespace/"+t),_=(t,e)=>n["a"].Put("/api/v1/namespace/"+t,e),k=t=>n["a"].Post("/auth/login",t),N=t=>n["a"].Post("/auth/logout",t)},"74f9":function(t,e,a){"use strict";a("9a58")},"799b":function(t,e,a){},"7dbb":function(t,e,a){},"8f43":function(t,e,a){"use strict";a("7dbb")},"9a58":function(t,e,a){},a9da:function(t,e,a){"use strict";a("1360")},e4be:function(t,e,a){},ed2a:function(t,e,a){"use strict";a("e4be")},f744:function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t._self._c;return e("div",[e("el-row",{directives:[{name:"loading",rawName:"v-loading",value:t.dialogVisible,expression:"dialogVisible"}],staticStyle:{padding:"20px 0","margin-left":"0","margin-right":"0"},attrs:{gutter:20,"element-loading-text":'启用【渲染功能】会把过境数据"渲染为人类可读"的信息；若不开启则只会将数据"原封不动"的送达至目标客户端！',"element-loading-spinner":"el-icon-info","element-loading-background":"rgba(0, 0, 0, 0.95)"}},[e("el-col",{attrs:{span:12}},[e("DataFormat",{staticClass:"shadow"}),e("br"),!1===t.configMapType?[e("DataMap",{staticClass:"shadow"})]:[e("DataMap2",{staticClass:"shadow"})]],2),e("el-col",{attrs:{span:12}},[e("CTemplate",{staticClass:"shadow"})],1)],1)],1)},s=[],l=function(){var t=this,e=t._self._c;return e("el-card",{attrs:{shadow:"always"}},[e("div",{attrs:{slot:"header"},slot:"header"},[e("span",{staticStyle:{"padding-left":"100px"}},[t._v("劫持【 "),e("span",{staticStyle:{"font-size":"18px",color:"#3de1ad","font-weight":"900"}},[t._v(t._s(t.getStoreNamespace))]),t._v(" 】通道中的最新过境数据")]),e("el-button",{staticStyle:{float:"right",padding:"3px 0","margin-left":"10px"},attrs:{type:"text"},on:{click:function(e){return t.getServerData()}}},[t._v(" 劫持数据 ")]),e("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(e){return t.copyCode()}}},[t._v("一键复制")])],1),e("div",[e("pre",{attrs:{id:"CodeStyle"}},[e("code",{attrs:{id:"CodeContent"}},[t._v(t._s(t.codeJsonContent))])])])])},i=[],r=a("6cac"),o={name:"cCodeFormat",data(){return{codeJsonContent:"点击 <劫持数据> 可以看到推送进GoMessage的原始数据...\n\n每次 <劫持数据> 都可以实时抓取到最新的一条数据...\n\n此处对数据进行了格式化与对齐，您可以把数据 <一键复制> 到其它地方使用...",codeUpdateTime:null,arr2:[]}},computed:{getStoreNamespace:function(){return this.$store.getters.getNamespace}},methods:{copyCode:function(){const t=document.getElementById("CodeContent");window.getSelection().selectAllChildren(t),document.execCommand("Copy"),this.$message("复制成功...")},getServerData:function(){Object(r["s"])(this.$store.getters.getNamespace,null).then(t=>{const e=t.data.result["request_json"],a=t.data.result["request_time"];null===e||0===e.length?this.$message({showClose:!1,message:"没有数据进入GoMessage服务，无法向您展示数据。"}):(this.codeJsonContent=JSON.stringify(e,null,2),this.codeUpdateTime=this.dateFormat(new Date(a)))}).catch((function(t){console.log(t)}))}}},c=o,u=(a("a9da"),a("2877")),p=Object(u["a"])(c,l,i,!1,null,"5e781ea1",null),d=p.exports,m=function(){var t=this,e=t._self._c;return e("el-card",{attrs:{shadow:"always"}},[e("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[e("span",{staticStyle:{"padding-left":"50px"}},[t._v("自定义变量映射")]),e("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:t.PushVarData}},[t._v("保存变量")])],1),e("el-table",{staticStyle:{width:"100%"},attrs:{border:!0,data:t.configList}},[e("el-table-column",{attrs:{label:"Key（自定义变量名）",prop:"key"}}),e("el-table-column",{attrs:{label:"Value（原始数据中的索引）",prop:"value"}}),e("el-table-column",{attrs:{fixed:"right",label:"操作",width:"100"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("el-button",{attrs:{size:"small",type:"text"},nativeOn:{click:function(e){return e.preventDefault(),t.deleteRow(a.$index,t.configList)}}},[t._v("移除 ")])]}}])})],1),e("br"),e("el-form",{ref:"newMap",attrs:{model:t.newMap,rules:t.newMapRules}},[e("el-row",{attrs:{gutter:20}},[e("el-col",{attrs:{span:12}},[e("div",{staticClass:"DataMapGridContent"},[e("el-form-item",{attrs:{prop:"mapKey"}},[e("el-input",{attrs:{placeholder:"纯字符串，不能有符号"},model:{value:t.newMap.mapKey,callback:function(e){t.$set(t.newMap,"mapKey",e)},expression:"newMap.mapKey"}},[e("template",{slot:"prepend"},[t._v("Key:")])],2)],1)],1)]),e("el-col",{attrs:{span:12}},[e("div",{staticClass:"DataMapGridContent"},[e("el-form-item",{attrs:{prop:"mapValue"}},[e("el-input",{attrs:{placeholder:"Json索引路径"},model:{value:t.newMap.mapValue,callback:function(e){t.$set(t.newMap,"mapValue",e)},expression:"newMap.mapValue"}},[e("template",{slot:"prepend"},[t._v("Value:")])],2)],1)],1)])],1)],1),e("br"),e("el-button",{attrs:{size:"mini",type:"primary"},on:{click:t.addTableData}},[t._v("新增变量")])],1)},f=[],g=(a("14d9"),a("61f7")),h={name:"cConfigMap",data(){return{newMap:{mapKey:"",mapValue:""},newMapRules:{mapKey:[{required:!0,message:"不能为空",trigger:"blur"},{validator:g["b"],trigger:"blur"},{validator:g["c"],trigger:"blur"}],mapValue:[{required:!0,message:"不能为空",trigger:"blur"},{validator:g["b"],trigger:"blur"}]},configList:[{key:"{{ .N1 }} ",value:"alerts.#.labels.alertname"},{key:"{{ .N2 }}",value:"alerts.#.labels.severity"},{key:"{{ .N3 }}",value:"alerts.#.labels.instance"},{key:"{{ .N4 }}",value:"alerts.#.annotations.description"},{key:"{{ .N5 }}",value:"status"},{key:"{{ .N6 }}",value:"alerts.#.startsAt"},{key:"{{ .N7 }}",value:"alerts.#.endsAt"}]}},computed:{},methods:{addTableData:function(){this.$refs["newMap"].validate(t=>{if(t){let t=this.newMap.mapKey,e=this.newMap.mapValue;if(0===t.length||0===e.length)this.$message.error("输入框不能为空...");else{let a={key:"{{ ."+t+" }}",value:e};this.configList.push(a),this.newMap.mapKey="",this.newMap.mapValue="",this.$message.success("添加成功..."),this.PushVarData()}}})},deleteRow(t,e){e.splice(t,1),this.PushVarData()},PushVarData:function(t=!0){const e=[];for(let a=0;a<this.configList.length;a++){let t=this.configList[a].key,n=t.split(" ")[1].split(".")[1],s=this.configList[a].value,l={};l[n]=s,e.push(l)}Object(r["n"])(this.$store.getters.getNamespace,{key_value_list:e}).then(e=>{console.log(e.data),t&&this.$message.success("数据库更新成功...")}).catch(t=>{console.log(t)})},PullVarData:function(){Object(r["t"])(this.$store.getters.getNamespace,null).then(t=>{if(0===t.data.result.length)this.PushVarData(!1);else{let e=[];t.data.result.forEach(t=>{let a={key:"{{ ."+t["key"]+" }}",value:t["value"]};e.push(a)}),this.configList=e}}).catch(t=>{console.log(t)})}},created(){this.PullVarData()}},v=h,b=(a("23d6"),Object(u["a"])(v,m,f,!1,null,"7ca0a878",null)),y=b.exports,w=function(){var t=this,e=t._self._c;return e("el-card",{attrs:{shadow:"always"}},[e("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[e("span",{staticStyle:{"padding-left":"50px"}},[t._v("分解过境数据")]),e("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:t.getServerData}},[t._v("刷新")])],1),e("el-table",{staticStyle:{width:"100%"},attrs:{border:!0,data:t.configList}},[e("el-table-column",{attrs:{label:"Key",prop:"key"}}),e("el-table-column",{attrs:{label:"Value",prop:"value"},scopedSlots:t._u([{key:"default",fn:function(a){return[e("div",{staticStyle:{"white-space":"nowrap",overflow:"hidden","text-overflow":"ellipsis"}},[t._v(t._s(a.row.value))])]}}])})],1)],1)},_=[],k={name:"cConfigMap2",data(){return{configList:[{key:"{{ .N1 }} ",value:"alerts.#.labels.alertname"},{key:"{{ .N2 }}",value:"alerts.#.labels.severity"},{key:"{{ .N3 }}",value:"alerts.#.labels.instance"},{key:"{{ .N4 }}",value:"alerts.#.annotations.description"},{key:"{{ .N5 }}",value:"status"},{key:"{{ .N6 }}",value:"alerts.#.startsAt"},{key:"{{ .N7 }}",value:"alerts.#.endsAt"}]}},computed:{},methods:{getServerData:function(){Object(r["r"])(this.$store.getters.getNamespace,null).then(t=>{const e=t.data.result["key_value_map"],a=t.data.result["request_time"];if(null===e||0===e.length)this.$message({showClose:!1,message:"没有数据进入GoMessage服务，无法向您展示数据。"});else{let t=[],n=Object.keys(e);n.forEach(a=>{let n={key:"{{ ."+a+" }}",value:e[a]};t.push(n)}),this.configList=t,this.codeUpdateTime=this.dateFormat(new Date(a))}}).catch((function(t){console.log(t)}))},pullMapData:function(){Object(r["t"])(this.$store.getters.getNamespace,null).then(t=>{let e=[];0===t.data.result.length?console.log("数据库里没有数据"):(t.data.result.forEach(t=>{let a={key:"{{ ."+t["key"]+" }}",value:t["value"]};e.push(a)}),this.configList=e)}).catch(t=>{console.log(t)})}},created(){this.getServerData()}},N=k,x=(a("8f43"),Object(u["a"])(N,w,_,!1,null,"54066553",null)),D=x.exports,C=function(){var t=this,e=t._self._c;return e("el-card",{attrs:{shadow:"always"}},[e("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[e("el-switch",{staticStyle:{float:"left",padding:"3px 0"},attrs:{"inactive-text":"聚合发送"},on:{change:t.PushTemplateData},model:{value:t.template.template_is_merge,callback:function(e){t.$set(t.template,"template_is_merge",e)},expression:"template.template_is_merge"}}),e("el-tooltip",{staticClass:"item",staticStyle:{float:"left","margin-left":"30px","padding-top":"3px"},attrs:{content:'跳转到新页面：查看"消息模板"编写教程',effect:"dark",placement:"bottom"}},[e("el-link",{attrs:{underline:!1,type:"primary"},on:{click:t.newTagPage}},[e("span",[e("i",{staticClass:"el-icon-info"})])])],1),e("span",{staticStyle:{"margin-right":"70px"}},[t._v("消息渲染模板")]),e("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(e){return t.PushTemplateData()}}},[t._v("保存模板")])],1),e("div",[e("el-input",{attrs:{autosize:{minRows:10,maxRows:200},placeholder:"请输入Golang语法的模板内容",resize:"none",type:"textarea"},model:{value:t.template.template_content,callback:function(e){t.$set(t.template,"template_content",e)},expression:"template.template_content"}})],1)])},S=[],M={name:"cTemplate",data(){return{template:{template_is_merge:!1,template_content:"{{ if eq .N5 \"firing\" }}\n\n## <font color='#FF0000'>【报警中】服务器{{ .N3 }}</font>\n\n{{ else if eq .N5 \"resolved\" }}\n\n## <font color='#20B2AA'>【已恢复】服务器{{ .N3 }}</font>\n\n{{ else }}\n\n## 标题：信息通知\n\n{{ end }}\n\n====================\n\n**告警规则**：{{ .N1 }}\n\n**告警级别**：{{ .N2 }}\n\n**主机名称**：{{ .N3 }} \n\n**告警详情**：{{ .N4 }}\n\n**告警状态**：{{ .N5 }}\n\n**触发时间**：{{ .N6 }}\n\n**发送时间**：{{ .N7 }}\n\n**规则详情**：[Prometheus控制台](https://www.baidu.com)\n\n**报警详情**：[Alertmanager控制台](https://www.baidu.com)\n"}}},components:{},methods:{newTagPage:function(){let t="https://github.com/gomessage/gomessage#gomessage";window.open(t)},PushTemplateData:function(t=!0){Object(r["m"])(this.$store.getters.getNamespace,this.template).then(e=>{console.log(e.data.result),t&&this.$message.success("数据库更新成功...")}).catch(t=>{console.log(t)})},PullTemplateData:function(){Object(r["g"])(this.$store.getters.getNamespace,null).then(t=>{if(0===t.data.result.length)this.PushTemplateData(!1);else{let e=t.data.result[0]["template_content"],a=t.data.result[0]["template_is_merge"];this.template.template_content=e,this.template.template_is_merge=a}}).catch(t=>{console.log(t)})}},created(){this.$store.commit("updateStepsActive",2),this.PullTemplateData()}},$=M,P=(a("74f9"),Object(u["a"])($,C,S,!1,null,"57ef3e70",null)),V=P.exports,T={name:"ViewRequestData",data(){return{thisRenders:!0,dialogVisible:!1,configMapType:!1}},components:{DataFormat:d,DataMap:y,CTemplate:V,DataMap2:D},computed:{getThisRenders:function(){return this.$store.getters.getNamespaceInfo["is_renders"]}},methods:{getNamespaceRenders:function(){let t=this.$store.getters.getNamespaceInfo;Object(r["f"])(t.id,null).then(t=>{this.thisRenders=t.data.result.is_renders,this.$store.commit("updateNamespaceInfo",t.data.result)})},updateNamespaceRenders:function(){let t=this.$store.getters.getNamespaceInfo;t["is_renders"]=this.thisRenders,Object(r["q"])(t.id,t).then(t=>{this.thisRenders=t.data.result["is_renders"],this.$store.commit("updateNamespaceInfo",t.data.result)})}},watch:{thisRenders:{immediate:!0,handler(t,e){console.log(t,e),this.dialogVisible=!t}}},created(){this.getNamespaceRenders(),this.$store.commit("updateStepsActive",1)}},j=T,O=(a("ed2a"),Object(u["a"])(j,n,s,!1,null,"66538b6a",null));e["default"]=O.exports}}]);
//# sourceMappingURL=chunk-6f03c842.3f51e0d1.js.map