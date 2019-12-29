import Vue from 'vue';
import Vuetify, { VSnackbar, VBtn, VIcon } from 'vuetify/lib'
import VuetifyToast from "vuetify-toast-snackbar";
import pl from 'vuetify/src/locale/pl';
import en from 'vuetify/src/locale/en';

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
    },
    lang: {
        locales: {
            en,pl
        },
        current: 'pl'
    }

});
