<template>
  <div class="container web-3">
    <form action="">
      <div class="wb-3">
        <label for="" class="form-label">
          Page Title
        </label>
        <input 
          type="text"
          class="form-control"
          :value="pageTitle"
          @input="(e) => pageTitle = e.target.value"
        >
      </div>
      <div class="wb-3">
        <label 
          class="form-label" 
          for=""
        >
          Content
        </label>
        <textarea 
          rows="5" 
          class="form-control"
          type="text"
        ></textarea>
      </div>
      <div class="wb-3">
        <button 
          class="btn btn-primary"
          @click.prevent="pageCreated({pageTitle})"
        >Create Page</button>
      </div>
    </form>
  </div>
</template>

<script>

export default {
  props: ['pageCreated'],
  data() {
    return {
      pageTitle: ''
    }
  },
}
</script>
