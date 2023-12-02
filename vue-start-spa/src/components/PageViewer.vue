<template>
  <div class="container">
    <h1>{{ page.pageTitle }}</h1>
    <p>{{ page.Content }</p>
  </div>
</template>

<script>
export default {
  props: ['page'],
}
</script>
