<template>
<v-container fluid grid-list-md>
  <v-layout v-if="fixtures" row wrap>
    <v-flex
      xs12 lg4
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
      fixturesAsync = api.soccer.getFixtures('2017-10-14', '2017-10-15')
    } catch (ex) {
      // The request failed. Handle failures
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
