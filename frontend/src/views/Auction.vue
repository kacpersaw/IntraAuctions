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
                    <v-toolbar-title>Auction</v-toolbar-title>
                    <v-spacer/>
                    <v-chip
                            class="ma-2"
                            color="red"
                            text-color="white"
                    >
                        End within
                        <v-icon right>mdi-clock</v-icon>
                    </v-chip>
                    <v-menu bottom left>
                        <template v-slot:activator="{ on }">
                            <v-btn
                                    dark
                                    icon
                                    v-on="on"
                            >
                                <v-icon>more_vert</v-icon>
                            </v-btn>
                        </template>

                        <v-list>
                            <v-list-item @click="logout()">
                                <v-list-item-title>Log out</v-list-item-title>
                            </v-list-item>
                        </v-list>
                    </v-menu>
                </v-toolbar>
                <v-container>
                    <v-row>
                        <v-col cols="6">
                            <v-list-item>
                                <v-list-item-content>
                                    <v-list-item-title class="headline">{{auction.title}}</v-list-item-title>
                                    <v-list-item-subtitle>{{auction.short_description}}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-card-text>
                                {{auction.description}}
                            </v-card-text>
                        </v-col>
                        <v-col cols="6">
                            <v-row>
                                <v-col
                                        v-for="img in auction.images"
                                        :key="img.id"
                                        class="d-flex child-flex"
                                        cols="4"
                                >
                                    <v-card flat tile class="d-flex">
                                        <v-img
                                                :src="`${apiUrl}/${img.url}`"
                                                :lazy-src="`${apiUrl}/${img.url}`"
                                                aspect-ratio="1"
                                                class="grey lighten-2"
                                                @click="open(`${apiUrl}/${img.url}`)"
                                                style="cursor: pointer;"
                                        >
                                            <template v-slot:placeholder>
                                                <v-row
                                                        class="fill-height ma-0"
                                                        align="center"
                                                        justify="center"
                                                >
                                                    <v-progress-circular indeterminate
                                                                         color="grey lighten-5">
                                                    </v-progress-circular>
                                                </v-row>
                                            </template>
                                        </v-img>
                                    </v-card>
                                </v-col>
                            </v-row>
                        </v-col>
                    </v-row>
                    <v-divider></v-divider>
                    <v-row justify="center">
                        <v-col align="center">
                            <p class="subtitle-1 font-weight-light mb-0">Actual price</p>
                            <p class="display-1 mb-0">{{auction.actual_price}} zł</p>
                        </v-col>
                    </v-row>
                    <v-row justify="center">
                        <v-col md="4">
                            <v-text-field
                                    label="Your offer"
                                    outlined
                                    suffix="zł"
                            ></v-text-field>
                        </v-col>
                        <v-col md="2">
                            <v-btn block
                                   color="primary"
                                   height="55"
                            >Bid</v-btn>
                        </v-col>
                    </v-row>
                </v-container>
            </v-card>
        </v-col>
        <vue-easy-lightbox
                moveDisabled
                :visible="visible"
                :index="0"
                :imgs="img"
                @hide="visible = false"
        >
        </vue-easy-lightbox>
    </v-row>
</template>

<script lang="ts">
    import Vue from 'vue';
    import moment from "moment";

    export default Vue.extend({
        name: 'Auctions',

        data: () => ({
            auction: {},
            img: 'http://localhost:9090/images/1VXURDlR25r9AcXRc5SWjVJlzOg.jpg',
            visible: false,
        }),

        computed: {
            apiUrl() {
                return this.$apiUrl;
            },
        },

        mounted() {
            this.getActive()
        },

        methods: {
            getActive() {
                this.$http.get('/v1/auction/active').then((res: any) => {
                    this.auction = res.data;
                })
            },
            open(url: string) {
                // eslint-disable-next-line no-console
                console.log(url);
                this.img = url;
                this.visible = true;
            },
            logout() {
                this.$auth.logout();
            }
        },
        filters: {
            formatDate: function (value: string) {
                if (value) {
                    return moment(value).format('DD-MM-YYYY HH:mm')
                }
            }
        }
    });
</script>