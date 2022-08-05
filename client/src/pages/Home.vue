<template>
    <GMapMap :center="{ lat: 51.093048, lng: 6.842120 }" :zoom="7" map-type-id="terrain"
        style="width: 100vw; height: 900px">
        <GMapMarker :key="marker.id" v-for="marker in markers" :position="marker.position" />
    </GMapMap>
</template>

<script >
import axios from 'axios'

export default {
    name: 'App',
    setup() {
        const center = { lat: 51.093048, lng: 6.842120 }
        const markers = [
            {
                id: 'dfsldjl3r',
                position: {
                    lat: 51.093048, lng: 6.842120
                },
            }
        ]

        return {
            center,
            markers,
        }
    }
}
</script>
