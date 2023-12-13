let nodes

export default data => {
  nodes = []
  processNode(data)
  return nodes
}

const processChildrenNodes = parent =>
  (parent.children || []).map(child => {
    processNode(child, parent)
  })

const processNode = (child, parent = {}) => {
  datapointFor(child, parent)
  processChildrenNodes(child)
}

const datapointFor = (child, parent) => {
  const node = {
    label: child.identifier,
    type: child.type,
    identifier: child.identifier,
    line: lineFor(child),
    icon: borderIconFor(child.icon),
    parent: {
      type: parent.type,
      identifier: parent.identifier,
    },
  }
  nodes.push(node)
}

const borderIconFor = (url) => {
   `<svg width="50" height="50" xmlns="http://www.w3.org/2000/svg">
  <image href="${url}" height="50" width="50" />
</svg>`
return url
}

const lineFor = node => {
  if (node.type == 'root') {
    return {}
  } else if (node.type == 'frame') {
    return {
      width: 10,
      color: '#999',
    }
  } else if (node.type == 'frame-connect') {
    return {
      width: 10,
      dash: [10, 10],
      color: '#CCC',
    }
  } else if (node.type == 'shape') {
    return {
      width: 5,
      color: '#999',
    }
  } else if (node.type == 'shape-connect') {
    return {
      width: 5,
      dash: [5, 5],
      color: '#CCC',
    }
  } else if (node.type == 'shape-embed') {
    return {
      width: 2,
      color: '#999',
    }
  }
}
