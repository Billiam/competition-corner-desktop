<script setup>

import Scores from './components/Scores.vue'
import Menubar from './components/Menubar.vue'

import './style.scss'

</script>

<template>
  <Menubar />
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
