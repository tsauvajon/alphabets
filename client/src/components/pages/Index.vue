<template>
<div>
    <div style="height: 20px"></div>
    <h2 class="text-xs-center accent--text">Upcoming matches</h2>
    <fixtures></fixtures>
</div>
</template>

<script>
import Fixtures from '@/components/layout/Fixtures'

export default {
  name: 'index',

  components: {
    Fixtures
  }
}
</script>
