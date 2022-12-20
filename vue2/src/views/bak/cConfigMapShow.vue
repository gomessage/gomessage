<template>
  <el-card shadow="always">

    <!--卡片标题-->
    <div slot="header" class="clearfix">
      <span>查看模板变量</span>
      <el-button style="float: right; padding: 3px 0" type="text" v-on:click="pullMapData">拉取更新</el-button>
    </div>

    <!--Table表格-->
    <el-table
        :data="configList"
        style="width: 100%"
        border="true"
        stripe="true"
        :header-cell-style="{background:'#2f2f35',color:'#fff'}">
      <el-table-column prop="key" label="Key（自定义变量名）"></el-table-column>
      <el-table-column prop="value" label="Value（请求数据中的字段名）"></el-table-column>
    </el-table>
  </el-card>
</template>

<script>
import {queryVars} from '@/service/requests'

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
          key: '{{ .alertName }} ',
          value: 'alerts.0.labels.alertname',
        }, {
          key: '{{ .aaa }}',
          value: 'alerts.0.status',
        }, {
          key: '{{ .bbb }}',
          value: 'alerts.0.annotations.description',
        }, {
          key: '{{ .N1 }}',
          value: 'alerts.0.annotations.description',
        }, {
          key: '{{ .N2 }}',
          value: 'alerts.0.annotations.description',
        }, {
          key: '{{ .N3 }}',
          value: 'alerts.0.annotations.description',
        },

      ],
    };
  },
  computed: {},
  methods: {
    //拉取数据更新
    pullMapData: function () {
      queryVars().then(response => {
        let tmpDataList = [];
        response.data.forEach(iData => {
          let tmpDict = {
            "key": "{{ ." + iData["Key"] + " }}",
            "value": iData["Value"],
          };
          tmpDataList.push(tmpDict);
        });
        this.configList = tmpDataList;
      }).catch(err => {
        console.log(err);
      });
    }
  }, //methods结束
  created() {
    this.pullMapData();
  }
}
</script>

<style scoped>

</style>
