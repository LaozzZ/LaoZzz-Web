<script setup>
const date = defineModel()
</script>

<template>
  <el-date-picker
    class="rotation-date-picker"
    v-model="date"
    type="date"
    value-format="YYYY-MM-DD"
    style="width: 150px"
  />
</template>

<style scoped>
  .rotation-date-picker {
    width: 100px;
    flex-shrink: 0;
  }
</style>
