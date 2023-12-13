export default topology => {
  interconnect = window.settings.graphInterconnect
  emptyDataObject()
  organizeChartData(topology)
  populateNodes()
  populateNodeIndicies()
  populateEdges()
  return data
}

let data, nodes, interconnect

const emptyDataObject = () => {
  data = {
    nodes: [],
    links: [],
    icons: [],
    lines: {
      colors: [],
      widths: [],
      dashes: [],
    },
  }
}

const findNodeIndexFor = (type, identifier) => {
  const result = data.nodes.findIndex(
    datum =>
      (type == 'root' && datum.type == 'root') ||
      (datum.type == type && datum.identifier == identifier)
  )
  return result
}

const organizeChartData = topology => {
  const dataRoot = topology[0]
  const dataNodes = topology.slice(1)
  let seed = Number(window.localStorage.graphShuffleSeed || '0')
  let currentIndex = dataNodes.length
  let randomIndex
  while (currentIndex > 0) {
    randomIndex = Math.floor(Math.abs(Math.sin(seed++)) * currentIndex)
    currentIndex--
    ;[dataNodes[currentIndex], dataNodes[randomIndex]] = [
      dataNodes[randomIndex],
      dataNodes[currentIndex],
    ]
  }
  nodes = [dataRoot, ...dataNodes]
}

const populateNodes = () => {
  nodes.forEach(datum => {
    const connection = datum.type.match(/connect/)
    if (!interconnect || !connection) {
      data.nodes.push(datum)
      const icon = new Image()
      icon.src = `/api${datum.icon}`
      data.icons.push(icon)
    }
  })
}

const populateNodeIndicies = () => {
  nodes.forEach(datum => {
    if (datum.type != 'root') {
      const connection = datum.type.match(/connect/)
      if (interconnect && connection) {
        const targetType = datum.type.match(/[^-]+/)[0]
        datum.target = findNodeIndexFor(targetType, datum.identifier)
      } else {
        datum.target = findNodeIndexFor(datum.type, datum.identifier)
      }
      datum.source = findNodeIndexFor(
        datum.parent.type,
        datum.parent.identifier
      )
    }
  })
}

const populateEdges = () => {
  nodes.forEach(datum => {
    if (datum.type != 'root') {
      data.lines.colors.push(datum.line.color)
      data.lines.widths.push(datum.line.width)
      data.lines.dashes.push((datum.line.dash || []).join(' '))
      data.links.push({
        label: '',
        target: datum.target,
        source: datum.source,
      })
    }
  })
}
