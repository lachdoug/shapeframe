import objectValidation from './schemaValidation/objectValidation'
// import keysValidation from './schemaValidation/keysValidation'

export default function shapeValidation(v) {
  let object

  // Is valid JSON.
  try {
    object = JSON.parse(v)
  } catch (e) {
    return e.message
  }

  if (window.settings.framesSchemaValidationSkip) return true

  return (
    // The value for top-level is an object.
    objectValidation('top-level', object) ||
    // // The keys for the top-level are valid.
    // keysValidation('top-level', object, [
    //   'identifier',
    //   'about*',
    //   'bindings',
    //   'system_packages',
    //   'bundled_packages',
    //   'managed_packages',
    // ]) ||
    // // The value for "about" is an object.
    // objectValidation('"about"', object.about) ||
    // // The keys for "about" are valid.
    // keysValidation('"about"', object.about, ['title*', 'explanation*']) ||
    true
  )
}
