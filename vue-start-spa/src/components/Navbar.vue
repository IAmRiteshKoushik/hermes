<template>

  <nav 
      :class="[`navbar-${theme}`, `bg-${theme}`, 'navbar', 'navbar-expand-lg']"
    >
      <div class="container-fluid">
        <a href="#" class="navbar-brand">My Vue</a>
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li v-for="(page, index) in page" class="nav-item" :key="index">
            <navbar-link
              :page="page"
              :isActive="activePage == index"
              @click.prevent="navLinkClick(index)"
            ></navbar-link>
          </li>
        </ul>
        <form class="d-flex">
          <button
            class="btn btn-primary"
            @click.prevent="changeTheme()"         
          >Toggle Navbar</button>
        </form>
      </div>
    </nav>

</template>

<script>

import NavbarLink from '../components/NavbarLink.vue'

export default {
  components : {
    NavbarLink
  },
  props: ['page', 'activePage', 'navLinkClick'],

  data(){
    return{
      theme: "light"
    }
  },
  // The toggle method is specific to the navigation bar 
  created() {
    this.getThemeSetting()
  },
  methods: {
    changeTheme(){
      let theme = "light";
      if(this.theme == "light") {
        theme = "dark";
      }
      this.theme = theme;
      this.storeThemeSetting();
    },
    storeThemeSetting() {
      // localStorage is used to store data locally to the browser 
      // it is similar to caching/cookies 
      localStorage.setItem('theme', this.theme);
    },
    getThemeSetting() {
      let theme = localStorage.getItem('theme');

      if (theme) {
        this.theme = theme;
      }
    },
  }
}

</script>
