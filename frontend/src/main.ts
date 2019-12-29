import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import vuetify from './plugins/vuetify';
import i18n from "@/i18n/i18n";

import axios, {AxiosStatic} from "axios";
import auth, {AuthInterface} from './auth';
// @ts-ignore
import DatetimePicker from 'vuetify-datetime-picker'
import Lightbox from 'vue-easy-lightbox';
// @ts-ignore
import vueAwesomeCountdown from 'vue-awesome-countdown';
import moment from "moment";

Vue.use(DatetimePicker);
Vue.use(Lightbox);
Vue.use(vueAwesomeCountdown);

Vue.config.productionTip = false;

axios.defaults.baseURL = process.env.VUE_APP_API_URL;
Vue.prototype.$http = axios;
Vue.prototype.$auth = auth;
Vue.prototype.$apiUrl = process.env.VUE_APP_API_URL;

declare module 'vue/types/vue' {
    interface Vue {
        $http: AxiosStatic;
        $auth: AuthInterface;
        $apiUrl: string;
    }
}

auth.checkAuthentication();

moment.locale('pl');

new Vue({
    router,
    store,
    vuetify,
    i18n,
    render: h => h(App)
}).$mount('#app');
