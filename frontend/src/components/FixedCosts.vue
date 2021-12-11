<template>
  <v-container>
    <v-row>
      <v-col>
        <v-skeleton-loader
          type="card-heading"
          :loading="!loaded"
          transition="scale-transition"
          class="mx-auto"
        >
          <v-banner sticky icon="fa-wallet" elevation="4">
            Aktuelle Bilanz (pro Monat):
            <strong
              :class="{ negative : currentBalance < 0 }"
            >{{ currentBalanceDisplay }}</strong>
          </v-banner>
        </v-skeleton-loader>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <v-skeleton-loader
          type="table-tbody"
          :loading="!loaded"
          transition="scale-transition"
          class="mx-auto"
        >
          <v-card no-body>
            <v-tabs v-model="tab" grow>
              <v-tabs-slider />
              <v-tab>Monatliche Kosten</v-tab>
              <v-tab-item>
                <fixed-costs-table
                  :entries="monthly"
                  :cols="monthlyCols"
                  formComponent="monthly-cost-edit-form"
                  @success="success"
                  @delete="deleteEntry"
                />
              </v-tab-item>
              <v-tab>Vierteljährliche Kosten</v-tab>
              <v-tab-item>
                <fixed-costs-table
                  :entries="quarterly"
                  :cols="quarterlyCols"
                  formComponent="quarterly-cost-edit-form"
                  @success="success"
                  @delete="deleteEntry"
                />
              </v-tab-item>
              <v-tab>Halbjährliche Kosten</v-tab>
              <v-tab-item>
                <fixed-costs-table
                  :entries="halfyearly"
                  :cols="halfyearlyCols"
                  formComponent="halfyearly-cost-edit-form"
                  @success="success"
                  @delete="deleteEntry"
                />
              </v-tab-item>
              <v-tab>Jährliche Kosten</v-tab>
              <v-tab-item>
                <fixed-costs-table
                  :entries="yearly"
                  :cols="yearlyCols"
                  formComponent="yearly-cost-edit-form"
                  @success="success"
                  @delete="deleteEntry"
                />
              </v-tab-item>
            </v-tabs>
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
import FixedCostsTable from "./FixedCostTable";
import LoadablePage from "./LoadablePage";
import {
  displayMonth,
  toCurrency,
  toQuarterlyDueDate,
  toHalfyearlyDueDate,
  toMonth
} from "./Utils";

const defaultCols = [
  { name: "name", label: "Bezeichnung" },
  { name: "amount", label: "Betrag", transformer: toCurrency },
  {
    name: "from",
    label: "Gültig ab",
    transformer: displayMonth,
    hide: true
  },
  { name: "to", label: "Gültig bis", transformer: displayMonth, hide: true }
];

function cols(additionalCols = false) {
  if (!additionalCols) {
    return defaultCols;
  }
  const cols = [...defaultCols];
  cols.splice(1, 0, ...additionalCols);
  return cols;
}

const quarterlyCols = cols([
  { name: "dueMonth", label: "Fällig in", transformer: toQuarterlyDueDate }
]);

const halfyearlyCols = cols([
  { name: "dueMonth", label: "Fällig in", transformer: toHalfyearlyDueDate }
]);

const yearlyCols = cols([
  { name: "dueMonth", label: "Fällig im", transformer: toMonth }
]);


export default {
  mixins: [LoadablePage],
  components: {
    FixedCostsTable
  },
  data() {
    return {
      tab: null,

      monthly: [],
      quarterly: [],
      halfyearly: [],
      yearly: [],

      currentBalance: -1,

      monthlyCols: cols(),
      quarterlyCols,
      halfyearlyCols,
      yearlyCols,

      snackbar: false,
      successMsg: ""
    };
  },
  computed: {
    currentBalanceDisplay() {
      return `${this.currentBalance} €`;
    }
  },
  methods: {
    success: async function({name, cost, created}) {
      this.snackbar = true;
      this.successMsg = 
        `${name} '${cost.name}' erfolgreich ${created ? "hinzugefügt" : "geändert"}`;

      this.loadData();
    },
    deleteEntry: async function({id, name}) {
      await fetch("/api/costs/" + id, {
        method: "DELETE"
      });

      this.snackbar = true;
      this.successMsg = `'${name}' wurde erfolgreich gelöscht`;

      this.loadData();
    },
    loadData: async function() {
      this.monthly = [];
      this.quarterly = [];
      this.halfyearly = [];
      this.yearly = [];

      const data = await this.fetchData("/api/costs");

      this.monthly = data.monthly;
      this.quarterly = data.quarterly;
      this.halfyearly = data.halfyearly;
      this.yearly = data.yearly;

      this.currentBalance = data.currentBalance; 
    }
  },

  created: async function() {
    this.loadData();
  }
};
</script>

<style scoped>
strong.negative {
  color: indianred;
}
.tabs {
  width: 100%;
}
</style>