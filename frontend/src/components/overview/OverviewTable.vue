<template>
  <v-card>
    <v-simple-table fixed-header :class="{ 'tight-table': $vuetify.breakpoint.xs }">
      <template v-slot:default>
        <thead>
          <tr>
            <th>Monat</th>
            <th>Fixe Kosten</th>
            <th>Sonder Kosten</th>
            <th>Saldo</th>
            <th />
          </tr>
        </thead>
        <tbody>
          <tr :key="index" v-for="(entry, index) in entries">
            <td>{{ entry | displayMonth }}</td>
            <td>{{ entry.sumFixedCosts | currency | responsive }}</td>
            <td>{{ entry.sumSpecialCosts | currency | responsive }}</td>
            <td
              class="amount"
              :class="{ 'negative-amount': entry.currentAmount < 0}"
            >{{ entry.currentAmount | currency | responsive }}</td>
            <td align="right" class="action-cell">
              <overview-details 
                v-if="!entry.empty" 
                :detail="{ ...entry, index }" 
                @success="success"
              />
              <special-cost-form 
                :cost="newEmptyCost(index, entry.yearMonth)" 
                icon="fa-plus-square"
                @success="success"
              />
            </td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
    <v-snackbar 
      v-model="snackbar" 
      bottom 
      color="success" 
      :timeout="7000">
        {{ successMsg }}
    </v-snackbar>
  </v-card>
</template>
<script>
import OverviewDetails from "./OverviewDetails";
import { displayMonth } from "../Utils";
import SpecialCostForm from "../editform/SpecialCostForm";

export default {
  props: ["entries"],
  components: {
    OverviewDetails,
    SpecialCostForm
  },
  data() {
    return {
      month: "No Month",
      
      snackbar: false,
      successMsg: ""
    };
  },
  filters: {
    displayMonth: ({ yearMonth }) => displayMonth(yearMonth),

    responsive: function(value) {
      if (value.length > 5 && window.innerWidth < 768) {
        var tmp = value.split(" ");
        var rest = tmp[0].substring(0, tmp[0].length - 2);
        return rest + " T" + tmp[1];
      } else {
        return value;
      }
    }
  },
  methods: {
    newEmptyCost(index, dueDate) {
      return {
        index,
        name: "",
        amount: 0,
        dueDate
      };
    },
    
    success: async function({name, cost, created}) {
      console.log('success', {name, cost, created})

      this.snackbar = true;
      this.successMsg = 
        `${name} '${cost.name}' erfolgreich ${created ? "hinzugefügt" : "geändert"}`;
      
      this.$emit("changed");
    },
  }
};
</script>