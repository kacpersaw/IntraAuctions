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
                        End {{timer}}
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
                                                :src="getImageUrl(img.url)"
                                                :lazy-src="getImageUrl(img.url)"
                                                aspect-ratio="1"
                                                class="grey lighten-2"
                                                @click="open(getImageUrl(img.url))"
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
                    <v-form v-model="bidValid">
                        <v-row justify="center">
                            <v-col md="4">
                                <v-text-field
                                        label="Your offer"
                                        outlined
                                        suffix="zł"
                                        type="number"
                                        step="0.1"
                                        :rules="bidRules"
                                        v-model="bid"
                                >
                                </v-text-field>
                            </v-col>
                            <v-col md="2">
                                <v-btn block
                                       color="primary"
                                       height="55"
                                       :disabled="!bidValid"
                                       @click="bidAuction()"
                                >Bid
                                </v-btn>
                            </v-col>
                        </v-row>
                    </v-form>
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
    import {Auction} from '@/types/api';

    export default Vue.extend({
        name: 'Auctions',

        data: () => ({
            auction: {} as Auction,
            img: '',
            visible: false,
            bidValid: true,
            bid: 0,
            interval: 0,
            now: new Date(),
        }),

        computed: {
            bidRules() {
                return [
                    (val: number) => !!val || 'Required',
                    (val: number) => val >= this.auction.minimal_bid || 'Minimal bid is ' + this.auction.minimal_bid
                ]
            },
            timer() {
                // @ts-ignore
                let c = this.countdown(this.now, this.auction.end_date);
                return moment(this.now).to(this.auction.end_date) + ' (' + c + ')';
            }
        },

        mounted() {
            this.getActive();

            this.interval = setInterval(() => {
                this.now = new Date()
            }, 1000)
        },

        methods: {
            getActive() {
                this.$http.get('/v1/auction/active').then((res) => {
                    this.auction = res.data;
                    this.auction.end_date = moment(res.data.end_date).toDate();
                    this.auction.start_date = moment(res.data.start_date).toDate();
                    this.bid = this.auction.minimal_bid
                })
            },
            open(url: string) {
                this.img = url;
                this.visible = true;
            },
            logout() {
                this.$auth.logout();
            },
            getImageUrl(path: string): string {
                return this.$apiUrl + '/' + path;
            },
            bidAuction() {
                this.$http.put('/v1/auction/' + this.auction.id + '/bid', {
                    bid: this.bid as number
                }).then((res) => {
                    this.$toast('Bid made successfully!', {
                        color: 'success',
                    });

                    this.getActive()
                }).catch((err) => {
                    this.$toast('Something went wrong :(', {
                        color: 'red',
                    })
                })
            },
            countdown(now: any, then: any): string {
                return 0 < then - now
                    ? moment.utc(then - now).format('HH:mm:ss')
                    : '- ' + moment.utc(now - then).format('HH:mm:ss')
            }
        },
        filters: {
            formatDate: function (value: string): string {
                if (value) {
                    return moment(value).format('DD-MM-YYYY HH:mm')
                }
                return ""
            }
        },
        beforeDestroy(): void {
            clearInterval(this.interval)
        }
    });
</script>