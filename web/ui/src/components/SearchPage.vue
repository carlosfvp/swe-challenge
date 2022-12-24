<template>
  <h1>MamuroEmail</h1>
  <input class="w-32" type="text" @input="search" v-model="searchTerm" placeholder="type to search by to, from, subject, content" />
  <h1 v-if="error">{{ error }}</h1>
  <div class="flex flex-row">
    <table>
      <tr>
        <th>Subject</th>
        <th>From</th>
        <th>To</th>
      </tr>
      <tr v-for="mail in mails" :key="mail" @click="selectMail(mail.Body)">
        <td>{{ mail.Subject }}</td>
        <td>{{ mail.From }}</td>
        <td>{{ mail.To }}</td>
      </tr>
    </table>
    {{ body }}
  </div>
</template>

<script>
export default {
  name: 'SearchPage',
  data() {
    return {
      searchTerm: '',
      mails: [],
      error: '',
      timeoutHandler: null,
      body: ''
    }
  },
  methods: {
    selectMail(body) {
      this.body = body;
    },
    search() {
      if (this.searchTerm.trim() == '') return;
      if (this.timeoutHandler) clearTimeout(this.timeoutHandler);
      const getResults = () => {
        fetch(`http://localhost:3000/api/search/${this.searchTerm}`)
          .then(response => response.json())
          .then(json => {
            this.mails = json.Matches;
          })
          .catch(error => {
            this.error = error;
          });
      }
      this.timeoutHandler = setTimeout(getResults, 200);
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
