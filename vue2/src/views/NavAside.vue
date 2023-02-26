<template>
  <el-menu
      :default-active="getStoreNamespace"
      active-text-color="#ffd04b"
      background-color="#000"
      style="height: 100%"
      text-color="#fff"
  >

    <!--左侧logo-->
    <el-menu-item index="999">
      <router-link to="/">
        <img height="90%" src="../assets/logo.svg">
        &nbsp;&nbsp;
        <span style="font-size: 15px;color: #fff;">GoMessage · 消息转发器</span>
      </router-link>
    </el-menu-item>

    <br>

    <!--for循环命名空间-->
    <el-menu-item
        v-for="(oneNs,index) in namespaceList"
        :key="index"
        :index="oneNs.name"
        style="text-align: left"
        @click="updateNamespace(oneNs,$event)">
      <i class="el-icon-menu"></i>
      <span slot="title">{{ oneNs.name }}</span>
    </el-menu-item>

    <br>

    <!--添加一个新的namespace-->
    <el-menu-item index="998" style="text-align: left">
      <el-button
          icon="el-icon-star-off"
          plain
          size="mini"
          @click="dialogFormVisible222 = true"
      >新增命名空间
      </el-button>
    </el-menu-item>


    <!--利用对话框，添加命名空间-->
    <el-dialog :visible.sync="dialogFormVisible222" style="text-align: left" title="新增命名空间">
      <!--这是一个form表单，对话框支持填充表单-->
      <el-form :model="namespaceForm" style="width: 70%">
        <el-form-item label="名称" label-width="200px">
          <el-input v-model="namespaceForm.name" autocomplete="off"></el-input>
        </el-form-item>

        <!--<el-form-item label="状态" label-width="200px">-->
        <!--    <el-radio-group v-model="namespaceForm.is_active">-->
        <!--        <el-radio label="激活"></el-radio>-->
        <!--        <el-radio label="不激活"></el-radio>-->
        <!--    </el-radio-group>-->
        <!--</el-form-item>-->

        <el-form-item label="描述" label-width="200px">
          <el-input v-model="namespaceForm.description" autocomplete="off" type="textarea"></el-input>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible222 = false">取 消</el-button>
        <el-button type="primary" @click="addNamespace">确 定</el-button>
      </div>
    </el-dialog>


  </el-menu>
</template>

<script>
import {getNamespace, postNamespace} from "@/service/requests";

export default {
  name: "NavAside",
  data() {
    return {
      namespaceList: [
        {"name": "default"},
        {"name": "prometheus"},
        {"name": "alertmanager"},
        {"name": "zabbix"},
        {"name": "elasticsearch"},
        {"name": "grafana"},
      ],
      dialogFormVisible222: false,
      namespaceForm: {
        name: '',
        description: '',
        is_active: false,
      },
    }
  },
  methods: {
    pullNamespace: function () {
      getNamespace().then(response => {
        if (response.data.code === 1) {
          this.namespaceList = response.data.result;
          // this.$message.success("成功");
        }
      }).catch(err => {
        console.log(err);
      })
    },
    updateNamespace: function (item, event) {
      let namespace = item.name;
      console.log(namespace, event);
      this.$store.commit("updateNamespace", namespace);
      //刷新当前页
      location.reload();
    },
    addNamespace: function () {
      // 关闭对话框视图
      this.dialogFormVisible222 = false;
      // 保证所有的命名空间永远都是激活的
      this.namespaceForm.is_active = true
      // 发送post请求
      postNamespace(this.namespaceForm).then(response => {
        console.log(response)
        location.reload();
      })
    }
  }, computed: {
    // 计算属性：动态获取vuex中的值
    getStoreNamespace: function () {
      return this.$store.getters.getNamespace
    }
  },
  //created生命周期
  created() {
    this.pullNamespace();
  }
}
</script>

<style scoped>
#NavAside-sty {
  height: 100%;
  background-color: aqua;
}

.router-link-active {
  text-decoration: none;
}
</style>
