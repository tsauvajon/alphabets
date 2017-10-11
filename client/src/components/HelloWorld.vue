<template>
  <div class="hello">
    <!-- <h1>Bienvenue sur <strong>AlphaBets</strong></h1> -->
    <br><br>
    <v-btn class="medium" color="primary">Primary</v-btn><br>
    <v-btn class="medium" color="warning">Warning</v-btn><br>
    <v-btn class="medium" color="secondary">Secondary</v-btn><br>
    <v-btn class="medium" color="accent white--text">Accent</v-btn><br>
    <v-btn class="medium" color="error">Error</v-btn><br>
    <v-btn class="medium" color="info">Info</v-btn><br>
    <v-btn class="medium" color="success">Success</v-btn><br>
    <br><br>
    <v-speed-dial bottom right fixed hover>
      <v-btn fab large color="white" slot="activator"><img src="../assets/bet.red.black.svg" height="35px" width="50px" /></v-btn>
      <v-btn fab large color="white"><img src="../assets/bet.svg"height="35px" width="50px" /></v-btn>
      <v-btn fab large color="primary"><img src="../assets/bet.red.svg" height="35px" width="50px" /></v-btn>
      <v-btn fab large color="primary"><img src="../assets/bet.red.white.svg" height="35px" width="50px" /></v-btn>
      <v-btn fab large color="primary"><img src="../assets/bet.white.svg" height="35px" width="50px" /></v-btn>
      <!-- <v-btn fab large color="primary"><img src="../assets/bet.white.red.svg" height="35px" width="50px" /></v-btn> -->
      <!-- <v-btn fab large color="accent"><img src="../assets/bet.black.white.svg" height="35px" width="50px" /></v-btn> -->
      <v-btn fab large color="accent"><img src="../assets/bet.white.svg" height="35px" width="50px" /></v-btn>
      <!-- <v-btn fab large color="accent"><img src="../assets/bet.white.black.svg" height="35px" width="50px" /></v-btn> -->
      <!-- <v-btn fab large color="accent"><img src="../assets/bet.black.svg" height="35px" width="50px" /></v-btn> -->
    </v-speed-dial>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld'
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #C52B2B;
}

.medium {
  width: 200px;
}
</style>
