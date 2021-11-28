<template>
  <cost-edit-form
    :title="title('Monatliche Kosten')"
    :changed="changed"
    :btn-text="btnText"
    @save="saveCost"
  >
    <v-row>
      <v-col>
        <name-text-field v-model="form.name" />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <currency-input label="Betrag" v-model="form.amount" />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <incoming-select v-model="form.incoming" />
      </v-col>
    </v-row>
    <v-row>
      <from-to-date-fields v-model="form.fromTo" />
    </v-row>
  </cost-edit-form>
</template>
<script>
import CurrencyInput from "./CurrencyInput";
import { baseFormToCost, CommonForm, monthlyCostToForm } from "../Utils";
import CostEditForm from "./CostEditForm";
import NameTextField from "./NameTextField";
import FromToDateFields from "./FromToDateFields";
import IncomingSelect from "./IncomingSelect";

const formToCost = form => ({
  ...baseFormToCost(form),
  ...form.fromTo
})

export default {
  mixins: [CommonForm(monthlyCostToForm, formToCost, 'Monatliche Kosten', '/api/costs/monthly')],
  components: {
    CostEditForm,
    NameTextField,
    CurrencyInput,
    FromToDateFields,
    IncomingSelect
  },
  props: ["btnText"]
};
</script>