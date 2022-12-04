//
// custom-artillery-functions.js
//
module.exports = {
  logHeaders: logHeaders
}

function logHeaders(requestParams, response, context, ee, next) {
  // console.log(response.headers);
  return next(); // MUST be called for the scenario to continue
}