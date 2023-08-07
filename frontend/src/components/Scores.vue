<script setup>
import { ref, onMounted, onBeforeUnmount, computed } from 'vue'
import { GetUser } from '../../wailsjs/go/main/App.js'
import Splash from './Splash.vue'
import Score from './Score.vue'
import Countdown from './Countdown.vue'

const data = ref({})
const me = ref("")

let reloadTimer
let loaded = ref(false)

const fetchScoreData = async () => {
  const response = await fetch('https://virtualpinballchat.com:6080/api/v1/currentWeek?channelName=competition-corner')
  return response.json()
}

const loadUser = async () => {
  me.value = await GetUser()
}

const loadFakeScores = () => {
  const newData = {...data.value}
  newData.scores = data.value.scores.slice().map(v => ({ v, sort: Math.random() })).sort((a, b) => a.sort - b.sort).map(u => u.v)
  data.value = newData
  reloadTimer = setTimeout(loadFakeScores, 15*1000)
}

const loadScores = async () => {
  try {
    data.value = await fetchScoreData()
    loaded.value = true
  } catch (e) {
    console.error(e)
  }
  reloadTimer = setTimeout(loadScores, 10*60*1000)
  // reloadTimer = setTimeout(loadFakeScores, 15*1000)
}

const scores = computed(() =>
  data.value.scores || []
)

// strip manufacturer/year from title
const tableName = computed(() =>
  data.value.table?.replace(/ ?\(.*\)$/, '')
)
const tableTitle = computed(() =>
  `${data.value.table}\n${data.value.periodStart} - ${data.value.periodEnd}\nSeason ${data.value.season}, Week ${data.value.currentSeasonWeekNumber}`
)
const tableUrl = computed(() =>
  `https://virtual-pinball-spreadsheet.web.app/game/${data.value.vpsId}`
)

onMounted(() => {
  loadScores()
  loadUser()
})
onBeforeUnmount(() => {
  clearTimeout(reloadTimer)
})
</script>

<template>
  <TransitionGroup name="fade">
    <Splash v-if="!loaded" />
    <div class="highscores" v-if="loaded">
      <div class="table-info">
        <h1 class="table">
          <a :title="tableTitle" :href="tableUrl" style="--wails-draggable:none">{{ tableName }}</a>
        </h1>
        <Countdown :date="data.periodEnd"></Countdown>
        <div class="divider">
          <div class="diamond"></div>
        </div>
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
            :date="score.posted"
            :me="me"
            :image="score.userAvatarUrl"
            :rank="index + 1"
          ></Score>
        </TransitionGroup>
      </table>
    </div>
  </TransitionGroup>
</template>

<style lang="scss" scoped>
.table-info {
  line-height: 1.1;
  text-align: center;
  padding-top: 0.5rem;

  .divider {
    position: relative;
    border-bottom: 1px solid #888;
    margin: 20px 35px;

    &:after,
    &:before {
      content: '';
      position: absolute;
      right: -30px;
      top: -3px;
      width: 20px;
      height: 7px;
      background-image: url('/src/assets/images/arrow.svg');
      background-repeat: no-repeat;
    }
    &:before {
      top: -4px;
      left: -30px;
      right: auto;
      transform: rotate(180deg);
    }
  }

  .diamond {
    position: absolute;
    left: 50%;
    top: -4px;
    transform: translateX(-50%) translateY(-50%);
    border: 3px solid transparent;
    border-bottom: 5px solid #ccc;

    &:after {
      content: '';
      position: absolute;
      left: -3px;
      top: 5px;
      width: 0;
      height: 0;
      border: 3px solid transparent;
      border-top: 5px solid #ccc;
    }
  }
}

.table {
  font-size: 1.2rem;
  margin-bottom: 0.4rem;

  a {
    text-decoration: none;
    color: #fff;
    transition: color 0.2s;
    &:hover {
      color: #ffefd4;
    }
  }
}

.scores {
  width: 100%;
  border-collapse: collapse;

  thead {
    display: none;
  }

  :deep(td),
  :deep(th) {
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

.fade-enter-active,
.fade-leave-active {
  opacity: 1;
  transition: opacity 0.5s;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
