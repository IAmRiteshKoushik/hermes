<template>
    <form action="" class="container mb-3">
    <div class="row">
      <div class="col-md-8">
      <div class="mb-3">
        <label for="" class="form-label">
          Page Title
        </label>
        <!-- The @input is the same as oninput in Vanilla javascript -->
        <!-- Here, we are setting pageTitle to the value of targetted event-element-->
        <input 
          type="text"
          class="form-control"
          v-model="pageTitle"
        >
          <!-- The bottom way does the same thing as v-model -->

        <!-- <input  -->
        <!--   type="text" -->
        <!--   class="form-control" -->
        <!--   v-model="pageTitle" -->
        <!--   :value="pageTitle" -->
        <!--   @input="(e) => pageTitle = e.target.value" -->
        <!-- > -->

      </div>
      <div class="mb-3">
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
          v-model="content"
        ></textarea>
      </div>
      </div>
      <div class="col">
        <div class="mb-3">
          <label for="" class="form-label">
            Link Text
          </label>
          <input 
            type="text" 
            class="form-control"
            v-model="linkText"
          />
        </div>
        <div class="mb-3">
          <label for="" class="form-label">
            Link URL
          </label>
          <input 
            type="text" 
            class="form-control"
            v-model="linkUrl"
          />
        </div>
        <div class="row wb-3">
          <div class="fore-check">
            <input type="checkbox" class="form-check-input">
            <label for="gridCheck1" class="form-check-label">
              Published
            </label>
          </div>
        </div>
      </div>
    </div>

      <!-- Form submission button -->
     <div class="mb-3">
        <button 
          class="btn btn-primary"
          :disabled="isFormInvalid"
          @click.prevent="submitForm"
        >Create Page</button>
      </div>
    </form>
</template>

<script>

export default {
  props: ['pageCreated'],
  computed: {
    isFormInvalid(){
      return (!this.pageTitle || !this.content || !this.linkText || !this.linkUrl)
    },
  },
  data() {
    return {
      pageTitle: '',
      content: '',
      linkText: '',
      linkUrl: '',
      published: true 
    }
  },
  methods: {
    // Method to disable to button when unless all the form data is filled
    submitForm() {
      if (!this.pageTitle || !this.content || !this.linkText || !this.linkUrl) {
        alert('Please fill the entire form');
        return;
      }

      this.pageCreated({
        pageTitle : this.pageTitle,
        content : this.content,
        link: {
          text: this.linkText,
          url: this.linkUrl
        },
        published: this.published
      });

      this.pageTitle = "";
      this.content = "";
      this.likText = "";
      this.linkUrl = "";
      this.published = true;
    }
  },
  // Enabling watchers in Vue.js
  watch: {
    pageTitle(newTitle, oldTitle) {
      if(this.linkText == oldTitle){
        this.linkText = newTitle;
      }
    }
  },
}
</script>
