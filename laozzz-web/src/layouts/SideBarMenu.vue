<script setup>
  import { ref } from 'vue'
  import LinkButton from './components/LinkButton.vue';
  import IconHome from './components/IconHome.vue';
  import IconApex from './components/IconApex.vue';

  const spinning = ref(false)
</script>

<template>
  <header>
    <img
      src="/headerPic.png"
      width="120"
      :class="{ spinning }"
      @click="spinning = !spinning"
    />
    <span class="headerTitle">醪糟的个人站点</span>
  </header>

  <nav>
    <span class="menuTitle">MENU</span>
    <LinkButton name="Home" linkName="home" :icon="IconHome"/>
    <LinkButton name="Apex地图查询" linkName="ApexMap" :icon="IconApex"/>
    <LinkButton name="Home" linkName="home" :icon="IconHome"/>
    <LinkButton name="Home" linkName="home" :icon="IconHome"/>
    <LinkButton name="Home" linkName="home" :icon="IconHome"/>
    <LinkButton name="Home" linkName="home" :icon="IconHome"/>
    <LinkButton name="Home" linkName="home" :icon="IconHome"/>
  </nav>
</template>

<style scoped>
  header {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 200px;
  }
  
  img {
    cursor: pointer;
    animation: spin 1.2s linear infinite;
    animation-play-state: paused;
  }

  .headerTitle {
    font-family: 'Microsoft YaHei';
    font-weight: bold;
    font-size: 22px;
    color: rgb(55, 55, 55);
    cursor: default;
  }

  .spinning {
    animation: spin 1.2s linear infinite;
  }

  @keyframes spin {
    from { transform: rotate(0deg); }
    to   { transform: rotate(360deg); }
  }

  nav {
    display: flex;
    flex-direction: column;
  }

  .menuTitle {
    margin-top: 10px;
    margin-left: 10px;
    justify-content: center;
    font-family: 'Gill Sans MT';
    font-weight: bold;
    font-size: 14px;
    color: rgb(150, 150, 150);
    cursor: default;
  }

</style>