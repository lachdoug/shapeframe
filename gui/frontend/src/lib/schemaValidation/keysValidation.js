// Check that an object has the right keys.
// label:String The key label to show in any message.
// object:Object The object to validate.
// validKeys:Array-of-strings A list of valid keys, with required keys appended with *.

export default function keysValidation(label, object, validKeys) {
  const permittedKeys = validKeys.map(k => k.match(/(?:\w+)/)[0])

  const requiredKeys = validKeys
    .filter(k => k.match(/\*/))
    .map(k => k.match(/(\w+)\*/)[1])

  const keys = Object.entries(object).map(kv => kv[0])

  for (const k of keys) {
    if (!permittedKeys.includes(k)) {
      return `Invalid  key "${k}" in ${label}.`
    }
  }

  for (const k of requiredKeys) {
    if (!keys.includes(k)) {
      return `Missing key "${k}" in ${label}.`
    }
  }
}
