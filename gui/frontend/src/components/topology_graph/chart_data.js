export default topology => {
  interconnect = window.settings.graphInterconnect
  emptyDataObject()
  organizeChartData(topology)
  return populateNodes().then((nodesWithIcons) => {
    nodesWithIcons.filter(n=>n).forEach(n => {
      data.nodes.push(n[0])
      data.icons.push(n[1])
    })
    populateNodeIndicies()
    populateEdges()
    return data
  })
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
    datum => {
      return (type == 'root' && datum.type == 'root') ||
      (datum.type == type && datum.identifier == identifier)
    }
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
  const nodeWithIcon = datum => {
    return new Promise((resolve) => {
      const icon = new Image()
      icon.onload = () => {
        const connection = datum.type.match(/connect/)
        if (!interconnect || !connection) {
          const canvas = document.createElement('canvas');
          canvas.width = 54;
          canvas.height = 54;
          let context = canvas.getContext('2d');
          context.fillStyle = "white";
          context.fillRect(0, 0, 54, 54);
          context.beginPath();
          context.lineWidth = 3;
          context.strokeStyle = '#666';
          context.strokeRect(0,0,54,54)
          context.closePath();
          context.drawImage(icon, 3,3);
          resolve([datum, canvas])
        } else {
          resolve(null)
        }
      }
      icon.src = `/api${datum.icon}`
    });
  }
  return Promise.all(nodes.map(nodeWithIcon))
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
      data.lines.dashes.push(datum.line.dash)
      data.links.push({
        label: '',
        target: datum.target,
        source: datum.source,
      })
    }
  })
}
