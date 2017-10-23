<template>
<div class="background-container" :style="`background-image: url('${backgroundUrl}')`">
  <div class="login-container">
    <v-card hover dark class="login-card">
    <!-- <v-card hover dark style="background: rgba(51, 51, 51, .7); padding: 80px"> -->
      <v-card-media @click.stop="refreshLogo">
        <smaller-logo v-if="logo" style="height: 190px; width: 230px; margin: auto"></smaller-logo>
        <div v-else style="height: 190px"></div>
      </v-card-media>
      <v-card-text v-if="logo" style="margin-top: -20px">
        <h3 class="alphabets">AlphaBets</h3>
      </v-card-text>
      <v-card-text>
        <v-text-field @keyup.enter="connect" v-model="username" dark label="Login"></v-text-field><br>
        <v-text-field @keyup.enter="connect" v-model="password" type="password" dark label="Password"></v-text-field>
      </v-card-text>
      <v-card-text v-if="failed">
        <span class="white--text">Connection failed. Invalid login or password.</span>
      </v-card-text>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn @click="connect" flat dark>
          <span>Connection</span>
          <v-icon v-if="!connecting" right>chevron_right</v-icon>
          <v-progress-circular indeterminate style="margin-left: 5px" v-else color="white"></v-progress-circular>
        </v-btn>
      </v-card-actions>
    </v-card>
    <div style="height: 300px"></div>
  </div>
</div>
</template>

<script>
import SmallerLogo from '@/components/SmallerLogo'

const backgrounds = [
  'https://static.pexels.com/photos/4417/black-and-white-people-bar-men.jpg'
]

let rand = null

export default {
  name: 'login',

  components: {
    SmallerLogo
  },

  data: () => ({
    logo: true,
    username: null,
    password: null,
    failed: false,
    connecting: false
  }),

  computed: {
    backgroundUrl () {
      return backgrounds[rand]
    }
  },

  methods: {
    refreshLogo () {
      this.logo = false

      // Différé au prochain tick pour que la page soit re-render
      this.$nextTick(() => { this.logo = true })
    },

    changeBackground () {
      rand = Math.floor(Math.random() * backgrounds.length)
    },

    async connect () {
      if (!this.username) {
        return
      }

      if (!this.password) {
        return
      }

      this.connecting = true

      const success = await this.$store.dispatch('connect', {
        username: this.username,
        password: this.password
      })

      this.connecting = false
      this.failed = !success
    }
  },

  created () {
    this.changeBackground()
    this.failed = false
  }
}
</script>

<style>
.login-container {
  padding-top: 100px; 
  max-width: 500px;
  margin: auto;
}

.background-container {
  height: 100%;
  min-height: 600px;
  background: no-repeat center;
  background-size: 100%;
}

@media screen and (max-width: 1500px) {
  .background-container {
    background-size: cover;
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.application .login-container .login-card {
  background-color: rgba(197, 43, 43, .8);
  padding: 60px 80px;
}
</style>

<style scoped>
.alphabets {
  color: rgba(255, 255, 255, .8);
  text-align: center;
  
  animation-delay: 1.8s;
  animation-duration: 1.6s;
  animation-fill-mode: both;
  animation-name: fadeIn;
}
</style>
