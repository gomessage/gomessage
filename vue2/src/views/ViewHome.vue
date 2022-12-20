<template>
    <div>
        <el-row style="margin-top: 20px" class="shadow">
            <el-col span="10" offset="7">
                <Domain></Domain>
            </el-col>
        </el-row>
        <el-row>
            <el-col span="10" offset="7" style="margin-top: 150px">
                <el-button @click="simulation" round>模拟 AlertManager 的报警推送</el-button>
            </el-col>
        </el-row>
    </div>
</template>

<script>
// @ is an alias to /src
import Domain from "@/components/cDomain";
import {postDemoData} from '@/service/requests'

export default {
    name: 'ViewHome',
    components: {
        Domain,
    }, methods: {
        simulation: function () {
            let jsonData = require('./bbb.json');
            postDemoData(jsonData).then(response => {
                console.log(response)
                this.$message.success("模拟消息推送成功...")
            }).catch(err => {
                console.log(err)
            });
        }
    },
    created() {
        //修改步骤条的值
        this.$store.commit("updateStepsActive", 0);
    }

}
</script>
<style scoped>

</style>
