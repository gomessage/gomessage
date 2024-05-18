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
        path: '/main',
        name: '1-1',
        component: () => import("@/views/main/ViewHome.vue")
    },
    {
        path: '/renders',
        name: '2-1',
        component: () => import("@/views/main/ViewJson.vue")
    },
    {
        path: '/clients',
        name: '3-1',
        component: () => import( '@/views/main/ViewClient.vue')
    },
    {
        path: '/crontab',
        name: '4-1',
        component: () => import( '@/views/cron/ViewCrontab.vue')
    },
    {
        //登录
        path: '/login',
        name: 'login',
        component: () => import("@/views/login.vue")
    },

    // {
    //     //重定向
    //     path: '/',
    //     redirect: '/main/'
    // },
    // {
    //     //登录
    //     path: '/login',
    //     name: 'login',
    //     component: () => import("@/views/login.vue")
    // },
    // {
    //     //首页
    //     path: '/main/',
    //     component: () => import("@/views/main/main.vue"),
    //     children: [
    //         {
    //             path: '',
    //             name: '1-1',
    //             component: () => import("@/views/main/ViewHome.vue")
    //         },
    //         {
    //             path: 'renders',
    //             name: '2-1',
    //             component: () => import("@/views/main/ViewJson.vue")
    //         },
    //         {
    //             path: 'clients',
    //             name: '3-1',
    //             component: () => import( '@/views/main/ViewClient.vue')
    //         },
    //         {
    //             path: 'crontab',
    //             name: '4-1',
    //             component: () => import( '@/views/cron/ViewCrontab.vue')
    //         }
    //     ]
    // },


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
    // 这个模式可以不带#号
    // mode: 'history',
    routes
})

router.beforeEach(async (to, from, next) => {

    let isAuthenticated = false;
    let token = store.getters.getToken;
    if (token === "") {
        //路由的优先级大于App.vue的create钩子，因此不管在其它什么地方设定的往sessionStorage中写入数据；甭管其他地方有没有取出来指定的值
        //在路由这里都重新读取一遍sessionStorage，以防路由逻辑判定token===""
        //只要能成功拿到store中的信息，就能把路由跳转到指定的目的地，否则将把路由跳转到登录页面
        if (sessionStorage.getItem("store")) {
            await store.replaceState(Object.assign({}, store.state, JSON.parse(sessionStorage.getItem("store"))))
            isAuthenticated = true
        }
    } else {
        isAuthenticated = true
    }

    if (to.path !== "/login" && !isAuthenticated) {
        if (from.path !== "/login") {
            next('/login');
        } else {
            // 在 Vue Router 中，next(false) 用于取消当前的导航。这意味着当前导航将不会完成，用户将停留在原来的位置，当前路由不会发生改变。
            // 具体来说，当你在导航守卫中调用 next(false) 时，意味着你显式地指示 Vue Router 不要进行这次导航，保持用户在当前路径。如果有其他导航操作（如重定向或者条件触发的导航），你可以利用这个机制控制导航流程。
            next(false);
        }
    } else {
        next();
    }
})

export default router
