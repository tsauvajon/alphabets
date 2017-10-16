<template>
<div>
<v-navigation-drawer v-model="drawer" temporary light>
  <v-toolbar flat>
    <v-list>
      <v-list-tile>
        <v-list-tile-title class="title">
          Application
        </v-list-tile-title>
      </v-list-tile>
    </v-list>
  </v-toolbar>
  <v-divider></v-divider>
</v-navigation-drawer>
<v-toolbar @click="drawer = !drawer" app dense color="transparent" class="elevation-0 abtoolbar">
  <v-toolbar-title>
    <a href="#/" style="text-decoration: none">
      <img style="height: 35px; padding: 5px 0 0 5px" src="../../assets/bet.animated.svg">
    </a>
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

<style scoped>
.abtoolbar {
  padding-top: 0.8em;
}
</style>
