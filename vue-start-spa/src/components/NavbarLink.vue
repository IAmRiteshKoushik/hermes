<template>
  <a 
    class="nav-link"
    :class="activeClasses"
    aria-current="page"
    :href="page.link.url"
    :title="`This link goes to the ${page.link.text} page`"
  >{{ page.link.text }}</a>
</template>

<script>

export default {
  // props is an array of properties
  // computed is a JSON of functions
  props: ["page", "isActive"],
  computed: {
    activeClasses() {
      return {
        active: this.isActive,
        emphasize: this.isActive
      }
    }
  }
}

</script>

<!-- Usually the exported style ends up getting applied to the entire file 
hence, we use style scoped to apply styling only to particular components-->
<style scoped>

.emphasize {
  text-decoration: underline !important;
}
</style>
