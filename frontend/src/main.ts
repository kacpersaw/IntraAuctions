import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';
import axios, {AxiosStatic} from "axios";
import auth, {AuthInterface} from './auth';

Vue.config.productionTip = false

axios.defaults.baseURL = process.env.VUE_APP_API_URL;
Vue.prototype.$http = axios;
declare module 'vue/types/vue' {
  interface Vue {
    $http: AxiosStatic;
  }
}

Vue.prototype.$auth = auth;
declare module 'vue/types/vue' {
  interface Vue {
    $auth: AuthInterface
  }
}

auth.checkAuthentication()

new Vue({
  router,
  store,
  vuetify,
  render: h => h(App)
}).$mount('#app')
