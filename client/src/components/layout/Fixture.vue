<template>
<v-card flat hover tile>
  <v-card-title>
    <span class="date">2017/10/15 6pm</span>
  </v-card-title>
  <v-card-text>
    <div class="fixture">
      <v-layout row wrap>
        <v-flex xs4 class="text-xs-center team" @click="$router.push(`teams/${home.id}`)" v-if="home">
          <img :src="home.logo_path" class="logo" />
          <span class="team-name accent--text">{{ home.name }}</span>
        </v-flex>
        <v-flex xs4>
          <v-layout v-if="odds" row wrap class="text-xs-center odds">
            <v-flex xs12>
              <v-btn class="bet" color="primary">
                <v-icon left dark>chevron_left</v-icon>
                <span>{{ odds.home.toFixed(2) }}</span>
              </v-btn>
            </v-flex><v-flex xs12>
              <v-btn class="bet" color="primary">
                <v-icon left dark>remove</v-icon>
                <span>{{ odds.draw.toFixed(2) }}</span>
              </v-btn>
            </v-flex>
            <v-flex xs12>
              <v-btn class="bet" color="primary">
                <v-icon left dark>chevron_right</v-icon>
                <span>{{ odds.away.toFixed(2) }}</span>
              </v-btn>
            </v-flex>
          </v-layout>
          <v-progress-circular v-else style="margin: auto" :size="80" indeterminate color="primary"></v-progress-circular>
        </v-flex>
        <v-flex xs4 class="text-xs-center team" @click="$router.push(`teams/${away.id}`)" v-if="away">
          <img :src="away.logo_path" class="logo" />
          <span class="team-name accent--text">{{ away.name }}</span>    
        </v-flex>
      </v-layout>
    </div>
  </v-card-text>
  <v-card-text>
  </v-card-text>
</v-card>
</template>

<script>
import api from '@/helpers/url'

const calculateOdds = (odds) => {
  try {
    const threeWayResult = odds.find(type => type.name === '3Way Result')

    const result = threeWayResult.bookmaker.data.reduce((acc, val) => {
      console.log(acc)
      return {
        home: parseInt(acc.home) + parseInt(val.odds.data[0].value),
        away: parseInt(acc.away) + parseInt(val.odds.data[1].value),
        draw: parseInt(acc.draw) + parseInt(val.odds.data[2].value)
      }
    }, {
      home: 0,
      away: 0,
      draw: 0
    })

    return {
      home: result.home / threeWayResult.bookmaker.data.length,
      away: result.away / threeWayResult.bookmaker.data.length,
      draw: result.draw / threeWayResult.bookmaker.data.length
    }
  } catch (e) {
    // console.log('Impossible to calculate bets', e)
    return {
      home: 1,
      away: 1,
      draw: 1
    }
  }
}

export default {
  name: 'fixture',

  data: () => ({
    odds: null
  }),

  computed: {
    home () {
      if (!this.fixture) {
        return null
      }

      return this.fixture.LocalTeam
    },

    away () {
      if (!this.fixture) {
        return null
      }

      return this.fixture.VisitorTeam
    }
  },

  props: ['fixture'],

  async created () {
    if (!this.fixture) {
      return
    }

    const { id } = this.fixture

    let oddsAsync

    try {
      oddsAsync = api.soccer.getOddsByFixture(id)
    } catch (ex) {
      console.error(ex)
      return
    }

    const odds = await oddsAsync

    if (odds.status !== 200) {
      // The request failed. Handle failures
      return
    }

    this.odds = calculateOdds(odds.data)
  }
}
</script>

<style>
.team-name, .date {
  /* font-weight: 700; */
  font-size: 1.8em;
  font-family: 'Maven Pro', 'Helvetica','Arial', 'Roboto', sans-serif;
}

.team {
  cursor: pointer;
  padding: 25px !important;
}

.odds {
  padding: auto;
  height: 100%;
}

.odd {
  font-size: 1.4em;
}

.hyphen {
  font-weight: 700;
  font-size: 2em;
}

.fixture {
  margin: auto;
}

.logo {
  height: auto;
  width: 100%;
}

.bet {
  width: 30%;
  margin-left: auto;
  margin-right: auto;
}
</style>
