<template>
    <v-app>
        <v-content>
            <v-toolbar v-if="admin">
                <v-toolbar-title>{{$t('admin.title')}}</v-toolbar-title>

                <v-spacer></v-spacer>

                <v-toolbar-items>
                    <v-btn text to="/admin/auctions">{{$t('auction.auctions')}}</v-btn>
                </v-toolbar-items>
            </v-toolbar>
            <v-container
                    fluid
            >
                <router-view></router-view>
            </v-container>
        </v-content>
    </v-app>
</template>

<script lang="ts">
    import Vue from 'vue';
    import moment from "moment";

    export default Vue.extend({
        name: 'App',

        data: () => ({
            //
        }),

        computed: {
            loggedIn() {
                return this.$auth.logged;
            },
            admin() {
                return this.$auth.admin;
            }
        }
    });
</script>
