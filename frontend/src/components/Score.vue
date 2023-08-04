<script setup>
const props = defineProps({
  score: {
    type: Number,
    required: true
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
  }
})
</script>
<template>
  <tr class="score-info">
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
<style lang="scss">
.score-info {
  word-break: break-word;
}
.score {
  white-space: nowrap;
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
</style>
