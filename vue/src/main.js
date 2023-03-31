import Vue from 'vue'
import './plugins/axios'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'

// import * as Sentry from '@sentry/browser';
// import * as Integrations from '@sentry/integrations';
//
// Sentry.init({
//   dsn: 'http://0b1fbfe6d1e041d6b66c85650f80ca5e@172.20.3.195:9000/3',
//   // integrations: [new Integrations.Vue({Vue, attachProps: true})],
// });

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')

//定义全局函数：dateFormat
Vue.prototype.dateFormat = function (date) {
  console.log(date)
  const daters = date
  if (daters != null) {
    const dateMat = new Date(daters);
    const year = dateMat.getFullYear();
    const month = dateMat.getMonth() + 1;
    const day = dateMat.getDate();
    const hh = dateMat.getHours();
    const mm = dateMat.getMinutes();
    const ss = dateMat.getSeconds();
    return year + "-" + (month < 10 ? "0" : "") + month + "-" + (day < 10 ? "0" : "") + day + (hh < 10 ? "0" : "") + " " + hh + ":" + (mm < 10 ? "0" : "") + mm + ":" + (ss < 10 ? "0" : "") + ss;
  }
}
