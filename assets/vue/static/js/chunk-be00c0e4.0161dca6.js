(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-be00c0e4"],{"37fb":function(t,e,l){"use strict";l.r(e);var n=function(){var t=this,e=t.$createElement,l=t._self._c||e;return l("el-row",{staticStyle:{"margin-top":"20px","margin-bottom":"20px"}},[l("el-col",{attrs:{offset:4,span:16}},[l("CDrawer",{attrs:{getClientList:t.getClientList}}),l("CDrawerOneDataInfo",{attrs:{oneClientInfo:t.clientOneInfo,thisClose:t.thisClose,visibleStatus:t.visibleStatus}}),l("el-card",{staticStyle:{"box-shadow":"#ccc 0px 30px 30px"}},[l("div",{staticClass:"clearfix",attrs:{slot:"header"},slot:"header"},[l("span",{staticStyle:{"padding-left":"50px"}},[t._v("客户端 · 列表")]),l("el-button",{staticStyle:{float:"right",padding:"3px 0"},attrs:{type:"text"},on:{click:t.getClientList}},[t._v("刷新列表")])],1),l("el-table",{staticStyle:{width:"100%"},attrs:{border:!0,data:t.clientList,"header-cell-style":{background:"#2f2f35",color:"#fff"},stripe:!0}},[l("el-table-column",{attrs:{label:"客户端名称",prop:"client_name"}}),l("el-table-column",{attrs:{label:"客户端描述",prop:"client_description"}}),l("el-table-column",{attrs:{label:"客户端类型",prop:"client_annotation"}}),l("el-table-column",{attrs:{label:"是否激活",prop:"is_active"},scopedSlots:t._u([{key:"default",fn:function(e){return[l("el-checkbox",{on:{change:t.saveActiveStatus},model:{value:e.row.is_active,callback:function(l){t.$set(e.row,"is_active",l)},expression:"scope.row.is_active"}},[t._v("激活")])]}}])}),l("el-table-column",{attrs:{fixed:"right",label:"操作",width:"100"},scopedSlots:t._u([{key:"default",fn:function(e){return[l("el-button",{attrs:{size:"small",type:"text"},nativeOn:{click:function(l){return l.preventDefault(),t.getOneClient(e.$index,t.clientList)}}},[t._v("详情 ")]),l("el-button",{attrs:{size:"small",type:"text"},nativeOn:{click:function(l){return l.preventDefault(),t.deleteOneClient(e.$index,t.clientList)}}},[t._v(" 移除 ")])]}}])})],1),l("br"),l("br"),l("br"),l("el-button",{attrs:{round:"",type:"primary"},on:{click:t.addClient}},[t._v("添加新客户端")])],1)],1)],1)},i=[],o=(l("d3b7"),l("159b"),l("a434"),function(){var t=this,e=t.$createElement,l=t._self._c||e;return l("el-drawer",{attrs:{title:"添加客户端",visible:t.getDrawerStatus,"before-close":t.handleClose,"destroy-on-close":!0}},[l("el-tabs",{staticStyle:{margin:"20px"},attrs:{"tab-position":t.tabPosition}},[l("el-tab-pane",{attrs:{label:"钉钉"}},[l("clientDingtalk",{attrs:{getClientList:t.getClientList}})],1),l("el-tab-pane",{attrs:{label:"企业微信"}},[l("clientWechat",{attrs:{getClientList:t.getClientList}})],1),l("el-tab-pane",{attrs:{label:"飞书"}},[l("clientFeishu",{attrs:{getClientList:t.getClientList}})],1),l("el-tab-pane",{staticStyle:{"text-align":"left"},attrs:{label:"其它"}},[l("clientOther",{attrs:{getClientList:t.getClientList}})],1)],1)],1)}),c=[],a=function(){var t=this,e=t.$createElement,l=t._self._c||e;return l("el-form",{staticStyle:{"text-align":"left"},attrs:{"label-position":t.labelPosition,model:t.client,rules:t.clientRules,"label-width":"100px"}},[l("el-form-item",{attrs:{label:"客户端名称:",prop:"client_name"}},[l("el-input",{attrs:{placeholder:""},model:{value:t.client.client_name,callback:function(e){t.$set(t.client,"client_name",e)},expression:"client.client_name"}})],1),l("el-form-item",{attrs:{label:"客户端描述:"}},[l("el-input",{attrs:{placeholder:""},model:{value:t.client.client_description,callback:function(e){t.$set(t.client,"client_description",e)},expression:"client.client_description"}})],1),l("el-form-item",{attrs:{label:"客户端类型:"}},[l("el-input",{attrs:{disabled:""},model:{value:t.client.typeDescription,callback:function(e){t.$set(t.client,"typeDescription",e)},expression:"client.typeDescription"}})],1),l("el-divider",{attrs:{"content-position":"center"}},[t._v("机器人")]),l("el-form-item",{attrs:{label:"放行关键字:"}},[l("el-input",{attrs:{placeholder:"要和机器人界面设置的放行关键词保持一直"},model:{value:t.client.client_info.robot_keyword,callback:function(e){t.$set(t.client.client_info,"robot_keyword",e)},expression:"client.client_info.robot_keyword"}})],1),t._l(t.client.client_info.robot_url_list,(function(e,n){return l("div",{key:n},[l("el-form-item",{attrs:{label:"机器人URL:"}},[l("el-input",{staticStyle:{width:"85%"},attrs:{placeholder:"从钉钉上粘贴而来的机器人URL地址"},model:{value:e.url,callback:function(l){t.$set(e,"url",l)},expression:"list.url"}}),l("el-button",{staticStyle:{"margin-left":"10px"},attrs:{circle:"",icon:"el-icon-delete",size:"mini",type:"danger"},on:{click:function(e){return t.del(n)}}})],1)],1)})),l("div",{staticStyle:{"text-align":"center"}},[l("el-button",{attrs:{icon:"el-icon-plus",size:"mini",type:"text"},on:{click:t.add}},[t._v("再添加一个机器人")])],1),l("p",{attrs:{id:"textStype"}},[t._v(" 此处可以添加多个机器人，推送消息时会从中随机挑选一个URL来使用，可以避免单个机器人消息推送时（每分钟）的次数限制，避免重要报警信息被漏送的可能。")]),l("br"),l("br"),l("el-form-item",[l("el-button",{attrs:{type:"primary"},on:{click:t.onSubmit}},[t._v("立即创建")]),l("el-button",{on:{click:t.closeDrawer}},[t._v("取消")])],1)],2)},s=[],r=l("6cac"),u={name:"clientDingtalk",data:function(){return{labelPosition:"right",client:{client_name:"",client_description:"",client_type:"dingtalk",is_active:!1,client_info:{robot_keyword:"",robot_url_list:[{url:""}]},typeDescription:"钉钉·机器人"},clientRules:{client_name:[{required:!0,message:"name不能为空",trigger:"blur"}]}}},props:{getClientList:Function},methods:{onSubmit:function(){var t=this;Object(r["g"])(this.$store.getters.getNamespace,this.client).then((function(e){1===e.data.code?(t.$message.success("添加成功..."),t.getClientList(),t.$store.commit("updateDrawerStatus",!1)):(t.$message.error("添加失败..."),t.getClientList(),t.$store.commit("updateDrawerStatus",!1))})).catch((function(t){console.log(t)}))},add:function(){var t={url:""};this.client.client_info.robot_url_list.push(t);for(var e=0;e<this.client.client_info.robot_url_list.length;e++)console.log(this.client.client_info.robot_url_list[e])},del:function(t){this.client.client_info.robot_url_list.splice(t,1)},closeDrawer:function(){this.getClientList(),this.$store.commit("updateDrawerStatus",!1)}}},f=u,d=(l("5ed7"),l("2877")),p=Object(d["a"])(f,a,s,!1,null,"efd07848",null),_=p.exports,b=function(){var t=this,e=t.$createElement,l=t._self._c||e;return l("el-form",{staticStyle:{"text-align":"left"},attrs:{"label-position":t.labelPosition,"label-width":"100px",model:t.client,rules:t.wechatRules}},[l("el-form-item",{attrs:{label:"客户端名称:",prop:"client_name"}},[l("el-input",{attrs:{placeholder:""},model:{value:t.client.client_name,callback:function(e){t.$set(t.client,"client_name",e)},expression:"client.client_name"}})],1),l("el-form-item",{attrs:{label:"客户端描述:"}},[l("el-input",{attrs:{placeholder:""},model:{value:t.client.client_description,callback:function(e){t.$set(t.client,"client_description",e)},expression:"client.client_description"}})],1),l("el-form-item",{attrs:{label:"客户端类型:"}},[l("el-input",{attrs:{disabled:""},model:{value:t.client.typeDescription,callback:function(e){t.$set(t.client,"typeDescription",e)},expression:"client.typeDescription"}})],1),l("el-divider",{attrs:{"content-position":"center"}},[t._v("应用号")]),l("el-form-item",{attrs:{label:"企业ID:"}},[l("el-input",{attrs:{placeholder:"请联系企业微信管理员获取"},model:{value:t.client.client_info.corp_id,callback:function(e){t.$set(t.client.client_info,"corp_id",e)},expression:"client.client_info.corp_id"}})],1),l("el-form-item",{attrs:{label:"应用AgentId:"}},[l("el-input",{attrs:{placeholder:"请联系企业微信管理员获取"},model:{value:t.client.client_info.agent_id,callback:function(e){t.$set(t.client.client_info,"agent_id",e)},expression:"client.client_info.agent_id"}})],1),l("el-form-item",{attrs:{label:"应用Secret:"}},[l("el-input",{attrs:{placeholder:"请联系企业微信管理员获取","show-password":""},model:{value:t.client.client_info.secret,callback:function(e){t.$set(t.client.client_info,"secret",e)},expression:"client.client_info.secret"}})],1),l("el-form-item",{attrs:{label:"接收用户:"}},[l("el-input",{attrs:{type:"textarea",autosize:{minRows:4,maxRows:6},placeholder:"可以填写多个用户账号，用 | 分割开 （例如：aaa|bbb）"},model:{value:t.client.client_info.touser,callback:function(e){t.$set(t.client.client_info,"touser",e)},expression:"client.client_info.touser"}})],1),l("el-form-item",[l("el-button",{attrs:{type:"primary"},on:{click:t.onSubmit}},[t._v("立即创建")]),l("el-button",{on:{click:t.closeDrawer}},[t._v("取消")])],1)],1)},m=[],h={name:"clientWechat",data:function(){return{labelPosition:"right",client:{client_name:"",client_description:"",client_type:"wechat",is_active:!1,client_info:{corp_id:"",agent_id:"",secret:"",touser:""},typeDescription:"企业微信·应用号"},wechatRules:{client_name:[{required:!0,message:"name不能为空",trigger:"blur"}]}}},props:{getClientList:Function},methods:{onSubmit:function(){var t=this,e=this.client.client_info.secret;e.length<=15?this.$message.error("应用Secret的输入长度不符合要求..."):Object(r["g"])(this.client).then((function(e){e.data.result?(t.$message.success("添加成功..."),t.getClientList(),t.$store.commit("updateDrawerStatus",!1)):(t.$message.error("添加失败..."),t.getClientList(),t.$store.commit("updateDrawerStatus",!1))})).catch((function(t){console.log(t)}))},closeDrawer:function(){this.getClientList(),this.$store.commit("updateDrawerStatus",!1)}}},v=h,g=Object(d["a"])(v,b,m,!1,null,"385c7a78",null),y=g.exports,C=function(){var t=this,e=t.$createElement,l=t._self._c||e;return l("div",[l("br"),l("p",{staticClass:"authorStatement"},[t._v("GoMessage永久开源且免费，任何人都可以进行二次开发与封装...")]),l("br"),l("p",{staticClass:"authorStatement"},[t._v("平时只能用业余时间开发，更新进度可能比较慢，希望小伙伴们理解万岁...")]),l("br"),l("p",{staticClass:"authorStatement"},[t._v(" 如果您喜欢GoMessage，并且希望GoMessage能兼容对接到其它客户端，可直接联系我们进行反馈，我们会尽快满足小伙伴们的各种实际使用场景...")]),l("br"),l("p",{staticClass:"authorStatement"},[t._v("有问题直接提，互相学习，共同进步哈~ (*^__^*)")]),l("br"),l("el-divider",{attrs:{"content-position":"center"}},[t._v("常用联系方式")]),l("el-table",{attrs:{data:t.author}},[l("el-table-column",{attrs:{prop:"name",label:"名称",width:"100px"}}),l("el-table-column",{attrs:{prop:"wechat",label:"微信"}}),l("el-table-column",{attrs:{prop:"qq",label:"QQ"}})],1)],1)},k=[],w={name:"clientOther",data:function(){return{author:[{name:"古寒飞",wechat:"SPE3SRU3STAY",qq:"237663347"},{name:"yesuu",wechat:"ACFOLPTPZXIZ",qq:"791581230"}]}}},x=w,$=(l("f86e"),Object(d["a"])(x,C,k,!1,null,"8d4e50f0",null)),S=$.exports,I=function(){var t=this,e=t.$createElement,l=t._self._c||e;return l("el-form",{staticStyle:{"text-align":"left"},attrs:{"label-position":t.labelPosition,"label-width":"100px",model:t.client,rules:t.feishuRules}},[l("el-form-item",{attrs:{label:"客户端名称:",prop:"client_name"}},[l("el-input",{attrs:{placeholder:""},model:{value:t.client.client_name,callback:function(e){t.$set(t.client,"client_name",e)},expression:"client.client_name"}})],1),l("el-form-item",{attrs:{label:"客户端描述:"}},[l("el-input",{attrs:{placeholder:""},model:{value:t.client.client_description,callback:function(e){t.$set(t.client,"client_description",e)},expression:"client.client_description"}})],1),l("el-form-item",{attrs:{label:"客户端类型:"}},[l("el-input",{attrs:{disabled:""},model:{value:t.client.typeDescription,callback:function(e){t.$set(t.client,"typeDescription",e)},expression:"client.typeDescription"}})],1),l("el-divider",{attrs:{"content-position":"center"}},[t._v("机器人")]),l("el-form-item",{attrs:{label:"标题名称:"}},[l("el-input",{attrs:{placeholder:"标题中要包含机器人的放行关键词"},model:{value:t.client.client_info.robot_keyword,callback:function(e){t.$set(t.client.client_info,"robot_keyword",e)},expression:"client.client_info.robot_keyword"}})],1),l("el-form-item",{attrs:{label:"标题颜色:"}},[l("el-select",{attrs:{placeholder:"请选择"},model:{value:t.client.client_info.title_color,callback:function(e){t.$set(t.client.client_info,"title_color",e)},expression:"client.client_info.title_color"}},t._l(t.all_title_color,(function(t){return l("el-option",{key:t.color,attrs:{label:t.label,value:t.color}})})),1)],1),t._l(t.client.client_info.robot_url_list,(function(e,n){return l("div",{key:n},[l("el-form-item",{attrs:{label:"机器人URL:"}},[l("el-input",{staticStyle:{width:"85%"},attrs:{placeholder:"从飞书上粘贴而来的机器人URL地址"},model:{value:e.url,callback:function(l){t.$set(e,"url",l)},expression:"list.url"}}),l("el-button",{staticStyle:{"margin-left":"10px"},attrs:{type:"danger",icon:"el-icon-delete",circle:"",size:"mini"},on:{click:function(e){return t.del(n)}}})],1)],1)})),l("div",{staticStyle:{"text-align":"center"}},[l("el-button",{attrs:{type:"text",icon:"el-icon-plus",size:"mini"},on:{click:t.add}},[t._v("再添加一个机器人")])],1),l("p",{attrs:{id:"textStype"}},[t._v(" 此处可以添加多个机器人，推送消息时会从中随机挑选一个URL来使用，可以避免单个机器人消息推送时（每分钟）的次数限制，避免重要报警信息被漏送的可能。")]),l("br"),l("br"),l("el-form-item",[l("el-button",{attrs:{type:"primary"},on:{click:t.onSubmit}},[t._v("立即创建")]),l("el-button",{on:{click:t.closeDrawer}},[t._v("取消")])],1)],2)},D=[],L={name:"clientFeishu",data:function(){return{labelPosition:"right",client:{client_name:"",client_description:"",client_type:"feishu",is_active:!1,client_info:{robot_keyword:"",robot_url_list:[{url:""}],title_color:"grey"},typeDescription:"飞书·机器人"},feishuRules:{client_name:[{required:!0,message:"name不能为空",trigger:"blur"}]},all_title_color:[{color:"grey",label:"灰色（默认）"},{color:"blue",label:"蓝色"},{color:"wathet",label:"浅蓝色"},{color:"turquoise",label:"松石绿"},{color:"green",label:"绿色"},{color:"yellow",label:"黄色"},{color:"orange",label:"橘色"},{color:"red",label:"红色"},{color:"carmine",label:"胭脂红"},{color:"violet",label:"紫罗兰色"},{color:"purple",label:"紫色"},{color:"indigo",label:"靛青"}]}},props:{getClientList:Function},methods:{onSubmit:function(){var t=this;Object(r["g"])(this.client).then((function(e){e.data.result?(t.$message.success("添加成功..."),t.getClientList(),t.$store.commit("updateDrawerStatus",!1)):(t.$message.error("添加失败..."),t.getClientList(),t.$store.commit("updateDrawerStatus",!1))})).catch((function(t){console.log(t)}))},add:function(){var t={url:""};this.client.client_info.robot_url_list.push(t);for(var e=0;e<this.client.client_info.robot_url_list.length;e++)console.log(this.client.client_info.robot_url_list[e])},del:function(t){this.client.client_info.robot_url_list.splice(t,1)},closeDrawer:function(){this.getClientList(),this.$store.commit("updateDrawerStatus",!1)}}},O=L,R=(l("83b3"),Object(d["a"])(O,I,D,!1,null,"0e58b1cf",null)),j=R.exports,P={name:"cDrawer",data:function(){return{tabPosition:"top"}},props:{getClientList:Function},components:{clientDingtalk:_,clientWechat:y,clientOther:S,clientFeishu:j},computed:{getDrawerStatus:function(){return this.$store.state.DrawerStatus}},methods:{handleClose:function(t){var e=this;this.$confirm("确认关闭？").then((function(l){console.log(l),e.$store.commit("updateDrawerStatus",!1),t()})).catch((function(t){console.log(t)}))}}},q=P,E=Object(d["a"])(q,o,c,!1,null,"f859285e",null),A=E.exports,F=function(){var t=this,e=t.$createElement,l=t._self._c||e;return l("el-drawer",{attrs:{title:"客户端信息查看",visible:t.visibleStatus,"destroy-on-close":!0,"before-close":t.thisClose},on:{"update:visible":function(e){t.visibleStatus=e}}},[l("el-form",{staticStyle:{"text-align":"left"},attrs:{"label-position":"right","label-width":"100px",model:t.oneClientInfo}},[l("el-form-item",{attrs:{label:"客户端名称:"}},[l("el-input",{attrs:{disabled:""},model:{value:t.oneClientInfo.client_name,callback:function(e){t.$set(t.oneClientInfo,"client_name",e)},expression:"oneClientInfo.client_name"}})],1),l("el-form-item",{attrs:{label:"客户端描述:"}},[l("el-input",{attrs:{disabled:""},model:{value:t.oneClientInfo.client_description,callback:function(e){t.$set(t.oneClientInfo,"client_description",e)},expression:"oneClientInfo.client_description"}})],1),l("el-form-item",{attrs:{label:"客户端类型:"}},[l("el-input",{attrs:{disabled:""},model:{value:t.oneClientInfo.client_type,callback:function(e){t.$set(t.oneClientInfo,"client_type",e)},expression:"oneClientInfo.client_type"}})],1),l("el-form-item",{attrs:{label:"激活状态:"}},[l("el-input",{attrs:{disabled:""},model:{value:t.oneClientInfo.is_active,callback:function(e){t.$set(t.oneClientInfo,"is_active",e)},expression:"oneClientInfo.is_active"}})],1),l("el-divider",{attrs:{"content-position":"center"}},[t._v("客户端详情")]),"dingtalk"===t.oneClientInfo.client_type?l("div",[l("el-form-item",{attrs:{label:"放行关键字:"}},[l("el-input",{attrs:{placeholder:"需要后端发数据给我",disabled:""},model:{value:t.oneClientInfo.client_info.robot_keyword,callback:function(e){t.$set(t.oneClientInfo.client_info,"robot_keyword",e)},expression:"oneClientInfo.client_info.robot_keyword"}})],1),t._l(t.oneClientInfo.client_info.robot_url_list,(function(e,n){return l("div",{key:n},[l("el-form-item",{attrs:{label:"机器人URL:"}},[l("el-input",{attrs:{placeholder:"需要后端发数据给我",disabled:""},model:{value:e.url,callback:function(l){t.$set(e,"url",l)},expression:"list.url"}})],1)],1)}))],2):"wechat"===t.oneClientInfo.client_type?l("div",[l("el-form-item",{attrs:{label:"企业ID:"}},[l("el-input",{attrs:{placeholder:"需要后端发数据给我",disabled:""},model:{value:t.oneClientInfo.client_info.corp_id,callback:function(e){t.$set(t.oneClientInfo.client_info,"corp_id",e)},expression:"oneClientInfo.client_info.corp_id"}})],1),l("el-form-item",{attrs:{label:"应用AgentId:"}},[l("el-input",{attrs:{placeholder:"需要后端发数据给我",disabled:""},model:{value:t.oneClientInfo.client_info.agent_id,callback:function(e){t.$set(t.oneClientInfo.client_info,"agent_id",e)},expression:"oneClientInfo.client_info.agent_id"}})],1),l("el-form-item",{attrs:{label:"应用Secret:"}},[l("el-input",{attrs:{placeholder:"需要后端发数据给我",disabled:""},model:{value:t.oneClientInfo.client_info.secret,callback:function(e){t.$set(t.oneClientInfo.client_info,"secret",e)},expression:"oneClientInfo.client_info.secret"}})],1),l("el-form-item",{attrs:{label:"接收用户:"}},[l("el-input",{attrs:{type:"textarea",autosize:{minRows:4,maxRows:6},placeholder:"可以填写多个用户账号，用 | 分割开 （例如：aaa|bbb）",disabled:""},model:{value:t.oneClientInfo.client_info.touser,callback:function(e){t.$set(t.oneClientInfo.client_info,"touser",e)},expression:"oneClientInfo.client_info.touser"}})],1)],1):"feishu"===t.oneClientInfo.client_type?l("div",[l("el-form-item",{attrs:{label:"标题名称:"}},[l("el-input",{attrs:{placeholder:"需要后端发数据给我",disabled:""},model:{value:t.oneClientInfo.client_info.robot_keyword,callback:function(e){t.$set(t.oneClientInfo.client_info,"robot_keyword",e)},expression:"oneClientInfo.client_info.robot_keyword"}})],1),l("el-form-item",{attrs:{label:"标题颜色:"}},[l("el-input",{attrs:{placeholder:"需要后端发数据给我",disabled:""},model:{value:t.oneClientInfo.client_info.title_color,callback:function(e){t.$set(t.oneClientInfo.client_info,"title_color",e)},expression:"oneClientInfo.client_info.title_color"}})],1),t._l(t.oneClientInfo.client_info.robot_url_list,(function(e,n){return l("div",{key:n},[l("el-form-item",{attrs:{label:"机器人URL:"}},[l("el-input",{attrs:{placeholder:"需要后端发数据给我",disabled:""},model:{value:e.url,callback:function(l){t.$set(e,"url",l)},expression:"list.url"}})],1)],1)}))],2):l("div",[t._v(" 不显示 ")])],1)],1)},U=[],z={name:"cDrawerOneDataInfo",data:function(){return{clientInfo:{id:null,client_name:null,client_description:null,is_active:null,client_type:null,client_info:{}}}},props:{oneClientInfo:Object,visibleStatus:Boolean,thisClose:Function}},N=z,G=Object(d["a"])(N,F,U,!1,null,"49ef7c62",null),M=G.exports,W={name:"ViewClient",data:function(){return{visibleStatus:!1,clientList:[{id:1,client_name:"示例客户端1",client_description:"示例数据，随时可删~",is_active:!1,client_type:"dingtalk",client_annotation:"钉钉·机器人"},{id:2,client_name:"示例客户端2",client_description:"示例数据，随时可删~",is_active:!1,client_type:"wechat",client_annotation:"企业微信·应用号"},{id:3,client_name:"示例客户端3",client_description:"示例数据，随时可删~",is_active:!0,client_type:"feishu",client_annotation:"飞书·机器人"}],clientOneInfo:{id:null,client_name:null,client_description:null,is_active:null,client_type:"dingtalk",client_annotation:"",client_info:{}}}},components:{CDrawer:A,CDrawerOneDataInfo:M},methods:{getAnnotation:function(t){var e={dingtalk:"钉钉·机器人",wechat:"企业微信·应用号",feishu:"飞书·机器人"};return e[t]},saveActiveStatus:function(){var t=this;this.clientList.forEach((function(e){Object(r["l"])(t.$store.getters.getNamespace,e.id,e).then((function(t){console.log(t.data)})).catch((function(t){console.log(t)}))})),this.$message.success("激活状态保存成功...")},addClient:function(){this.$store.commit("updateDrawerStatus",!0)},thisClose:function(){this.visibleStatus=!1},deleteOneClient:function(t,e){var l=this,n=e[t].id;Object(r["a"])(this.$store.getters.getNamespace,n,null).then((function(n){1===n.data.code?(l.$message.success("删除一行数据成功..."),e.splice(t,1)):l.$message.error("删除数据失败...")})).catch((function(t){console.log(t)}))},getOneClient:function(t,e){var l=this,n=e[t].id;Object(r["d"])(this.$store.getters.getNamespace,n,null).then((function(t){console.log(t.data.result),l.clientOneInfo=t.data.result})).catch((function(t){console.log(t)})),this.visibleStatus=!0},getClientList:function(){var t=this;Object(r["c"])(this.$store.getters.getNamespace,null).then((function(e){0===e.data.result.length?(console.log("数据库里没有数据"),t.$message.error("数据库中没有数据，显示Demo数据...")):(t.clientList=e.data.result,t.clientList.forEach((function(e){e.client_annotation=t.getAnnotation(e.client_type)})))})).catch((function(t){console.log(t)}))}},created:function(){this.$store.commit("updateStepsActive",3),this.getClientList()}},J=W,Q=(l("ed1a"),Object(d["a"])(J,n,i,!1,null,"053d8507",null));e["default"]=Q.exports},"5ed7":function(t,e,l){"use strict";l("c3a3")},"77ba":function(t,e,l){},"82f9":function(t,e,l){},"83b3":function(t,e,l){"use strict";l("82f9")},c3a3:function(t,e,l){},d4eb:function(t,e,l){},ed1a:function(t,e,l){"use strict";l("77ba")},f86e:function(t,e,l){"use strict";l("d4eb")}}]);
//# sourceMappingURL=chunk-be00c0e4.0161dca6.js.map