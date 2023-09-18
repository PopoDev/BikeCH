package data

// Attribute represents the key-value pairs.
type Attribute struct {
	Ordinal  int    // Ordinal number
	Key      string // Key
	Value    string // Value
	KeyValue string // Key=Value
}

// Constants for Attribute values.
var (
	// Highways
	HIGHWAY_SERVICE       = Attribute{0, "highway", "service", "highway=service"}
	HIGHWAY_TRACK         = Attribute{1, "highway", "track", "highway=track"}
	HIGHWAY_RESIDENTIAL   = Attribute{2, "highway", "residential", "highway=residential"}
	HIGHWAY_FOOTWAY       = Attribute{3, "highway", "footway", "highway=footway"}
	HIGHWAY_PATH          = Attribute{4, "highway", "path", "highway=path"}
	HIGHWAY_UNCLASSIFIED  = Attribute{5, "highway", "unclassified", "highway=unclassified"}
	HIGHWAY_TERTIARY      = Attribute{6, "highway", "tertiary", "highway=tertiary"}
	HIGHWAY_SECONDARY     = Attribute{7, "highway", "secondary", "highway=secondary"}
	HIGHWAY_STEPS         = Attribute{8, "highway", "steps", "highway=steps"}
	HIGHWAY_PRIMARY       = Attribute{9, "highway", "primary", "highway=primary"}
	HIGHWAY_CYCLEWAY      = Attribute{10, "highway", "cycleway", "highway=cycleway"}
	HIGHWAY_MOTORWAY      = Attribute{11, "highway", "motorway", "highway=motorway"}
	HIGHWAY_PEDESTRIAN    = Attribute{12, "highway", "pedestrian", "highway=pedestrian"}
	HIGHWAY_TRUNK         = Attribute{13, "highway", "trunk", "highway=trunk"}
	HIGHWAY_LIVING_STREET = Attribute{14, "highway", "living_street", "highway=living_street"}
	HIGHWAY_ROAD          = Attribute{15, "highway", "road", "highway=road"}

	// Roads with motorway-like restrictions
	MOTORROAD_YES = Attribute{16, "motorroad", "yes", "motorroad=yes"}

	// Track type
	TRACKTYPE_GRADE1 = Attribute{17, "tracktype", "grade1", "tracktype=grade1"}
	TRACKTYPE_GRADE2 = Attribute{18, "tracktype", "grade2", "tracktype=grade2"}
	TRACKTYPE_GRADE3 = Attribute{19, "tracktype", "grade3", "tracktype=grade3"}
	TRACKTYPE_GRADE4 = Attribute{20, "tracktype", "grade4", "tracktype=grade4"}
	TRACKTYPE_GRADE5 = Attribute{21, "tracktype", "grade5", "tracktype=grade5"}

	// Surface
	SURFACE_ASPHALT       = Attribute{22, "surface", "asphalt", "surface=asphalt"}
	SURFACE_UNPAVED       = Attribute{23, "surface", "unpaved", "surface=unpaved"}
	SURFACE_GRAVEL        = Attribute{24, "surface", "gravel", "surface=gravel"}
	SURFACE_PAVED         = Attribute{25, "surface", "paved", "surface=paved"}
	SURFACE_GROUND        = Attribute{26, "surface", "ground", "surface=ground"}
	SURFACE_CONCRETE      = Attribute{27, "surface", "concrete", "surface=concrete"}
	SURFACE_COMPACTED     = Attribute{28, "surface", "compacted", "surface=compacted"}
	SURFACE_PAVING_STONES = Attribute{29, "surface", "paving_stones", "surface=paving_stones"}
	SURFACE_GRASS         = Attribute{30, "surface", "grass", "surface=grass"}
	SURFACE_DIRT          = Attribute{31, "surface", "dirt", "surface=dirt"}
	SURFACE_FINE_GRAVEL   = Attribute{32, "surface", "fine_gravel", "surface=fine_gravel"}
	SURFACE_PEBBLESTONE   = Attribute{33, "surface", "pebblestone", "surface=pebblestone"}
	SURFACE_SETT          = Attribute{34, "surface", "sett", "surface=sett"}
	SURFACE_WOOD          = Attribute{35, "surface", "wood", "surface=wood"}
	SURFACE_SAND          = Attribute{36, "surface", "sand", "surface=sand"}
	SURFACE_COBBLESTONE   = Attribute{37, "surface", "cobblestone", "surface=cobblestone"}

	// One-way roads
	ONEWAY_YES         = Attribute{38, "oneway", "yes", "oneway=yes"}
	ONEWAY_M1          = Attribute{39, "oneway", "-1", "oneway=-1"}
	ONEWAY_BICYCLE_YES = Attribute{40, "oneway:bicycle", "yes", "oneway:bicycle=yes"}
	ONEWAY_BICYCLE_NO  = Attribute{41, "oneway:bicycle", "no", "oneway:bicycle=no"}

	// Vehicle access
	VEHICLE_NO      = Attribute{42, "vehicle", "no", "vehicle=no"}
	VEHICLE_PRIVATE = Attribute{43, "vehicle", "private", "vehicle=private"}

	// General access
	ACCESS_YES        = Attribute{44, "access", "yes", "access=yes"}
	ACCESS_NO         = Attribute{45, "access", "no", "access=no"}
	ACCESS_PRIVATE    = Attribute{46, "access", "private", "access=private"}
	ACCESS_PERMISSIVE = Attribute{47, "access", "permissive", "access=permissive"}

	// Bicycle lanes
	CYCLEWAY_OPPOSITE       = Attribute{48, "cycleway", "opposite", "cycleway=opposite"}
	CYCLEWAY_OPPOSITE_LANE  = Attribute{49, "cycleway", "opposite_lane", "cycleway=opposite_lane"}
	CYCLEWAY_OPPOSITE_TRACK = Attribute{50, "cycleway", "opposite_track", "cycleway=opposite_track"}

	// Bicycle access
	BICYCLE_YES          = Attribute{51, "bicycle", "yes", "bicycle=yes"}
	BICYCLE_NO           = Attribute{52, "bicycle", "no", "bicycle=no"}
	BICYCLE_DESIGNATED   = Attribute{53, "bicycle", "designated", "bicycle=designated"}
	BICYCLE_DISMOUNT     = Attribute{54, "bicycle", "dismount", "bicycle=dismount"}
	BICYCLE_USE_SIDEPATH = Attribute{55, "bicycle", "use_sidepath", "bicycle=use_sidepath"}
	BICYCLE_PERMISSIVE   = Attribute{56, "bicycle", "permissive", "bicycle=permissive"}
	BICYCLE_PRIVATE      = Attribute{57, "bicycle", "private", "bicycle=private"}

	// Bicycle route
	ICN_YES = Attribute{58, "icn", "yes", "icn=yes"}
	NCN_YES = Attribute{59, "ncn", "yes", "ncn=yes"}
	RCN_YES = Attribute{60, "rcn", "yes", "rcn=yes"}
	LCN_YES = Attribute{61, "lcn", "yes", "lcn=yes"}
)

