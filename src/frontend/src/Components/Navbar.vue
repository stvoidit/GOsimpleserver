<template>
  <nav class="uk-navbar-container" uk-navbar>
    <div class="uk-navbar-left">
      <ul class="uk-navbar-nav">
        <li v-for="link in URLS" :key="link.url">
          <a :href="link.url">{{link.name}}</a>
        </li>
      </ul>
    </div>

    <div class="uk-navbar-right">
      <ul class="uk-navbar-nav">
        <li>
          <a href="/logout">Logout</a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {
  name: "Menu",
  data() {
    return {
      URLS: [
        { url: "/login", name: "Login" },
        { url: "/MyVieos", name: "MyVieos" }
      ]
    };
  }
};
</script>
