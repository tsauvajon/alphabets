<template>
<div>
  <div style="height: 20px"></div>
  <v-card v-if="team" class="team">
    <v-card-title>
      <h2 class="accent--text">{{ team.name }}</h2>
    </v-card-title>
    <v-card-media>
      <img :src="team.logo_path" style="height: 100%; width: auto;" />
    </v-card-media>
    <v-card-text>
      <template v-if="team.twitter && team.twitter !== ''">
        <strong>Twitter : </strong><a target="_blank" :href="`https://twitter.com/${team.twitter.substring(1)}`">{{ team.twitter }}</a><br>
      </template>
      <strong>Founded : </strong><span>{{ team.founded }}</span>
    </v-card-text>
  </v-card>
  <center v-else>
    <logo></logo>
  </center>
</div>
</template>

<script>
import { mapGetters } from 'vuex'

import Logo from '@/components/Logo'
import api from '@/helpers/url'

export default {
  name: 'team',

  components: {
    Logo
  },

  computed: mapGetters(['teams']),

  data: () => ({ team: null }),

  async created () {
    const { id } = this.$route.params

    if (!id || id === 0) {
      this.$router.push('/')
    }

    if (!this.teams) {
      console.log('teams not loaded yet')
    }

    console.log('id:', id)

    let team = this.teams.find(t => {
      console.log(t.id)
      return t.id === id
    })

    if (!team) {
      // get team
      let teamAsync

      try {
        teamAsync = api.soccer.getTeamById(id)
      } catch (ex) {
        // The request failed. Handle failures
        console.error(ex)
        return
      }

      team = await teamAsync
    }

    this.team = team.data
  }
}
</script>

<style scoped>
.team {
  max-width: 500px;
  margin: auto;
}
</style>
