<template>
  <cost-edit-form
    :title="title('Jährliche Kosten')"
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
        <v-select :items="items" v-model="form.month" label="Fällig im" />
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
import {
  CommonForm,
  monthlyCostToForm,
  toSelectItems,
  monthMap,
  baseFormToCost
} from "../Utils";
import CostEditForm from "./CostEditForm";
import NameTextField from "./NameTextField";
import FromToDateFields from "./FromToDateFields";
import IncomingSelect from "./IncomingSelect";

const formToCost = form => ({
  ...baseFormToCost(form),
  ...form.fromTo,
  dueMonth: form.month
})

const costToForm = cost => {
  const form = monthlyCostToForm(cost);
  return !cost
    ? {
        ...form,
        month: null
      }
    : {
        ...form,
        month: cost.dueMonth - 1
      };
};

export default {
  mixins: [CommonForm(costToForm, formToCost, 'Jährliche Kosten', '/api/costs/yearly')],
  components: {
    CostEditForm,
    NameTextField,
    CurrencyInput,
    FromToDateFields,
    IncomingSelect
  },
  props: ["btnText"],
  data() {
    return {
      items: toSelectItems(monthMap)
    };
  }
};
</script>