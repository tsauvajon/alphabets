<template>
<v-card flat hover tile>
  <v-card-title>
    <span class="date">2017/10/15 6pm</span>
  </v-card-title>
  <v-card-text>
    <div class="fixture">
      <!-- <v-container fluid> -->
        <v-layout row wrap>
          <v-flex xs4 class="text-xs-center team" @click="$router.push('teams/44')">
            <img src="https://cdn.sportmonks.com/images/soccer/teams/12/44.png" class="logo" />
            <span class="team-name accent--text">OM</span>
          </v-flex>
          <v-flex xs4>
            <v-layout row wrap class="text-xs-center odds">
              <v-flex xs12>
                <v-btn class="bet" color="primary">
                  <v-icon left dark>chevron_left</v-icon>
                  <span>3.5</span>
                </v-btn>
              </v-flex><v-flex xs12>
                <v-btn class="bet" color="primary">
                  <v-icon left dark>remove</v-icon>
                  <span>3</span>
                </v-btn>
              </v-flex>
              <v-flex xs12>
                <v-btn class="bet" color="primary">
                  <v-icon left dark>chevron_right</v-icon>
                  <span>1.2</span>
                </v-btn>
              </v-flex>
            </v-layout>
          </v-flex>
          <v-flex xs4 class="text-xs-center team" @click="$router.push('teams/591')">
            <img src="https://cdn.sportmonks.com/images/soccer/teams/15/591.png" class="logo" />
            <span class="team-name accent--text">PSG</span>    
          </v-flex>
        </v-layout>
      <!-- </v-container> -->
    </div>
  </v-card-text>
  <!-- <v-card-actions>
    <v-spacer></v-spacer>
    <v-btn flat color="primary">Infos</v-btn>
    <v-btn flat color="accent">Bet</v-btn>
  </v-card-actions> -->
</v-card>
</template>

<script>
export default {
  name: 'fixture'
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
