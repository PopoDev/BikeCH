import Openrouteservice from "openrouteservice-js";

const API_KEY = process.env.NEXT_PUBLIC_ORS_API_KEY;

let orsDirections = new Openrouteservice.Directions({ api_key: API_KEY });

/**
 * Compute the route between two points given with coordinates (lng, lat) !!!
 * @param {[Number, Number]} start [lng, lat]
 * @param {[Number, Number]} end [lng, lat]
 * @returns
 */
const getRoute = async (start, end) => {
  try {
    console.log("Key: ", API_KEY, typeof API_KEY);
    let response = await orsDirections.calculate({
      coordinates: [start, end],
      profile: "cycling-regular",
      elevation: true,
      format: "json",
    });
    // Add your own result handling here
    let polyline = decodePolyline(response.routes[0].geometry, true);
    let points = polyline.points;
    let elevations = polyline.elevations;
    let summary = response.routes[0].summary;

    console.log("response: ", response);
    console.log("points: ", points);
    console.log("elevations: ", elevations);

    let distances = [0];
    let sum = 0;
    for (let i = 0; i < points.length - 2; i++) {
      sum += calculateDistance(points[i], points[i + 1]);
      distances.push(sum);
    }
    console.log("distances: ", distances);

    return {
      points: points,
      elevations: elevations,
      distances: distances,
      distance: summary.distance,
      duration: summary.duration,
    };
  } catch (err) {
    console.log("An error occurred: " + err.status);
    console.error(await err.response.json());
  }
};

export default getRoute;

/**
 * Decode an x,y or x,y,z encoded polyline
 * @param {*} encodedPolyline
 * @param {Boolean} includeElevation - true for x,y,z polyline
 * @returns {Array} of coordinates
 */
const decodePolyline = (encodedPolyline, includeElevation) => {
  // array that holds the points
  let points = [];
  let elevations = [];
  let index = 0;
  const len = encodedPolyline.length;
  let lat = 0;
  let lng = 0;
  let ele = 0;
  while (index < len) {
    let b;
    let shift = 0;
    let result = 0;
    do {
      b = encodedPolyline.charAt(index++).charCodeAt(0) - 63; // finds ascii
      // and subtract it by 63
      result |= (b & 0x1f) << shift;
      shift += 5;
    } while (b >= 0x20);

    lat += (result & 1) !== 0 ? ~(result >> 1) : result >> 1;
    shift = 0;
    result = 0;
    do {
      b = encodedPolyline.charAt(index++).charCodeAt(0) - 63;
      result |= (b & 0x1f) << shift;
      shift += 5;
    } while (b >= 0x20);
    lng += (result & 1) !== 0 ? ~(result >> 1) : result >> 1;

    if (includeElevation) {
      shift = 0;
      result = 0;
      do {
        b = encodedPolyline.charAt(index++).charCodeAt(0) - 63;
        result |= (b & 0x1f) << shift;
        shift += 5;
      } while (b >= 0x20);
      ele += (result & 1) !== 0 ? ~(result >> 1) : result >> 1;
    }
    try {
      let location = [lat / 1e5, lng / 1e5];
      points.push(location);
      if (includeElevation) elevations.push(ele / 100);
    } catch (e) {
      console.log(e);
    }
  }
  return {
    points: points,
    elevations: elevations,
  };
};

// Function to calculate the distance between two coordinates
const calculateDistance = (coord1, coord2) => {
  const [lat1, lon1] = coord1;
  const [lat2, lon2] = coord2;
  const R = 6371e3; // Radius of the Earth in meters
  const φ1 = (lat1 * Math.PI) / 180;
  const φ2 = (lat2 * Math.PI) / 180;
  const Δφ = ((lat2 - lat1) * Math.PI) / 180;
  const Δλ = ((lon2 - lon1) * Math.PI) / 180;

  const a =
    Math.sin(Δφ / 2) * Math.sin(Δφ / 2) +
    Math.cos(φ1) * Math.cos(φ2) * Math.sin(Δλ / 2) * Math.sin(Δλ / 2);
  const c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1 - a));

  return R * c;
};
