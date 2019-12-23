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
                    <v-toolbar-title>Aukcja #1</v-toolbar-title>
                    <v-spacer/>
                    <v-chip
                            class="ma-2"
                            color="red"
                            text-color="white"
                    >
                        Koniec za 10 dni
                        <v-icon right>mdi-clock</v-icon>
                    </v-chip>
                </v-toolbar>
                <v-container>
                    <v-row>
                        <v-col>
                            <v-list-item>
                                <v-list-item-content>
                                    <v-list-item-title class="headline">Lorem ipsum</v-list-item-title>
                                    <v-list-item-subtitle>Lorem ipsum dolor sit emet
                                    </v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                            <v-card-text>
                                Lorem ipsum dolor sit amet, consectetur adipiscing elit. In non eros
                                ullamcorper, convallis tortor aliquam, iaculis erat. Nulla commodo, ipsum
                                convallis ultrices faucibus, quam sapien placerat quam, bibendum tincidunt
                                nibh dui ac sem. Integer bibendum velit quis lacus eleifend, eget congue ex
                                interdum.
                            </v-card-text>
                        </v-col>
                        <v-col>
                            <v-row style="padding-right: 16px">
                                <v-col
                                        v-for="n in 6"
                                        :key="n"
                                        class="d-flex child-flex"
                                        cols="4"
                                >
                                    <v-card flat tile class="d-flex">
                                        <v-img
                                                :src="`https://picsum.photos/500/300?image=${n * 5 + 10}`"
                                                :lazy-src="`https://picsum.photos/10/6?image=${n * 5 + 10}`"
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
                                                                         color="grey lighten-5"></v-progress-circular>
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
                            <p class="subtitle-1 font-weight-light mb-0">Aktualna cena</p>
                            <p class="display-1 mb-0">25 843 zł</p>
                        </v-col>
                    </v-row>
                    <v-row justify="center">
                        <v-col md="4">
                            <v-text-field
                                    label="Twoja oferta"
                                    outlined
                                    suffix="zł"
                            ></v-text-field>
                        </v-col>
                        <v-col md="2">
                            <v-btn block
                                   color="primary"
                                   height="55"
                            >Licytuj</v-btn>
                        </v-col>
                    </v-row>
                </v-container>
            </v-card>
        </v-col>
    </v-row>
</template>