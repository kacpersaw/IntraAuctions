import VueI18n from "vue-i18n";
import en from "@/i18n/en.json";
import pl from "@/i18n/pl.json";
import Vue from 'vue';

const messages = {
    en: en,
    pl: pl
};

Vue.use(VueI18n);

export default new VueI18n({
    locale: 'pl',
    messages
});
