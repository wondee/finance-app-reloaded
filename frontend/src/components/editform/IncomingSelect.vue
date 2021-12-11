<template>
  <v-checkbox
    :input-value="checked"
    @change="changed"
    hint="Anhacken, wenn es sich bei dem Betrag um einen eingehenden (z.B. Gehalt) handelt."
    persistent-hint
    label="Eingehend"
    color="primary"
  />
</template>
<script>
export default {
  props: ["value"],
  data() {
    return {
      checked: !!this.value
    }
  }, 
  methods: {
    changed(e) {
      this.$emit("input", e);
    }
  }
};
</script>


    