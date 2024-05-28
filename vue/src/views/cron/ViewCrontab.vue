<template>
  <!--  布局撑满el-main容器 -->
  <el-row style="top: 0;left: 0;right: 0;bottom: 0;position: absolute; display: flex;">

    <!--    左侧col-->
    <el-col :span="12">

      <el-card style="margin: 20px;padding: 10px;text-align: left;">

        <div slot="header" class="clearfix">
          <span>发送定时消息（Crontab类型的时间规则定义）</span>
        </div>

        <el-form :ref="formData2" :model="formData2" label-width="90px">
          <el-form-item label="名称：">
            <el-input v-model="formData2.crontab_name" placeholder="下班打卡提醒"></el-input>
          </el-form-item>

          <el-form-item label="时间规则：">
            <el-input v-model="formData2.crontab_rule" placeholder="*/5  *  *  *  * （分、时、日、月、周）"></el-input>
          </el-form-item>

          <!--          <el-form-item label="发送通道：">-->
          <!--            <el-checkbox-group v-model="form.crontab_namespace">-->
          <!--              <el-checkbox v-for="item in namespaceOptions" :label="item.label" :key="item.label" name="type"></el-checkbox>-->
          <!--            </el-checkbox-group>-->
          <!--          </el-form-item>-->

          <el-form-item label="发送通道：">
            <el-select
                v-model="formData2.crontab_namespace"
                multiple
                placeholder="请选择"
                style="width: 70%;"
                @visible-change="visible => { if (visible) listNamespace() }">
              <el-option
                  v-for="item in namespaceOptions"
                  :key="item.label"
                  :label="item.label"
                  :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否启用：">
<!--            {{formData2.crontab_activate}}-->
            <el-switch v-model="formData2.crontab_activate"></el-switch>
          </el-form-item>


          <el-form-item label="消息内容：">
            <el-input
                type="textarea"
                v-model="formData2.crontab_content"
                :autosize="{ minRows: 10, maxRows: 200}"
            ></el-input>
          </el-form-item>

          <el-form-item>
            <el-button type="primary" @click="onSubmit">创建或更新</el-button>
            <el-button>取消</el-button>
          </el-form-item>

        </el-form>
      </el-card>
    </el-col>

    <!--    ===========================================================================================-->
    <!--    ===========================================================================================-->
    <!--    ===========================================================================================-->
    <!--    ===========================================================================================-->
    <!--    ===========================================================================================-->
    <!--    右侧col-->
    <el-col :span="12">
      <el-card style="margin: 20px;">
        <el-table
            :data="ListData2"
            border
            :header-cell-style="{background:'#2f2f35',color:'#fff'}"
            @row-click="handleRowClick"
            highlight-current-row
            style="width: 100%;">

          <el-table-column
              prop="crontab_name"
              show-overflow-tooltip
              sortable
              label="名称">
          </el-table-column>

          <el-table-column
              prop="crontab_rule"
              width="110"
              label="时间规则">
          </el-table-column>

          <el-table-column
              prop="crontab_namespace"
              label="发送通道"
              show-overflow-tooltip
              :formatter="formatNamespace">
          </el-table-column>

          <el-table-column
              width="110"
              prop="crontab_activate"
              :formatter="formatStatus"
              sortable
              label="是否启用">
          </el-table-column>

          <el-table-column
              width="80"
              label="操作">
            <template slot-scope="scope">
              <el-button @click="delete_one_data(scope.row)" type="text" size="small">删除</el-button>
              <!--              <el-button type="text" size="small">编辑</el-button>-->
            </template>
          </el-table-column>


        </el-table>
      </el-card>
    </el-col>
  </el-row>
</template>


<script>
import {
  deleteCrontabOne,
  getCrontab,
  getCrontabOne,
  getNamespace,
  postCrontab,
  putCrontabOne
} from "@/service/requests";

