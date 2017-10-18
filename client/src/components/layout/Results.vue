<template>
<v-container fluid grid-list-md>
  <v-layout v-if="results" row wrap>
    <v-flex
      xs12 md4
      v-for="fixture in results"
      :key="fixture.id"
    >
      <fixture :fixture="fixture" :results="true"></fixture>
    </v-flex>
  </v-layout>
  <center v-else>
    <logo></logo>
  </center>
</v-container>
</template>

<script>
import { mapGetters } from 'vuex'

import Fixture from '@/components/layout/Fixture'
import Logo from '@/components/Logo'
import api from '@/helpers/url'

export default {
  name: 'results',

  components: {
    Fixture,
    Logo
  },

  computed: mapGetters(['results']),

  async created () {
    // Avoid spamming the sportmonks API
    if (this.results !== null) {
      return
    }

    let resultsAsync

    try {
      resultsAsync = api.soccer.getFixtures('2017-10-14', '2017-10-15')
    } catch (ex) {
      // The request failed. Handle failures
      return
    }

    const results = await resultsAsync

    if (results.status !== 200) {
      // The request failed. Handle failures
      return
    }

    this.$store.dispatch('setResults', { results: results.data })
  }
}
</script>
