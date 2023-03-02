<template>
  <div>
    <br>
    <p class="authorStatement">GoMessage永久开源且免费，任何人都可以进行二次开发与封装...</p>
    <br>
    <p class="authorStatement">平时只能用业余时间开发，更新进度可能比较慢，希望小伙伴们理解万岁...</p>
    <br>
    <p class="authorStatement">
      如果您喜欢GoMessage，并且希望GoMessage能兼容对接到其它客户端，可直接联系我们进行反馈，我们会尽快满足小伙伴们的各种实际使用场景...</p>
    <br>
    <p class="authorStatement">有问题直接提，互相学习，共同进步哈~ (*^__^*)</p>
    <br>

    <el-divider content-position="center">常用联系方式</el-divider>

    <el-table :data="author">
      <el-table-column prop="name" label="名称" width="100px"></el-table-column>
      <el-table-column prop="wechat" label="微信"></el-table-column>
      <el-table-column prop="qq" label="QQ"></el-table-column>

    </el-table>
  </div>
</template>

<script>
export default {
  name: "clientOther",
  data() {
    return {
      author: [
        {
          name: "古寒飞",
          wechat: "SPE3SRU3STAY",
          qq: "237663347"
        },
        {
          name: "yesuu",
          wechat: "ACFOLPTPZXIZ",
          qq: "791581230"
        },

      ],
    }
  }
}
</script>

<style scoped>
.authorStatement {
  font-size: 14px;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  line-height: 1.7;
  color: #C0C4CC;
}
</style>
