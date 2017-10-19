<template>
<div>
  <div style="height: 20px"></div>
  <h2 class="text-xs-center accent--text">My Bets</h2>
  <v-container v-if="bets" fluid grid-list-md>
    <v-layout row wrap>
      <v-flex xs12 md4 v-for="bet in bets" :key="bet.assetId">
        <v-card hover>
          <v-card-title>
            <h4 class="primary--text">{{ new Date(bet.dateBet).toLocaleString() }}</h4>
          </v-card-title>
          <v-card-text>
            <span class="bet-label accent--text">Amount : </span><span class="bet-value">{{ bet.bet }}</span><img style="height: 30px; margin-left: 10px;" src="../../assets/bet.animated.svg" />
          </v-card-text>
          <v-card-text>
            <span class="bet-label accent--text">Bet : </span><span class="bet-value">{{ bet.choice }}</span>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn large v-if="!bet.paid" flat color="primary">Already claimed</v-btn>
            <v-btn large v-else flat color="accent">Claim<img style="height: 25px; margin-left: 10px; cursor: pointer;" src="../../assets/bet.red.svg"></v-btn>
          </v-card-actions>
        </v-card>
      </v-flex>
    </v-layout>
  </v-container>
  <center v-else>
    <logo></logo>
  </center>
  <v-tooltip left>
    <v-btn slot="activator" bottom right fixed fab large color="white"><img src="../../assets/bet.red.black.svg" height="35px" width="50px" /></v-btn>
    <h6>Claim everything</h6>
  </v-tooltip>
</div>
</template>

<script>
import { mapGetters } from 'vuex'

import Logo from '@/components/Logo'
import api from '@/helpers/url'

export default {
  name: 'nets',

  components: {
    Logo
  },

  computed: mapGetters(['bets', 'id']),

  data: () => ({ team: null }),

  async created () {
    let betsAsync

    try {
      betsAsync = api.blockchain.getBets()
    } catch (e) {
      console.error(e)
    }

    const bets = await betsAsync

    if (bets.status !== 200) {
      return
    }

    let filteredBets

    try {
      filteredBets = bets.data.filter(bet => bet.userId === `resource:world.alphabets.User#${this.id}`)
    } catch (e) {
      console.error(e)
    }

    const choice = bet => {
      switch (bet) {
        case '1':
        case 1:
          return 'Local'
        case '2':
        case 2:
          return 'Visitor'
        default:
          return 'Draw'
      }
    }

    const mappedBets = filteredBets.map(b => ({
      ...b,
      choice: choice(b.choice)
    }))

    this.$store.dispatch('setBets', { bets: mappedBets })
  }
}
</script>

<style>
.bet-label {
  font-weight: 700;
  font-family: 'Maven Pro', 'Helvetica','Arial', 'Roboto', sans-serif;
  font-size: 2em;
}

.bet-value {
  font-size: 2em;
}
</style>
