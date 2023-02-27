<template>
  <el-form :label-position="labelPosition" label-width="100px" :model="client" style="text-align: left">
    <el-form-item label="客户端名称:">
      <el-input v-model="client.client_name" placeholder=""></el-input>
    </el-form-item>

    <el-form-item label="客户端描述:">
      <el-input v-model="client.client_description" placeholder=""></el-input>
    </el-form-item>

    <el-form-item label="客户端类型:">
      <el-input v-model="client.typeDescription" disabled></el-input>
    </el-form-item>

    <el-divider content-position="center">机器人</el-divider>

    <el-form-item label="标题名称:">
      <el-input v-model="client.client_info.robot_keyword" placeholder="标题中要包含机器人的放行关键词"></el-input>
    </el-form-item>

    <el-form-item label="标题颜色:">
      <el-select v-model="client.client_info.title_color" placeholder="请选择">
        <el-option v-for="item in all_title_color" :key="item.color" :label="item.label" :value="item.color">
        </el-option>
      </el-select>
    </el-form-item>

    <div v-for="(list, index) in client.client_info.robot_url_list" :key="index">
      <el-form-item label="机器人URL:">
        <el-input v-model="list.url" placeholder="从飞书上粘贴而来的机器人URL地址" style="width: 85%"></el-input>
        <el-button type="danger" icon="el-icon-delete" circle size="mini" v-on:click="del(index)"
                   style="margin-left: 10px"></el-button>
      </el-form-item>
    </div>
    <div style="text-align: center">
      <el-button type="text" icon="el-icon-plus" size="mini" @click="add">再添加一个机器人</el-button>
    </div>
    <p id="textStype">
      此处可以添加多个机器人，推送消息时会从中随机挑选一个URL来使用，可以避免单个机器人消息推送时（每分钟）的次数限制，避免重要报警信息被漏送的可能。</p>

    <br><br>
    <el-form-item>
      <el-button type="primary" @click="onSubmit">立即创建</el-button>
      <el-button @click="closeDrawer">取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
import {postClient} from '@/service/requests'

export default {
  name: "clientFeishu",
  data() {
    return {
      labelPosition: "right", //表格标题右对齐

      client: {
        client_name: "",
        client_description: "",
        client_type: "feishu", //客户端类型，与后端的约定，禁止修改
        is_active: false,
        client_info: {
          robot_keyword: "",
          robot_url_list: [
            {
              url: "",
            },
          ],
          title_color: "grey",
        },
        typeDescription: "飞书·机器人",
      },
      all_title_color: [
        {
          color: "grey",
          label: "灰色（默认）"
        },
        {
          color: "blue",
          label: "蓝色"
        },
        {
          color: "wathet",
          label: "浅蓝色"
        },
        {
          color: "turquoise",
          label: "松石绿"
        },
        {
          color: "green",
          label: "绿色"
        },
        {
          color: "yellow",
          label: "黄色"
        },
        {
          color: "orange",
          label: "橘色"
        },
        {
          color: "red",
          label: "红色"
        },
        {
          color: "carmine",
          label: "胭脂红"
        },
        {
          color: "violet",
          label: "紫罗兰色"
        },
        {
          color: "purple",
          label: "紫色"
        },
        {
          color: "indigo",
          label: "靛青"
        },

      ],
    }
  }, props: {
    getClientList: Function,
  },
  methods: {
    onSubmit: function () {
      postClient(this.client).then(response => {
        // console.log(response.data);
        if (response.data.result) {
          this.$message.success("添加成功...")
          this.getClientList();
          this.$store.commit("updateDrawerStatus", false);
        } else {
          this.$message.error("添加失败...");
          this.getClientList();
          this.$store.commit("updateDrawerStatus", false);
        }
      }).catch(err => {
        console.log(err);
      });
    },
    // 添加输入框
    add: function () {
      let cope = {
        url: "",
      };
      this.client.client_info.robot_url_list.push(cope);
      // console.log(this.mapList2);
      for (let i = 0; i < this.client.client_info.robot_url_list.length; i++) {
        console.log(this.client.client_info.robot_url_list[i]);
      }
    },
    del: function (index) {
      this.client.client_info.robot_url_list.splice(index, 1);
      // console.log(this.mylist);
    },
    closeDrawer: function () {
      this.getClientList();
      this.$store.commit("updateDrawerStatus", false);
    },
  }
}
</script>

<style scoped>
#textStype {
  font-size: 11px;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  line-height: 1.7;
  color: #C0C4CC;
}
</style>
