(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-ef80ac28"],{"20db":function(t,e,n){},"3ea0":function(t,e,n){"use strict";n("20db")},4886:function(t,e,n){"use strict";n.r(e);var a=function(){var t=this,e=t._self._c;return e("div",[e("el-row",{staticClass:"shadow",staticStyle:{"margin-top":"40px"}},[e("el-col",{attrs:{offset:7,span:10}},[e("Domain")],1)],1),e("el-row",[e("el-col",{staticStyle:{"margin-top":"150px"},attrs:{offset:7,span:10}},[e("el-button",{attrs:{round:""},on:{click:t.simulation}},[t._v("模拟 AlertManager 推送一条报警信息")])],1)],1)],1)},o=[],s=function(){var t=this,e=t._self._c;return e("el-card",{staticClass:"box-card",staticStyle:{width:"100%"},attrs:{shadow:"always",id:"yinying"}},[e("div",{attrs:{slot:"header"},slot:"header"},[e("span",{staticStyle:{"padding-left":"80px"}},[t._v("消息推送地址")]),e("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:function(e){return t.copyCode()}}},[t._v("一键复制")])],1),e("div",[e("pre",{attrs:{id:"DomainUrlStyle"}},[e("code",{attrs:{id:"DomainUrlContent"}},[t._v(t._s(t.myDomain)+t._s(t.showNamespace))])])])])},i=[],r={name:"cDomain",data(){return{myDomain:""}},computed:{getStoreNamespace:function(){return this.$store.getters.getNamespace},showNamespace:function(){return"default"===this.getStoreNamespace?"message":this.getStoreNamespace}},methods:{getDomain:function(){let t=window.location.href,e=t.split("#/")[0];console.log(e),this.myDomain=e+"go/"},copyCode:function(){const t=document.getElementById("DomainUrlContent");window.getSelection().selectAllChildren(t),document.execCommand("Copy"),this.$message("复制成功...")}},created(){this.getDomain()}},c=r,u=(n("3ea0"),n("2877")),p=Object(u["a"])(c,s,i,!1,null,"b2f5cb96",null),l=p.exports,m=n("6cac"),d={name:"ViewHome",components:{Domain:l},methods:{simulation:function(){let t=n("cf4f");Object(m["k"])(this.$store.getters.getNamespace,t).then(t=>{console.log(t),this.$message.success("模拟消息推送成功...")}).catch(t=>{console.log(t)})}},created(){this.$store.commit("updateStepsActive",0)}},f=d,g=Object(u["a"])(f,a,o,!1,null,"8c966164",null);e["default"]=g.exports},"6cac":function(t,e,n){"use strict";n.d(e,"k",(function(){return o})),n.d(e,"s",(function(){return s})),n.d(e,"t",(function(){return i})),n.d(e,"n",(function(){return r})),n.d(e,"r",(function(){return c})),n.d(e,"g",(function(){return u})),n.d(e,"m",(function(){return p})),n.d(e,"c",(function(){return l})),n.d(e,"j",(function(){return m})),n.d(e,"d",(function(){return d})),n.d(e,"p",(function(){return f})),n.d(e,"o",(function(){return g})),n.d(e,"a",(function(){return h})),n.d(e,"e",(function(){return b})),n.d(e,"l",(function(){return v})),n.d(e,"f",(function(){return y})),n.d(e,"b",(function(){return w})),n.d(e,"q",(function(){return D})),n.d(e,"h",(function(){return x})),n.d(e,"i",(function(){return _}));var a=n("be3b");const o=(t,e)=>a["a"].Post("/go/"+t,e),s=(t,e)=>a["a"].Get("/api/v1/"+t+"/json",e),i=(t,e)=>a["a"].Get("/api/v1/"+t+"/vars",e),r=(t,e)=>a["a"].Post("/api/v1/"+t+"/vars",e),c=(t,e)=>a["a"].Get("/api/v1/"+t+"/flattening",e),u=(t,e)=>a["a"].Get("/api/v1/"+t+"/template",e),p=(t,e)=>a["a"].Post("/api/v1/"+t+"/template",e),l=(t,e)=>a["a"].Get("/api/v1/"+t+"/client",e),m=(t,e)=>a["a"].Post("/api/v1/"+t+"/client",e),d=(t,e,n)=>a["a"].Get("/api/v1/"+t+"/client/"+e,n),f=(t,e,n)=>a["a"].Put("/api/v1/"+t+"/client/"+e,n),g=(t,e,n)=>a["a"].Put("/api/v1/"+t+"/client-info/"+e,n),h=(t,e,n)=>a["a"].Delete("/api/v1/"+t+"/client/"+e,n),b=t=>a["a"].Get("/api/v1/namespace",t),v=t=>a["a"].Post("/api/v1/namespace",t),y=(t,e)=>a["a"].Get("/api/v1/namespace/"+t,e),w=t=>a["a"].Delete("/api/v1/namespace/"+t),D=(t,e)=>a["a"].Put("/api/v1/namespace/"+t,e),x=t=>a["a"].Post("/auth/login",t),_=t=>a["a"].Post("/auth/logout",t)},cf4f:function(t){t.exports=JSON.parse('{"receiver":"dingding2","status":"firing","alerts":[{"status":"firing","labels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","hostname":"sh2-fat-kubectl-1","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","ping":"172.20.5.10","severity":"严重","yunwei":"test"},"annotations":{"description":"服务器 (172.20.5.10) 存活探测失败,可能已经宕机,请尽快查看!","summary":"存活探测失败"},"startsAt":"2021-10-20T03:04:22.615Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://sh3-tools-prometheus-1:9090/graph?g0.expr=probe_success%7Bconsul_type%3D%22blackbox-icmp%22%7D+%21%3D+1&g0.tab=1","fingerprint":"5582f114464c5dc3"},{"status":"firing","labels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","hostname":"sh2-fat-kubectl-2","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","ping":"172.20.5.10","severity":"严重","yunwei":"test"},"annotations":{"description":"服务器 (172.20.5.10) 存活探测失败,可能已经宕机,请尽快查看!","summary":"存活探测失败"},"startsAt":"2021-10-20T03:04:22.615Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://sh3-tools-prometheus-1:9090/graph?g0.expr=probe_success%7Bconsul_type%3D%22blackbox-icmp%22%7D+%21%3D+1&g0.tab=1","fingerprint":"5582f114464c5dc3"},{"status":"firing","labels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","hostname":"sh2-fat-kubectl-3","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","ping":"192.168.35.14","severity":"严重","yunwei":"test"},"annotations":{"description":"服务器 (192.168.35.14) 存活探测失败,可能已经宕机,请尽快查看!","summary":"存活探测失败"},"startsAt":"2021-10-20T03:04:22.615Z","endsAt":"0001-01-01T00:00:00Z","generatorURL":"http://sh3-tools-prometheus-1:9090/graph?g0.expr=probe_success%7Bconsul_type%3D%22blackbox-icmp%22%7D+%21%3D+1&g0.tab=1","fingerprint":"e1755cd18e72b740"}],"groupLabels":{"alertname":"服务器存活探测"},"commonLabels":{"alertname":"服务器存活探测","consul_type":"blackbox-icmp","group":"yunwei-ping","instance":"192.168.35.18:9115","job":"Autodiscover-icmp","severity":"严重","yunwei":"test"},"commonAnnotations":{"summary":"存活探测失败"},"externalURL":"http://sh3-tools-prometheus-1:9093","version":"4","groupKey":"{}:{alertname=\'服务器存活探测\'}","truncatedAlerts":0}')}}]);
//# sourceMappingURL=chunk-ef80ac28.8033b6ec.js.map