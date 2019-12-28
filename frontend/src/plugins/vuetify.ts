import Vue from 'vue';
import Vuetify, { VSnackbar, VBtn, VIcon } from 'vuetify/lib'
import VuetifyToast from "vuetify-toast-snackbar";

Vue.use(Vuetify, {
    components: {
        VSnackbar,
        VBtn,
        VIcon
    }
});
Vue.use(VuetifyToast, {
    x: 'top',
    y: 'top',
    showClose: true,
    dismissable: true,
    queueable: true,
});

export default new Vuetify({
    icons: {
        iconfont: 'md',
    }
});
