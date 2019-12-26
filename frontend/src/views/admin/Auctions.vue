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
                    <v-toolbar-title>Auction list</v-toolbar-title>
                    <v-spacer/>
                    <v-btn text v-if="selected.length > 0" @click="setActive()">Set active</v-btn>
                    <v-btn text @click="addAuctionDialog()">Add auction</v-btn>
                </v-toolbar>
                <v-container>
                    <v-data-table
                            :headers="headers"
                            :items="auctions"
                            :items-per-page="5"
                            show-select
                            v-model="selected"
                            :single-select="true"
                    >
                        <template v-slot:item.start_date="{ item }">
                            {{item.start_date | formatDate}}
                        </template>
                        <template v-slot:item.end_date="{ item }">
                            {{item.end_date | formatDate}}
                        </template>
                        <template v-slot:item.active="{ item }">
                            <v-chip color="green" v-if="item.active" dark>Yes</v-chip>
                            <v-chip color="red" v-if="!item.active" dark>No</v-chip>
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
                        <v-alert
                                :value="error !== ''"
                                type="error"
                        >
                            {{error}}
                        </v-alert>
                        <v-row>
                            <v-col cols="12">
                                <v-text-field
                                        label="Title"
                                        required
                                        v-model="form.title"
                                >
                                </v-text-field>
                            </v-col>
                            <v-col cols="12">
                                <v-text-field
                                        label="Short description"
                                        required
                                        v-model="form.short_description"
                                >
                                </v-text-field>
                            </v-col>
                            <v-col cols="12">
                                <v-textarea
                                        label="Description"
                                        required
                                        v-model="form.description"
                                >
                                </v-textarea>
                            </v-col>
                            <v-col cols="12">
                                <v-text-field
                                        type="number"
                                        label="Start price"
                                        required
                                        v-model="form.start_price"
                                >
                                </v-text-field>
                            </v-col>
                            <v-col cols="12">
                                <v-text-field
                                        type="number"
                                        label="Minimum bid"
                                        required
                                        v-model="form.minimum_bid"
                                >
                                </v-text-field>
                            </v-col>

                            <v-col cols="12">
                                <v-datetime-picker
                                        v-model="form.start_date"
                                        :time-picker-props="timePickerProps"
                                >
                                </v-datetime-picker>
                            </v-col>
                            <v-col cols="12">
                                <v-datetime-picker
                                        v-model="form.end_date"
                                        :time-picker-props="timePickerProps"
                                >
                                </v-datetime-picker>
                            </v-col>
                            <v-col cols="12">
                                <v-file-input
                                        accept="image/*"
                                        multiple
                                        label="Pictures"
                                        chips
                                        v-model="form.images">
                                </v-file-input>
                            </v-col>
                        </v-row>
                    </v-container>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="dialog = false">Close</v-btn>
                    <v-btn color="blue darken-1" text @click="addAuction()">Add</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-row>
</template>

<script lang="ts">
    import Vue from 'vue';
    import moment from "moment";

    export default Vue.extend({
        name: 'Auctions',

        data: () => ({
            auctions: [],
            dialog: false,
            selected: [],
            error: '',
            form: {
                title: '',
                short_description: '',
                description: '',
                start_price: 0,
                minimum_bid: 0,
                start_date: new Date(),
                end_date: new Date(),
                images: []
            },
            timePickerProps: {
                format: '24h'
            },
            headers: [
                {
                    text: '#',
                    align: 'left',
                    sortable: false,
                    value: 'id'
                },
                {
                    text: 'Title',
                    value: 'title',
                },
                {
                    text: 'Start date',
                    value: 'start_date',
                },
                {
                    text: 'End date',
                    value: 'end_date',
                },
                {
                    text: 'Active',
                    value: 'active'
                }
            ]
        }),

        mounted() {
            this.list()
        },

        methods: {
            list() {
                this.$http.get('/v1/auction').then((res: any) => {
                    this.auctions = res.data;
                })
            },
            addAuctionDialog() {
                this.dialog = true;
                this.form = {
                    title: '',
                    short_description: '',
                    description: '',
                    start_price: 0,
                    minimum_bid: 0,
                    start_date: new Date(),
                    end_date: new Date(),
                    images: []
                };
            },
            addAuction() {
                let data = new FormData();
                data.set('auction', JSON.stringify({
                    title: this.form.title,
                    short_description: this.form.short_description,
                    description: this.form.description,
                    start_price: +this.form.start_price,
                    minimum_bid: +this.form.minimum_bid,
                    start_date: this.form.start_date.toISOString(),
                    end_date: this.form.end_date.toISOString(),
                }));

                this.form.images.forEach((img) => {
                    data.append('images', img)
                });

                this.$http.post('/v1/auction', data, {
                    headers: {'Content-Type': 'multipart/form-data'}
                }).then((res: any) => {
                    this.dialog = false;
                    this.list()
                }).catch((res: any) => {
                    this.error = 'Invalid error occurred. Try again.'
                })
            },
            setActive() {
                let selected:any = this.selected[0];
                this.$http.put(`/v1/auction/${selected!.id}/active`).then((res) => {
                    this.list()
                })
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