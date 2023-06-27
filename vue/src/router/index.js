import Vue from 'vue'
import VueRouter from 'vue-router'
import store from "@/store";


Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        redirect: '/main/'
    },
    {
        path: '/login',
        name: 'login',
        component: () => import("@/views/login.vue")
    },
    {
        path: '/main/',
        component: () => import("@/views/main.vue"),
        children: [
            {
                path: '',
                name: '1-1',
                component: () => import("@/views/ViewHome.vue")
            },
            {
                path: 'renders',
                name: '2-1',
                component: () => import("@/views/ViewJson.vue")
            },
            {
                path: 'clients',
                name: '3-1',
                component: () => import( '@/views/ViewClient.vue')
            }
        ]
    },


    // {
    //   //数据格式化页面
    //   path: '/data_format',
    //   name: 'DataFormat',
    //   component: () => import( '@/components/cCodeFormat.vue')
    // },
    // {
    //     //消息模板页面
    //     path: '/data_template',
    //     name: 'MessageTemplate',
    //     component: () => import( '../views/ViewTemplate.vue')
    // },

]

const router = new VueRouter({
    // mode: 'history', //这个模式可以不带#号
    routes
})

router.beforeEach((to, form, next) => {
    if (to.path === "/login") {
        return next();
    } else {
        let token = store.getters.getToken;
        if (token === "") {
            return next('/login')
        } else {
            return next()
        }
    }
})

export default router
