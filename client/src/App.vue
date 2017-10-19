<template>
<v-app>
  <login v-if="!connected"></login>
  <template v-else>
    <ab-toolbar></ab-toolbar>
    <main>
      <v-content>
        <router-view></router-view>
      </v-content>
    </main>
    <ab-footer></ab-footer>
  </template>
</v-app>
</template>

<script>
import { mapGetters } from 'vuex'

import Toolbar from '@/components/layout/Toolbar'
import ABFooter from '@/components/layout/ABFooter'
import Login from '@/components/pages/Login'

export default {
  name: 'app',

  components: {
    'ab-toolbar': Toolbar,
    'ab-footer': ABFooter,
    Login
  },

  computed: mapGetters(['connected'])
}
</script>

<style lang="stylus">
  @require './stylus/main'
</style>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #333;
}

strong {
  color: #C52B2B;
}

h1, h2, h3 {
  font-family: 'Maven Pro', 'Helvetica','Arial', 'Roboto', sans-serif;
}
</style>
