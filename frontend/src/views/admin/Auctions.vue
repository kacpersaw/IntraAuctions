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
                    <v-toolbar-title>{{$t('auction.list')}}</v-toolbar-title>
                    <v-spacer/>
                    <v-btn class="ma-2" tile color="green" v-if="selected.length > 0" @click="setActive()">{{$t('auction.setActive')}}</v-btn>
                    <v-btn class="ma-2" tile color="red" v-if="selected.length > 0" @click="deleteAuction()">{{$t('general.delete')}}</v-btn>
                    <v-divider vertical></v-divider>
                    <v-btn class="ma-2" tile color="blue" @click="addAuctionDialog()">{{$t('auction.add')}}</v-btn>
                </v-toolbar>
                <v-container>
                    <v-data-table
                            :headers="headers"
                            :items="auctions"
                            :items-per-page="5"
                            show-select
                            v-model="selected"
                            :single-select="true"
                            @click:row="goToAuction"
                    >
                        <template v-slot:item.start_date="{ item }">
                            {{item.start_date | formatDate}}
                        </template>
                        <template v-slot:item.end_date="{ item }">
                            {{item.end_date | formatDate}}
                        </template>
                        <template v-slot:item.active="{ item }">
                            <v-chip color="green" v-if="item.active" dark>{{$t('general.yes')}}</v-chip>
                            <v-chip color="red" v-if="!item.active" dark>{{$t('general.no')}}</v-chip>
                        </template>
                    </v-data-table>
                </v-container>
            </v-card>
        </v-col>

        <v-dialog v-model="dialog" persistent max-width="600px">
            <v-card>
                <v-card-title>
                    <span class="headline">New auction</span>
                </v-card-title>
                <v-card-text>
                    <v-container>
                        <v-row>
                            <v-col cols="12">
                                <v-text-field
                                        :label="$t('auction.title')"
                                        required
                                        v-model="form.title"
                                >
                                </v-text-field>
                            </v-col>
                            <v-col cols="12">
                                <v-text-field
                                        :label="$t('auction.shortDescription')"
                                        required
                                        v-model="form.short_description"
                                >
                                </v-text-field>
                            </v-col>
                            <v-col cols="12">
                                <v-textarea
                                        :label="$t('auction.description')"
                                        required
                                        v-model="form.description"
                                >
                                </v-textarea>
                            </v-col>
                            <v-col cols="12">
                                <v-text-field
                                        type="number"
                                        :label="$t('auction.startPrice')"
                                        required
                                        v-model="form.start_price"
                                >
                                </v-text-field>
                            </v-col>
                            <v-col cols="12">
                                <v-text-field
                                        type="number"
                                        :label="$t('auction.minimalBid')"
                                        required
                                        v-model="form.minimal_bid"
                                >
                                </v-text-field>
                            </v-col>

                            <v-col cols="12">
                                <v-datetime-picker
                                        :label="$t('auction.startDate')"
                                        v-model="form.start_date"
                                        :time-picker-props="timePickerProps"
                                >
                                </v-datetime-picker>
                            </v-col>
                            <v-col cols="12">
                                <v-datetime-picker
                                        :label="$t('auction.endDate')"
                                        v-model="form.end_date"
                                        :time-picker-props="timePickerProps"
                                >
                                </v-datetime-picker>
                            </v-col>
                            <v-col cols="12">
                                <v-file-input
                                        accept="image/*"
                                        multiple
                                        :label="$t('auction.pictures')"
                                        chips
                                        v-model="form.images">
                                </v-file-input>
                            </v-col>
                        </v-row>
                    </v-container>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="dialog = false">{{$t('general.close')}}</v-btn>
                    <v-btn color="blue darken-1" text @click="addAuction()">{{$t('general.add')}}</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-row>
</template>

<script lang="ts">
    import Vue from 'vue';
    import moment from "moment";
    import {Auction} from "@/types/api";

    export default Vue.extend({
        name: 'Auctions',

        data: () => ({
            auctions: [] as Auction[],
            dialog: false,
            selected: [] as Auction[],
            form: {
                title: '',
                short_description: '',
                description: '',
                start_price: 0,
                minimal_bid: 0,
                start_date: new Date(),
                end_date: new Date(),
                images: []
            },
            timePickerProps: {
                format: '24h'
            },
        }),

        computed: {
            headers() {
                return [
                    {
                        text: this.$t('auction.title'),
                        value: 'title',
                        align: 'left',
                    },
                    {
                        text: this.$t('auction.startDate'),
                        value: 'start_date',
                    },
                    {
                        text: this.$t('auction.endDate'),
                        value: 'end_date',
                    },
                    {
                        text: this.$t('auction.active'),
                        value: 'active'
                    }
                ]
            }
        },

        mounted(): void {
            this.listAuctions()
        },

        methods: {
            listAuctions(): void {
                this.$http.get('/v1/auction').then((res) => {
                    this.auctions = res.data;
                })
            },
            addAuctionDialog(): void {
                this.dialog = true;
                this.form = {
                    title: '',
                    short_description: '',
                    description: '',
                    start_price: 0,
                    minimal_bid: 0,
                    start_date: new Date(),
                    end_date: new Date(),
                    images: []
                };
            },
            addAuction(): void {
                let data = new FormData();
                data.set('auction', JSON.stringify({
                    title: this.form.title,
                    short_description: this.form.short_description,
                    description: this.form.description,
                    start_price: +this.form.start_price,
                    minimal_bid: +this.form.minimal_bid,
                    start_date: this.form.start_date.toISOString(),
                    end_date: this.form.end_date.toISOString(),
                }));

                this.form.images.forEach((img) => {
                    data.append('images', img)
                });

                this.$http.post('/v1/auction', data, {
                    headers: {'Content-Type': 'multipart/form-data'}
                }).then((res) => {
                    this.$toast(this.$tc('auction.toasts.created'), {
                        color: 'success',
                    });
                    this.dialog = false;
                    this.listAuctions()
                }).catch((res: any) => {
                    this.$toast(this.$tc('general.error.generic'), {
                        color: 'red',
                    })
                })
            },
            setActive() {
                let selected:any = this.selected[0];
                this.$http.put(`/v1/auction/${selected!.id}/active`).then((res) => {
                    this.$toast(this.$tc('auction.toasts.active'), {
                        color: 'success',
                    });
                    this.selected = [];
                    this.listAuctions()
                })
            },
            deleteAuction(): void {
                let selected:any = this.selected[0];
                this.$http.delete(`/v1/auction/${selected!.id}`).then((res) => {
                    this.$toast(this.$tc('auction.toasts.removed'), {
                        color: 'success',
                    });
                    this.selected = [];
                    this.listAuctions()
                })
            },
            goToAuction(item: Auction): void {
                this.$router.push("/admin/auction/" + item.id);
            }
        },
        filters: {
            formatDate(value: string): string {
                if (value) {
                    return moment(value).format('DD-MM-YYYY HH:mm')
                }
                return ""
            }
        }
    });
</script>
