<template>
  <cost-edit-form
    :title="title('Sonderkosten Kosten')"
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
      <v-col>
        <month-date-picker 
          v-model="form.dueDate" 
          label="Fällig am" 
          required="true"
        />
      </v-col>
    </v-row>
  </cost-edit-form>
</template>
<script>
import CostEditForm from "./CostEditForm";
import NameTextField from "./NameTextField";
import CurrencyInput from "./CurrencyInput";
import IncomingSelect from "./IncomingSelect";
import MonthDatePicker from "./MonthDatePicker";
import { monthlyCostToForm, CommonForm, baseFormToCost } from "../Utils";

const costToForm = cost => {
  const form = monthlyCostToForm(cost);
  return !cost
    ? {
        ...form,
        dueDate: null
      }
    : {
        ...form,
        dueDate: cost.dueDate
      };
};

const formToCost = form => ({
  ...baseFormToCost(form),
  dueDate: form.dueDate
})

export default {
  mixins: [CommonForm(costToForm, formToCost, 'Sonderkosten', '/api/specialcosts')],
  components: {
    CostEditForm,
    NameTextField,
    CurrencyInput,
    IncomingSelect,
    MonthDatePicker
  },
  props: ["btnText"]
};
</script>