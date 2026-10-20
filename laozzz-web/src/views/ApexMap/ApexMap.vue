<script setup>
  import mainHeader from '@/components/main-header.vue'
  import MapRotation from './components/MapRotation.vue';

</script>

<template>
  <div class="apexMap view">
    <mainHeader title="地图轮换" describe="实时查看Apex每日排位地图轮换" />
    <MapRotation />
  </div>
</template>