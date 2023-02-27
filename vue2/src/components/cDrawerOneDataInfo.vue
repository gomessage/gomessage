<template>
  <el-drawer
      title="客户端信息查看"
      :visible.sync="visibleStatus"
      :destroy-on-close="true"
      :before-close="thisClose">
    <el-form label-position="right" label-width="100px" :model="oneClientInfo" style="text-align: left">
      <el-form-item label="客户端名称:">
        <el-input v-model="oneClientInfo.client_name" disabled></el-input>
      </el-form-item>

      <el-form-item label="客户端描述:">
        <el-input v-model="oneClientInfo.client_description" disabled></el-input>
      </el-form-item>

      <el-form-item label="客户端类型:">
        <el-input v-model="oneClientInfo.client_type" disabled></el-input>
      </el-form-item>

      <el-form-item label="激活状态:">
        <el-input v-model="oneClientInfo.is_active" disabled></el-input>
      </el-form-item>

      <el-divider content-position="center">客户端详情</el-divider>

      <div v-if="oneClientInfo.client_type==='dingtalk'">

        <el-form-item label="放行关键字:">
          <el-input v-model="oneClientInfo.client_info.robot_keyword" placeholder="需要后端发数据给我"
                    disabled></el-input>
        </el-form-item>

        <div v-for="(list, index) in oneClientInfo.client_info.robot_url_list" :key="index">
          <el-form-item label="机器人URL:">
            <el-input v-model="list.url" placeholder="需要后端发数据给我" disabled></el-input>
          </el-form-item>
        </div>
      </div>

      <div v-else-if="oneClientInfo.client_type==='wechat'">
        <el-form-item label="企业ID:">
          <el-input v-model="oneClientInfo.client_info.corp_id" placeholder="需要后端发数据给我" disabled></el-input>
        </el-form-item>

        <el-form-item label="应用AgentId:">
          <el-input v-model="oneClientInfo.client_info.agent_id" placeholder="需要后端发数据给我" disabled></el-input>
        </el-form-item>

        <el-form-item label="应用Secret:">
          <el-input v-model="oneClientInfo.client_info.secret" placeholder="需要后端发数据给我" disabled></el-input>
        </el-form-item>

        <el-form-item label="接收用户:">
          <el-input
              type="textarea"
              :autosize="{ minRows: 4, maxRows: 6}"
              v-model="oneClientInfo.client_info.touser"
              placeholder="可以填写多个用户账号，用 | 分割开 （例如：aaa|bbb）"
              disabled>
          </el-input>
        </el-form-item>
      </div>

      <div v-else-if="oneClientInfo.client_type==='feishu'">

        <el-form-item label="标题名称:">
          <el-input v-model="oneClientInfo.client_info.robot_keyword" placeholder="需要后端发数据给我"
                    disabled></el-input>
        </el-form-item>

        <el-form-item label="标题颜色:">
          <el-input v-model="oneClientInfo.client_info.title_color" placeholder="需要后端发数据给我"
                    disabled></el-input>
        </el-form-item>

        <div v-for="(list, index) in oneClientInfo.client_info.robot_url_list" :key="index">
          <el-form-item label="机器人URL:">
            <el-input v-model="list.url" placeholder="需要后端发数据给我" disabled></el-input>
          </el-form-item>
        </div>
      </div>

      <div v-else>
        不显示
      </div>

    </el-form>

  </el-drawer>
</template>

<script>
export default {
  name: "cDrawerOneDataInfo",
  data() {
    return {
      clientInfo: {
        id: null,
        client_name: null,
        client_description: null,
        is_active: null,
        client_type: null,
        client_info: {}
      }
    }
  },
  props: {
    oneClientInfo: Object,
    visibleStatus: Boolean,
    thisClose: Function,
  }
}
</script>

<style scoped>

</style>
