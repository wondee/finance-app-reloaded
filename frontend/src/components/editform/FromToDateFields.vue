<template>
  <v-row>
    <v-col>
      <month-date-picker label="Gültig ab" v-model="value.from" :max="value.to" />
    </v-col>
    <v-col>
      <month-date-picker label="Gültig bis" v-model="value.to" :rules="toDateRules" :min="value.from" />
    </v-col>
  </v-row>
</template>
<script>
import MonthDatePicker from './MonthDatePicker'
import { createDateString } from "../Utils";

export default {
  props: ['value'],
  components: {
    MonthDatePicker
  },
  data() {
    return {
      toDateRules: [
        d =>
          d &&
          this.value.from &&
          new Date(createDateString(d)) < new Date(createDateString(this.value.from)) &&
          "'Gültig bis' darf nicht kleiner als 'Gültig ab' sein"
      ]
    }
  }
}
</script>
