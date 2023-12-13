<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import {
  mdiGraphOutline,
  mdiShapeOutline,
  mdiVectorSquare,
  mdiSourceRepository,
  mdiKeyChain,
  mdiStrategy,
  mdiCogOutline,
} from '@mdi/js'

const drawer = ref(true)
const items = [
  {
    path: '/',
    title: 'Graph',
    icon: mdiGraphOutline,
  },
  '-',
  {
    path: '/shapes',
    title: 'Shapes',
    icon: mdiShapeOutline,
  },
  {
    path: '/frames',
    title: 'Frames',
    icon: mdiVectorSquare,
  },
  '-',
  {
    path: '/strategies',
    title: 'Strategies',
    icon: mdiStrategy,
  },
  '-',
  {
    path: '/repositories',
    title: 'Repositories',
    icon: mdiSourceRepository,
  },
  {
    path: '/keys',
    title: 'Keys',
    icon: mdiKeyChain,
  },
  '-',
  {
    path: '/settings',
    title: 'Settings',
    icon: mdiCogOutline,
  },
]

const router = useRouter()
const back = () => router.push('/')
</script>

<template>
  <v-app>
    <v-layout>
      <v-app-bar color="primary">
        <v-app-bar-nav-icon
          variant="text"
          @click="drawer = !drawer"
        ></v-app-bar-nav-icon>
        <v-toolbar-title style="cursor: pointer" @click="back"
          >Shapeframe</v-toolbar-title
        >
      </v-app-bar>
      <v-navigation-drawer v-model="drawer">
        <v-list>
          <div v-for="(item, i) in items" :key="i">
            <v-divider v-if="item == '-'"></v-divider>
            <v-list-item
              v-else
              :key="item.text"
              router
              :to="item.path"
              :prepend-icon="item.icon"
              :title="item.title"
            >
            </v-list-item>
          </div>
        </v-list>
      </v-navigation-drawer>

      <v-main>
        <v-container fill-height fluid>
          <slot></slot>
        </v-container>
      </v-main>
    </v-layout>
  </v-app>
</template>
