<script setup>
  import MapCard from './MapCard.vue';

  import { ref, onMounted } from 'vue';
  import axios from 'axios';

  const mapStatus = ref([
    { id: 1, NameCN: "残月", NameEN: "BrokenMoon", mapId: 3, startTime: "00:00", endTime: "04:30"}
  ])


</script>

<template>
  <div class="map-group">
    <h3>排位赛</h3>
    <div class="rotation-list">
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

  h3 {
    width: fit-content;
    padding-bottom: 0.5rem;
    border-bottom: 2px solid rgb(33, 152, 37);
    margin: 0;
    margin-bottom: 16px;
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

  /* Webkit 悬停：滑块显形，轨道依旧透明 */
  .rotation-list:hover::-webkit-scrollbar-thumb {
    background-color: rgb(176, 196, 214); 
  }
</style>