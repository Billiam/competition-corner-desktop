<script setup>
import { ExitApp } from '../../wailsjs/go/main/App.js'

const close = () => {
  ExitApp()
}
</script>

<style scoped lang="scss">
.close {
  position: absolute;
  z-index: 1000;
  top: 0;
  right: 0;
  width: 20px;
  height: 20px;
  padding: 0;
  margin: 0;
  border: 0;
  color: white;
  opacity: 0;
  transition: opacity 0.1s;
  background-color: #000;

  .icon {
    width: 100%;
    height: 100%;
  }
}
body:hover {
  .close {
    opacity: 1;
  }
}
.icon {
  display: inline-block;
}
</style>


<template>
  <button @click="close" class="close button" style="--wails-draggable:none"><span class="visually-hidden">Close</span>
    <div class="icon">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth={1.5} stroke="currentColor" className="w-6 h-6">
        <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
      </svg>
    </div>
  </button>
</template>

