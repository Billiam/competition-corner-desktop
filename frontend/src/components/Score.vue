<script setup>
import { computed } from 'vue'
const props = defineProps({
  score: {
    type: Number,
    required: true
  },
  date: {
    type: String,
  },
  rank: {
    type: Number,
    required: true
  },
  user: {
    type: String,
    required: true
  },
  image: {
    type: String
  },
  me: {
    type: String
  }
})

const localeDate = computed(() => {
  if (props.date) {
    return new Date(`${props.date} UTC`).toLocaleString()
  }
})
</script>
<template>
  <tr :title="localeDate" class="score-info" :class="{ me: me === user }">
    <td>{{ rank }}</td>
    <td class="user"><img class="avatar" :src="image" :alt="user" :title="user" />{{ user }}</td>
    <td class="score">{{ score.toLocaleString() }}</td>
  </tr>
</template>

<style scoped lang="scss">
.score-info:hover {
  .avatar {
    filter: saturate(1);
  }
  background-color: rgba(20,20,20,20);
}
</style>
<style scoped lang="scss">
.me {
  color: #fff;
  text-shadow: 0 0 3px #ffffffcc;
}
.score-info {
  word-break: break-word;
}
.score {
  white-space: nowrap;
  color: #ccc;
}
.avatar {
  filter: saturate(0);
  transition: filter 0.5s;
  border-radius: 50%;
  display: inline;
  vertical-align: middle;
  margin-right: 5px;
  max-width: 30px;
}
@media only screen and (max-width: 250px) {
  .avatar {
    display: none;
  }
}
</style>
