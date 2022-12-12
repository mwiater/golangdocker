//
// custom-artillery-functions.js
//

module.exports = {
  logHeaders: logHeaders,
  customMetrics: customMetrics
}

function logHeaders(requestParams, response, context, events, next) {
  // console.log(response.headers);
  return next();
}

function customMetrics(requestParams, response, context, events, next) {
  const latency = parseServerTimingLatency(response.headers["server-timing"], "route");
  const url = new URL(requestParams.url);
  const routePath = url.pathname.replaceAll("/", "_")
  events.emit("histogram", "route_latency"+routePath.trim(), latency);
  return next();
}

function parseServerTimingLatency(header, timingMetricName) {
  const serverTimings = header.split(",");

  for (let timing of serverTimings) {
    const timingDetails = timing.split(";");
    if (timingDetails[0] === timingMetricName) {
      return parseFloat(timingDetails[1].split("=")[1]);
    }
  }
}
