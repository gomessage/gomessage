(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d22497a"],{e191:function(t,a,e){"use strict";e.r(a);var o=function(){var t=this,a=t._self._c;return a("el-row",{staticStyle:{top:"0",left:"0",right:"0",bottom:"0",position:"absolute",display:"flex"}},[a("el-col",{attrs:{span:12}},[a("el-card",{staticStyle:{margin:"20px",padding:"10px","text-align":"left"}},[a("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[a("span",[t._v("发送定时消息（Crontab类型的时间规则定义）")])]),a("el-form",{ref:t.formData2,attrs:{model:t.formData2,"label-width":"90px"}},[a("el-form-item",{attrs:{label:"名称："}},[a("el-input",{attrs:{placeholder:"下班打卡提醒"},model:{value:t.formData2.crontab_name,callback:function(a){t.$set(t.formData2,"crontab_name",a)},expression:"formData2.crontab_name"}})],1),a("el-form-item",{attrs:{label:"时间规则："}},[a("el-input",{attrs:{placeholder:"*/5  *  *  *  * （分、时、日、月、周）"},model:{value:t.formData2.crontab_rule,callback:function(a){t.$set(t.formData2,"crontab_rule",a)},expression:"formData2.crontab_rule"}})],1),a("el-form-item",{attrs:{label:"发送通道："}},[a("el-select",{staticStyle:{width:"70%"},attrs:{multiple:"",placeholder:"请选择"},on:{"visible-change":a=>{a&&t.listNamespace()}},model:{value:t.formData2.crontab_namespace,callback:function(a){t.$set(t.formData2,"crontab_namespace",a)},expression:"formData2.crontab_namespace"}},t._l(t.namespaceOptions,(function(t){return a("el-option",{key:t.label,attrs:{label:t.label,value:t.value}})})),1)],1),a("el-form-item",{attrs:{label:"是否启用："}},[a("el-switch",{model:{value:t.formData2.crontab_activate,callback:function(a){t.$set(t.formData2,"crontab_activate",a)},expression:"formData2.crontab_activate"}})],1),a("el-form-item",{attrs:{label:"消息内容："}},[a("el-input",{attrs:{type:"textarea",autosize:{minRows:10,maxRows:200}},model:{value:t.formData2.crontab_content,callback:function(a){t.$set(t.formData2,"crontab_content",a)},expression:"formData2.crontab_content"}})],1),a("el-form-item",[a("el-button",{attrs:{type:"primary"},on:{click:t.onSubmit}},[t._v("创建或更新")]),a("el-button",[t._v("取消")])],1)],1)],1)],1),a("el-col",{attrs:{span:12}},[a("el-card",{staticStyle:{margin:"20px"}},[a("el-table",{staticStyle:{width:"100%"},attrs:{data:t.ListData2,border:"","header-cell-style":{background:"#2f2f35",color:"#fff"},"highlight-current-row":""},on:{"row-click":t.handleRowClick}},[a("el-table-column",{attrs:{prop:"crontab_name","show-overflow-tooltip":"",sortable:"",label:"名称"}}),a("el-table-column",{attrs:{prop:"crontab_rule",width:"110",label:"时间规则"}}),a("el-table-column",{attrs:{prop:"crontab_namespace",label:"发送通道","show-overflow-tooltip":"",formatter:t.formatNamespace}}),a("el-table-column",{attrs:{width:"110",prop:"crontab_activate",formatter:t.formatStatus,sortable:"",label:"是否启用"}}),a("el-table-column",{attrs:{width:"80",label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){return[a("el-button",{attrs:{type:"text",size:"small"},on:{click:function(a){return t.delete_one_data(e.row)}}},[t._v("删除")])]}}])})],1)],1)],1)],1)},l=[],r=e("6cac"),n={name:"ViewCrontab",data(){return{componentKey:0,ListData2:[{id:"",crontab_name:"提醒小伙伴下班打卡",crontab_rule:"*/5 * 2 * *",crontab_namespace:["default","ttpai","alarmdog"],crontab_activate:!1,crontab_content:"113123123123123123123113123123123123123123113123123123123123123113123123123123123123113123123123123123123113123123123123123123"}],formData2:{id:-1,crontab_name:"",crontab_rule:"",crontab_namespace:[],crontab_activate:!1,crontab_content:""},namespaceOptions:[{label:"default",value:"default"}]}},methods:{handleRowClick(t,a,e){console.log("点击的行数据:",t),console.log("点击的列:",a),console.log("点击事件:",e),Object(r["g"])(t.id).then(t=>{1===t.data.code&&(this.formData2=t.data.result)}).catch(t=>{console.log(t)})},onSubmit(){-1===this.formData2.id?(this.formData2.id=null,Object(r["n"])(this.formData2).then(t=>{1===t.data.code?this.refresh_self():alert(t.data.msg)}).catch(t=>{alert("数据添加失败，请检查数据格式是否正确！"),console.log(t)})):Object(r["u"])(this.formData2.id,this.formData2).then(t=>{1===t.data.code&&this.refresh_self()}).catch(t=>{console.log(t)})},formatStatus(t,a,e){return e?"启用":"未启用"},formatNamespace(t,a,e){return e.join(", ")},listCrontab:function(){Object(r["f"])().then(t=>{1===t.data.code&&(this.ListData2=t.data.result)}).catch(t=>{console.log(t)})},listNamespace:function(){Object(r["h"])().then(t=>{1===t.data.code&&(this.namespaceOptions=t.data.result.map(t=>({label:t.name,value:t.name})))}).catch(t=>{console.log(t)})},delete_one_data:function(t){confirm(`确定要删除名称为 '${t.crontab_name}' 的定时任务吗？`)&&Object(r["b"])(t.id).then(t=>{1===t.data.code?(this.refresh_self(),alert("删除成功！")):(this.refresh_self(),alert("删除失败："+t.data.message))}).catch(t=>{console.error("删除失败：",t),alert("删除失败，请检查网络或联系管理员。")})},refresh_self:function(){this.formData2={id:-1,crontab_name:"",crontab_rule:"",crontab_namespace:[],crontab_activate:!1,crontab_content:""},this.listCrontab()}},created(){this.listCrontab()}},c=n,s=e("2877"),i=Object(s["a"])(c,o,l,!1,null,"3b018904",null);a["default"]=i.exports}}]);
//# sourceMappingURL=chunk-2d22497a.ec6d63f2.js.map