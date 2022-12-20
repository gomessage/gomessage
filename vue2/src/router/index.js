import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/ViewHome'


Vue.use(VueRouter)

const routes = [
    {
        //首页
        path: '/',
        name: '1-1',
        component: Home
    },
    {
        //数据格式化页面
        path: '/data_format',
        name: 'DataFormat',
        component: () => import( '@/components/cCodeFormat.vue')
    },
    // {
    //     //消息模板页面
    //     path: '/data_template',
    //     name: 'MessageTemplate',
    //     component: () => import( '../views/ViewTemplate.vue')
    // },
    {
        //客户端配置页面
        path: '/request_data',
        name: '2-1',
        component: () => import( '../views/ViewRequestData.vue')
    },
    {
        //客户端配置页面
        path: '/data_client',
        name: '3-1',
        component: () => import( '../views/ViewClient.vue')
    }
]

const router = new VueRouter({
    // mode: 'history', //这个模式可以不带#号
    routes
})

export default router
