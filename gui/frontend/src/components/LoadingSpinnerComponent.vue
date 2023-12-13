<script setup>
defineProps({
  placeholder: {
    type: String,
    default: 'Loading',
  },
})
</script>

<template>
  <div class="loading-spinner">
    <v-progress-circular indeterminate />
    {{ placeholder }}
  </div>
</template>

<style scoped>
.loading-spinner {
  padding: 10px;
}
</style>
