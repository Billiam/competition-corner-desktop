<script setup>
import Close from '@/components/Close.vue'
import Minimize from '@/components/Minimize.vue'
</script>

<style lang="scss" scoped>
.menubar {
  position: absolute;
  top: 0;
  right: 0;
  opacity: 0;
  padding: 1px 1px 0 0;
  margin: 0;
  white-space: nowrap;
  transition: opacity 0.2s;
  z-index: 1000;
  line-height: 16px;
}
body:hover {
  .menubar {
    opacity: 1;
  }
}

:deep(button) {
  cursor: pointer;
  display: inline-block;
  text-align: top;
  width: 18px;
  height: 18px;
  padding: 0;
  margin: 0 0 0 2px;
  color: #ddd;
  transition: opacity 0.1s, border-color 0.1s, color 0.1s;
  background-color: #000;
  border: 1px solid #ffffff55;
  overflow: hidden;

  .icon {
    width: 100%;
    height: 100%;
  }
  &:hover {
    color: #fff;
    border-color: #ffffffaa;
  }
}

</style>

<template>
  <div class="menubar" style="--wails-draggable:none">
    <Minimize />
    <Close />
  </div>
</template>

