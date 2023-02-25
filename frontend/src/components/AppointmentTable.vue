<template>
  <table>
    <thead>
    <tr>
      <th>Время</th>
      <th>Имя</th>
      <th>Телефон</th>
    </tr>
    </thead>
    <tbody>
    <tr v-for="appointment in appointments" :key="appointment.id">
      <td>{{ appointment.time }}</td>
      <td>{{ appointment.name }}</td>
      <td>{{ appointment.phone }}</td>
    </tr>
    </tbody>
  </table>
</template>
<script>
import axios from 'axios';

export default {
  name: 'AppointmentTable',
  props: {
    selectedDate: String
  },
  data() {
    return {
      appointments: []
    };
  },
};
</script>