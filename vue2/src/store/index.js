import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        StepsActive: 0, //步骤
        DrawerStatus: false, //抽屉状态
        Namespace: "default", //命名空间
    },
    getters: {
        getNamespace: state => state.Namespace
    },
    mutations: {
        updateStepsActive(state, num) {
            state.StepsActive = num
        },
        updateDrawerStatus(state, status) {
            state.DrawerStatus = status
        },
        updateNamespace(state, namespace) {
            state.Namespace = namespace
        },
    },
    actions: {},
    modules: {}
})
