<template>
  <el-card shadow="always">

    <!--卡片标题-->
    <div slot="header" class="clearfix">
      <span style="padding-left: 50px">分解过境数据</span>
      <el-button style="float: right; padding: 3px 0" type="text" v-on:click="getServerData">刷新</el-button>
    </div>

    <el-table
      :border="true"
      :data="configList"
      style="width: 100%">
      <el-table-column label="Key (把下面的值，粘贴到右侧模板内使用，充当占位符)" prop="key"></el-table-column>
      <el-table-column label="Value (最终发出去的消息中，就会被渲染成下列的值)" prop="value"></el-table-column>
    </el-table>
  </el-card>
</template>

<script>

import {queryFlattening, queryVars} from '@/service/requests'


export default {
  name: "cConfigMap2",
  data() {
    return {
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
    //获取服务端数据（服务端已经在后台劫持好"过路"数据了）
    getServerData: function () {
      // 这里如果使用funciton(response){}这个种模式，那么this指的就不是vue实例了，
      // 如果下面这样使用箭头函数，就可以实现上下文穿透，this就可以代指vue实例本身了，然后就有$router方法了
      queryFlattening(this.$store.getters.getNamespace, null).then(response => {
        // const jsonData = response.data.result["key_value_list"];
        const jsonData = response.data.result["key_value_map"];
        const t = response.data.result["request_time"];

        if (jsonData === null || jsonData.length === 0) {
          this.$message({
            showClose: false,
            message: "没有数据进入GoMessage服务，无法向您展示数据。"
          });
        } else {
          let tmpDataList = []
          // jsonData.forEach(data => {
          //   let tmpDict = {
          //     "key": "{{ ." + data["key"] + " }}",
          //     "value": data["value"],
          //   };
          //   //把这个临时的字典追加到tmpDataList的list中
          //   tmpDataList.push(tmpDict);
          // })

          let key_list = Object.keys(jsonData)

          key_list.forEach(key => {
            let tmpDict = {
              "key": "{{ ." + key + " }}",
              "value": jsonData[key],
            };
            tmpDataList.push(tmpDict);
          })

          // jsonData.forEach(data => {
          //   let key = Object.keys(data)[0]
          //   console.log(key)
          //   let tmpDict = {
          //     "key": "{{ ." + key + " }}",
          //     "value": jsonData[key],
          //   };
          //   tmpDataList.push(tmpDict);
          // })


          this.configList = tmpDataList;
          this.codeUpdateTime = this.dateFormat(new Date(t));
        }
      }).catch(function (error) {
        console.log(error);
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
    // this.pullMapData();
    this.getServerData();
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
