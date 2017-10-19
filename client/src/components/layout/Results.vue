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
import moment from 'moment'

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
      const startDate = moment().subtract(6, 'days').format('YYYY-MM-DD')
      const endDate = moment().format('YYYY-MM-DD')

      resultsAsync = api.soccer.getFixtures(startDate, endDate)
    } catch (ex) {
      // The request failed. Handle failures
      console.error(ex)
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
