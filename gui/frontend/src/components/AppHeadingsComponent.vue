<script setup>
import { watch } from 'vue'
import { ref } from 'vue'
import { useRoute } from 'vue-router'

const headings = ref(null)
const route = useRoute()

const decomposedPaths = () =>
  route.path
    .split('/')
    .slice(1)
    .reduce((paths, segment) => {
      paths.push(`${paths[paths.length - 1] || ''}/${segment}`)
      return paths
    }, [])

const headingForPath = path => {
  const segment = path.match(/[\w.:@-]*$/)[0]
  if (segment[0] == '@') return segment
  return toTitle(segment)
}

const toTitle = string =>
  string.replace(/./, first => first.toUpperCase()).replace(/_/g, ' ')

const setHeadings = () =>
  (headings.value = decomposedPaths().map(path => ({
    path: path,
    title: headingForPath(path),
  })))

setHeadings()
watch(route, setHeadings)
</script>

<template>
  <div
    v-if="!(headings[0].path == '/' || route.params.noPage)"
    class="headings"
  >
    <span v-for="(heading, i) of headings" :key="i">
      <h2 v-if="i == headings.length - 1">
        {{ heading.title }}
      </h2>
      <router-link v-else :to="heading.path">
        <h2>
          {{ heading.title }}
        </h2>
      </router-link>
    </span>
  </div>
</template>

<style scoped>
.headings {
  margin-bottom: 0.3rem;
}

h2 {
  display: inline-block;
  margin-right: 0.75rem;
}

a h2:not(:hover) {
  color: rgba(var(--v-theme-on-surface));
}

a h2:hover {
  color: rgb(var(--v-theme-info));
  text-decoration: underline;
}
</style>
