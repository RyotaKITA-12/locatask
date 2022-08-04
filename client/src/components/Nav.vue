<template>
    <nav class="navbar navbar-expand-md navbar-dark bg-dark">
        <template v-if="auth">
            <ul class="navbar-nav mr-auto">
                <li class="nav-item">
                    <router-link to="/home" class="nav-link">
                        <h2>LocaTASK</h2>
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
                        <h2>LocaTASK</h2>
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
        </template>
    </nav>
</template>

<script>
export default {
    name: "Nav",
}
</script>
