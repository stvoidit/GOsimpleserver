<template>
  <nav class="uk-navbar-container" uk-navbar>
    <div class="uk-navbar-left">
      <ul class="uk-navbar-nav">
        <!-- class="uk-active" -->
        <li v-for="link in URLS" :key="link.url">
          <a :href="link.url">{{link.name}}</a>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {
  name: "menu",
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

<style>
</style>