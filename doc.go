/*
Package transit_realtime is a simple library for handling GTFS-realtime feeds.

GTFS-realtime is a feed specification that allows public transportation
agencies to provide realtime updates about their fleet to application
developers. It is an extension to GTFS (General Transit Feed Specification), an
open data format for public transportation schedules and associated geographic
information. GTFS-realtime was designed around ease of implementation, good
GTFS interoperability and a focus on passenger information.

     https://developers.google.com/transit/gtfs-realtime/

The GTFS-realtime data is encoded using protocol buffers, so this library contains
the generated Go class to be able to encode/decode this data.
*/
package transit_realtime
