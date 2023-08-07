<script setup>
import { ExitApp } from '../../wailsjs/go/main/App.js'

const close = () => {
  ExitApp()
}
</script>

<template>
  <button @click="close" class="close button"><span class="visually-hidden">Close</span>
    <div class="icon">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
        <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
      </svg>
    </div>
  </button>
</template>

