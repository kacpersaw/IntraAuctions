import VueI18n from "vue-i18n";
import en from "@/i18n/en.json";
import Vue from 'vue';

const messages = {
    en: en
};

Vue.use(VueI18n)

export default new VueI18n({
    locale: 'en',
    messages
});
