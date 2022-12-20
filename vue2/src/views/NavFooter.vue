<template>
    <el-menu id="NavFooter-Footer" class="el-menu-demo" mode="horizontal" @select="handleSelect"
             background-color="#303133" text-color="#fff" active-text-color="#ffd04b" :collapse-transition="true">

        <el-row type="flex" justify="center" id="my-row">

            <el-menu-item id="NavFooter-menu-item" disabled>


                <el-link href="https://github.com/gomessage/gomessage" type="primary">@ 2020~2022&nbsp;&nbsp;</el-link>
                <el-link href="https://opensource.org/licenses/MIT" type="primary">MIT License&nbsp;&nbsp;&nbsp;</el-link>
                <el-link href="https://github.com/gomessage/gomessage" type="primary">Github&nbsp;</el-link>
                or
                <el-link href="https://gitee.com/gomessage/gomessage" type="primary">&nbsp;Gitee&nbsp;&nbsp;</el-link>


                Give feedback and Suggestions to the author :
                <el-link href="mailto:taycc3223@gmail.com" type="primary">Email</el-link>


            </el-menu-item>

        </el-row>

    </el-menu>
</template>

<script>
// import Image from './image.vue'

export default {
    name: "NavFooter",
    data() {
        return {
            activeIndex: '1',
            activeIndex2: '1'
        };
    },
    methods: {
        handleSelect(key, keyPath) {
            console.log(key, keyPath);
        }
    },
    // components: {
    //     Image
    // }
}
</script>

<style scoped>

#NavFooter-Footer {
    padding: 0;
    margin: 0;
}

/*#NavFooter-menu-item{*/
/*    height: 60px;*/
/*}*/

</style>
