<template>
  <div class="m-10">
    <div class="flex flex-col items-center gap-2 m-2">
      <p class="text-2xl">MamuroEmail</p>
      <input class="w-80 border" type="text" @input="search" v-model="searchTerm" placeholder="type to search by to, from, subject, content" />
      Click any row to display the content
      <h1 v-if="error">{{ error }}</h1>
    </div>
    <div class="flex flex-row">
      <MailList v-bind:mails="mails" v-on:selectMail="updateSelectedMail" class="" />
      <div v-if="selectedMail" class="flex flex-col items-start border">
        <div>To: {{ selectedMail.To }}</div>
        <div>From: {{ selectedMail.From }} </div>
        <div>Subject: {{ selectedMail.Subject }}</div>
        <div class="w-full border">{{ selectedMail.Body }}</div>
      </div>
    </div>
  </div>
</template>

<script>
import MailList from './MailList.vue';

export default {
    name: "SearchPage",
    data() {
        return {
            searchTerm: "",
            mails: [],
            error: "",
            timeoutHandler: null,
            selectedMail: null
        };
    },
    methods: {
        updateSelectedMail(selectedMail) {
          console.log(selectedMail);
            this.selectedMail = selectedMail;
        },
        search() {
            this.error = "";
            if (this.searchTerm.trim() == "")
                return;
            if (this.timeoutHandler)
                clearTimeout(this.timeoutHandler);
            const getResults = () => {
                fetch(`http://localhost:3000/api/search/${this.searchTerm}`)
                    .then(response => response.json())
                    .then(json => {
                    this.mails = json.Matches;
                })
                    .catch(error => {
                    this.error = error;
                });
            };
            this.timeoutHandler = setTimeout(getResults, 200);
        }
    },
    components: { MailList }
}
</script>
