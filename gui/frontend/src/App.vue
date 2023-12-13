<script setup>
import AppLayoutComponent from './components/AppLayoutComponent.vue'
import AppHeadingsComponent from './components/AppHeadingsComponent.vue'
</script>

<template>
  <AppLayoutComponent>
    <AppHeadingsComponent></AppHeadingsComponent>
    <router-view></router-view>
  </AppLayoutComponent>
</template>
