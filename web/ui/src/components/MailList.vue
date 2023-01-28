<template>
  <div class="flex flex-row">
    <table class="w-[60rem] border">
      <thead>
        <tr class="text-center">
          <th class="flex flex-row items-start pl-4">
            Subject
          </th>
          <th class="w-[15rem]">From</th>
          <th class="w-[15rem]">To</th>
        </tr>
      </thead>
      <tbody v-if="mails.length > 0">
        <tr v-for="mail in mails" :key="mail" @click="selectMail(mail)" class="border w-full">
          <td class="flex flex-row items-start pl-4 truncate max-w-xs">{{ mail.Subject }}</td>
          <td>{{ mail.From }}</td>
          <td class="truncate max-w-sm">{{ mail.To }}</td>
        </tr>
      </tbody>
      <tbody v-else>
        <tr class="border">
          <td colspan="3" class="w-full">
            No mails to display
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
  
<script>
  export default {
    props: ['mails'],
    methods: {
      selectMail(mail) {
        this.$emit('selectMail', mail);
      }
    }
  }
</script>