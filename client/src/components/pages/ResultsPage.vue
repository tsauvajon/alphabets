<template>
<div>
  <div style="height: 20px"></div>
  <h2 class="text-xs-center accent--text">Results</h2>
  <results></results>
</div>
</template>

<script>
import Results from '@/components/layout/Results'

export default {
  name: 'index',

  components: {
    Results
  }
}
</script>
