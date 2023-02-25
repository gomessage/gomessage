<template>
  <el-card shadow="always">

    <!--卡片标题-->
    <div slot="header" class="clearfix">
      <span style="padding-left: 50px">用户变量</span>
      <el-button style="float: right; padding: 3px 0" type="text" v-on:click="post_data">保存变量</el-button>
    </div>

    <!--        <el-row style="color: #C0C4CC;font-size: 12px;line-height: 1.7;padding-bottom: 20px">-->
    <!--            <el-col span="24" style="text-align: right">查看文档</el-col>-->
    <!--        </el-row>-->

    <!--Table表格-->
    <el-table
        :border="true"
        :data="configList"
        style="width: 100%">
      <el-table-column label="Key（自定义变量名）" prop="key"></el-table-column>
      <el-table-column label="Value（原始数据中的索引）" prop="value"></el-table-column>
      <el-table-column fixed="right" label="操作" width="100">
        <template slot-scope="scope">
          <el-button size="small" type="text" @click.native.prevent="deleteRow(scope.$index, configList)">移除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <br>

    <!--输入框，是从最初的（可增加数量的输入框）转变过来的-->
    <div v-for="(list, index) in mapList2" :key="index">
      <!-- 填写服务器地址，col之间的间隔为20 -->
      <el-row :gutter="20">

        <!--Key-->
        <el-col :span="12">
          <div class="DataMapGridContent">
            <el-input v-model="list.mapKey" placeholder="纯字符串，不能有符号">
              <template slot="prepend">
                Key:
              </template>
            </el-input>
          </div>
        </el-col>

        <!--Value-->
        <el-col :span="12">
          <div class="DataMapGridContent">
            <el-input v-model="list.mapValue" placeholder="Json索引路径">
              <template slot="prepend">
                Value:
              </template>
            </el-input>
          </div>
        </el-col>
      </el-row>
    </div>

    <br>

    <!--添加新的输入框（+号按钮）-->
    <el-button size="mini" type="primary" v-on:click="addTableData">新增变量</el-button>

  </el-card>
</template>

<script>

import {postVars, queryVars} from '@/service/requests'

export default {
  name: "cConfigMap",
  data() {
    return {
      mapList2: [
        {
          mapKey: '',
          mapValue: ''
        }
      ],
      configList: [
        {
          key: '{{ .N1 }} ',
          value: 'alerts.#.labels.alertname',
        }, {
          key: '{{ .N2 }}',
          value: 'alerts.#.labels.severity',
        }, {
          key: '{{ .N3 }}',
          value: 'alerts.#.labels.hostname',
        }, {
          key: '{{ .N4 }}',
          value: 'alerts.#.labels.ping',
        }, {
          key: '{{ .N5 }}',
          value: 'alerts.#.annotations.description',
        }, {
          key: '{{ .N6 }}',
          value: 'status',
        }, {
          key: '{{ .N7 }}',
          value: 'alerts.#.startsAt',
        },
        {
          key: '{{ .N8 }}',
          value: 'alerts.#.endsAt',
        },

      ],
    };
  },
  computed: {},
  methods: {
    // 添加tableData的一行数据（和上面的添加输入框不是一回事）
    addTableData: function () {
      let mapKey = this.mapList2[0].mapKey;
      let mapValue = this.mapList2[0].mapValue;

      if (mapKey.length === 0 || mapValue.length === 0) {
        this.$message.error('输入框不能为空...');
      } else {
        let tableRow = {
          key: "{{ ." + mapKey + " }}",
          value: mapValue
        };
        this.configList.push(tableRow);
        this.mapList2[0].mapKey = "";
        this.mapList2[0].mapValue = "";

        this.$message.success("添加成功...");
        this.post_data()
      }
    },

    // 删除一行数据
    deleteRow(index, rows) {
      rows.splice(index, 1);
      this.post_data()
    },

    // 提交数据
    post_data: function () {
      const dataList = [];
      for (let i = 0; i < this.configList.length; i++) {
        let mapKey = this.configList[i].key;
        let mapKey2 = mapKey.split(" ")[1].split(".")[1]
        let mapValue = this.configList[i].value;

        //封装到临时字典中
        let newKeyValue = {};
        newKeyValue[mapKey2] = mapValue;

        //添加到List中
        dataList.push(newKeyValue);
      }

      postVars(this.$store.getters.getNamespace, {"key_value_list": dataList}).then(response => {
        console.log(response.data);
        this.$message.success("数据库更新成功...")
      }).catch(err => {
        console.log(err);
      });
    },

    //拉取后端数据库中的数据
    pullMapData: function () {
      queryVars(this.$store.getters.getNamespace, null).then(response => {
        //声明一个空的list便于下文存放数据
        let tmpDataList = [];
        //判断response的长度是否为0，如果为0，则代表数据库中没有查询到数据
        if (response.data.result.length === 0) {
          console.log("数据库里没有数据")
        } else {
          //遍历response中的数据
          response.data.result.forEach(iData => {
            //声明一个临时的字典
            let tmpDict = {
              "key": "{{ ." + iData["key"] + " }}",
              "value": iData["value"],
            };
            //把这个临时的字典追加到tmpDataList的list中
            tmpDataList.push(tmpDict);
          });
          //把tmpDataList赋值给该组件data中的configList属性
          this.configList = tmpDataList;
        }
      }).catch(err => {
        console.log(err);
      });
    }
  },

  //created生命周期
  created() {
    this.pullMapData();
  }
}
</script>

<style scoped>
/*#DataMapCard {*/
/*    width: 100%;*/
/*}*/


#subButton {
  width: 100%;
}


/*.box-card {*/
/*    width: 45rem;*/
/*    !* margin-top: 1.5rem; *!*/
/*    !* margin-left: 1.5rem; *!*/
/*}*/

.DataMapGridContent {
  /*圆角*/
  border-radius: 4px;
  /*最小高度*/
  min-height: 36px;
}


</style>
