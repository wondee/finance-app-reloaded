<template>
  <v-text-field
    :value="displayValue"
    @blur="focus = false;"
    @focus="focus = true"
    @change="inputChanged"
    required
    :label="label"
  />
</template>
<script>
import { toCurrency } from "../Utils";

export default {
  props: ["id", "label", "value"],
  data() {
    return {
      focus: false,
    }
  },
  
  computed: {
    displayValue() {
      return this.focus ? this.value : toCurrency(this.value);
    }
  },
  methods: {
    inputChanged(e) {
      const newValue = Number(e);
      
      this.$emit("input", isNaN(newValue) ? 0 : newValue);
    }
  }
};
</script>