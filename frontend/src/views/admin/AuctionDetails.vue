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
                    <v-btn icon class="hidden-xs-only" @click="goToAuctions()">
                        <v-icon>arrow_back</v-icon>
                    </v-btn>
                    <v-toolbar-title>{{$t('auction.details')}}</v-toolbar-title>
                    <v-spacer/>
                </v-toolbar>
                <v-container>
                    <v-row>
                        <v-col cols="6">
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.title}}</v-list-item-title>
                                    <v-list-item-subtitle>{{$t('auction.title')}}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.short_description}}</v-list-item-title>
                                    <v-list-item-subtitle>{{$t('auction.shortDescription')}}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <p>{{auction.description}}</p>
                                    <v-list-item-subtitle>{{$t('auction.description')}}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.start_price}}</v-list-item-title>
                                    <v-list-item-subtitle>{{$t('auction.startPrice')}}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.actual_price}}</v-list-item-title>
                                    <v-list-item-subtitle>{{$t('auction.actualPrice')}}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.minimal_bid}}</v-list-item-title>
                                    <v-list-item-subtitle>{{$t('auction.minimalBid')}}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.start_date | formatDate}}</v-list-item-title>
                                    <v-list-item-subtitle>{{$t('auction.startDate')}}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.end_date | formatDate}}</v-list-item-title>
                                    <v-list-item-subtitle>{{$t('auction.endDate')}}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                        </v-col>
                        <v-col cols="6">
                            <v-row style="padding-right: 16px">
                                <v-col
                                        v-for="img in auction.images"
                                        :key="img.id"
                                        class="d-flex child-flex"
                                        cols="3"
                                >
                                    <v-card flat tile class="d-flex">
                                        <v-img
                                                :src="getImageUrl(img.url)"
                                                :lazy-src="getImageUrl(img.url)"
                                                aspect-ratio="1"
                                                class="grey lighten-2"
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
                    <v-row>
                        <v-col cols="12">
                            <v-data-table
                                    :headers="headers"
                                    :items="auction.bids"
                                    :items-per-page="10"
                            >
                                <template v-slot:top>
                                    <v-toolbar flat color="white">
                                        <v-toolbar-title>{{$t('bid.bids')}}</v-toolbar-title>
                                    </v-toolbar>
                                </template>
                                <template v-slot:item.created_at="{ item }">
                                    {{item.created_at | formatDate}}
                                </template>
                            </v-data-table>
                        </v-col>
                    </v-row>
                </v-container>
            </v-card>
        </v-col>
    </v-row>
</template>

<script lang="ts">
    import Vue from 'vue';
    import moment from "moment";
    import {Auction} from "@/types/api";

    export default Vue.extend({
        name: 'AuctionDetails',

        data: () => ({
            auction: {} as Auction,
        }),

        computed: {
            headers() {
                return [
                    {
                        text: this.$t('auth.username'),
                        value: 'uid',
                        align: 'left',
                    },
                    {
                        text: this.$t('bid.bid'),
                        value: 'bid',
                    },
                    {
                        text: this.$t('general.date'),
                        value: 'created_at',
                    }
                ]
            }
        },

        mounted() {
            this.getAuction()
        },

        methods: {
            getAuction(): void {
                this.$http.get('/v1/auction/' + this.$route.params.id).then((res) => {
                    this.auction = res.data;
                })
            },
            goToAuctions(): void {
                this.$router.push("/admin/auctions");
            },
            getImageUrl(path: string): string {
                return this.$apiUrl + '/' + path;
            },
        },
        filters: {
            formatDate(value: string): string {
                if (value) {
                    return moment(value).format('DD-MM-YYYY HH:mm:ss')
                }
                return ""
            }
        }
    });
</script>