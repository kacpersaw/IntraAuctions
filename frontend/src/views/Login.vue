<template>
    <v-row
            align="center"
            justify="center"
    >
        <v-col
                cols="12"
                sm="8"
                md="8"
        >
            <v-card class="elevation-12">
                <v-toolbar
                        color="primary"
                        dark
                        flat
                >
                    <v-toolbar-title>{{ $t("login") }}</v-toolbar-title>
                </v-toolbar>
                <v-card-text>
                    <v-alert
                            :value="error !== ''"
                            type="error"
                    >
                        {{error}}
                    </v-alert>

                    <v-form
                            ref="form"
                            v-model="valid"
                            lazy-validation
                            @keyup.native.enter="valid && login()"
                    >
                        <v-text-field
                                :label="$t('username')"
                                prepend-icon="account_box"
                                type="text"
                                v-model="form.username"
                                :rules="usernameRules"
                                required
                        />

                        <v-text-field
                                :label="$t('password')"
                                prepend-icon="lock"
                                type="password"
                                v-model="form.password"
                                :rules="passwordRules"
                                required
                        />
                    </v-form>
                </v-card-text>
                <v-card-actions>
                    <v-spacer/>
                    <v-btn color="primary" :disabled="!valid" @click="login()">{{$t('loginButton')}}</v-btn>
                </v-card-actions>
            </v-card>
        </v-col>
    </v-row>
</template>

<script lang="ts">
    import Vue from 'vue';

    export default Vue.extend({
        name: 'Login',

        data: () => ({
            valid: true,
            form: {
                username: '',
                password: '',
            },
            error: '',
        }),

        computed: {
            usernameRules() {
                return [
                    (v: string) => {
                        return !!v || this.$t('usernameRequired');
                    },
                ]
            },
            passwordRules() {
                return [
                    (v: string) => {
                        return !!v || this.$t('passwordRequired')
                    }
                ]
            }
        },

        methods: {
            login() {
                this.error = '';
                this.$auth.login(this, this.form, "/")
            }
        }
    });
</script>