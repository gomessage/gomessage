import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    StepsActive: 0, //步骤
    DrawerStatus: false, //抽屉状态
    Namespace: "default", //命名空间
    NamespaceIsRenders: false, //命名空间是否开启渲染模式
  },
  getters: {
    getNamespace: state => state.Namespace,
    getNamespaceIsRenders: state => state.NamespaceIsRenders,
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
    updateNamespaceIsRenders(state, IsRenders) {
      state.NamespaceIsRenders = IsRenders
    },
  },
  actions: {},
  modules: {}
})
