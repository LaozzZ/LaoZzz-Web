<script setup>
  defineProps({
    name: {
      type: String,
      default: '链接',
    },
    linkName: {
      type: String,
      required: true,
    },
    icon: {
      type: Object,
    }
  })

</script>

<template>
  <router-link :to="{name: linkName}">
    <component :is="icon" class="menu-icon" />
    <span>{{name}}</span>
  </router-link>
</template>

<style scoped>
  a {  
    display: flex;
    height: 45px;
    align-items: center;
    color:rgb(110, 110, 110);
    margin: 2px;
    border-radius: 12px;
    font-family: 'Oswald Medium';
    font-size: 18px;
    text-decoration-line: none;
    transition: 0.4s;
  }

  a:hover {
    color: rgb(45, 119, 238);
    background-color: rgb(230, 230, 230);
  }

  .menu-icon {
    margin-left: 1rem;
    margin-right: 1rem;
  }
</style>