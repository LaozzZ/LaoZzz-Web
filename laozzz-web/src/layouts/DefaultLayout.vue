<script setup>
  import SideBarMenu from './SideBarMenu.vue';

</script>

<template>
  <div class="layout">
    <aside class="sidebar">
      <SideBarMenu />
    </aside>

      <main>
        <div class="app-shell">
          <router-view v-slot="{ Component }">
            <transition name="page" mode="out-in">
                <component :is="Component" />
            </transition>
          </router-view>
        </div>
      </main>
  </div>
</template>

<style scoped>
  .layout {
    display: flex;
    height: 100%;
    overflow: hidden;
  }

  .sidebar {
    width: 220px;
    overflow-y: auto;
    flex-shrink: 0;
    background-color: rgb(255, 255, 255);
    border-right: 1px solid rgb(220, 220, 220);
  }

  main {
    display: flex;
    flex: auto;
    overflow-y: auto;
    /* min-width: 0; */
  }

  .app-shell {
  flex: 1;
  align-self: flex-start;
  overflow: hidden;
}

  /* 页面切换过渡动画 */
  .page-enter-active,
  .page-leave-active {
    transition: all 0.12s ease;
  }

  .page-enter-from {
    opacity: 0;
    transform: translateY(20px);
  }

  .page-leave-to {
    opacity: 0;
  }

</style>
