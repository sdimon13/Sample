<template>
  <div class="overlay">
    <b-form class="form">
      <div class="con">
        <header class="head-form">
          <h2>Просмотр записей</h2>
        </header>
        <div class="field-set">
          <b-form-group id="input-group-1" label="Выберите дату">
            <b-form-datepicker v-model="selectedDate" @input="fetchAvailableDatesAndTimes()" :date-disabled-fn="dateDisabledFn" class="calendar-datepicker"></b-form-datepicker>
          </b-form-group>
          <b-table v-if="timesForSelectedDate.length" :items="timesForSelectedDate" :fields="fields" :per-page="perPage" class="calendar-available-times">
            <template #cell(time)="data">
              {{ data.value }}
            </template>
            <template #cell(name)="data">
              {{ data.value }}
            </template>
            <template #cell(phone)="data">
              {{ data.value }}
            </template>
          </b-table>
        </div>
      </div>
    </b-form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "AdminCalendarView",
  data() {
    return {
      selectedDate: '',
      selectedTime: '',
      name: '',
      phoneNumber: '',
      availableDates: [],
      availableTimes: {},
      dateCellClasses: {
        'available-date': date => !this.dateDisabledFn(date),
        'unavailable-date': date => this.dateDisabledFn(date),
      },
      fields: ['time', 'name', 'phone'],
      perPage: 10,
    };
  },
  computed: {
    timesForSelectedDate() {
      return this.availableTimes[this.selectedDate]?.times || [];
    }
  },
  created() {
    this.fetchAvailableDatesAndTimes().then(() => {
      this.selectedDate = this.availableDates[0] || '';
    });
  },
  methods: {
    async fetchAvailableDatesAndTimes() {
      try {
        const response = await axios.get('http://localhost:8682/api/admin/appointments');
        const availableDates = response.data.availableDates;
        const availableTimes = {};

        // Преобразование временных слотов в объект, содержащий массив времен для каждой даты
        for (const [date, { timeSlots }] of Object.entries(response.data.availableTimes)) {
          availableTimes[date] = { times: [] };
          timeSlots.forEach(({ time, name, phone }) => {
            availableTimes[date].times.push({ time, name, phone });
          });
        }

        this.availableDates = availableDates;
        this.availableTimes = availableTimes;
      } catch (error) {
        console.error(error);
      }
    },
    dateDisabledFn(date) {
      return this.availableDates.length && !this.availableDates.includes(date);
    },
    removeFormatting(phoneNumber) {
      return phoneNumber.replace(/\D/g, '');
    }
  }
};
</script>