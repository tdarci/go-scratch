This is a bare-bones project for writing up some ad-hoc code.

--------------------------------

open book

earthquake: https://www.notion.so/Felt-Fullstack-Challenge-Earthquake-64bf1c9e11e247eab1754cc63ff5b6a8

build a system that allows people to "subscribe" to that API with a given criteria

# part 1

- Poll earthquake endpoint to get current earthquake data. (`https://earthquake.usgs.gov/earthquakes/feed/v1.0/summary/1.0_hour.geojson`)
- If there is a new quake, call our "earthquake event" endpoint (will implement details in part 2)

## approach

- Create server.
- Set up polling.
    - record quakes.
    - all new requests get sent to notification mechanism.

# part 2

- endpoint: "subscribe for alerts" (parameter: minimum magnitude)
- flesh out endpoint: "earthquake event" to send notification to all subscribers

