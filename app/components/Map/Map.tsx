"use client";

import React, { useState, useEffect } from "react";
import "leaflet/dist/leaflet.css";
import {
  MapContainer,
  TileLayer,
  Marker,
  Popup,
  Polyline,
  useMapEvent,
} from "react-leaflet";
import getRoute from "../OpenRouteService";
import ElevationProfile from "../ElevationProfile";
import IconBike from "../BikeIcon";
import MarkerIcon from "../MarkerIcon";
import DraggableMarker from "../DraggableMarker";

export type OsmPoint = { lat: number; lng: number };
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
  const [route, setRoute] = useState<OsmRoute | null>(null);

  const [start, setStart] = useState<OsmPoint | null>(null);
  const [end, setEnd] = useState<OsmPoint | null>(null);

  function OnMapClick() {
    const map = useMapEvent("click", (e) => {
      console.log("Clicked at:", e.latlng.toString());

      if (!start) {
        setStart(e.latlng);
      } else if (!end) {
        setEnd(e.latlng);
      } else {
        setStart(e.latlng);
        setEnd(null);
        setRoute(null);
      }
    });
  }

  useEffect(() => ComputeRoute(), [start, end]);
  function ComputeRoute() {
    if (start && end) {
      getRoute(toOsrPoint(start), toOsrPoint(end))
        .then((data) => {
          setRoute(data as OsmRoute);
        })
        .catch((error) => {
          console.error("setRouteError:", error);
        });
    }
  }

  function DisplayRoute() {
    return (
      <div>
        {start && (
          <DraggableMarker
            init={start}
            icon={MarkerIcon("green")}
            setPos={setStart}
          />
        )}
        {end && (
          <DraggableMarker
            init={end}
            icon={MarkerIcon("blue")}
            setPos={setEnd}
          />
        )}
        {route ? (
          <div>
            <Polyline
              positions={route.points}
              pathOptions={{ color: "red", weight: 4 }}
            />
          </div>
        ) : (
          <div>
            {start && end && (
              <Popup position={start}>No route has been found</Popup>
            )}
          </div>
        )}
      </div>
    );
  }

  function RouteManager() {
    OnMapClick();
    return <DisplayRoute />;
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
          <RouteManager />
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
