<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { ForceDirectedGraphChart } from 'chartjs-chart-graph'
import ChartDataLabels from 'chartjs-plugin-datalabels'
import ChartDragData from 'chartjs-plugin-dragdata'
import { mdiDownloadOutline, mdiShuffle } from '@mdi/js'
import topologyDataFor from '/src/components/topology_graph/topology_data.js'
import loadChartDataFor from '/src/components/topology_graph/chart_data.js'

const viewScroller = ref(null)
const canvasWrapper = ref(null)
const canvasSizing = ref(null)
const canvas = ref(null)
const router = useRouter()
const graphFontSize = window.settings.graphFontSize * window.settings.graphScale

let chart, chartData, resizeTimer, topologyData

const props = defineProps({
  data: {
    type: Object,
    default: () => ({}),
  },
})

onMounted(() => {
  window.addEventListener('resize', resizeHandler)
  topologyData = topologyDataFor(props.data)
  loadDataAndRenderChart()
})

const loadDataAndRenderChart = () => {
  loadChartDataFor(topologyData).then(result => {
    chartData = result
    console.log(chartData)
    resizeCanvas()
    renderChart()
    setupCanvasEvents()
    // This next step is required, otherwise labels don't appear until animation has stopped.
    setTimeout(() => chart.update(), 0)
  })
}

onUnmounted(() => {
  window.removeEventListener('resize', resizeHandler)
})

const shufflePositions = () => {
  chart.getDatasetMeta(0).controller.stopLayout()
  chart.destroy()
  window.localStorage.graphShuffleSeed = Math.floor(Math.random() * 10000)
  loadDataAndRenderChart()
}

const resizeHandler = () => {
  clearTimeout(resizeTimer)
  resizeTimer = setTimeout(resizeCanvas, 100)
}

const dataset = () => ({
  pointStyle: chartData.icons,
  pointHitRadius: 50,
  data: chartData.nodes,
  edges: chartData.links,
  edgeLineBorderColor: ctx => chartData.lines.colors[ctx.index],
  edgeLineBorderWidth: ctx => chartData.lines.widths[ctx.index],
  edgeLineBorderDash: ctx => chartData.lines.dashes[ctx.index],
})

const renderChart = () => {
  chart = new ForceDirectedGraphChart(canvas.value.getContext('2d'), {
    plugins: [ChartDragData, ...(graphFontSize ? [ChartDataLabels] : [])],
    data: {
      datasets: [dataset(), dataset()],
    },
    options: {
      scales: {
        x: {
          min: -1.2,
          max: 1.2,
        },
        y: {
          min: -1.3,
          max: 1.2,
        },
      },
      maintainAspectRatio: false,
      tension: 0.5,
      simulation: {
        initialIterations: 50,
        forces: {
          collide: {
            radius: 25,
          },
        },
      },
      plugins: {
        legend: {
          display: false,
        },
        datalabels: {
          align: 'bottom',
          offset: 30,
          backgroundColor: '#FFF',
          borderRadius: 5,
          borderWidth: 1,
          borderColor: '#999',
          font: {
            color: '#000',
            size: graphFontSize,
          },
        },
      },
    },
  })
}

