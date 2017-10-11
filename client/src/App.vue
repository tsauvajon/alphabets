<template>
  <div id="app">
    <logo></logo>
    <router-view/>
  </div>
</template>

<script>
import Logo from '@/components/Logo'

export default {
  name: 'app',

  components: {
    Logo
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #333333;
  margin-top: 60px;
}

strong {
  color: #C52B2B;
}
</style>

<style lang="stylus">
  @require './stylus/main'
</style>
