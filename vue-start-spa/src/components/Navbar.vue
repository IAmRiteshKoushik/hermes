<template>
  <nav 
      :class="[`navbar-\${theme}`, `bg-${theme}`, 'navbar', 'navbar-expand-lg']"
    >
      <div class="container-fluid">
        <a href="#" class="navbar-brand">My Vue</a>
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li v-for="(page, index) in page" class="nav-item" :key="index">
            <a 
              :href="page.link.url" 
              class="nav-link" 
              :class="{active: activePage == index}"
              aria-current="page" 
              :title="`This link goes to the ${page.link.text} page`"
              @click.prevent="navLinkClick(index)"
            >{{ page.link.text }}</a>
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
export default {
  props: ['page', 'activePage', 'navLinkClick'],
  data(){
    return{
      theme: "light"
    }
  },
  // The toggle method is specific to the navigation bar 
  methods: {
    changeTheme(){
      let theme = "light";
      if(this.theme == "light") {
        theme = "dark";
      }
      this.theme = theme;
    }
  }
}


</script>
