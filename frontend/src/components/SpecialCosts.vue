<template>
  <v-container>
    <v-row>
      <v-col>
        <v-skeleton-loader
          type="table-tbody"
          :loading="!loaded"
          transition="scale-transition"
          class="mx-auto"
        >
          <v-card>
            <v-card-text>
              <cost-table
                :entries="entries"
                :cols="cols"
                @success="success"
                @delete="deleteEntry"
              >
                <template v-slot:edit-button="slotProps">
                  <special-cost-form 
                    :cost="slotProps.entry"
                    @success="success"
                  />    
                </template>
              </cost-table>
            </v-card-text>
            <v-card-actions>
              <special-cost-form 
                btn-text="Neue Sonderkosten Hinzufügen" 
                @success="success"
              />
            </v-card-actions>
          </v-card>
        </v-skeleton-loader>
        <v-snackbar 
          v-model="snackbar" 
          bottom 
          color="success" 
          :timeout="7000">
            {{ successMsg }}
        </v-snackbar>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import LoadablePage from "./LoadablePage";
import CostTable from "./CostTable";

import {
  toCurrency,
  displayMonth
} from "./Utils";
import SpecialCostForm from './editform/SpecialCostForm.vue';

const cols = [
  { name: "name", label: "Bezeichnung" },
  { name: "amount", label: "Betrag", transformer: toCurrency },
  { name: "dueDate", label: "Fällig am", transformer: displayMonth }
];

export default {
  mixins: [LoadablePage],
  components: {
    CostTable,
    SpecialCostForm
  },
  data() {
    return {
      entries: [],
      cols,

      snackbar: false,
      successMsg: ""
    };
  },
  created: async function() {
    this.loadData();
  }, 
  methods: {
    loadData: async function() {
      const data = await this.fetchData("/api/specialcosts");
      this.entries = data;
      this.loaded = true;
    }, 
    success: async function({name, cost, created}) {
      this.snackbar = true;
      this.successMsg = 
        `${name} '${cost.name}' erfolgreich ${created ? "hinzugefügt" : "geändert"}`;
      
      this.loadData();
    },
    deleteEntry: async function({id, name}) {
      await fetch("/api/specialcosts/" + id, {
        method: "DELETE"
      });

      this.snackbar = true;
      this.successMsg = `'${name}' wurde erfolgreich gelöscht`;

      this.loadData();
    },
  }
};
</script>