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
                    <v-toolbar-title>Auction Details</v-toolbar-title>
                    <v-spacer/>
                </v-toolbar>
                <v-container>
                    <v-row>
                        <v-col cols="6">
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.title}}</v-list-item-title>
                                    <v-list-item-subtitle>Title</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.short_description}}</v-list-item-title>
                                    <v-list-item-subtitle>Short description</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.description}}</v-list-item-title>
                                    <v-list-item-subtitle>Description</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.start_price}}</v-list-item-title>
                                    <v-list-item-subtitle>Start price</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.actual_price}}</v-list-item-title>
                                    <v-list-item-subtitle>Actual price</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.minimum_bid}}</v-list-item-title>
                                    <v-list-item-subtitle>Minimum bid</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.start_date | formatDate}}</v-list-item-title>
                                    <v-list-item-subtitle>Start date</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-list-item two-line>
                                <v-list-item-content>
                                    <v-list-item-title>{{auction.end_date | formatDate}}</v-list-item-title>
                                    <v-list-item-subtitle>End date</v-list-item-subtitle>
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
                                                :src="`${apiUrl}/${img.url}`"
                                                :lazy-src="`${apiUrl}/${img.url}`"
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
                                        <v-toolbar-title>Bids</v-toolbar-title>
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

    export default Vue.extend({
        name: 'AuctionDetails',

        data: () => ({
            auction: {},
            headers: [
                {
                    text: 'Username',
                    value: 'uid',
                    align: 'left',
                },
                {
                    text: 'Bid',
                    value: 'bid',
                },
                {
                    text: 'Date',
                    value: 'created_at',
                }
            ]
        }),

        mounted() {
            this.getAuction()
        },

        computed: {
            apiUrl() {
                return this.$apiUrl;
            },
        },

        methods: {
            getAuction() {
                this.$http.get('/v1/auction/' + this.$route.params.id).then((res: any) => {
                    this.auction = res.data;
                })
            },
            goToAuctions() {
                this.$router.push("/admin/auctions");
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