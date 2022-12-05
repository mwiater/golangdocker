//
// custom-artillery-functions.js
//
// Placeholder for custom Artiller function, just a middleware passthrough currently.
//

module.exports = {
  logHeaders: logHeaders
}

function logHeaders(requestParams, response, context, events, next) {
  // console.log(response.headers);
  return next();
}