export default {
  name: "ViewCrontab",
  data() {
    return {
      componentKey: 0,
      ListData2: [
        {
          id: "",
          crontab_name: '提醒小伙伴下班打卡',
          crontab_rule: '*/5 * 2 * *',
          crontab_namespace: ["default", "ttpai", "alarmdog"],
          crontab_activate: false,
          crontab_content: '113123123123123123123113123123123123123123113123123123123123123113123123123123123123113123123123123123123113123123123123123123',
        },
        // {
        //   id: "",
        //   crontab_name: '002',
        //   crontab_rule: '*/4 * * * *',
        //   crontab_namespace: ["default", "test"],
        //   crontab_activate: false,
        //   crontab_content: '113123123123123123123113123123123123123123113123123123123123113123123123123123123113123123123123123123113123123123123123123',
        // },
      ],
      formData2: {
        id: -1,
        crontab_name: '',
        crontab_rule: '',
        crontab_namespace: [],
        crontab_activate: false,
        crontab_content: '',
      },
      namespaceOptions: [
        {label: 'default', value: 'default'},
        // {label: 'test', value: 'test'},
        // {label: 'aaa', value: 'aaa'},
        // {label: 'bbb', value: 'bbb'}
      ]
    }
  },
  methods: {
    // validateCrontabRule(rule, value, callback) {
    //   console.log("调用了 validateCrontabRule 函数"); // 添加日志来确认函数调用
    //   const crontabRegex = /^(\*|([0-5]?[0-9](-[0-5]?[0-9])?)(\/[0-9]+)?)(\s+(\*|([01]?[0-9]|2[0-3])(-[01]?[0-9]|2[0-3])?)(\/[0-9]+)?)(\s+(\*|([1-9]|[12][0-9]|3[01])(-[1-9]|[12][0-9]|3[01])?)(\/[0-9]+)?)(\s+(\*|([1-9]|1[0-2])(-[1-9]|1[0-2])?)(\/[0-9]+)?)(\s+(\*|([0-6])(-[0-6])?)(\/[0-9]+)?)$/;
    //   if (!value) {
    //     callback(new Error('请输入时间规则'));
    //   } else if (!crontabRegex.test(value)) {
    //     callback(new Error('时间规则格式不正确'));
    //   } else {
    //     callback();
    //   }
    // },

    handleRowClick(row, column, event) {
      // row: 当前点击行的数据对象
      // column: 当前点击的列数据
      // event: 原生点击事件
      console.log('点击的行数据:', row);
      console.log('点击的列:', column);
      console.log('点击事件:', event);

      // 在这里可以执行你需要的逻辑，例如存储/处理该行数据
      // this.form = row;

      getCrontabOne(row.id).then(response => {
        if (response.data.code === 1) {
          this.formData2 = response.data.result;
        }
      }).catch(err => {
        console.log(err);
      })


    },

    //提交表单（或）更新表单
    onSubmit() {
      //如果data等于编号-1则新增数据，否则就只是更新数据
      if (this.formData2.id === -1) {
        this.formData2.id = null;
        postCrontab(this.formData2).then(response => {
          if (response.data.code === 1) {
            this.refresh_self();
          } else {
            alert(response.data.msg);
          }
        }).catch(err => {
          alert("数据添加失败，请检查数据格式是否正确！");
          console.log(err);
        })

      } else { //如果this.formData2.id的值不等于-1，则只更新数据
        putCrontabOne(this.formData2.id, this.formData2).then(response => {
          if (response.data.code === 1) {
            this.refresh_self();
          }
        }).catch(err => {
          console.log(err);
        })
      }
    },

    // 将bool值转换为中文
    formatStatus(row, column, cellValue) {
      return cellValue ? '启用' : '未启用'; // 将布尔值转换为文本
    },

    // 将数组转换为字符串
    formatNamespace(row, column, cellValue) {
      return cellValue.join(', '); // 将数组转换为逗号分隔的字符串
    },

    // 获取所有的crontab规则
    listCrontab: function () {
      getCrontab().then(response => {
        if (response.data.code === 1) {
          this.ListData2 = response.data.result;
        }
      }).catch(err => {
        console.log(err)
      })
    },

    // 获取所有的namespace用来给下拉菜单使用
    listNamespace: function () {
      getNamespace().then(response => {
        if (response.data.code === 1) {
          this.namespaceOptions = response.data.result.map(item => {
            return {label: item.name, value: item.name};
          });
        }
      }).catch(err => {
        console.log(err);
      })
    },

    // 删除一条数据（当前行选中的数据）
    delete_one_data: function (row) {
      if (confirm(`确定要删除名称为 '${row.crontab_name}' 的定时任务吗？`)) {
        deleteCrontabOne(row.id).then(response => {
          if (response.data.code === 1) {
            this.refresh_self();
            alert('删除成功！');
          } else {
            this.refresh_self();
            alert('删除失败：' + response.data.message);
          }
        }).catch(err => {
          console.error('删除失败：', err);
          alert('删除失败，请检查网络或联系管理员。');
        });
      }
    },


    // 刷新自身
    refresh_self: function () {
      this.formData2 = {
        id: -1,
        crontab_name: '',
        crontab_rule: '',
        crontab_namespace: [],
        crontab_activate: false,
        crontab_content: '',
      };
      this.listCrontab();
    }
  },

  //组件开局就先获取一下crontab的所有规则
  created() {
    this.listCrontab();
  }

}
</script>

<style scoped>

</style>