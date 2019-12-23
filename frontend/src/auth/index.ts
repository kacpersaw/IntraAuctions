import Vue from 'vue';
import axios from 'axios';
import router from '../router';
import jwt_decode from 'jwt-decode';

export interface AuthInterface {
    username: string;
    logged: boolean;
    login(context: any, data: any, redirect: any): any;
    logout(): any;
    checkAuthentication(): any;
    getAuthorizationHeader(): any;
}

export default new Vue({
    data: {
        email: '',
        username: '',
        logged: false,
    },

    methods: {
        login(context: any , data: any, redirect: any) {
            this.$http.post('/v1/auth/login', data).then((res: any) => {
                if (res.data.token) {
                    const decoded: any = jwt_decode(res.data.token);
                    this.username = decoded.username;

                    localStorage.setItem('token', res.data.token);
                    this.logged = true;

                    axios.defaults.headers.common.Authorization = this.getAuthorizationHeader();

                    if (redirect) {
                        router.push(redirect);
                    }
                }
            }).catch((e: any) => {
                if(!e.response) {
                    context.error = 'Network error';
                } else {
                    context.error = e.response.data.message;
                }
            });
        },
        logout() {
            this.logged = false;
            localStorage.removeItem('token');
            router.push('/login');
        },

        checkAuthentication() {
            const token = localStorage.getItem('token');
            if (token) {
                this.logged = true;
                const decoded: any = jwt_decode(token);
                this.username = decoded.username;
                this.email = decoded.email;
                axios.defaults.headers.common.Authorization = this.getAuthorizationHeader();
            } else {
                this.logged = false;
            }
        },

        getAuthorizationHeader() {
            return 'Bearer ' + localStorage.getItem('token');
        },
    },
});
