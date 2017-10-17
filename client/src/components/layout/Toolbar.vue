<template>
<div>
<v-navigation-drawer v-model="drawer" temporary light>
  <v-toolbar color="accent" flat>
    <v-list>
      <v-list-tile>
        <v-list-tile-title class="white--text title">
          AlphaBets
        </v-list-tile-title>
      </v-list-tile>
    </v-list>
  </v-toolbar>
  <v-divider></v-divider>
  <v-list>
    <v-list-tile href="#/mybets">
      <v-list-tile-title class="primary--text link">My Bets</v-list-tile-title>
    </v-list-tile>
    <v-list-tile href="#/">
      <v-list-tile-title class="accent--text link">Fixtures</v-list-tile-title>
    </v-list-tile>
    <v-list-tile href="#/results">
      <v-list-tile-title class="primary--text link">Results</v-list-tile-title>
    </v-list-tile>
    <v-list-tile href="#/standings">
      <v-list-tile-title class="primary--text link">Standings</v-list-tile-title>
    </v-list-tile>
  </v-list>
</v-navigation-drawer>
<v-toolbar dense color="transparent" class="elevation-0 abtoolbar">
  <v-toolbar-title @click="drawer = !drawer" app>
    <!-- <a href="#/" style="text-decoration: none"> -->
    <img style="height: 35px; padding: 5px 0 0 5px; cursor: pointer;" src="../../assets/bet.animated.svg">
    <!-- </a> -->
  </v-toolbar-title>
  <div style="width: 30px"></div>
  <v-select
  class="ma-0"
  append-icon="search"
  hide-details
  solo
  single-line
  ></v-select>
  <div style="width: 30px"></div>
</v-toolbar>
</div>
</template>

<script>
import Logo from '@/components/Logo'

export default {
  name: 'ab-toolbar',

  components: {
    Logo
  },

  data: () => ({
    drawer: false
  })
}
</script>

<style>
.abtoolbar {
  padding-top: 0.8em;
}

.link {
  font-size: 1.4em;
}

.title {
  font-size: 1.6em;
}
</style>
