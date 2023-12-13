// // Pass an object and get back nested query params

// export default function (object) {
//   let queryString = []
//   let property

//   for (property in object) {
//     if (object.hasOwnProperty(property)) {
//       let k = property,
//         v = object[property]
//       queryString.push(
//         v !== null && ax.is.object(v)
//           ? this.stringify(v, {
//               prefix: k,
//             })
//           : `${k}=${encodeURIComponent(v)}`
//       )
//     }
//   }
//   return queryString.join('&')
// }
