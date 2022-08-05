<template>
    <GMapMap :center="center" :zoom="10" map-type-id="terrain" style="width: 100vw; height: 100vh">
        <GMapMarker :key="marker.id" v-for="marker in markers" :position="marker.position" />
    </GMapMap>
</template>

<script >
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from "vuex"


export default {
    name: 'App',
    setup() {
        const center = { lng: 139.7671248, lat: 35.6812362 }
        const store = useStore()
        const markers = ref([])
        onMounted(async () => {
            try {
                const { data } = await axios.get('post')
                data.forEach(elem => {
                    var e = ref({})
                    e = {
                        id: elem.ID,
                        position: {
                            lat: elem.latitude, lng: elem.longitude
                        },
                    }
                    markers.value.push(e)
                })
                await store.dispatch('setAuth', true)
            } catch (e) {
                await store.dispatch('setAuth', false)
            }
        })

        return {
            center,
            markers,
        }
    }
}
</script>
