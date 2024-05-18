<template>
  <!--  布局撑满el-main容器 -->
  <el-row style="top: 0;left: 0;right: 0;bottom: 0;position: absolute; display: flex;">

    <!--    左侧col-->
    <el-col :span="12">

      <el-card style="margin: 20px;padding: 10px;text-align: left;">

        <div slot="header" class="clearfix">
          <span>发送定时消息（Crontab类型的时间规则定义）</span>
        </div>

        <el-form ref="form" :model="form" label-width="90px">
          <el-form-item label="名称：">
            <el-input v-model="form.crontab_name" placeholder="下班打卡提醒"></el-input>
          </el-form-item>

          <el-form-item label="时间规则：" >
            <el-input v-model="form.crontab_rule" placeholder="*/5  *  *  *  * （分、时、日、月、周）"></el-input>
          </el-form-item>

<!--          <el-form-item label="发送通道：">-->
<!--            <el-checkbox-group v-model="form.crontab_namespace">-->
<!--              <el-checkbox v-for="item in namespaceOptions" :label="item.label" :key="item.label" name="type"></el-checkbox>-->
<!--            </el-checkbox-group>-->
<!--          </el-form-item>-->

          <el-form-item label="发送通道：">
            <el-select
                v-model="form.crontab_namespace"
                multiple
                placeholder="请选择"
                style="width: 70%;">
              <el-option
                  v-for="item in namespaceOptions"
                  :key="item.label"
                  :label="item.label"
                  :value="item.value">
              </el-option>
            </el-select>
          </el-form-item>

          <el-form-item label="是否启用：">
            <el-switch v-model="form.crontab_activate"></el-switch>
          </el-form-item>


          <el-form-item label="消息内容：">
            <el-input
                type="textarea"
                v-model="form.crontab_content"
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
    <!--    右侧col-->
    <el-col :span="12">
      <el-card style="margin: 20px;">
        <el-table
            :data="tableData"
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

        </el-table>
      </el-card>
    </el-col>
  </el-row>
</template>


<script>
export default {
  name: "ViewCrontab",
  data() {
    return {
      tableData: [{
        crontab_name: '提醒小伙伴下班打卡',
        crontab_rule: '*/5 * 2 * *',
        crontab_namespace: ["default","ttpai","alarmdog"],
        crontab_activate: false,
        crontab_content: '113123123123123123123113123123123123123123113123123123123123123113123123123123123123113123123123123123123113123123123123123123',
      }, {
        crontab_name: '002',
        crontab_rule: '*/4 * * * *',
        crontab_namespace: ["default","alarmdog"],
        crontab_activate: false,
        crontab_content: '113123123123123123123113123123123123123123113123123123123123113123123123123123123113123123123123123123113123123123123123123',
      }, {
        crontab_name: '003',
        crontab_rule: '*/5 * 2 * *',
        crontab_namespace: ["alarmdog"],
        crontab_activate: true,
        crontab_content: '113123123123123123123113123123123123123123113123123123123123113123123123123123123113123123123123123123113123123123123123123',
      }, {
        crontab_name: '123',
        crontab_rule: '5 * * * *',
        crontab_namespace: ["default","ttpai","alarmdog","aaaa","bbb"],
        crontab_activate: true,
        crontab_content: '113123123123123123123113123123123123123123113123123123123123113123123123123123123113123123123123123123113123123123123123123',
      }],
      form: {
        crontab_name: '',
        crontab_rule: '',
        crontab_namespace: ["default","ttpai","alarmdog"],
        crontab_activate: false,
        crontab_content: '',
      },
      namespaceOptions: [
        { label: 'default', value: 'default' },
        { label: 'ttpai', value: 'ttpai' },
        { label: 'alarmdog', value: 'alarmdog' },
        { label: 'aaaa', value: 'aaaa' },
        { label: 'bbb', value: 'bbb' }
      ]
    }
  },
  methods: {
    handleRowClick(row, column, event) {
      // row: 当前点击行的数据对象
      // column: 当前点击的列数据
      // event: 原生点击事件
      console.log('点击的行数据:', row);
      console.log('点击的列:', column);
      console.log('点击事件:', event);
      // 在这里可以执行你需要的逻辑，例如存储/处理该行数据
      this.form = row;
    },
    onSubmit() {
      console.log('表单数据:', this.form);
      // 这里可以添加更多处理表单提交的逻辑
    },
    formatStatus(row, column, cellValue) {
      // 将布尔值转换为文本
      return cellValue ? '启用' : '未启用';
    },
    formatNamespace(row, column, cellValue) {
      // 将数组转换为逗号分隔的字符串
      return cellValue.join(', ');
    }
  }
}
</script>

<style scoped>

</style>