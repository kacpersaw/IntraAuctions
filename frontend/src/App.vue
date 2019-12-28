<template>
    <v-app>
        <v-content>
            <v-toolbar v-if="admin">
                <v-toolbar-title>Administration</v-toolbar-title>

                <v-spacer></v-spacer>

                <v-toolbar-items>
                    <v-btn text to="/admin/auctions">Auctions</v-btn>
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
