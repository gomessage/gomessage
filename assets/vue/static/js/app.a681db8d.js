(function(e){function t(t){for(var a,r,s=t[0],c=t[1],u=t[2],l=0,p=[];l<s.length;l++)r=s[l],Object.prototype.hasOwnProperty.call(o,r)&&o[r]&&p.push(o[r][0]),o[r]=0;for(a in c)Object.prototype.hasOwnProperty.call(c,a)&&(e[a]=c[a]);d&&d(t);while(p.length)p.shift()();return i.push.apply(i,u||[]),n()}function n(){for(var e,t=0;t<i.length;t++){for(var n=i[t],a=!0,r=1;r<n.length;r++){var s=n[r];0!==o[s]&&(a=!1)}a&&(i.splice(t--,1),e=c(c.s=n[0]))}return e}var a={},r={app:0},o={app:0},i=[];function s(e){return c.p+"static/js/"+({}[e]||e)+"."+{"chunk-0b923e88":"6ba1ecb4","chunk-1f1bd5cd":"3e6ac13a","chunk-94ca40d8":"75da7a5d"}[e]+".js"}function c(t){if(a[t])return a[t].exports;var n=a[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,c),n.l=!0,n.exports}c.e=function(e){var t=[],n={"chunk-0b923e88":1,"chunk-1f1bd5cd":1,"chunk-94ca40d8":1};r[e]?t.push(r[e]):0!==r[e]&&n[e]&&t.push(r[e]=new Promise((function(t,n){for(var a="static/css/"+({}[e]||e)+"."+{"chunk-0b923e88":"026d007d","chunk-1f1bd5cd":"f938e7a8","chunk-94ca40d8":"5dc3a687"}[e]+".css",o=c.p+a,i=document.getElementsByTagName("link"),s=0;s<i.length;s++){var u=i[s],l=u.getAttribute("data-href")||u.getAttribute("href");if("stylesheet"===u.rel&&(l===a||l===o))return t()}var p=document.getElementsByTagName("style");for(s=0;s<p.length;s++){u=p[s],l=u.getAttribute("data-href");if(l===a||l===o)return t()}var d=document.createElement("link");d.rel="stylesheet",d.type="text/css",d.onload=t,d.onerror=function(t){var a=t&&t.target&&t.target.src||o,i=new Error("Loading CSS chunk "+e+" failed.\n("+a+")");i.code="CSS_CHUNK_LOAD_FAILED",i.request=a,delete r[e],d.parentNode.removeChild(d),n(i)},d.href=o;var m=document.getElementsByTagName("head")[0];m.appendChild(d)})).then((function(){r[e]=0})));var a=o[e];if(0!==a)if(a)t.push(a[2]);else{var i=new Promise((function(t,n){a=o[e]=[t,n]}));t.push(a[2]=i);var u,l=document.createElement("script");l.charset="utf-8",l.timeout=120,c.nc&&l.setAttribute("nonce",c.nc),l.src=s(e);var p=new Error;u=function(t){l.onerror=l.onload=null,clearTimeout(d);var n=o[e];if(0!==n){if(n){var a=t&&("load"===t.type?"missing":t.type),r=t&&t.target&&t.target.src;p.message="Loading chunk "+e+" failed.\n("+a+": "+r+")",p.name="ChunkLoadError",p.type=a,p.request=r,n[1](p)}o[e]=void 0}};var d=setTimeout((function(){u({type:"timeout",target:l})}),12e4);l.onerror=l.onload=u,document.head.appendChild(l)}return Promise.all(t)},c.m=e,c.c=a,c.d=function(e,t,n){c.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},c.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},c.t=function(e,t){if(1&t&&(e=c(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(c.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)c.d(n,a,function(t){return e[t]}.bind(null,a));return n},c.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return c.d(t,"a",t),t},c.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},c.p="",c.oe=function(e){throw console.error(e),e};var u=window["webpackJsonp"]=window["webpackJsonp"]||[],l=u.push.bind(u);u.push=t,u=u.slice();for(var p=0;p<u.length;p++)t(u[p]);var d=l;i.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"0342":function(e,t,n){"use strict";n("2b3d")},"034f":function(e,t,n){"use strict";n("85ec")},"2b3d":function(e,t,n){},"2b4c":function(e,t,n){"use strict";n("f4a2")},"56d7":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var a=n("2b0e"),r=(n("be3b"),function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("el-container",{attrs:{id:"MyContainer-Container"}},[n("el-header",{attrs:{id:"MyContainer-Header"}},[n("NavHeader")],1),n("el-container",[n("el-aside",{attrs:{id:"MyContainer-Aside",width:"275px"}},[n("NavAside")],1),n("el-main",{attrs:{id:"MyContainer-Main"}},[n("router-view")],1)],1),n("el-footer",{attrs:{id:"MyContainer-Footer"}},[n("NavFooter")],1)],1)],1)}),o=[],i=(n("e9c4"),function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("el-menu",{staticStyle:{height:"100%"},attrs:{"collapse-transition":!0,"default-active":e.$route.name,"active-text-color":"#ffd04b","background-color":"#161823",mode:"horizontal",router:"","text-color":"#fff"},on:{select:e.handleSelect}},[a("el-menu-item",[a("router-link",{attrs:{to:"/"}},[a("img",{attrs:{height:"90%",src:n("9b19")}}),e._v("    "),a("span",{staticStyle:{"font-size":"15px"}},[e._v("GoMessage · 消息转发器")])])],1),a("el-menu-item",{staticStyle:{padding:"0"}},[a("router-link",{attrs:{to:"/"}},[a("el-button",{staticStyle:{color:"#fff","margin-left":"10px"},attrs:{type:"text",icon:"el-icon-location"}},[e._v(" "+e._s(e.getStoreNamespace)+" ")])],1)],1),a("el-row",{attrs:{justify:"end",type:"flex"}},[a("router-link",{attrs:{to:"/"}},[a("el-menu-item",{attrs:{index:"1-1"}},[a("i",{staticClass:"el-icon-position"}),e._v("消息入口")])],1),a("router-link",{attrs:{to:"/request_data"}},[a("el-menu-item",{attrs:{index:"2-1"}},[a("i",{staticClass:"el-icon-set-up"}),e._v("数据渲染")])],1),a("router-link",{attrs:{to:"/data_client"}},[a("el-menu-item",{attrs:{index:"3-1"}},[a("i",{staticClass:"el-icon-chat-dot-square"}),e._v("接收客户端")])],1),a("el-menu-item",[a("el-link",{attrs:{underline:!1,href:"https://github.com/gomessage/gomessage#gomessage",target:"_blank",type:"primary"}},[a("i",{staticClass:"el-icon-document"}),e._v("文档 ")])],1)],1)],1)}),s=[],c={name:"NavHeader",data:function(){return{}},components:{},computed:{getStoreNamespace:function(){return this.$store.getters.getNamespace}},methods:{handleSelect:function(e,t){console.log(e,t)}}},u=c,l=(n("0342"),n("2877")),p=Object(l["a"])(u,i,s,!1,null,"83955578",null),d=p.exports,m=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-menu",{staticClass:"el-menu-demo",attrs:{id:"NavFooter-Footer",mode:"horizontal","background-color":"#161823","text-color":"#fff","active-text-color":"#ffd04b","collapse-transition":!0},on:{select:e.handleSelect}},[n("el-row",{attrs:{type:"flex",justify:"center",id:"my-row"}},[n("el-menu-item",{attrs:{id:"NavFooter-menu-item",disabled:""}},[n("el-link",{attrs:{href:"https://github.com/gomessage/gomessage",type:"primary"}},[e._v("@ 2020~2023  ")]),n("el-link",{attrs:{href:"https://opensource.org/licenses/MIT",type:"primary"}},[e._v("MIT License   ")]),n("el-link",{attrs:{href:"https://github.com/gomessage/gomessage",type:"primary"}},[e._v("Github ")]),e._v(" or "),n("el-link",{attrs:{href:"https://gitee.com/gomessage/gomessage",type:"primary"}},[e._v(" Gitee  ")]),e._v(" Version: "),n("el-link",{attrs:{href:"https://github.com/gomessage/gomessage/releases",type:"primary"}},[e._v("v2.x  ")]),e._v(" Give feedback and Suggestions to the author: "),n("el-link",{attrs:{href:"mailto:taycc3223@gmail.com",type:"primary"}},[e._v("Email")])],1)],1)],1)},f=[],g={name:"NavFooter",data:function(){return{activeIndex:"1",activeIndex2:"1"}},methods:{handleSelect:function(e,t){console.log(e,t)}}},h=g,b=(n("f9ab"),Object(l["a"])(h,m,f,!1,null,"560092b6",null)),v=b.exports,y=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-menu",{staticStyle:{height:"100%","border-right":"solid 0 red"},attrs:{"default-active":e.getStoreNamespace,"active-text-color":"#ffd04b","background-color":"#41555d","text-color":"#fff"}},[n("br"),e._l(e.namespaceList,(function(t,a){return n("el-menu-item",{key:a,staticStyle:{"text-align":"left"},attrs:{index:t.name},on:{click:function(n){return e.updateNamespace(t,n)}}},[n("i",{staticClass:"el-icon-menu"}),n("span",{attrs:{slot:"title"},slot:"title"},[e._v(e._s(t.name))])])})),n("br"),n("el-menu-item",{staticStyle:{"text-align":"left"},attrs:{index:"998"}},[n("el-button",{staticStyle:{"padding-right":"35px"},attrs:{icon:"el-icon-setting",plain:"",size:"mini"},on:{click:function(t){e.dialogFormVisible222=!0}}},[e._v("管理推送通道 ")])],1),n("el-dialog",{attrs:{visible:e.dialogFormVisible222,title:"消息推送通道",modal:"",width:"60%",top:"10vh","lock-scroll":""},on:{"update:visible":function(t){e.dialogFormVisible222=t}}},[n("el-table",{staticStyle:{width:"100%"},attrs:{data:e.namespaceList,border:"","header-cell-style":{background:"#2f2f35",color:"#fff"}}},[n("el-table-column",{attrs:{prop:"name",label:"名称",width:"180"}}),n("el-table-column",{attrs:{prop:"description",label:"描述"}}),n("el-table-column",{attrs:{fixed:"right",label:"操作",width:"100"},scopedSlots:e._u([{key:"default",fn:function(t){return[n("el-button",{attrs:{size:"small",type:"danger"},nativeOn:{click:function(n){return n.preventDefault(),e.deleteOneNamespace(t.$index,e.namespaceList)}}},[e._v("删除 ")])]}}])})],1),n("br"),n("p",{staticClass:"authorStatement2"},[e._v(' GoMessage v2 版本支持多通道并发，可以同时存在多个"通道"，互不干扰的进行消息转发。 ')]),n("br"),n("el-divider",{attrs:{"content-position":"left"}},[n("i",{staticClass:"el-icon-circle-plus-outline"},[e._v(" 新增通道")])]),n("el-form",{ref:"namespaceForm",staticStyle:{width:"60%"},attrs:{model:e.namespaceForm,rules:e.namespaceRules}},[n("el-form-item",{attrs:{label:"通道名称","label-width":"105px",prop:"name"}},[n("el-input",{attrs:{autocomplete:"off",placeholder:"请输入通道名称（只能是纯英文名称，不限大小写）"},model:{value:e.namespaceForm.name,callback:function(t){e.$set(e.namespaceForm,"name",t)},expression:"namespaceForm.name"}})],1),n("el-form-item",{attrs:{label:"通道描述","label-width":"105px"}},[n("el-input",{attrs:{autocomplete:"off",type:"textarea",rows:3,placeholder:"请输入推送通道的描述信息"},model:{value:e.namespaceForm.description,callback:function(t){e.$set(e.namespaceForm,"description",t)},expression:"namespaceForm.description"}})],1)],1),n("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[n("el-button",{on:{click:function(t){e.dialogFormVisible222=!1}}},[e._v("取 消")]),n("el-button",{attrs:{type:"primary"},on:{click:e.addNamespace}},[e._v("确 定")])],1)],1)],2)},_=[],x=(n("d3b7"),n("159b"),n("b0c0"),n("a434"),n("6cac")),k=n("61f7"),w={name:"NavAside",data:function(){return{namespaceList:[{name:"default",description:"default",is_active:!0,is_renders:!1}],dialogFormVisible222:!1,namespaceForm:{name:"",description:"",is_active:!1},namespaceRules:{name:[{required:!0,message:"name不能为空",trigger:"blur"},{validator:k["a"],trigger:"blur"},{validator:k["b"],trigger:"blur"},{validator:k["c"],trigger:"blur"}]}}},methods:{pullNamespace:function(){var e=this;Object(x["e"])().then((function(t){if(1===t.data.code){e.namespaceList=t.data.result;var n=e.$store.getters.getNamespace;e.namespaceList.forEach((function(t){t["name"]===n&&e.$store.commit("updateNamespaceInfo",t)}))}})).catch((function(e){console.log(e)}))},updateNamespace:function(e,t){var n=e.name;console.log(n,t),this.$store.commit("updateNamespace",n),this.$store.commit("updateNamespaceInfo",e),location.reload()},addNamespace:function(){var e=this;this.$refs["namespaceForm"].validate((function(t){t&&(e.dialogFormVisible222=!1,e.namespaceForm.is_active=!0,Object(x["j"])(e.namespaceForm).then((function(e){console.log(e),location.reload()})))}))},deleteOneNamespace:function(e,t){var n=this,a=t[e].id;Object(x["b"])(a).then((function(a){1===a.data.code?(n.$message.success("删除一行数据成功..."),t.splice(e,1)):n.$message.error("删除数据失败...")})).catch((function(e){console.log(e)}))},activeNamespace:function(){this.namespaceList.forEach((function(e){Object(x["n"])(e.id,e).then((function(e){console.log(e.data.result)})).catch((function(e){console.log(e)}))})),this.$message.success("数据更新成功...")}},computed:{getStoreNamespace:function(){return this.$store.getters.getNamespace}},created:function(){this.pullNamespace()}},S=w,N=(n("2b4c"),Object(l["a"])(S,y,_,!1,null,"6e676bf8",null)),j=N.exports,C={name:"app",components:{NavHeader:d,NavFooter:v,NavAside:j},created:function(){var e=this;sessionStorage.getItem("store")&&this.$store.replaceState(Object.assign({},this.$store.state,JSON.parse(sessionStorage.getItem("store")))),window.addEventListener("beforeunload",(function(){sessionStorage.setItem("store",JSON.stringify(e.$store.state))}))}},D=C,$=(n("034f"),Object(l["a"])(D,r,o,!1,null,null,null)),O=$.exports,A=(n("3ca3"),n("ddb0"),n("8c4f")),F=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("el-row",{staticClass:"shadow",staticStyle:{"margin-top":"40px"}},[n("el-col",{attrs:{offset:7,span:10}},[n("Domain")],1)],1),n("el-row",[n("el-col",{staticStyle:{"margin-top":"150px"},attrs:{offset:7,span:10}},[n("el-button",{attrs:{round:""},on:{click:e.simulation}},[e._v("模拟 AlertManager 推送一条报警信息")])],1)],1)],1)},L=[],P=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-card",{staticClass:"box-card",staticStyle:{width:"100%"},attrs:{shadow:"always",id:"yinying"}},[n("div",{attrs:{slot:"header"},slot:"header"},[n("span",{staticStyle:{"padding-left":"80px"}},[e._v("消息推送地址")]),n("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(t){return e.copyCode()}}},[e._v("一键复制")])],1),n("div",[n("pre",{attrs:{id:"DomainUrlStyle"}},[n("code",{attrs:{id:"DomainUrlContent"}},[e._v(e._s(e.myDomain)+e._s(e.showNamespace))])])])])},E=[],M=(n("ac1f"),n("1276"),{name:"cDomain",data:function(){return{myDomain:""}},computed:{getStoreNamespace:function(){return this.$store.getters.getNamespace},showNamespace:function(){return"default"===this.getStoreNamespace?"message":this.getStoreNamespace}},methods:{getDomain:function(){var e=window.location.href,t=e.split("#/")[0];console.log(t),this.myDomain=t+"go/"},copyCode:function(){var e=document.getElementById("DomainUrlContent");window.getSelection().selectAllChildren(e),document.execCommand("Copy"),this.$message("复制成功...")}},created:function(){this.getDomain()}}),T=M,I=(n("da68"),Object(l["a"])(T,P,E,!1,null,"b2f5cb96",null)),G=I.exports,U={name:"ViewHome",components:{Domain:G},methods:{simulation:function(){var e=this,t=n("9b9d");Object(x["i"])(this.$store.getters.getNamespace,t).then((function(t){console.log(t),e.$message.success("模拟消息推送成功...")})).catch((function(e){console.log(e)}))}},created:function(){this.$store.commit("updateStepsActive",0)}},V=U,q=Object(l["a"])(V,F,L,!1,null,"54d45d7b",null),B=q.exports;a["default"].use(A["a"]);var H=[{path:"/",name:"1-1",component:B},{path:"/data_format",name:"DataFormat",component:function(){return n.e("chunk-1f1bd5cd").then(n.bind(null,"7d3c"))}},{path:"/request_data",name:"2-1",component:function(){return n.e("chunk-0b923e88").then(n.bind(null,"f744"))}},{path:"/data_client",name:"3-1",component:function(){return n.e("chunk-94ca40d8").then(n.bind(null,"37fb"))}}],R=new A["a"]({routes:H}),Z=R,z=n("2f62");a["default"].use(z["a"]);var J=new z["a"].Store({state:{StepsActive:0,DrawerStatus:!1,Namespace:"default",NamespaceInfo:{}},getters:{getNamespace:function(e){return e.Namespace},getNamespaceInfo:function(e){return e.NamespaceInfo}},mutations:{updateStepsActive:function(e,t){e.StepsActive=t},updateDrawerStatus:function(e,t){e.DrawerStatus=t},updateNamespace:function(e,t){e.Namespace=t},updateNamespaceInfo:function(e,t){e.NamespaceInfo=t}},actions:{},modules:{}}),K=n("5c96"),Y=n.n(K);n("0fae");a["default"].use(Y.a),a["default"].config.productionTip=!1,new a["default"]({router:Z,store:J,render:function(e){return e(O)}}).$mount("#app"),a["default"].prototype.dateFormat=function(e){console.log(e);var t=e;if(null!=t){var n=new Date(t),a=n.getFullYear(),r=n.getMonth()+1,o=n.getDate(),i=n.getHours(),s=n.getMinutes(),c=n.getSeconds();return a+"-"+(r<10?"0":"")+r+"-"+(o<10?"0":"")+o+(i<10?"0":"")+" "+i+":"+(s<10?"0":"")+s+":"+(c<10?"0":"")+c}}},"61f7":function(e,t,n){"use strict";n.d(t,"b",(function(){return a})),n.d(t,"a",(function(){return r})),n.d(t,"c",(function(){return o}));n("ac1f"),n("00b4"),n("d3b7"),n("25f0");function a(e,t,n){var a=/^[0-9]+$/,r=a.test(t.toString());r&&n("不能以数字开头"),n()}function r(e,t,n){t.length<3&&n("长度不能小于3个字符"),n()}function o(e,t,n){var a=/^[A-Za-z0-9]+$/,r=a.test(t.toString());r||n("只能是字母或数字，不能包含特殊符号"),n()}},"6cac":function(e,t,n){"use strict";n.d(t,"i",(function(){return r})),n.d(t,"o",(function(){return o})),n.d(t,"p",(function(){return i})),n.d(t,"l",(function(){return s})),n.d(t,"g",(function(){return c})),n.d(t,"k",(function(){return u})),n.d(t,"c",(function(){return l})),n.d(t,"h",(function(){return p})),n.d(t,"d",(function(){return d})),n.d(t,"m",(function(){return m})),n.d(t,"a",(function(){return f})),n.d(t,"e",(function(){return g})),n.d(t,"j",(function(){return h})),n.d(t,"f",(function(){return b})),n.d(t,"b",(function(){return v})),n.d(t,"n",(function(){return y}));var a=n("be3b"),r=function(e,t){return a["a"].Post("/go/"+e,t)},o=function(e,t){return a["a"].Get("/api/v1/"+e+"/json",t)},i=function(e,t){return a["a"].Get("/api/v1/"+e+"/vars",t)},s=function(e,t){return a["a"].Post("/api/v1/"+e+"/vars",t)},c=function(e,t){return a["a"].Get("/api/v1/"+e+"/template",t)},u=function(e,t){return a["a"].Post("/api/v1/"+e+"/template",t)},l=function(e,t){return a["a"].Get("/api/v1/"+e+"/client",t)},p=function(e,t){return a["a"].Post("/api/v1/"+e+"/client",t)},d=function(e,t,n){return a["a"].Get("/api/v1/"+e+"/client/"+t,n)},m=function(e,t,n){return a["a"].Put("/api/v1/"+e+"/client/"+t,n)},f=function(e,t,n){return a["a"].Delete("/api/v1/"+e+"/client/"+t,n)},g=function(e){return a["a"].Get("/api/v1/namespace",e)},h=function(e){return a["a"].Post("/api/v1/namespace",e)},b=function(e,t){return a["a"].Get("/api/v1/namespace/"+e,t)},v=function(e){return a["a"].Delete("/api/v1/namespace/"+e)},y=function(e,t){return a["a"].Put("/api/v1/namespace/"+e,t)}},"84b6":function(e,t,n){},"85ec":function(e,t,n){},"9b19":function(e,t,n){e.exports=n.p+"static/img/logo.45c29d20.svg"},"9b9d":function(e){e.exports=JSON.parse('{"receiver":"dingding2","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","hostname":"sh2-fat-kubectl-1","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","ping":"172.20.5.10","severity":"严重","yunwei":"test"},"annotations":{"description":"服务器 (172.20.5.10) 存活探测失败,可能已经宕机,请尽快查看!","summary":"存活探测失败"},"startsAt":"2021-10-20T03:04:22.615Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://sh3-tools-prometheus-1:9090/graph?g0.expr=probe_success%7Bconsul_type%3D%22blackbox-icmp%22%7D+%21%3D+1&g0.tab=1","fingerprint":"5582f114464c5dc3"},{"status":"firing","labels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","hostname":"sh2-fat-kubectl-2","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","ping":"172.20.5.10","severity":"严重","yunwei":"test"},"annotations":{"description":"服务器 (172.20.5.10) 存活探测失败,可能已经宕机,请尽快查看!","summary":"存活探测失败"},"startsAt":"2021-10-20T03:04:22.615Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://sh3-tools-prometheus-1:9090/graph?g0.expr=probe_success%7Bconsul_type%3D%22blackbox-icmp%22%7D+%21%3D+1&g0.tab=1","fingerprint":"5582f114464c5dc3"},{"status":"firing","labels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","hostname":"sh2-fat-kubectl-3","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","ping":"192.168.35.14","severity":"严重","yunwei":"test"},"annotations":{"description":"服务器 (192.168.35.14) 存活探测失败,可能已经宕机,请尽快查看!","summary":"存活探测失败"},"startsAt":"2021-10-20T03:04:22.615Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://sh3-tools-prometheus-1:9090/graph?g0.expr=probe_success%7Bconsul_type%3D%22blackbox-icmp%22%7D+%21%3D+1&g0.tab=1","fingerprint":"e1755cd18e72b740"}],"groupLabels":{"alertname":"服务器存活探测"},"commonLabels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","severity":"严重","yunwei":"test"},"commonAnnotations":{"summary":"存活探测失败"},"externalURL":"http://sh3-tools-prometheus-1:9093","version":"4","groupKey":"{}:{alertname=\'服务器存活探测\'}","truncatedAlerts":0}')},b52f:function(e,t,n){},be3b:function(e,t,n){"use strict";n("d3b7");var a=n("bc3a"),r=n.n(a),o=r.a.create({baseURL:"",timeout:6e4});o.interceptors.request.use((function(e){return e.headers={"Content-Type":"application/json"},e}),(function(e){return Promise.reject(e)})),o.interceptors.response.use((function(e){return e}),(function(e){return Promise.reject(e)})),t["a"]={Get:function(e,t,n){return o({method:"get",url:e,headers:n,params:t})},Delete:function(e,t,n){return o({method:"delete",url:e,headers:n,params:t})},Post:function(e,t,n){return o({method:"post",url:e,headers:n,data:t})},Put:function(e,t,n){return o({method:"put",url:e,headers:n,data:t})}}},da68:function(e,t,n){"use strict";n("b52f")},f4a2:function(e,t,n){},f9ab:function(e,t,n){"use strict";n("84b6")}});
//# sourceMappingURL=app.a681db8d.js.map