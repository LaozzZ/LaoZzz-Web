<script setup>
  defineProps({
    mapNameCN: {
      type: String,
      required: true,
    },
    mapNameEN: {
      type: String,
      required: true,
    },
    imageUrl: {
      type: String,
      required: true,
    },
    startTime: {
      type: String,
      required: true,
    },
    endTime: {
      type: String,
      required: true,
    },
  })
</script>

<template>
  <div class="map-card">
      <img :src="imageUrl" class="card-image">
    <div class="card-content">
      <h3 class="card-title">
        {{mapNameEN}} <span class="card-subtitle">{{mapNameCN}}</span>
      </h3>
      <p class="card-time">{{startTime}} - {{endTime}}</p>
    </div>
  </div>
</template>

<style scoped>
  .map-card {
    display: flex;
    flex-direction: column;
    flex-shrink: 0;
    width: 300px;
    margin-right: 16px;
    border-radius: 16px; /* 圆角 */
    overflow: hidden; /* 让图片也遵循圆角 */
    background: #fff;
    border: 1px solid #e0e0e0;
    font-family: sans-serif;
    transition: all 0.2s;
  }

  .map-card:hover {
    box-shadow: 0 4px 16px rgba(0,0,0,0.2);
  }

  .card-image {
    width: 300px;
    height: 100%;
    object-fit: cover; /* 填充裁剪 */
    display: block;
  }

  .card-content {
    padding: 10px;
  }

  .card-title {
    display: flex;
    align-items: baseline; /* 让英文和中文底部对齐 */
    gap: 8px; /* 英文和中文之间的间距 */
    margin: 0;
  }

  .card-subtitle {
    font-size: 16px;
    font-weight: normal;
    color: #666;
  }
  .card-time, .card-countdown {
    margin: 0;
    font-size: 14px;
    color: #666;
    line-height: 1.5;
  }
</style>