<template>
  <!--#303133-->
  <!--#41555d-->
  <el-menu
    :default-active="getStoreNamespace"
    active-text-color="#ffd04b"
    background-color="#41555d"
    style="height: 100%;border-right: solid 0 red"
    text-color="#fff"
  >

    <!--左侧logo-->
    <!--<el-menu-item index="999">-->
    <!--  <router-link to="/">-->
    <!--    <img height="90%" src="../assets/logo.svg">-->
    <!--    &nbsp;&nbsp;-->
    <!--    <span style="font-size: 15px;color: #fff;">GoMessage · 消息转发器</span>-->
    <!--  </router-link>-->
    <!--</el-menu-item>-->

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
        icon="el-icon-setting"
        plain
        size="mini"
        @click="dialogFormVisible222 = true"
        style="padding-right: 35px"
      >管理推送通道
      </el-button>
    </el-menu-item>


    <!--利用对话框，添加命名空间-->
    <el-dialog
      :visible.sync="dialogFormVisible222"
      title="消息推送通道"
      modal
      width="60%"
      top="10vh"
      lock-scroll
    >
      <!--表格-->
      <el-table
        :data="namespaceList"
        border
        :header-cell-style="{background:'#2f2f35',color:'#fff'}"
        style="width: 100%"
      >

        <el-table-column
          prop="name"
          label="名称"
          width="180">
        </el-table-column>

        <el-table-column
          prop="description"
          label="描述">
        </el-table-column>

        <!--<el-table-column-->
        <!--    prop="is_active"-->
        <!--    label="是否激活"-->
        <!--    width="150">-->
        <!--    <template slot-scope="scope">-->
        <!--        <el-checkbox v-model="scope.row.is_active" @change="activeNamespace">激活</el-checkbox>-->
        <!--    </template>-->
        <!--</el-table-column>-->

        <el-table-column fixed="right" label="操作" width="100">
          <template slot-scope="scope">
            <el-button size="small" type="danger"
                       @click.native.prevent="deleteOneNamespace(scope.$index, namespaceList)">删除
            </el-button>
          </template>
        </el-table-column>
      </el-table>


      <br>
      <p class="authorStatement2">
        GoMessage v2 版本支持多通道并发，可以同时存在多个"通道"，互不干扰的进行消息转发。
      </p>
      <br>

      <!--分割线-->
      <el-divider content-position="left"><i class="el-icon-circle-plus-outline"> 新增通道</i></el-divider>
      <!--<br>-->

      <!--这是一个form表单，对话框支持填充表单-->
      <el-form ref="namespaceForm" :model="namespaceForm" :rules="namespaceRules" style="width: 60%;">
        <el-form-item label="通道名称" label-width="105px" prop="name">
          <el-input v-model="namespaceForm.name" autocomplete="off" placeholder="请输入通道名称（只能是纯英文名称，不限大小写）"></el-input>
        </el-form-item>

        <el-form-item label="通道描述" label-width="105px">
          <el-input v-model="namespaceForm.description" autocomplete="off" type="textarea" :rows="3"
                    placeholder="请输入推送通道的描述信息"></el-input>
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
import {deleteNamespaceOne, getNamespace, postNamespace, putNamespaceOne} from "@/service/requests";
import {isLength, isStringOrNumber, isNumberStart} from "@/utils/validate";

export default {
  name: "NavAside",
  data() {
    return {
      namespaceList: [
        {
          "name": "default",
          "description": "default",
          "is_active": true
        },
      ],
      dialogFormVisible222: false,
      namespaceForm: {
        name: '',
        description: '',
        is_active: false,
      },
      namespaceRules: {
        name: [
          {required: true, message: "name不能为空", trigger: "blur"},
          {validator: isLength, trigger: "blur"},
          {validator: isNumberStart, trigger: "blur"},
          {validator: isStringOrNumber, trigger: "blur"},
        ],
      }
    }
  },
  methods: {
    pullNamespace: function () {
      getNamespace().then(response => {
        if (response.data.code === 1) {
          this.namespaceList = response.data.result;
          // this.$message.success("成功");

          // 根据当前通道名称来缓存通道数据
          let thisNamespaceName = this.$store.getters.getNamespace;
          this.namespaceList.forEach((item) => {
            if (item["name"] === thisNamespaceName) {
              this.$store.commit("updateNamespaceInfo", item)
            }
          })

        }
      }).catch(err => {
        console.log(err);
      })
    },
    updateNamespace: function (item, event) {
      let namespace = item.name;
      console.log(namespace, event);
      this.$store.commit("updateNamespace", namespace);
      this.$store.commit("updateNamespaceInfo", item)
      //刷新当前页
      location.reload();
    },
    addNamespace: function () {
      this.$refs["namespaceForm"].validate(valid => {
        if (valid) {
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
      });
    },
    // 删除一行数据：跟后端交互，然后刷新表格
    deleteOneNamespace(index, rows) {
      let id = rows[index].id;
      deleteNamespaceOne(id).then(response => {
        if (response.data.code === 1) {
          this.$message.success("删除一行数据成功...");
          rows.splice(index, 1);
        } else {
          this.$message.error("删除数据失败...");
        }
      }).catch(err => {
        console.log(err);
      });
    },
    activeNamespace: function () {
      this.namespaceList.forEach(namespace => {
        putNamespaceOne(namespace.id, namespace).then(response => {
          console.log(response.data.result)
        }).catch(err => {
          console.log(err)
        })
      })
      this.$message.success("数据更新成功...")
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

.authorStatement2 {
  text-align: left;
  font-size: 8px;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  line-height: 1.7;
  color: #C0C4CC;
}
</style>
