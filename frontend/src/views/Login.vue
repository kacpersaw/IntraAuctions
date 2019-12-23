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
                    <v-toolbar-title>Login</v-toolbar-title>
                </v-toolbar>
                <v-card-text>
                    <v-form
                    ref="form"
                    v-model="valid"
                    lazy-validation
                    >
                        <v-text-field
                                label="Username"
                                prepend-icon="account_box"
                                type="text"
                                v-model="username"
                                :rules="usernameRules"
                                required
                        />

                        <v-text-field
                                label="Password"
                                prepend-icon="lock"
                                type="password"
                                v-model="password"
                                :rules="passwordRules"
                                required
                        />
                    </v-form>
                </v-card-text>
                <v-card-actions>
                    <v-spacer/>
                    <v-btn color="primary" :disabled="!valid">Login</v-btn>
                </v-card-actions>
            </v-card>
        </v-col>
    </v-row>
</template>

<script lang="ts">
    import Vue from 'vue';

    export default Vue.extend({
        name: 'App',

        data: () => ({
            valid: true,
            username: '',
            password: '',
            usernameRules: [
                (v: string) => {
                    return !!v || 'Username is required';
                },
            ],
            passwordRules: [
                (v: string) => {
                    return !!v || 'Password is required'
                }
            ]
        }),

        methods: {
        }
    });
</script>