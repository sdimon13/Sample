<template>
  <div>
    <b-table :items="services" :fields="fields" @row-clicked="selectService">
      <template v-slot:cell(select)="data">
        <b-form-checkbox v-model="data.item.selected"></b-form-checkbox>
      </template>
    </b-table>
    <b-button @click="submitServices" class="mt-3">Submit Selected Services</b-button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  "name": "ServicesView",
  data() {
    return {
      services: [],
      fields: [
        { key: 'name', label: 'Название услуги' },
        { key: 'price', label: 'Цена' },
        { key: 'time', label: 'Время' },
        { key: 'select', label: '', sortable: false }
      ]
    }
  },
  created() {
    axios.get('http://localhost:8682/api/services')
        .then(response => {
          this.services = response.data.list;
        })
        .catch(error => {
          console.log(error);
        });
  },
  methods: {
    selectService(item) {
      item.selected = !item.selected;
    },
    submitServices() {
      let selectedServices = this.services.filter(service => service.selected);
      axios.post('/api/submit-services', { services: selectedServices })
          .then(response => {
            console.log(response);
          })
          .catch(error => {
            console.log(error);
          });
    }
  }
}
</script>