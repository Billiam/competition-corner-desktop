<script setup>
import { computed, ref } from 'vue'
import { DateTime } from 'luxon'

const props = defineProps({
  date: {
    type: String,
    required: true
  }
})
const endDate = computed(() => {
  if (!props.date) {
    return
  }

  return DateTime.fromISO(`${props.date}T00:00:00`, { zone: 'America/Los_Angeles'}).plus({ days: 1 }).startOf('day')
})

const plural = (word, num = 0) =>  `${word.replace(/s$/, '')}${num == 1 ? '' : 's'}`
const now = ref(DateTime.now())

const updateTime = () => {
  now.value = DateTime.now()
  setTimeout(updateTime, nextDelay())
}

const nextDelay = () => {
  if (countdown.value.seconds < 0 && countdown.value.minutes == null) {
    return 1000*60*10
  }

  if (countdown.value.seconds != null) {
    return 1000
  }
  if (countdown.value.minutes != null) {
    return 5000
  }

  if (countdown.value.hours != null) {
    return 1000*60
  }
  return 1000*60*10
}

const countdown = computed(() => {
  if (!endDate.value) {
    return
  }
  const diff = endDate.value.diff(now.value, ['seconds'])

  if (diff.seconds > 60*60*24*2) {
    return diff.shiftTo('days')
  }
  if (diff.seconds > 60*60*24) {
    return diff.shiftTo('days', 'hours')
  }
  if (diff.seconds > 60*60*2) {
    return diff.shiftTo('hours')
  }
  if (diff.seconds > 60*60) {
    return diff.shiftTo('hours', 'minutes')
  }
  if (diff.seconds > 60) {
    return diff.shiftTo('minutes')
  }
  return diff.shiftTo('seconds')
})

const diffString = computed(() => {
  if (!countdown.value) {
    return
  }

  const diff = countdown.value.toObject()

  return (['days', 'hours', 'minutes', 'seconds']).map(interval => {
    if (diff[interval] == null) {
      return
    }

    const val = Math.max(0, Math.floor(diff[interval]))
    return `${val} ${plural(interval, val)}`
  }).filter(Boolean).join(', ')
})

updateTime()

</script>
<template>
  <span class="countdown" v-if="diffString">{{ diffString }} left</span>
</template>

<style scoped>
  .countdown {
    font-style: italic;
    color: #999;
  }
</style>
