<template>
<div>
  <div style="height: 50px"></div>
  <v-card color="transparent" flat class="embed-width-container">
    <v-card-title>
      <h2 class="accent--text">Video highlights</h2>
    </v-card-title>
  </v-card>
  <v-card-text
    v-for="h in highlights"
    :key="h.fixture_id"
    class="embed-container"
  >
    <iframe style="height: 100%; width: 100%;" :src="h.location"></iframe>
  </v-card-text>
</div>
</template>

<script>
import api from '@/helpers/url'

export default {
  name: 'fixture-detail',

  data: () => ({
    fixture: null,
    highlights: null
  }),

  async created () {
    const { id } = this.$route.params

    if (!id || id === 0) {
      this.$router.push('/')
    }

    let highlightsAsync

    try {
      highlightsAsync = api.soccer.getHighlightsByFixture(id)
    } catch (e) {
      console.error(e)
    }

    const highlights = await highlightsAsync

    this.highlights = highlights.data
  }
}
</script>

<style>
.embed-container {
  height: 810px;
  width: 1440px;
  margin: auto;
}

.embed-width-container {
  width: 1440px;
  margin: auto;
}

/* 1440 * 810 */
/* 1024 * 576 */

</style>
