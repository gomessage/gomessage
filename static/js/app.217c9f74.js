(function(e){function t(t){for(var o,r,s=t[0],c=t[1],u=t[2],l=0,d=[];l<s.length;l++)r=s[l],Object.prototype.hasOwnProperty.call(a,r)&&a[r]&&d.push(a[r][0]),a[r]=0;for(o in c)Object.prototype.hasOwnProperty.call(c,o)&&(e[o]=c[o]);p&&p(t);while(d.length)d.shift()();return i.push.apply(i,u||[]),n()}function n(){for(var e,t=0;t<i.length;t++){for(var n=i[t],o=!0,r=1;r<n.length;r++){var s=n[r];0!==a[s]&&(o=!1)}o&&(i.splice(t--,1),e=c(c.s=n[0]))}return e}var o={},r={app:0},a={app:0},i=[];function s(e){return c.p+"static/js/"+({}[e]||e)+"."+{"chunk-137cbf81":"5f8dc5da","chunk-2560bf0a":"55a4db06","chunk-613b4440":"40bd4fe6","chunk-6a350716":"b84fc1ad"}[e]+".js"}function c(t){if(o[t])return o[t].exports;var n=o[t]={i:t,l:!1,exports:{}};return e[t].call(n.exports,n,n.exports,c),n.l=!0,n.exports}c.e=function(e){var t=[],n={"chunk-137cbf81":1,"chunk-2560bf0a":1,"chunk-613b4440":1,"chunk-6a350716":1};r[e]?t.push(r[e]):0!==r[e]&&n[e]&&t.push(r[e]=new Promise((function(t,n){for(var o="static/css/"+({}[e]||e)+"."+{"chunk-137cbf81":"90d7aaf3","chunk-2560bf0a":"2623e17a","chunk-613b4440":"b261c614","chunk-6a350716":"beebf1d2"}[e]+".css",a=c.p+o,i=document.getElementsByTagName("link"),s=0;s<i.length;s++){var u=i[s],l=u.getAttribute("data-href")||u.getAttribute("href");if("stylesheet"===u.rel&&(l===o||l===a))return t()}var d=document.getElementsByTagName("style");for(s=0;s<d.length;s++){u=d[s],l=u.getAttribute("data-href");if(l===o||l===a)return t()}var p=document.createElement("link");p.rel="stylesheet",p.type="text/css",p.onload=t,p.onerror=function(t){var o=t&&t.target&&t.target.src||a,i=new Error("Loading CSS chunk "+e+" failed.\n("+o+")");i.code="CSS_CHUNK_LOAD_FAILED",i.request=o,delete r[e],p.parentNode.removeChild(p),n(i)},p.href=a;var f=document.getElementsByTagName("head")[0];f.appendChild(p)})).then((function(){r[e]=0})));var o=a[e];if(0!==o)if(o)t.push(o[2]);else{var i=new Promise((function(t,n){o=a[e]=[t,n]}));t.push(o[2]=i);var u,l=document.createElement("script");l.charset="utf-8",l.timeout=120,c.nc&&l.setAttribute("nonce",c.nc),l.src=s(e);var d=new Error;u=function(t){l.onerror=l.onload=null,clearTimeout(p);var n=a[e];if(0!==n){if(n){var o=t&&("load"===t.type?"missing":t.type),r=t&&t.target&&t.target.src;d.message="Loading chunk "+e+" failed.\n("+o+": "+r+")",d.name="ChunkLoadError",d.type=o,d.request=r,n[1](d)}a[e]=void 0}};var p=setTimeout((function(){u({type:"timeout",target:l})}),12e4);l.onerror=l.onload=u,document.head.appendChild(l)}return Promise.all(t)},c.m=e,c.c=o,c.d=function(e,t,n){c.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},c.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},c.t=function(e,t){if(1&t&&(e=c(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(c.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var o in e)c.d(n,o,function(t){return e[t]}.bind(null,o));return n},c.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return c.d(t,"a",t),t},c.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},c.p="",c.oe=function(e){throw console.error(e),e};var u=window["webpackJsonp"]=window["webpackJsonp"]||[],l=u.push.bind(u);u.push=t,u=u.slice();for(var d=0;d<u.length;d++)t(u[d]);var p=l;i.push([0,"chunk-vendors"]),n()})({0:function(e,t,n){e.exports=n("56d7")},"034f":function(e,t,n){"use strict";n("85ec")},"397d":function(e){e.exports=JSON.parse('{"receiver":"dingding2","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","hostname":"sh2-fat-kubectl-1","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","ping":"172.20.5.10","severity":"严重","yunwei":"test"},"annotations":{"description":"服务器 (172.20.5.10) 存活探测失败,可能已经宕机,请尽快查看!","summary":"存活探测失败"},"startsAt":"2021-10-20T03:04:22.615Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://sh3-tools-prometheus-1:9090/graph?g0.expr=probe_success%7Bconsul_type%3D%22blackbox-icmp%22%7D+%21%3D+1&g0.tab=1","fingerprint":"5582f114464c5dc3"},{"status":"firing","labels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","hostname":"sh3-tools-EMR-Core-1","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","ping":"192.168.35.14","severity":"严重","yunwei":"test"},"annotations":{"description":"服务器 (192.168.35.14) 存活探测失败,可能已经宕机,请尽快查看!","summary":"存活探测失败"},"startsAt":"2021-10-20T03:04:22.615Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://sh3-tools-prometheus-1:9090/graph?g0.expr=probe_success%7Bconsul_type%3D%22blackbox-icmp%22%7D+%21%3D+1&g0.tab=1","fingerprint":"e1755cd18e72b740"}],"groupLabels":{"alertname":"服务器存活探测"},"commonLabels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","severity":"严重","yunwei":"test"},"commonAnnotations":{"summary":"存活探测失败"},"externalURL":"http://sh3-tools-prometheus-1:9093","version":"4","groupKey":"{}:{alertname=\'服务器存活探测\'}","truncatedAlerts":0}')},4656:function(e,t,n){"use strict";n("a5e7")},"56d7":function(e,t,n){"use strict";n.r(t);n("e260"),n("e6cf"),n("cca6"),n("a79d");var o=n("2b0e"),r=(n("d3b7"),n("bc3a")),a=n.n(r),i={timeout:6e4},s=a.a.create(i);s.interceptors.request.use((function(e){return e.headers={"Content-Type":"application/json"},e}),(function(e){return Promise.reject(e)})),s.interceptors.response.use((function(e){return e}),(function(e){return Promise.reject(e)})),Plugin.install=function(e,t){console.log(t),e.axios=s,window.axios=s,Object.defineProperties(e.prototype,{axios:{get:function(){return s}},$axios:{get:function(){return s}}})},o["default"].use(Plugin);Plugin;var c=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",{attrs:{id:"app"}},[n("el-container",{attrs:{id:"MyContainer-Container"}},[n("el-header",{attrs:{id:"MyContainer-Header"}},[n("NavHeader")],1),n("el-main",{attrs:{id:"MyContainer-Main"}},[n("router-view")],1),n("el-footer",{attrs:{id:"MyContainer-Footer"}},[n("NavFooter")],1)],1)],1)},u=[],l=function(){var e=this,t=e.$createElement,o=e._self._c||t;return o("el-menu",{staticStyle:{height:"100%"},attrs:{"default-active":e.$route.name,router:"",mode:"horizontal","background-color":"#303133","text-color":"#fff","active-text-color":"#ffd04b","collapse-transition":"true"},on:{select:e.handleSelect}},[o("el-menu-item",[o("router-link",{attrs:{to:"/"}},[o("img",{attrs:{src:n("9b19"),height:"90%"}}),e._v("    "),o("span",{staticStyle:{"font-size":"15px"}},[e._v("GoMessage · 消息转发器")])])],1),o("el-row",{attrs:{type:"flex",justify:"end"}},[o("router-link",{attrs:{to:"/"}},[o("el-menu-item",{attrs:{index:"1-1"}},[o("i",{staticClass:"el-icon-position"}),e._v("消息入口")])],1),o("router-link",{attrs:{to:"/request_data"}},[o("el-menu-item",{attrs:{index:"2-1"}},[o("i",{staticClass:"el-icon-set-up"}),e._v("模板编写")])],1),o("router-link",{attrs:{to:"/data_client"}},[o("el-menu-item",{attrs:{index:"3-1"}},[o("i",{staticClass:"el-icon-chat-dot-square"}),e._v("接收客户端")])],1),o("el-menu-item",[o("el-link",{attrs:{href:"https://github.com/gomessage/gomessage#gomessage",underline:!1,type:"primary",target:"_blank"}},[o("i",{staticClass:"el-icon-document"}),e._v("文档")])],1)],1)],1)},d=[],p={name:"NavHeader",data:function(){return{}},components:{},methods:{handleSelect:function(e,t){console.log(e,t)}}},f=p,m=(n("627e"),n("2877")),h=Object(m["a"])(f,l,d,!1,null,"e6756bee",null),g=h.exports,b=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-menu",{staticClass:"el-menu-demo",attrs:{id:"NavFooter-Footer",mode:"horizontal","background-color":"#303133","text-color":"#fff","active-text-color":"#ffd04b","collapse-transition":"true"},on:{select:e.handleSelect}},[n("el-row",{attrs:{type:"flex",justify:"center",id:"my-row"}},[n("el-menu-item",{attrs:{id:"NavFooter-menu-item",disabled:""}},[n("el-link",{attrs:{href:"https://gitee.com/ergou-open-source",type:"primary"}},[e._v("@ 2021 EGOS")]),e._v("   "),n("el-link",{attrs:{href:"https://opensource.org/licenses/MIT",type:"primary"}},[e._v("MIT License")]),e._v("   Give feedback and Suggestions to the author： "),n("el-link",{attrs:{href:"mailto:tay3223@qq.com",type:"primary"}},[e._v("Email")])],1)],1)],1)},v=[],y={name:"NavFooter",data:function(){return{activeIndex:"1",activeIndex2:"1"}},methods:{handleSelect:function(e,t){console.log(e,t)}}},_=y,x=(n("4656"),Object(m["a"])(_,b,v,!1,null,"298d89d4",null)),w=x.exports,k={name:"app",components:{NavHeader:g,NavFooter:w}},S=k,C=(n("034f"),Object(m["a"])(S,c,u,!1,null,null,null)),j=C.exports,D=(n("3ca3"),n("ddb0"),n("8c4f")),A=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("div",[n("el-row",{staticClass:"shadow",staticStyle:{"margin-top":"20px"}},[n("el-col",{attrs:{span:"10",offset:"7"}},[n("Domain")],1)],1),n("el-row",[n("el-col",{staticStyle:{"margin-top":"150px"},attrs:{span:"10",offset:"7"}},[n("el-button",{attrs:{round:""},on:{click:e.simulation}},[e._v("模拟 AlertManager 的报警推送")])],1)],1)],1)},O=[],E=function(){var e=this,t=e.$createElement,n=e._self._c||t;return n("el-card",{staticClass:"box-card",staticStyle:{width:"100%"},attrs:{shadow:"always",id:"yinying"}},[n("div",{attrs:{slot:"header"},slot:"header"},[n("span",{staticStyle:{"padding-left":"80px"}},[e._v("消息推送地址")]),n("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(t){return e.copyCode()}}},[e._v("一键复制")])],1),n("div",[n("pre",{attrs:{id:"DomainUrlStyle"}},[n("code",{attrs:{id:"DomainUrlContent"}},[e._v(e._s(e.myDomain))])])])])},M=[],P=(n("ac1f"),n("1276"),{name:"cDomain",data:function(){return{myDomain:""}},methods:{getDomain:function(){var e=window.location.href,t=e.split("#/")[0];console.log(t),this.myDomain=t+"go/message"},copyCode:function(){var e=document.getElementById("DomainUrlContent");window.getSelection().selectAllChildren(e),document.execCommand("Copy"),this.$message("复制成功...")}},created:function(){this.getDomain()}}),T=P,N=(n("c6ae"),Object(m["a"])(T,E,M,!1,null,"6ec14945",null)),F=N.exports,L={name:"ViewHome",components:{Domain:F},methods:{simulation:function(){var e=this,t=n("397d");this.axios.post("/go/message",t).then((function(t){console.log(t),e.$message.success("模拟消息推送成功...")})).catch((function(e){console.log(e)}))}},created:function(){this.$store.commit("updateStepsActive",0)}},$=L,q=Object(m["a"])($,A,O,!1,null,"59b8d427",null),H=q.exports;o["default"].use(D["a"]);var U=[{path:"/",name:"1-1",component:H},{path:"/data_format",name:"DataFormat",component:function(){return n.e("chunk-137cbf81").then(n.bind(null,"f1f3"))}},{path:"/data_template",name:"MessageTemplate",component:function(){return n.e("chunk-2560bf0a").then(n.bind(null,"f4f5"))}},{path:"/request_data",name:"2-1",component:function(){return n.e("chunk-6a350716").then(n.bind(null,"b01b"))}},{path:"/data_client",name:"3-1",component:function(){return n.e("chunk-613b4440").then(n.bind(null,"37fb"))}}],B=new D["a"]({routes:U}),I=B,R=n("2f62");o["default"].use(R["a"]);var Z=new R["a"].Store({state:{StepsActive:0,DrawerStatus:!1},mutations:{updateStepsActive:function(e,t){e.StepsActive=t},updateDrawerStatus:function(e,t){e.DrawerStatus=t}},actions:{},modules:{}}),z=n("5c96"),G=n.n(z);n("0fae");o["default"].use(G.a),o["default"].config.productionTip=!1,new o["default"]({router:I,store:Z,render:function(e){return e(j)}}).$mount("#app"),o["default"].prototype.dateFormat=function(e){console.log(e);var t=e;if(null!=t){var n=new Date(t),o=n.getFullYear(),r=n.getMonth()+1,a=n.getDate(),i=n.getHours(),s=n.getMinutes(),c=n.getSeconds();return o+"-"+(r<10?"0":"")+r+"-"+(a<10?"0":"")+a+(i<10?"0":"")+" "+i+":"+(s<10?"0":"")+s+":"+(c<10?"0":"")+c}}},"627e":function(e,t,n){"use strict";n("b05d")},"85ec":function(e,t,n){},"9b19":function(e,t,n){e.exports=n.p+"static/img/logo.3a5f80f9.svg"},a5e7:function(e,t,n){},b05d:function(e,t,n){},c6ae:function(e,t,n){"use strict";n("e40c")},e40c:function(e,t,n){}});
//# sourceMappingURL=app.217c9f74.js.map