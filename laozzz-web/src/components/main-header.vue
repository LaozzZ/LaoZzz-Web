<script setup>
  import { ref, onMounted } from 'vue'

  const today = ref('')

  onMounted(() => {
    const date = new Date()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    const day = String(date.getDate()).padStart(2, '0')
    today.value = `${date.getFullYear()}.${month}.${day}`
  })

  defineProps({
    title: {
      type: String,
      default: "",
    },
    describe: {
      type: String,
      default: "",
    }
  })
</script>

<template>
  <header>
    <div class="header-top">
      <span>醪糟的个人站点</span>
      <span class="time">{{ today }}</span>
    </div>
    <h1>{{title}}</h1>
    <p>{{describe}}</p>
  </header>
</template>

<style scoped>
  header {
    padding-top: 18px;
    padding-bottom: 18px;
    padding-left: 24px;
    padding-right: 24px;
    margin-bottom: 16px;
    border-radius: 16px;
    border-top: 3px solid var(--border-color, rgb(33, 152, 37));
    background-color: rgba(255, 255, 255, 0.8);
  }

  .header-top {
    display: flex;
    font-weight: bold;
    color: var(--top-color, rgb(60, 113, 95));
  }

  .header-top .time {
    margin-left: auto;
  }

  h1 {
    margin: 0;
    font-size: 30px;
  }

  p {
    font-weight: 700;
    color: var(--describe-color, rgb(150, 150, 150));
    margin: 0;
  }
</style>