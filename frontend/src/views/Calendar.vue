<template>
  <div class="overlay">
    <b-form @submit.prevent="submitForm" class="form">
      <div class="con">
        <header class="head-form">
          <h2>Онлайн запись</h2>
        </header>
        <div class="field-set">
          <b-form-group id="input-group-1" label="Выберите дату">
            <b-form-datepicker v-model="selectedDate" @input="fetchAvailableDatesAndTimes()" :date-disabled-fn="dateDisabledFn" class="calendar-datepicker"></b-form-datepicker>
          </b-form-group>
          <b-form-group id="input-group-2" v-if="timesForSelectedDate.length" label="Выберите время"
                        class="calendar-available-times">
            <b-btn-group class="d-flex">
              <b-btn class="button-time" v-for="time in timesForSelectedDate" :key="time.id"
                     :class="{ 'btn-active': selectedTime === time }"
                     @click="selectedTime = time">
                {{ time.time }}
              </b-btn>
            </b-btn-group>
          </b-form-group>
          <b-form-group id="input-group-3">
            <b-form-input class="calendar-input" type="text" v-model="name" placeholder="Имя"></b-form-input>
          </b-form-group>
          <b-form-group id="input-group-4">
            <b-form-input class="calendar-input" type="tel" v-model="phoneNumber" placeholder="Номер телефона" v-mask="'+7 (###) ###-##-##'">
              <vue-the-mask mask="+7 (###) ###-##-##" />
            </b-form-input>
          </b-form-group>
        </div>
        <div class="other">
          <b-button class="calendar-button" type="submit" variant="primary">Записаться</b-button>
        </div>
      </div>
    </b-form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: "CalendarView",
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
      }
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
        const response = await axios.get('http://localhost:8682/api/appointments');
        const availableDates = response.data.availableDates;
        const availableTimes = {};

        // Преобразование временных слотов в объект, содержащий массив времен для каждой даты
        for (const [date, { timeSlots }] of Object.entries(response.data.availableTimes)) {
          availableTimes[date] = { times: [] };
          timeSlots.forEach(({ id, time }) => {
            availableTimes[date].times.push({ id, time });
          });
        }

        this.availableDates = availableDates;
        this.availableTimes = availableTimes;
      } catch (error) {
        console.error(error);
      }
    },
    async submitForm() {
      const appointment = {
        appointmentId: this.selectedTime.id,
        date: this.selectedDate,
        time: this.selectedTime.time,
        name: this.name,
        phoneNumber: this.removeFormatting(this.phoneNumber)
      };
      try {
        await axios.patch(`http://localhost:8682/api/appointments/${appointment.appointmentId}`, null, {
          params: {
            date: appointment.date,
            time: appointment.time,
            name: appointment.name,
            phoneNumber: appointment.phoneNumber
          }
        });
        alert('Вы успешно записаны!');
        await this.fetchAvailableDatesAndTimes();
      } catch (error) {
        alert('Упс, кажется уже кто-то занял Ваше время, попробуйте выбрать другое.');
        await this.fetchAvailableDatesAndTimes();
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