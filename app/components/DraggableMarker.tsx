import React, { useState, useMemo, useRef } from "react";
import { Marker } from "react-leaflet";
import { Icon, Marker as LeafletMarker } from "leaflet";
import { OsmPoint } from "./Map/Map";

interface DraggableMarkerProps {
  init: OsmPoint;
  icon: Icon;
  setPos: (point: OsmPoint) => void;
}

export default function DraggableMarker({
  init,
  icon,
  setPos,
}: DraggableMarkerProps) {
  const [position, setPosition] = useState(init);
  const markerRef = useRef<LeafletMarker>(null);
  const eventHandlers = useMemo(
    () => ({
      dragend() {
        const marker = markerRef.current;
        if (marker != null) {
          setPosition(marker.getLatLng());
          setPos(marker.getLatLng());
        }
      },
    }),
    []
  );

  return (
    <Marker
      draggable={true}
      eventHandlers={eventHandlers}
      position={position}
      ref={markerRef}
      icon={icon}
    ></Marker>
  );
}
