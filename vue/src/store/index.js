import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
    state: {
        StepsActive: 0, //步骤
        DrawerStatus: false, //抽屉状态
        Namespace: "default", //命名空间，默认命名空间为default
        // NamespaceIsRenders: false, //命名空间是否开启渲染模式
        NamespaceInfo: {}, //命名空间的所有信息
        Token: "",
        UserID: "",
    },
    getters: {
        getNamespace: state => state.Namespace,
        // getNamespaceIsRenders: state => state.NamespaceIsRenders,
        getNamespaceInfo: state => state.NamespaceInfo,
        getToken: state => state.Token,
        getUserID: state => state.UserID,
    },
    mutations: {
        updateStepsActive(state, num) {
            state.StepsActive = num
        },
        updateToken(state, token) {
            state.Token = token
        },
        updateDrawerStatus(state, status) {
            state.DrawerStatus = status
        },
        updateNamespace(state, namespace) {
            state.Namespace = namespace
        },
        // updateNamespaceIsRenders(state, IsRenders) {
        //   state.NamespaceIsRenders = IsRenders
        // },
        updateNamespaceInfo(state, namespaceInfo) {
            state.NamespaceInfo = namespaceInfo
        },
        updateUserID(state, userID) {
            state.UserID = userID
        }
    },
    actions: {},
    modules: {}
})
