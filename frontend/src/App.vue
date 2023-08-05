<script setup>
import Scores from './components/Scores.vue'
import Close from './components/Close.vue'
import './style.scss'
</script>

<template>
  <close />
  <main>
    <Scores />
  </main>
</template>

<style scoped>
body {
  font-size: 16px;
  height: 100vh;
}
main {
  padding: 5px;
  font-size: .75em;
  color: #efefef;
  position: relative;
  height: 100%;
}
header {
  line-height: 1.5;
}
</style>
