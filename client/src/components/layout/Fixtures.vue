<template>
<v-container fluid grid-list-md>
  <v-layout row wrap>
    <v-flex
      xs12 md4
      v-for="n in 9"
      :key="n"
    >
      <fixture></fixture>
    </v-flex>
  </v-layout>
</v-container>
</template>

<script>
import Fixture from '@/components/layout/Fixture'
export default {
  name: 'fixtures',

  components: {
    Fixture
  }
}
</script>
