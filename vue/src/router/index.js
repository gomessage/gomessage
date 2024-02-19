import Vue from 'vue'
import VueRouter from 'vue-router'
import store from "@/store";


Vue.use(VueRouter)

const routes = [
    {
        //重定向
        path: '/',
        redirect: '/main/'
    },
    {
        //登录
        path: '/login',
        name: 'login',
        component: () => import("@/views/login.vue")
    },
    {
        //首页
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
    //这个模式可以不带#号
    // mode: 'history',
    routes
})

router.beforeEach((to, form, next) => {
    //如果是去登录页面，则直接往下走
    if (to.path === "/login") {
        return next();
    } else {
        let token = store.getters.getToken;
        if (token === "") {
            //路由的优先级大于App.vue的create钩子，因此不管在其它什么地方设定的往sessionStorage中写入数据；甭管其他地方有没有取出来指定的值
            //在路由这里都重新读取一遍sessionStorage，以防路由逻辑判定token===""
            //只要能成功拿到store中的信息，就能把路由跳转到指定的目的地，否则将把路由跳转到登录页面
            if (sessionStorage.getItem("store")) {
                store.replaceState(Object.assign({}, store.state, JSON.parse(sessionStorage.getItem("store"))))
                return next()
            } else {
                return next('/login')
            }

        } else {
            return next()
        }
    }
})

export default router
