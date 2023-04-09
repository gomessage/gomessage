<template>
  <el-card shadow="always">
    <div slot="header" class="clearfix">

      <!--是否聚合发送信息-->
      <el-switch v-model="template.template_is_merge" inactive-text="聚合发送" style="float: left; padding: 3px 0;" @change="pushTemplateData"></el-switch>

      <!--教程提示按钮，跳转代码-->
      <el-tooltip class="item" content='跳转到新页面：查看"消息模板"编写教程' effect="dark" placement="bottom" style="float: left;margin-left: 30px;padding-top: 3px;">
        <el-link :underline="false" type="primary" @click="newTagPage">
          <span><i class="el-icon-info"></i></span>
        </el-link>
      </el-tooltip>

      <!--标题-->
      <span style="margin-right: 70px">消息渲染模板</span>

      <!--保存模板-->
      <el-button style="float: right; padding: 3px 0;" type="text" v-on:click="pushTemplateData()">保存模板</el-button>
    </div>

    <!--<el-row style="color: #C0C4CC;font-size: 12px;line-height: 1.7;padding-bottom: 20px">-->
    <!--    <el-col span="24" style="text-align: right">-->
    <!--        <el-link type="primary">查看文档</el-link>-->
    <!--    </el-col>-->
    <!--</el-row>-->

    <div>
      <el-input
        v-model="template.template_content"
        :autosize="{ minRows: 10, maxRows: 200}"
        placeholder="请输入Golang语法的模板内容"
        resize="none"
        type="textarea">
      </el-input>
    </div>
  </el-card>
</template>

<script>

import {getTemplate, postTemplate} from '@/service/requests'

export default {
  name: "cTemplate",
  data() {
    return {
      template: {
        template_is_merge: false,
        template_content: '{{ if eq .N6 "firing" }}\n' +
          '\n' +
          '## <font color=\'#FF0000\'>【报警中】服务器{{ .N3 }}</font>\n' +
          '\n' +
          '{{ else if eq .N6 "resolved" }}\n' +
          '\n' +
          '## <font color=\'#20B2AA\'>【已恢复】服务器{{ .N3 }}</font>\n' +
          '\n' +
          '{{ else }}\n' +
          '\n' +
          '## 标题：信息通知\n' +
          '\n' +
          '{{ end }}\n' +
          '\n' +
          '========================\n' +
          '\n' +
          '**告警规则**：{{ .N1 }}\n' +
          '\n' +
          '**告警级别**：{{ .N2 }}\n' +
          '\n' +
          '**主机名称**：{{ .N3 }} \n' +
          '\n' +
          '**主机地址**：{{ .N4 }}\n' +
          '\n' +
          '**告警详情**：{{ .N5 }}\n' +
          '\n' +
          '**告警状态**：{{ .N6 }}\n' +
          '\n' +
          '**触发时间**：{{ .N7 }}\n' +
          '\n' +
          '**发送时间**：{{ .N8 }}\n' +
          '\n' +
          '**规则详情**：[Prometheus控制台](https://www.baidu.com)\n' +
          '\n' +
          '**报警详情**：[Alertmanager控制台](https://www.baidu.com)\n',
      }
    }
  },
  components: {},
  methods: {
    newTagPage: function () {
      let url = "https://github.com/gomessage/gomessage#gomessage"
      window.open(url)
    },

    //保存模板数据
    pushTemplateData: function () {
      postTemplate(this.$store.getters.getNamespace, this.template).then(response => {
        console.log(response.data.result);
        this.$message.success("数据库更新成功...")
      }).catch(err => {
        console.log(err);
      });
    },

    //拉取数据
    pullTemplateData: function () {
      getTemplate(this.$store.getters.getNamespace, null).then(response => {
        let temp = response.data.result[0]["template_content"]
        let isMerge = response.data.result[0]["template_is_merge"]

        if (temp === undefined || temp === null || temp === "") {
          console.log("数据库里没有数据")
        } else {
          this.template.template_content = temp;
          this.template.template_is_merge = isMerge;
        }
      }).catch(err => {
        console.log(err);
      });
    }
  },

  created() {
    //修改步骤条的值
    this.$store.commit("updateStepsActive", 2);
    //拉取数据
    this.pullTemplateData();
  }
}
</script>

<style scoped>
/*#MessageTemplateContent {*/
/*    width: 100%;*/
/*    !*margin-left: 20%;*!*/
/*}*/

/*.templateContentClass >>> input {*/
/*    background-color: #2c3e50;*/
/*}*/
</style>