// AllAttributes contains all defined attributes.
var AllAttributes = []Attribute{
	HIGHWAY_SERVICE, HIGHWAY_TRACK, HIGHWAY_RESIDENTIAL, HIGHWAY_FOOTWAY,
	HIGHWAY_PATH, HIGHWAY_UNCLASSIFIED, HIGHWAY_TERTIARY, HIGHWAY_SECONDARY,
	HIGHWAY_STEPS, HIGHWAY_PRIMARY, HIGHWAY_CYCLEWAY, HIGHWAY_MOTORWAY,
	HIGHWAY_PEDESTRIAN, HIGHWAY_TRUNK, HIGHWAY_LIVING_STREET, HIGHWAY_ROAD,
	MOTORROAD_YES,
	TRACKTYPE_GRADE1, TRACKTYPE_GRADE2, TRACKTYPE_GRADE3, TRACKTYPE_GRADE4,
	TRACKTYPE_GRADE5,
	SURFACE_ASPHALT, SURFACE_UNPAVED, SURFACE_GRAVEL, SURFACE_PAVED,
	SURFACE_GROUND, SURFACE_CONCRETE, SURFACE_COMPACTED, SURFACE_PAVING_STONES,
	SURFACE_GRASS, SURFACE_DIRT, SURFACE_FINE_GRAVEL, SURFACE_PEBBLESTONE,
	SURFACE_SETT, SURFACE_WOOD, SURFACE_SAND, SURFACE_COBBLESTONE,
	ONEWAY_YES, ONEWAY_M1, ONEWAY_BICYCLE_YES, ONEWAY_BICYCLE_NO,
	VEHICLE_NO, VEHICLE_PRIVATE, ACCESS_YES, ACCESS_NO, ACCESS_PRIVATE,
	ACCESS_PERMISSIVE,
	CYCLEWAY_OPPOSITE, CYCLEWAY_OPPOSITE_LANE, CYCLEWAY_OPPOSITE_TRACK,
	BICYCLE_YES, BICYCLE_NO, BICYCLE_DESIGNATED, BICYCLE_DISMOUNT,
	BICYCLE_USE_SIDEPATH, BICYCLE_PERMISSIVE, BICYCLE_PRIVATE, ICN_YES, NCN_YES,
	RCN_YES, LCN_YES,
}

// COUNT is the number of defined attributes.
var COUNT = len(AllAttributes)
