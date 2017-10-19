<template>
<div>
<v-navigation-drawer v-model="drawer" temporary light>
  <v-toolbar extended color="accent" flat>
    <v-list>
      <v-list-tile>
        <v-list-tile-title class="white--text title">
          <span>AlphaBets</span>
        </v-list-tile-title>
      </v-list-tile>
      <v-list-tile>
        <v-list-tile-title class="white--text">
          <span style="font-weight: 700">{{ username }}</span>
        </v-list-tile-title>
        <v-list-tile-action @click="$store.dispatch('disconnect')">
          <v-btn flat icon><v-icon color="white">exit_to_app</v-icon></v-btn>
        </v-list-tile-action>
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
    <v-list-tile href="#/standing/6405">
      <v-list-tile-title class="primary--text link">Standings</v-list-tile-title>
    </v-list-tile>
  </v-list>

  <v-footer fixed style="height: 100px;" flat class="accent">
    <div style="width: 100%; margin-left: auto; margin-right: auto; text-align: center">
    <span style="font-size: 40px;" class="white--text">{{ wallet }}<img style="height: 35px; padding: 5px 0 0 5px;" src="../../assets/bet.white.svg" /></span>
    </div>
  </v-footer>
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
import { mapGetters } from 'vuex'

import Logo from '@/components/Logo'

export default {
  name: 'ab-toolbar',

  components: {
    Logo
  },

  computed: mapGetters(['wallet', 'username']),

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
