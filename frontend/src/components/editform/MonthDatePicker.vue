<template>
  <v-menu
    v-model="menu"
    :close-on-content-click="false"
    :nudge-right="40"
    transition="scale-transition"
    offset-y
    min-width="290px"
  >
    <template v-slot:activator="{ on }">
      <v-text-field
        :value="displayDate"
        :label="label"
        prepend-icon="fa-calendar-week"
        :error-messages="errorMessages"
        :clearable="!required"
        clear-icon="mdi-close"
        readonly
        v-on="on"
      />
    </template>
    <v-date-picker
      :value="date"
      @input="input"
      type="month"
      color="blue"
      :min="minDate"
      :max="maxDate"
    />
  </v-menu>
</template>
<script>
import { displayMonth, createDateString } from "../Utils";

const nowDate = new Date();

const requiredRule = v => !!v || 'Datum darf nicht leer sein';

export default {
  props: ["value", "min", "max", "label", "rules", "required"],

  data() {
    return {
      menu: false,
      now: `${nowDate.getFullYear()}-${nowDate.getMonth() + 1}`,
      errorMessages: []
    };
  },
  computed: {
    displayDate() {
      return displayMonth(this.value, false, null);
    },
    date() {
      return this.value ? createDateString(this.value) :  null;
    },
    minDate() {
      return this.min ? createDateString(this.min) : this.now;
    },
    maxDate() {
      return this.max ? createDateString(this.max) : "2100-12";
    },
    internalRules() {
      return [(this.required ? requiredRule : () => false), ... (this.rules || [])]
    }
  },
  methods: {
    input(e) {
      console.log('input', this.internalRules)
      const createYearMonth = () => {
        const elems = e.split('-')
        return {year: Number(elems[0]), month: Number(elems[1])}
      }

      const inputYearMonth = e ? createYearMonth() : e;
      this.$emit("input", inputYearMonth);
      this.menu = false;

      this.errorMessages = [];

      if (!this.internalRules) return;

      this.internalRules.forEach(rule => {
        const result = rule(e);
        if (typeof result === "string") {
          this.errorMessages.push(result);
        }
      });
    }
  }
};
</script>