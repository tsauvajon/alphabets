<template>
<v-container fluid grid-list-md>
  <v-layout v-if="fixtures" row wrap>
    <v-flex
      xs12 md4
      v-for="fixture in fixtures"
      :key="fixture.id"
    >
      <fixture :fixture="fixture"></fixture>
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
  name: 'fixtures',

  components: {
    Fixture,
    Logo
  },

  computed: mapGetters(['fixtures']),

  async created () {
    // Avoid spamming the sportmonks API
    if (this.fixtures !== null) {
      return
    }

    let fixturesAsync

    try {
      const startDate = moment().add(1, 'days').format('YYYY-MM-DD')
      const endDate = moment().add(7, 'days').format('YYYY-MM-DD')

      fixturesAsync = api.soccer.getFixtures(startDate, endDate)
    } catch (ex) {
      // The request failed. Handle failures
      console.error(ex)
      return
    }

    const fixtures = await fixturesAsync

    if (fixtures.status !== 200) {
      // The request failed. Handle failures
      return
    }

    this.$store.dispatch('setFixtures', { fixtures: fixtures.data })
  }
}
</script>
