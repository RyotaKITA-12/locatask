<template>
    <div class="container text-center">
        <br>
        <span v-html=message></span>
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from "vuex"

export default {
    name: "Index",
    setup() {
        const message = ref('<p class="h5"><b>ログインは<a href="/login" style="color:#58E;">こちら</b></a></p>')
        const store = useStore()

        onMounted(async () => {
            try {
                const _ = await axios.get('user')
                message.value = '<p class="h5"><b>ホームは<a href="/home" style="color:#58E;">こちら</b></a></p>'
                await store.dispatch('setAuth', true)
            } catch (e) {
                await store.dispatch('setAuth', false)
            }
        })

        return { message }
    }
}
</script>

