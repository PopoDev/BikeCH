"use client";

import React, { useState } from "react";
import "leaflet/dist/leaflet.css";
import {
  MapContainer,
  TileLayer,
  Marker,
  Popup,
  Polyline,
  useMapEvent,
} from "react-leaflet";
import Icon from "../Icon";
import getRoute from "../OpenRouteService";
import ElevationProfile from "../ElevationProfile";
import IconBike from "../IconBike";

type OsmPoint = { lat: number; lng: number };
const DEFAULT_CENTER: OsmPoint = { lat: 46.948669, lng: 7.455841 };
const DEFAULT_ZOOM = 12;

interface MapProps {
  center?: [number, number];
  zoom?: number;
  scrollWheelZoom?: boolean;
}

function toOsrPoint(point: OsmPoint): [number, number] {
  return [point.lng, point.lat];
}

export default function Map(props: MapProps) {
  const {
    center = DEFAULT_CENTER,
    zoom = DEFAULT_ZOOM,
    scrollWheelZoom = true,
  } = props;

  type OsmRoute = {
    points: OsmPoint[];
    elevations?: number[];
    distances: number[];
    distance: number;
    duration: number;
  };
  const [route, setRoute] = useState<OsmRoute>({
    points: [],
    elevations: [],
    distances: [],
    distance: 0,
    duration: 0,
  });

  const [position, setPosition] = useState(DEFAULT_CENTER);

  function MarkerOnMapClick() {
    const map = useMapEvent("click", (e) => {
      console.log("Clicked at:", e.latlng.toString());
      setPosition(e.latlng);

      getRoute(toOsrPoint(DEFAULT_CENTER), toOsrPoint(e.latlng))
        .then((data) => {
          setRoute(data as OsmRoute);
        })
        .catch((error) => {
          console.error("setRouteError:", error);
        });
    });
    return (
      <div>
        {route ? (
          <div>
            <Polyline
              positions={route.points}
              pathOptions={{ color: "red", weight: 4 }}
            />
            <Marker position={position} icon={Icon} />
          </div>
        ) : (
          <div>
            <Popup position={position}>
              No point found around this location
            </Popup>
          </div>
        )}
      </div>
    );
  }

  return (
    <div className="flex-grow flex flex-col w-full" style={{ height: "300px" }}>
      <div className="h-3/4">
        <MapContainer
          style={{ width: "100%", height: "100%" }}
          center={center}
          zoom={zoom}
          scrollWheelZoom={scrollWheelZoom}
        >
          <TileLayer
            url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
            attribution='&copy; <a href="http://osm.org/copyright">OpenStreetMap</a> contributors'
          />
          <Marker position={center} icon={Icon}>
            <Popup>
              A pretty CSS3 popup. <br /> Easily customizable.
            </Popup>
          </Marker>
          <MarkerOnMapClick />
        </MapContainer>
      </div>
      {route && route.points.length > 0 && (
        <div className="h-1/4 flex flex-col">
          <div className="route-info flex flex-row text-center justify-center font-mono">
            <IconBike
              className="inline-block"
              style={{ width: "24px", height: "24px", padding: "4px" }}
            />
            <p className="text-base px-[4px]">
              {formatTime(route.duration)}
              {" ("}
              {route.distance < 1000
                ? route.distance.toFixed(0) + "m"
                : (route.distance / 1000).toFixed(1) + "km"}
              {")"}
            </p>
          </div>
          {route.elevations && route.elevations.length > 0 && (
            <div
              className="route-elevation flex-grow"
              style={{ minHeight: "100px" }}
            >
              <ElevationProfile
                elevations={route.elevations}
                distances={route.distances}
              />
            </div>
          )}
        </div>
      )}
    </div>
  );
}

function formatTime(seconds: number) {
  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  const remainingSeconds = Math.floor(seconds % 60);

  const formattedHours = hours.toString();
  const formattedMinutes = minutes.toString();
  const formattedSeconds = remainingSeconds.toString();

  if (hours > 0) {
    return `${formattedHours}h 
    ${formattedMinutes.padStart(2, "0")}min 
    ${formattedSeconds.padStart(2, "0")}s`;
  } else if (minutes > 0) {
    return `${formattedMinutes}min ${formattedSeconds.padStart(2, "0")}s`;
  } else {
    return `${formattedSeconds}s`;
  }
}
