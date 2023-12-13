export default (type, value) =>
  ({
    array: v => v instanceof Array,
    boolean: v => v === 'boolean',
    class: v => ('' + v).slice(0, 5) === 'class',
    false: v => v === false,
    function: v => typeof v === 'function',
    node: v => v instanceof Node,
    nodelist: v => v instanceof NodeList,
    null: v => v === null,
    number: v => typeof v === 'number',
    object: v => v.constructor === Object,
    promise: v => v instanceof Promise,
    string: v => typeof v === 'string',
    true: v => v === true,
    undefined: v => v === void 0,
  })[type.toLowerCase()](value)
