<template>
<h1 class="ml5">
  <span class="text-wrapper">
    <span class="line line1"></span>
    <span class="letters letters-left">Alpha</span>
    <span class="letters logo"><img src="../assets/bet.svg" height="35px" /></span>
    <span class="letters letters-right">Bets</span>
    <span class="line line2"></span>
  </span>
</h1>
</template>

<script>
import anime from 'animejs'

export default {
  name: 'logo',

  mounted () {
    const logoAnimation = anime.timeline()

    logoAnimation.add({
      targets: '.ml5 .line',
      opacity: [0.5, 1],
      scaleX: [0, 1],
      easing: 'easeInOutExpo',
      duration: 700,
      delay: 600
    })

    logoAnimation.add({
      targets: '.ml5 .line',
      duration: 600,
      easing: 'easeOutExpo',
      translateY: function (e, i, l) {
        var offset = -0.625 + 0.625 * 2 * i
        return offset + 'em'
      }
    })

    logoAnimation.add({
      targets: '.ml5 .logo',
      opacity: [0, 1],
      scaleY: [0.5, 1],
      easing: 'easeOutExpo',
      duration: 600,
      offset: '-=600'
    })

    logoAnimation.add({
      targets: '.ml5 .letters-left',
      opacity: [0, 1],
      translateX: ['0.5em', 0],
      easing: 'easeOutExpo',
      duration: 600,
      offset: '-=300'
    })

    logoAnimation.add({
      targets: '.ml5 .letters-right',
      opacity: [0, 1],
      translateX: ['-0.5em', 0],
      easing: 'easeOutExpo',
      duration: 600,
      offset: '-=600'
    })
  }
}
</script>

<style>
.ml5 {
  position: relative;
  font-weight: 300;
  font-size: 4.5em;
  color: #333333;
}

.ml5 .text-wrapper {
  position: relative;
  display: inline-block;
  padding-top: 0.1em;
  padding-right: 0.05em;
  padding-bottom: 0.15em;
  line-height: 1em;
}

.ml5 .line {
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  margin: auto;
  height: 3px;
  width: 100%;
  background-color: #333333;
  transform-origin: 0.5 0;
}

.ml5 .letters {
  display: inline-block;
  opacity: 0;
}
</style>
