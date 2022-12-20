import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        StepsActive: 0,
        DrawerStatus: false,
    },
    mutations: {
        updateStepsActive(state, num) {
            state.StepsActive = num
        },
        updateDrawerStatus(state, status) {
            state.DrawerStatus = status
        }
    },
    actions: {},
    modules: {}
})
