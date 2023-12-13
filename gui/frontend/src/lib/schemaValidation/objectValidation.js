// Check that the object is a plain object and not an array or null.
// label:String The key label to show in any message.
// object:Object The object to validate.

import isType from '../isType'

export default function objectValidation(label, object) {
  if (!isType('object', object)) {
    return `Value for ${label} must be an object.`
  }
  if (isType('array', object)) {
    return `Value for ${label} must not be an array.`
  }
  if (isType('null', object)) {
    return `Value for ${label} must not be null.`
  }
}
