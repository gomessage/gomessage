<template>
  <el-card shadow="always">
    <div slot="header" class="clearfix">

      <!--是否聚合发送信息-->
      <el-switch v-model="template.template_is_merge" inactive-text="聚合发送" style="float: left; padding: 3px 0;" @change="PushTemplateData"></el-switch>

      <!--教程提示按钮，跳转代码-->
      <el-tooltip class="item" content='跳转到新页面：查看"消息模板"编写教程' effect="dark" placement="bottom" style="float: left;margin-left: 30px;padding-top: 3px;">
        <el-link :underline="false" type="primary" @click="newTagPage">
          <span><i class="el-icon-info"></i></span>
        </el-link>
      </el-tooltip>

      <!--标题-->
      <span style="margin-right: 70px">消息渲染模板</span>

      <!--保存模板-->
      <el-button style="float: right; padding: 3px 0;" type="text" v-on:click="PushTemplateData()">保存模板</el-button>
    </div>

    <!--<el-row style="color: #C0C4CC;font-size: 12px;line-height: 1.7;padding-bottom: 20px">-->
    <!--    <el-col span="24" style="text-align: right">-->
    <!--        <el-link type="primary">查看文档</el-link>-->
    <!--    </el-col>-->
    <!--</el-row>-->

    <div class="system-template-panel">
      <div class="system-template-toolbar">
        <span class="system-template-label">系统模板</span>
        <el-select
            v-model="selectedSystemTemplateId"
            class="system-template-select"
            filterable
            placeholder="选择适用的通道模板">
          <el-option
              v-for="item in systemTemplates"
              :key="item.id"
              :label="item.platform + ' · ' + item.name"
              :value="item.id">
            <span>{{ item.platform }}</span>
            <span class="system-template-option-name">{{ item.name }}</span>
          </el-option>
        </el-select>
        <el-button
            type="primary"
            plain
            :disabled="!selectedSystemTemplate"
            @click="replaceWithSystemTemplate">应用到当前模板
        </el-button>
      </div>
      <template v-if="selectedSystemTemplate">
        <div class="system-template-description">
          {{ selectedSystemTemplate.description }}
        </div>
        <div class="system-template-preview-header">
          <span>模板内容预览</span>
          <el-tag size="mini" type="info">{{ selectedSystemTemplate.platform }}</el-tag>
        </div>
        <el-input
            :value="selectedSystemTemplate.content"
            :autosize="{ minRows: 8, maxRows: 18 }"
            class="system-template-preview"
            readonly
            resize="none"
            type="textarea">
        </el-input>
      </template>
      <div v-else class="system-template-description">
        选择后只会替换下方编辑区，确认内容后请点击“保存模板”。
      </div>
    </div>

    <div class="current-template-header">当前模板（保存时提交此处内容）</div>
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
import {commonAlertTemplate, SYSTEM_MESSAGE_TEMPLATES} from '@/utils/systemMessageTemplates'

export default {
  name: "cTemplate",
  data() {
    return {
      template: {
        template_is_merge: false,
        template_content: commonAlertTemplate,
      },
      systemTemplates: SYSTEM_MESSAGE_TEMPLATES,
      selectedSystemTemplateId: ''
    }
  },
  components: {},
  computed: {
    selectedSystemTemplate() {
      return this.systemTemplates.find(item => item.id === this.selectedSystemTemplateId)
    }
  },
  methods: {
    newTagPage: function () {
      let url = "https://github.com/gomessage/gomessage#gomessage"
      window.open(url)
    },

    // 用系统样例替换编辑区，不自动保存，避免误覆盖现有模板
    replaceWithSystemTemplate: function () {
      if (!this.selectedSystemTemplate) {
        this.$message.warning("请先选择系统模板...")
        return
      }

      this.$confirm(
          '将用“' + this.selectedSystemTemplate.platform + ' · ' + this.selectedSystemTemplate.name + '”替换当前编辑区内容。替换后不会自动保存，你可以继续修改。',
          '替换消息模板',
          {
            confirmButtonText: '确认替换',
            cancelButtonText: '取消',
            type: 'warning'
          }
      ).then(() => {
        this.template.template_content = this.selectedSystemTemplate.content
        this.$message.success('已替换编辑区，请检查后保存模板...')
      }).catch(() => {
      })
    },

    //保存模板数据
    PushTemplateData: function (updateData = true) {
      postTemplate(this.$store.getters.getNamespace, this.template).then(response => {
        console.log(response.data.result);
        if (updateData) {
          this.$message.success("数据库更新成功...")
        }
      }).catch(err => {
        console.log(err);
      });
    },

    //拉取数据
    PullTemplateData: function () {
      getTemplate(this.$store.getters.getNamespace, null).then(response => {
        if (response.data.result.length === 0) {
          this.PushTemplateData(false); //如果数据库中没有数据，则把本地的demo数据存储过去。
        } else {
          let temp = response.data.result[0]["template_content"]
          let isMerge = response.data.result[0]["template_is_merge"]
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
    this.PullTemplateData();
  }
}
</script>

<style scoped>
.system-template-panel {
  margin-bottom: 18px;
  padding: 14px 16px;
  border: 1px solid #ebeef5;
  border-radius: 4px;
  background: #fafafa;
}

.system-template-toolbar {
  display: flex;
  align-items: center;
}

.system-template-label {
  margin-right: 14px;
  color: #303133;
  font-weight: 500;
  white-space: nowrap;
}

.system-template-select {
  flex: 1;
  margin-right: 12px;
}

.system-template-option-name {
  float: right;
  margin-left: 24px;
  color: #8492a6;
  font-size: 13px;
}

.system-template-description {
  margin-top: 10px;
  color: #909399;
  font-size: 12px;
  line-height: 1.6;
  text-align: left;
}

.system-template-preview-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin: 14px 0 8px;
  color: #606266;
  font-size: 13px;
  font-weight: 500;
}

.system-template-preview >>> textarea {
  background: #f5f7fa;
  color: #606266;
  font-family: Menlo, Monaco, Consolas, "Courier New", monospace;
  font-size: 12px;
  line-height: 1.6;
}

.current-template-header {
  margin-bottom: 8px;
  color: #606266;
  font-size: 13px;
  font-weight: 500;
  text-align: left;
}

/*#MessageTemplateContent {*/
/*    width: 100%;*/
/*    !*margin-left: 20%;*!*/
/*}*/

/*.templateContentClass >>> input {*/
/*    background-color: #2c3e50;*/
/*}*/
</style>
