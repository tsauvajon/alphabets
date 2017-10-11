<template>
<v-app id="app">
    <main>
      <v-content>
        <router-view></router-view>
      </v-content>
    </main>
</v-app>
</template>

<script>
export default {
  name: 'app'
}
</script>

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

<style lang="stylus">
  @require './stylus/main'
</style>