const setupCanvasEvents = () => {
  const nearestNode = e => {
    const nearest = chart.getElementsAtEventForMode(e, 'nearest', {
      intersect: true,
    })[0]
    if (nearest) return chart.data.datasets[0].data[nearest.index]
  }

  canvas.value.onclick = e => {
    const node = nearestNode(e)
    if (node) {
      const match = node.type.match(/(shape|frame)/)[1]
      if (match && match == 'shape') {
        router.push(`/shapes/@${node.label}`)
      } else if (match && match == 'frame') {
        router.push(`/frames/@${node.label}`)
      }
    }
  }

  canvas.value.onmousemove = e => {
    const node = nearestNode(e)
    if (node) {
      e.target.style.cursor = 'pointer'
    } else {
      e.target.style.cursor = 'default'
    }
  }

  // let dragged

  // canvas.value.onmousedown = e => {
  //   const node = nearestNode(e)
  //   if (node) dragged = node
  // };

  // canvas.value.onmousemove = e => {
  //   const xRatio = e.layerX / e.target.clientWidth
  //   const yRatio = e.layerY / e.target.clientHeight
  //   console.log(e.layerX, e.target.clientWidth, e.layerY, e.target.clientHeight)
  //   if (dragged) {
  //     const x = 2 * e.layerX / e.target.clientWidth - 1
  //     const y = -2 * e.layerY / e.target.clientHeight + 1
  //     dragged.x = x
  //     dragged.y = y
  //     chart.update()
  //   }
  // };

  // canvas.value.onmouseup = e => {
  //   const node = nearestNode(e)
  //   // debugger
  //   if (!node) {
  //     // const x = e.clientX / e.target.clientWidth
  //     // const y = e.clientY / e.target.clientHeight
  //     dragged.x = -1
  //     dragged.y = 1
  //     // chart.update()
  //     // dragged.x = null
  //     // dragged.y = null
  //     // chart.update()
  //     // debugger
  //     // chart.getDatasetMeta(0).controller._simulation.restart()

  //   }
  // };
}

const exportImage = () => {
  const link = document.createElement('a')
  link.download = 'shapeframe-graph.png'
  link.href = canvas.value.toDataURL('image/png')
  link.click()
}

const resizeCanvas = () => {
  const scale = window.settings.graphScale
  const fit = window.settings.graphFit
  const canvasSizeX = window.settings.graphSizeX
  const canvasSizeY = window.settings.graphSizeY

  canvas.value.style.scale = 1 / scale
  if (fit) {
    canvasSizing.value.style.width = `${100 * scale}%`
    canvasSizing.value.style.height = `${100 * scale}%`
    canvasSizing.value.style.marginLeft = `-${(100 * (scale - 1)) / 2}%`
    canvasSizing.value.style.marginTop = `-${
      ((window.innerHeight - 64) * (scale - 1)) / 2
    }px`
    canvasWrapper.value.style.height = `100%`
  } else {
    canvasSizing.value.style.width = `${canvasSizeX * scale}px`
    canvasSizing.value.style.height = `${canvasSizeY * scale}px`
    canvasSizing.value.style.marginLeft = `-${
      (canvasSizeX * (scale - 1)) / 2
    }px`
    canvasSizing.value.style.marginTop = `-${(canvasSizeY * (scale - 1)) / 2}px`

    canvasWrapper.value.style.width = `${canvasSizeX}px`
    canvasWrapper.value.style.height = `${canvasSizeY}px`
    const verticalWhitespace =
      (viewScroller.value.clientHeight - canvasWrapper.value.clientHeight) / 2
    canvasWrapper.value.style.marginTop = `${
      verticalWhitespace < 0 ? 0 : verticalWhitespace
    }px`
  }
}
</script>

<template>
  <div ref="viewScroller" class="view-scroller">
    <div class="buttons">
      <v-btn title="Shuffle positions" @click="shufflePositions">
        <v-icon :icon="mdiShuffle" />
      </v-btn>
      <v-btn title="Download image" @click="exportImage">
        <v-icon :icon="mdiDownloadOutline" />
      </v-btn>
    </div>
    <div ref="canvasWrapper" class="canvas-wrapper">
      <div ref="canvasSizing" class="canvas-sizing">
        <canvas ref="canvas"></canvas>
      </div>
    </div>
  </div>
</template>

<style scoped>
.view-scroller {
  overflow: auto;
  margin: -16px;
  height: calc(100vh - 64px);
  /* Vuetify v-toolbar height is 64px (on desktop and 56px on mobile). v-container margins are 16px. */
}

.canvas-wrapper {
  overflow: hidden;
  margin: auto;
}

.canvas-sizing {
  position: relative;
}

.buttons {
  z-index: 2000;
  position: fixed;
  right: 0px;
  top: 65px;
}

.view-scroller .buttons {
  display: none;
}

.view-scroller:hover .buttons {
  display: unset;
}

.buttons > * {
  margin: 5px;
}
</style>
