<template>
  <v-card>
    <v-card-text>
      <cost-table :entries="entries" :cols="cols" v-on="$listeners">
        <template v-if="$vuetify.breakpoint.smAndDown" v-slot:header>
          <th />
        </template>
        <template v-if="$vuetify.breakpoint.smAndDown" v-slot:content="slotProps">
          <responsive-date-col :entry="slotProps.entry" />
        </template>
        <template v-slot:edit-button="slotProps">
          <component 
            :is="formComponent" 
            :cost="slotProps.entry" 
            v-on="$listeners"
          />
        </template>
      </cost-table>
    </v-card-text>

    <v-card-actions>
      <component 
        :is="formComponent" 
        btn-text="Neue Kosten Hinzufügen"
        v-on="$listeners" 
      />
    </v-card-actions>
  </v-card>
</template>
<script>
import CostTable from "./CostTable";
import ResponsiveDateCol from "./ResponsiveDateCol";
import MonthlyCostEditForm from "./editform/MonthlyCostEditForm";
import QuarterlyCostEditForm from "./editform/QuarterlyCostEditForm";
import HalfyearlyCostEditForm from "./editform/HalfyearlyCostEditForm";
import YearlyCostEditForm from "./editform/YearlyCostEditForm";


export default {
  props: ["entries", "cols", "formComponent"],

  components: {
    CostTable,
    ResponsiveDateCol,
    MonthlyCostEditForm,
    QuarterlyCostEditForm,
    HalfyearlyCostEditForm,
    YearlyCostEditForm
  }
};
</script>