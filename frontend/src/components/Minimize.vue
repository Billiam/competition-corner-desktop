<script setup>
import { WindowMinimise } from '../../wailsjs/runtime'

const minimize = () => {
  WindowMinimise()
}
</script>

<template>
  <button @click="minimize" class="minimize button"><span class="visually-hidden">Minimize</span>
    <div class="icon">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6">
        <path stroke-linecap="round" stroke-linejoin="round" d="M19.5 17h-15" />
      </svg>
    </div>
  </button>
</template>

