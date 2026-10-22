<script setup>
  import MapCard from './MapCard.vue';
  import DatePicker from './DatePicker.vue';

  import { ref, onMounted, watch } from 'vue';
  import axios from 'axios';

  const mapStatus = ref([]);
  const loading = ref(true);
  const error = ref('');

  function formatDate(date) {
    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0')
    const day = String(date.getDate()).padStart(2, '0')
    return `${year}-${month}-${day}`
  }

  function selectedDateParams(value) {
    const match = /^(\d{4})-(\d{2})-(\d{2})$/.exec(value)
    if (!match) {
      return null
    }

    return {
      year: Number(match[1]),
      month: Number(match[2]),
      day: Number(match[3]),
    }
  }

  const selectedDate = ref(formatDate(new Date()))
  let latestRequest = 0

  async function loadMapStatus() {
    const params = selectedDateParams(selectedDate.value)
    if (!params) {
      return
    }

    const request = ++latestRequest
    loading.value = true
    error.value = ''

    try {
      const { data } = await axios.get('/api/rotations', {
        params,
      });
      if (request !== latestRequest) {
        return
      }
      mapStatus.value = data;
    } catch (e) {
      if (request !== latestRequest) {
        return
      }
      error.value = '地图轮换加载失败';
      console.error(e);
    } finally {
      if (request === latestRequest) {
        loading.value = false;
      }
    }
  }

  watch(selectedDate, loadMapStatus)
  onMounted(loadMapStatus);
</script>

<template>
  <div class="map-group">
    <div class="map-group-header">
      <h3>排位赛</h3>
      <DatePicker v-model="selectedDate" />
    </div>
    <Transition name="rotation" mode="out-in">
      <p v-if="loading" key="loading" class="rotation-tip">加载中…</p>
      <p v-else-if="error" key="error" class="rotation-tip">{{ error }}</p>
      <div v-else :key="selectedDate" class="rotation-list">
        <MapCard 
          v-for="m in mapStatus"
          :key="m.id"
          :map-name-c-n="m.NameCN"
          :map-name-e-n="m.NameEN"
          :image-url="`/api/image/${m.mapId}`"
          :start-time="m.startTime"
          :end-time="m.endTime"
        />
      </div>
    </Transition>
  </div>
</template>

<style scoped>
  .map-group {
    padding-top: 18px;
    padding-bottom: 18px;
    padding-left: 24px;
    padding-right: 24px;
    border-radius: 16px;
    border: 1px solid rgba(33, 152, 37, 0.25);
    background-color: rgba(255, 255, 255, 0.8);
  }

  .map-group-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    gap: 16px;
    margin-bottom: 16px;
  }

  h3 {
    width: fit-content;
    padding-bottom: 0.5rem;
    border-bottom: 2px solid rgb(33, 152, 37);
    margin: 0;
  }

  .rotation-tip {
    margin: 0;
    font-size: 14px;
    color: #666;
  }

  .rotation-enter-active,
  .rotation-leave-active {
    transition: opacity 0.15s ease, transform 0.15s ease;
  }

  .rotation-enter-from {
    opacity: 0;
    transform: translateY(16px);
  }

  .rotation-leave-to {
    opacity: 0;
    transform: translateY(-8px);
  }

  .rotation-list {
    display: flex;
    flex-wrap: nowrap;
    overflow-x: auto;

    padding: 16px 16px 12px;        /* 处理阴影异常问题 */
    margin: -16px -16px 0;

    scrollbar-width: thin; /* 滚动条变细 */
    scrollbar-color: transparent transparent; 
  }

  .rotation-list:hover {
    scrollbar-color: rgb(176, 196, 214) transparent; 
  }

  /* Webkit 默认透明 */
  .rotation-list::-webkit-scrollbar {
    height: 6px; /* 滚动条细一点 */
  }
  .rotation-list::-webkit-scrollbar-track {
    background: transparent; /* 轨道永远透明 */
    border: none;
  }
  .rotation-list::-webkit-scrollbar-thumb {
    background-color: transparent; /* 滑块默认透明 */
    border-radius: 4px;
  }

  .rotation-list:hover::-webkit-scrollbar-thumb {
    background-color: rgb(176, 196, 214); 
  }

  @media (max-width: 520px) {
    .map-group-header {
      align-items: flex-start;
      flex-direction: column;
      gap: 10px;
    }
  }
</style>
