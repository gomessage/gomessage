<template>
    <el-row style="margin: 20px 0" :gutter="20">
        <el-col :span="8">
            <DataFormat class="shadow"></DataFormat>
        </el-col>

        <el-col :span="8">
            <DataMap class="shadow"></DataMap>
        </el-col>

        <el-col :span="8">
            <CTemplate class="shadow"></CTemplate>
        </el-col>
    </el-row>
</template>

<script>
import DataFormat from "@/components/cCodeFormat";
import DataMap from "@/components/cConfigMap";
import CTemplate from "@/components/cTemplate";

export default {
    name: "ViewRequestData",
    components: {
        DataFormat,
        DataMap,
        CTemplate,
    },
    created() {
        //修改步骤条的值
        this.$store.commit("updateStepsActive", 1);
    }
}
</script>

<style scoped>
.shadow {
    -webkit-box-shadow: #ccc 0px 10px 10px;
    -moz-box-shadow: #ccc 0px 10px 10px;
    box-shadow: #ccc 0px 10px 10px;
    /*transform: scale(1.01, 1.01);*/
}
</style>
