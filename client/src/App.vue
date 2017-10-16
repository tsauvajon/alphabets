<template>
<v-app>
  <ab-toolbar></ab-toolbar>
  <main>
    <v-content>
      <router-view></router-view>
    </v-content>
  </main>
  <ab-footer></ab-footer>
</v-app>
</template>

<script>
import Toolbar from '@/components/layout/Toolbar'
import ABFooter from '@/components/layout/ABFooter'

export default {
  name: 'app',

  components: {
    'ab-toolbar': Toolbar,
    'ab-footer': ABFooter
  }
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
</style>
