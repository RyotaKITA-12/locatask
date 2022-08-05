<template>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark">
        <template v-if="auth">
            <ul class="navbar-nav mr-auto">
                <li class="nav-item">
                    <router-link to="/home" class="nav-link">
                        <h2>　LocaTASK　</h2>
                    </router-link>
                </li>
            </ul>
            <ul class="navbar-nav my-2 my-lg-0 ms-auto">
                <li class="nav-item">
                    <router-link to="/login" class="nav-link" @click="logout">
                        <button class="btn btn-outline-danger">ログアウト</button>
                    </router-link>
                </li>
            </ul>
        </template>
        <template v-if="!auth">
            <ul class="navbar-nav mr-auto">
                <li class="nav-item">
                    <router-link to="/" class="nav-link">
                        <h2>　LocaTASK　</h2>
                    </router-link>
                </li>
            </ul>
            <ul class="navbar-nav my-2 my-lg-0 ms-auto">
                <li class="nav-item">
                    <router-link to="/login" class="nav-link">
                        <button class="btn btn-outline-primary">ログイン</button>
                    </router-link>
                </li>
                <li class="nav-item">
                    <router-link to="/register" class="nav-link">
                        <button class="btn btn-outline-success">新規登録</button>
                    </router-link>
                </li>
            </ul>
            <br>
        </template>
    </nav>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import axios from 'axios'

export default {
    name: "Nav",
    setup() {
        const store = useStore()
        const router = useRouter()
        const auth = computed(() => store.state.auth)
        const logout = async () => {
            await axios.post('logout', {})
            store.dispatch('setAuth', false)
            await router.push('/login')
        }

        return {
            auth,
            logout,
        }
    }
}
</script>

<style>
.form-profile {
    width: 100%;
    max-width: 500px;
    padding: 15px;
    margin: auto;
}

.form-profile .form-control {
    position: relative;
    box-sizing: border-box;
    height: auto;
    padding: 10px;
    font-size: 16px;
}

.form-profile .form-control:focus {
    z-index: 2;
}

/* .profile-container div {       */
/*     background-color: #F0F0F5; */
/* }                              */

.container p {
    display: inline;
    text-align: left;
}

.profile-container h2 {
    text-align: left;
}

.profile-rect svg {
    margin-left: 5px;
    margin-right: 5px;
}

.rect-left {
    float: left;
}

.bg-gray {
    background-color: #F5F5FF;
}

#overlay {
    z-index: 1;

    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);

    display: flex;
    align-items: center;
    justify-content: center;

}

#content {
    z-index: 2;
    width: 450px;
    padding: 1em;
    background: #fff;
}

input[type="date"] {
    padding: 0 10px;
    width: 200px;
    height: 40px;
    border: solid 1px #CCC;
    box-sizing: border-box;
    border-radius: 5px;
    font-size: 15px;
    color: #999;
    box-shadow: 0px;
}

.btn-circle {
    font-size: 100%;
    font-weight: bold;
    border: 1px solid #999;
    color: #999;
    display: flex;
    justify-content: center;
    align-items: center;
    border-radius: 100%;
    width: 1.5em;
    line-height: 1.3em;
    cursor: pointer;
    transition: .2s;
}

.shadow-bg {
    box-shadow: 0 10px 25px 0 rgba(80, 134, 135, 0.3);
}
</style>
