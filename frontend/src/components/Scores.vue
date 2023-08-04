<script setup>
import { ref, onMounted, onBeforeUnmount, computed } from 'vue'
import Score from './Score.vue'

const data = ref({})

let reloadTimer
let loaded = false

const fetchScoreData = async () => {
  const response = await fetch('https://virtualpinballchat.com:6080/api/v1/currentWeek?channelName=competition-corner')
  return response.json()
}

// const loadFakeScores = () => {
//   const newData = {...data.value}
//   newData.scores = data.value.scores.slice().map(v => ({ v, sort: Math.random() })).sort((a, b) => a.sort - b.sort).map(u => u.v)
//   data.value = newData
//   reloadTimer = setTimeout(loadFakeScores, 15*1000)
// }

const loadScores = async () => {
  try {
    data.value = await fetchScoreData()
    loaded = true
  } catch (e) {
    console.error(e)
  }
  reloadTimer = setTimeout(loadScores, 10*60*1000)

  // reloadTimer = setTimeout(loadFakeScores, 15*1000)
}

const scores = computed(() =>
  data.value?.scores || []
)
// strip manufacturer/year from title
const tableName = computed(() =>
  data.value?.table?.replace(/ ?\(.*\)$/, '')
)
onMounted(() => {
  loadScores()
})
onBeforeUnmount(() => {
  clearTimeout(reloadTimer)
})
</script>

<template>
  <div class="highscores" :class="{ hide: !loaded }">
  <div class="table-info">
    <h1 class="table">
      <a :title="data.table" :href="data.tableUrl" style="--wails-draggable:none">{{ tableName }}</a>
    </h1>
  </div>
  <table class="scores">
    <thead>
      <tr>
        <th>Rank</th>
        <th>User</th>
        <th>Score</th>
      </tr>
    </thead>
    <TransitionGroup name="score" tag="tbody">
      <Score
        v-for="(score, index) in scores"
        :key="score.username"
        :score="score.score"
        :user="score.username"
        :image="score.userAvatarUrl"
        :rank="index + 1"
      ></Score>
    </TransitionGroup>
  </table>
  </div>
</template>

<style lang="scss">
.table-info {
  line-height: 1.1;
  margin-bottom: 0.5rem;
  text-align: center;
  border-bottom: 3px #FF1186BA dotted;
  padding: 0.5rem 0;
}

.table {
  font-size: 1.2rem;

  a {
    text-decoration: none;
    color: #fff;
    transition: color 0.1s;
    &:hover {
      color: #baecff;
    }
  }
}
.scores {
  width: 100%;
  border-collapse: collapse;

  thead {
    display: none;
  }

  td, th {
    padding: 4px 6px;
    text-align: left;

    &:nth-child(1) {
      text-align: right;
      width: 1px;
      white-space: nowrap;
    }
    &:nth-child(3){
      width: 1px;
    }
  }
}

.score-move,
.score-enter-active,
.score-leave-active {
  transition: all 1s ease;
}
.score-enter-from,
.score-leave-to {
  opacity: 0;
  transform: translateX(30px);
}
.score-leave-active {
  position: absolute;
}

@media only screen and (max-width: 250px) {
  .avatar {
    display: none;
  }
}
</style>
<style scoped lang="scss">
  .highscores {
    transition: opacity 0.5s;
  }
  .hide {
    opacity: 0;
  }
</style>
