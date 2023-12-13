<script setup>
defineProps({
  text: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <p class="text-caption">{{ text }}</p>
</template>
